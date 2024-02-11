package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
)

type Config struct {
	remoteFile string
	savePath string
}

func getEnv(envName string, obligatory bool) string {
	// read environment variables
	env, ok := os.LookupEnv(envName)
	if !ok && obligatory {
		fmt.Println(envName, "not set")
		os.Exit(0)
	}
	return env
}

func configuration() (Config) {
	// prepare configuration
	savePath := getEnv("SAVE_PATH", false)
	if savePath == "" {
		savePath = "/tmp"
	}		
	config := Config{
		remoteFile: getEnv("REMOTE_FILE", true),
		savePath: savePath,
	}
	return config
}

func downloadFile(config Config) {
	// download file
	resp, err := http.Get(config.remoteFile)
	if err != nil {
		fmt.Println("Error downloading file:", err)
		return
	}
	defer resp.Body.Close()
	// filename from url
	filename := path.Base(config.remoteFile)
	if resp.StatusCode != http.StatusOK {
		fmt.Println("Error downloading file:", resp.Status)
		return
	}
	// save file
	file, err := os.Create(config.savePath + "/" + filename)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	_, err = io.Copy(file, resp.Body)
	if err != nil {
		fmt.Println("Error saving file:", err)
		return
	}
	fmt.Println("File saved:", config.savePath + "/" + filename)
}


func main() {
	// read environment variables
	config := configuration()
	fmt.Printf("remoteFile: %s\n", config.remoteFile)
	fmt.Printf("savePath: %s\n", config.savePath)
	// download file
	downloadFile(config)
}
