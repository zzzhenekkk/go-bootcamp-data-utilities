package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	onlyFiles := flag.Bool("f", false, "Find only files")
	onlyDirs := flag.Bool("d", false, "Find only directories")
	onlySymlinks := flag.Bool("sl", false, "Find only symbolic links")
	extension := flag.String("e", "", "Find with extension")

	flag.Parse()

	if !*onlyFiles && !*onlyDirs && !*onlySymlinks {
		*onlyFiles = true
		*onlyDirs = true
		*onlySymlinks = true
	}

	if flag.NArg() != 1 {
		fmt.Println("Usage: ./myFind [OPTION]... <pathFile>")
		return
	}

	rootPath := flag.Arg(0)

	err := filepath.Walk(rootPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if *onlyFiles && info.Mode().IsRegular() || *onlyDirs && info.Mode().IsDir() || *onlySymlinks && info.Mode()&os.ModeSymlink != 0 {
			if *extension != "" && filepath.Ext(info.Name()) != "."+*extension {
				return nil
			}
			fmt.Println(path)
		}
		return nil

	})

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}
