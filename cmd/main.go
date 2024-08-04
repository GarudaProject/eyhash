package main

import (
	"fmt"
	"os"
	"path"

	"github.com/GarudaProject/eyhash"
)

func usage() {
	fmt.Printf(`
usage: eyehash <filename/directory> -f/-d [file/directory]
  `)
}

func main() {
	if len(os.Args) < 2 {
		usage()
		return
	}

	file := os.Args[1]
	mode := os.Args[2]

	switch mode {
	case "-f":
		info, err := eyhash.FileInfo(file)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		fmt.Println()
		hashInfo(file, info)
		break
	case "-d":
		dirs, err := os.ReadDir(file)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		for _, dir := range dirs {
			filePath := path.Join(file, dir.Name())
			info, err := eyhash.FileInfo(filePath)
			if err != nil {
				fmt.Println(err.Error())
				return
			}
			fmt.Println()
			hashInfo(filePath, info)
		}
		break
	default:
		usage()
		break
	}
}

func hashInfo(file string, info *eyhash.Info) {
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
		fmt.Println(err.Error())
		return
	}

	fmt.Printf("FileName: %s, Size: %d bytes, ModTime: %s\n", info.Name, info.Size, info.ModTime.Local())
	fmt.Println("MD5:", md5hash)
	fmt.Println("SHA1:", sh1hash)
	fmt.Println("SHA256:", sh256hash)
	fmt.Println("SHA512:", sh512hash)
}
