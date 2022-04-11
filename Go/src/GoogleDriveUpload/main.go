package main

import (
	"GoogleDriveUpload/archivation"
	"fmt"
	"os"
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

	archivation.RecursiveZip(os.Args[1], os.Args[1])

	fmt.Println(Yellow + "--- Uploading finished ---" + Reset)
}
