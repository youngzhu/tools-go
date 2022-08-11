package main

import (
	"fmt"
	"github.com/youngzhu/godate"
	_ "github.com/youngzhu/godate"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

// Obsidian to Logseq

const (
	dir    = "E:\\temp"
	output = dir + "\\output"

	suffix = ".md"
)

func main() {

	fileInfos, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}

	// 必须先建目录
	os.Mkdir(output, 0777)

	for _, file := range fileInfos {
		if file.IsDir() {
			continue
		}

		fileName := strings.TrimSuffix(file.Name(), suffix)
		log.Println(fileName)

		newFileName := fileName

		date, err := godate.Parse(fileName)
		if err == nil {
			newFileName = date.Format("2006_01_02")
		}

		CopyFile(dir+"\\"+fileName+suffix, output+"\\"+newFileName+suffix)

		//upper2 = strings.ToUpper(fileName)
		//
		//if strings.Contains(upper2, upper1) {
		//	log.Println(fileName)
		//}

	}
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
