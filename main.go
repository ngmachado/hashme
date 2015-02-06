package main

import (
	"crypto/sha256"
	"encoding/hex"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"log"
)

//flags
var fpath string

//Working map - string = path ; bool = is the work done
//In this version we are not working with the value bool
var w = make(map[string]bool)

func init() {
	//Get the working directory
	pwd, err := os.Getwd()
	if err != nil {
		pwd = "/"
	}
	flag.StringVar(&fpath, "d", pwd, "usage : -d path")
}

func main() {

	flag.Parse()
	fpath += path.Clean(string(os.PathSeparator))

	err := filepath.Walk(fpath, savePath)
	if err != nil {
		panic(err)
	}

	fmt.Println("#Working Directories : " + fmt.Sprintln(len(w)))

	for key, _ := range w {

		files, err := ioutil.ReadDir(key)

		if err != nil {
			log.Fatal(err)
		}

		for _, f := range files {
			if !f.IsDir() {
				shex, err := Sha256Hash(key + string(os.PathSeparator) + f.Name())
				
				if err != nil {
					log.Fatal(err)
				}
				
				fmt.Println(f.Name() + " (sha256) : " + shex)
			}

		}
	}

}

//Save the directories/files to w map
func savePath(basepath string, info os.FileInfo, err error) error {

	if err != nil {
		return err
	}

	w[filepath.Dir(basepath)] = false
	return nil
}

//for a given path make a file hash
func Sha256Hash(basepath string) (string, error) {
	hasher := sha256.New()
	f, err := ioutil.ReadFile(basepath)
	if err != nil {
		return "",err
	}
	hasher.Write(f)
	return hex.EncodeToString(hasher.Sum(nil)), nil
}
