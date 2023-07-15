/*
	Modpack Updater
    Copyright (C) 2023  Marfjeh

    This program is free software: you can redistribute it and/or modify
    it under the terms of the GNU General Public License as published by
    the Free Software Foundation, either version 3 of the License, or
    (at your option) any later version.

    This program is distributed in the hope that it will be useful,
    but WITHOUT ANY WARRANTY; without even the implied warranty of
    MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
    GNU General Public License for more details.

    You should have received a copy of the GNU General Public License
    along with this program.  If not, see <https://www.gnu.org/licenses/>.
*/

package main

import (
	"archive/zip"
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"
)

type modpack struct {
	Name        string `json:"Name"`
	URL         string `json:"URL"`
	AutoUpdater bool   `json:"AutoUpdater"`
	Version     string `json:"Version"`
	Sha1sum     string `json:"sha1sum"`
}

type modpackupdate struct {
	Version string `json:"version"`
	Sha1    string `json:"sha1"`
}

func main() {
	var modpack modpack
  args := os.Args[1:]
  
  if len(args) != 0 {
	  if os.Args[1] == "-?" {
		  fmt.Println("Command arguments:")
		  fmt.Println("-f to force update.")
		  os.Exit(0)
	  }
  }

	//Read the local modpack file.
	data, err := ioutil.ReadFile("./modpack.json")
	err = json.Unmarshal(data, &modpack)

	//Print logo.
	fmt.Println("Modpack Updater Version 1.2\n" +
		"Github: https://github.com/Marfjeh/Modpack-Updater\n" +
		"-----------------------------------------------------------------------------\n" +
		"Modpack: " + modpack.Name + "\n" +
		"Modpack version: " + modpack.Version + "\n" +
		"-----------------------------------------------------------------------------\n" +
		"Copyright (C) 2023  Marfjeh\n" +
		"This program comes with ABSOLUTELY NO WARRANTY; for details see the git repo.\n" +
		"This is free software, and you are welcome to redistribute it\n" +
		"under certain conditions; see the git repo for details.\n" +
		"-----------------------------------------------------------------------------\n")

	//Init Update check
	fmt.Printf("Checking for updated version... ")

	if !modpack.AutoUpdater {
		fmt.Println(" [ FAILED ]")
		fmt.Println("Auto update is disabled in modpack.json exiting.")
		fmt.Println("If you want to re-enable auto update, please set AutoUpdate to true in modpack.json file.")
		os.Exit(0)
	}

	Modpackupdate := checkupdate(modpack.Sha1sum, modpack.Version)
	err = downloadFile("deployment.zip", modpack.URL)
	fmt.Println(" [ OK ]")

	fmt.Printf("Verifying checksum of deployment...")
	ziphash := getHash()
	if ziphash != Modpackupdate.Sha1 {
		fmt.Println(" [ FAILED ]")

		fmt.Println("Detected checksum: " + ziphash)
		fmt.Println("Expected checksum: " + ziphash)
		os.Exit(1)
	} else {
		fmt.Println(" [ OK ]")

		fmt.Println("Detected checksum: " + ziphash)
		fmt.Println("Expected checksum: " + ziphash)
	}

	fmt.Printf("Cleaning mods folder...")
	err = cleanFolders()
	fmt.Println(" [ OK ]")

	fmt.Printf("Extracting update package...")
	_, err = extractZip()
	fmt.Println(" [ OK ]")

	fmt.Printf("Cleaning up... ")
	err = os.Remove("deployment.zip")
	fmt.Println(" [ OK ]")

	fmt.Println("Updating is complete. you can now start up your modpack.")

	if err != nil {
		failExit(err)
	}

	os.Exit(0)
}

func failExit(err error) {
	log.Fatal("--------- ERROR ---------\n" + err.Error())
	os.Exit(1)
}

func checkupdate(url string, version string) modpackupdate {
  args := os.Args[1:]

	ModpackUpdate := modpackupdate{}
	getJSON(url, &ModpackUpdate)

  if len(args) != 0 {
	  if os.Args[1] == "-f" {
		  version = "force update"
	  }
  }

	if version == ModpackUpdate.Version {
		fmt.Println("[ Already up-to-date ]")
		fmt.Println("Exiting, you can now play the modpack!")
		os.Exit(0)
	} else {
		fmt.Println("[ Update found ]")
	}

	return ModpackUpdate
}

var myClient = &http.Client{Timeout: 60 * time.Second}

func getJSON(url string, target interface{}) error {
	r, err := myClient.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(target)
}

// downloadFile will download a url to a local file. It's efficient because it will
// write as it downloads and not load the whole file into memory.
func downloadFile(filepath string, url string) error {

	fmt.Printf("Downloading update. This can take some time depending on your internet speed...")
	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err
}

// Clean the mods folder so we can move in the new mods.
func cleanFolders() error {
	moddir, err := ioutil.ReadDir("mods/")
	for _, d := range moddir {
		err = os.RemoveAll(path.Join([]string{"mods", d.Name()}...))
		if err != nil {
			return err
		}
	}

	return nil
}

// Extract the zip containing the files.
func extractZip() ([]string, error) {
	var filenames []string

	r, err := zip.OpenReader("deployment.zip")
	if err != nil {
		failExit(err)
	}

	defer r.Close()

	for _, f := range r.File {
		dest, _ := os.Getwd()
		fpath := filepath.Join(string(dest), f.Name)

		// Check for ZipSlip. More Info: http://bit.ly/2MsjAWE
		if !strings.HasPrefix(fpath, filepath.Clean(dest)+string(os.PathSeparator)) {
			return filenames, fmt.Errorf("%s: illegal file path", fpath)
		}

		filenames = append(filenames, fpath)
		if f.FileInfo().IsDir() {
			os.MkdirAll(fpath, os.ModePerm)
			continue
		}

		if err = os.MkdirAll(filepath.Dir(fpath), os.ModePerm); err != nil {
			return filenames, err
		}
		outFile, err := os.OpenFile(fpath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
		if err != nil {
			return filenames, err
		}

		rc, err := f.Open()
		if err != nil {
			return filenames, err
		}

		_, err = io.Copy(outFile, rc)

		// Close the file without defer to close before next iteration of loop
		outFile.Close()
		rc.Close()

		if err != nil {
			return filenames, err
		}
	}
	return filenames, nil
}

// Calcuate the hash of the deployment.zip to verify that the zip is downloaded successfully.
// So people with awful internet connection wont get currupted files.
func getHash() string {
	f, err := os.Open("deployment.zip")
	if err != nil {
		fmt.Println(" [ Failed ]")
		failExit(err)
	}

	defer f.Close()

	h := sha1.New()
	if _, err := io.Copy(h, f); err != nil {
		fmt.Println(" [ Failed ]")
		failExit(err)
	}

	return hex.EncodeToString(h.Sum(nil))
}
