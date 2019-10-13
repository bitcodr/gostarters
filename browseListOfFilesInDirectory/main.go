package main

import (
	"net/http"
	"os"
)

func main() {
	dir, _ := os.Getwd()                                         //get current directory
	http.ListenAndServe(":3000", http.FileServer(http.Dir(dir))) //serve file and folders in the directory to port 3000 localhost
}
