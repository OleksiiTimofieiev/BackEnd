package main

import (
	"GoogleDriveUpload/archivation"
	"GoogleDriveUpload/configs"
	"fmt"
	"os"
	// "sync"
)

// TODO: one folder passed as command line argument, move binary to PATH
const (
	// pathBackEnd = "/home/otimofieiev/Desktop/BackEnd/"
	// TODO: use config for vars below with relative path
	pathWork  = "/home/otimofieiev/Desktop/"
	uploadDir = "upload"
)

var (
	// wg     sync.WaitGroup/
	config configs.Config
)

var (
	Blue   = "\033[34m"
	Yellow = "\033[33m"
	Reset  = "\033[0m"
)

// TODO: do not forget about n
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

	// zip archives
	// wg.Add(1)
	// TODO: use vars, not concatenation
	/*go*/
	archivation.RecursiveZip(pathWork+config.Work, uploadDir+"/"+config.Work)

	// for _, dir := range config.BackEnd {
	// 	wg.Add(1)
	// 	go archivation.RecursiveZip(&wg, pathBackEnd+dir, uploadDir+"/"+dir)
	// }

	// wg.Wait()

	fmt.Println(Yellow + "--- Uploading finished ---" + Reset)
}
