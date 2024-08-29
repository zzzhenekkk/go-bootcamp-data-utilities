package main

import (
	"archive/tar"
	"compress/gzip"
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sync"
)

func main() {
	destDirFlag := flag.String("a", "", "destination directory for archives")
	flag.Parse()
	destDir := *destDirFlag

	logFIles := flag.Args()

	if len(logFIles) == 0 {
		fmt.Println("No log files provided")
		os.Exit(1)
	}

	var wg sync.WaitGroup
	for _, file := range logFIles {
		wg.Add(1)
		go func(file string) {
			defer wg.Done()
			err := archiveLogFile(file, destDir)
			if err != nil {
				fmt.Printf("Error archiving %s: %s\n", file, err)
			}
		}(file)
	}
	wg.Wait()
}

func archiveLogFile(logFile, destDir string) error {
	fileInfo, err := os.Stat(logFile)
	if err != nil {
		return err
	}

	modTime := fileInfo.ModTime()
	timestamp := modTime.Unix()
	archiveName := fmt.Sprintf("%s_%d.tar.gz", filepath.Base(logFile), timestamp)

	if destDir != "" {
		archiveName = filepath.Join(destDir, archiveName)
	}
	//
	file, err := os.Create(archiveName)
	if err != nil {
		return err
	}
	defer file.Close()

	gw := gzip.NewWriter(file)
	defer gw.Close()

	tw := tar.NewWriter(gw)
	defer tw.Close()

	if err := addFileToTar(tw, logFile); err != nil {
		return err
	}

	fmt.Printf("Archived %s to %s\n", logFile, archiveName)
	return nil
}

func addFileToTar(tw *tar.Writer, filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	stat, err := file.Stat()
	if err != nil {
		return err
	}

	header := &tar.Header{
		Name:    filepath.Base(filename),
		Size:    stat.Size(),
		Mode:    int64(stat.Mode()),
		ModTime: stat.ModTime(),
	}

	if err := tw.WriteHeader(header); err != nil {
		return err
	}

	if _, err := io.Copy(tw, file); err != nil {
		return err
	}

	return nil
}
