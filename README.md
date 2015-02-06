# HashMe - Golang
 
HashMe is a small tool that for a given directory output a hash (sha256) for each file in that directory and also do the same for each subdirectories (recursively).
 
### Build
```sh
$ go build main.go
```
 
### Installation
```sh
$ go install
```
 
### How to run
```sh
hashme -d Users\Documents
```
 
### Save result to file
```sh
hashme -d Users\Documents > myFile.txt
```
 
License
----
 
MIT
