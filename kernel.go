package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

type modpack struct {
	Name        string
	URL         string
	AutoUpdater bool
	Version     string
	sha1sum     string
}

func main() {
	var modpack modpack
	data, err := ioutil.ReadFile("./modpack.json")
	if err != nil {
		failExit(err)
	}

	err = json.Unmarshal(data, &modpack)
	if err != nil {
		failExit(err)
	}

	fmt.Println("SweetNyanCraft Modpack Updater Version 1.0\n" +
		"Github: https://github.com/Marfjeh/Modpack-Updater\n" +
		"Licence: GNU Version 2.0\n" +
		"------------------------------------------\n" +
		"Modpack: " + modpack.Name + "\n" +
		"Modpack version: " + modpack.Version + "\n" +
		"------------------------------------------")

	fmt.Printf("Checking for updated version...")

	//if err := DownloadFile("deployment.zip", ""); err != nil {
	//failExit(err)
	//}
}

func failExit(err error) {
	fmt.Println("--------- ERROR ---------\n" + err.Error())
	os.Exit(1)
}

// DownloadFile will download a url to a local file. It's efficient because it will
// write as it downloads and not load the whole file into memory.
func DownloadFile(filepath string, url string) error {

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
