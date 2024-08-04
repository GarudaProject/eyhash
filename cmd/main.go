package main

import (
	"fmt"
	"log"
	"os"

	"github.com/GarudaProject/eyhash"
)

func usage() {
	fmt.Printf(`
    usage: eyehash <filename>
  `)
}

func main() {
	if len(os.Args) <= 1 {
		usage()
		return
	}

	file := os.Args[1]

	fmt.Println()

	info, err := eyhash.FileInfo(file)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	md5hash, err := eyhash.MD5File(file)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	sh1hash, err := eyhash.SHA1File(file)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	sh256hash, err := eyhash.SHA256File(file)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	sh512hash, err := eyhash.SHA512File(file)
	if err != nil {
		log.Println(err.Error())
		return
	}

	fmt.Printf("FileName: %s, Size: %d bytes, ModTime: %s\n", info.Name, info.Size, info.ModTime.Local())
	fmt.Println("MD5:", md5hash)
	fmt.Println("SHA1:", sh1hash)
	fmt.Println("SHA256:", sh256hash)
	fmt.Println("SHA512:", sh512hash)
}
