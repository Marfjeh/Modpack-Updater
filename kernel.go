package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type modpack struct {
	URI          string
	name         string
	version      string
	remoteserver string
	downloadurl  string
	autoupdate   bool
}

func main() {
	fmt.Println("SweetNyanCraft Modpack updater Version 0.2 \n")
	if err := DownloadFile("deployment.zip", ""); err != nil {
		failExit(err)
	}
}

func populateModpackStruct() {
	//Read the local modpack.json on the current disk.
	//Then popluate the modpack struct.
	var modpack modpack

}

func failExit(err error) {
	fmt.Println("----------------------\nERROR\n----------------------\n" + err.Error())
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
