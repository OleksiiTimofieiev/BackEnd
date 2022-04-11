package main

import (
	"GoogleDriveUpload/archivation"
	"GoogleDriveUpload/configs"
	"fmt"
	"os"
)

// TODO: one folder passed as command line argument, move binary to PATH

const (
	// TODO: use config for vars below with relative path
	pathWork  = "/home/user/Desktop/"
	uploadDir = "upload"
)

var (
	config configs.Config
)

var (
	Blue   = "\033[34m"
	Yellow = "\033[33m"
	Reset  = "\033[0m"
)

// TODO: https://www.youtube.com/watch?v=P9CjdN93iHE&t=241s
// TODO: add cron job
// TODO: remove zip file after work was done

func main() {
	fmt.Println(Blue + "--- Uploading started ---" + Reset)

	// create upload dir
	if err := os.Mkdir(uploadDir, os.ModePerm); err != nil {
		panic(err)
	}

	// read uploadable files
	configs.ReadConfigs(&config)

	archivation.RecursiveZip(pathWork+config.Work, uploadDir+"/"+config.Work)

	fmt.Println(Yellow + "--- Uploading finished ---" + Reset)
}
