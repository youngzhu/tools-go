package main

import (
	"fmt"
	"github.com/youngzhu/godate"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

// Obsidian to Logseq

const (
	pathSeparator = string(os.PathSeparator)

	dir      = "E:\\temp1"
	pages    = dir + pathSeparator + "pages"
	journals = dir + pathSeparator + "journals"

	suffix = ".md"
)

func init() {
	// 必须先建目录
	os.Mkdir(pages, 0777)
	os.Mkdir(journals, 0777)
}

func main() {

	fileInfos, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	counterJournals, counterPages := 0, 0

	for _, file := range fileInfos {
		if file.IsDir() || !strings.HasSuffix(file.Name(), suffix) {
			continue
		}

		fileName := strings.TrimSuffix(file.Name(), suffix)
		//log.Println(fileName)

		date, err := godate.Parse(fileName)
		if err == nil {
			newFileName := date.Format("2006_01_02")
			CopyFile(dir+pathSeparator+fileName+suffix, journals+pathSeparator+newFileName+suffix)
			counterJournals++
		} else {
			CopyFile(dir+pathSeparator+file.Name(), pages+pathSeparator+file.Name())
			counterPages++
		}

	}

	fmt.Printf("%d Journals, %d Pages, total: %d", counterJournals, counterPages, counterJournals+counterPages)
}

func CopyFile(sourcePath, destPath string) error {
	inputFile, err := os.Open(sourcePath)
	if err != nil {
		return fmt.Errorf("Couldn't open source file: %s", err)
	}
	outputFile, err := os.Create(destPath)
	if err != nil {
		inputFile.Close()
		return fmt.Errorf("Couldn't open dest file: %s", err)
	}
	defer outputFile.Close()
	_, err = io.Copy(outputFile, inputFile)
	inputFile.Close()
	if err != nil {
		return fmt.Errorf("Writing to output file failed: %s", err)
	}

	return nil
}

func MoveFile(sourcePath, destPath string) error {
	inputFile, err := os.Open(sourcePath)
	if err != nil {
		return fmt.Errorf("Couldn't open source file: %s", err)
	}
	outputFile, err := os.Create(destPath)
	if err != nil {
		inputFile.Close()
		return fmt.Errorf("Couldn't open dest file: %s", err)
	}
	defer outputFile.Close()
	_, err = io.Copy(outputFile, inputFile)
	inputFile.Close()
	if err != nil {
		return fmt.Errorf("Writing to output file failed: %s", err)
	}
	// The copy was successful, so now delete the original file
	err = os.Remove(sourcePath)
	if err != nil {
		return fmt.Errorf("Failed removing original file: %s", err)
	}
	return nil
}
