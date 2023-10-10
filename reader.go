package main

import (
	"bytes"
	"os"
	"strings"
)

func renameDir(currDir string, includeDir, includeHidden bool) error {
	dirContents, err := os.ReadDir(currDir)
	if err != nil {
		return err
	}

	fileNames := []string{}
	fileNamesBuff := bytes.Buffer{}

	for _, file := range dirContents {
		if !includeDir && file.IsDir() {
			continue
		}
		if !includeHidden && strings.HasPrefix(file.Name(), ".") {
			continue
		}

		fileNames = append(fileNames, file.Name())
		fileNamesBuff.WriteString(file.Name() + "\n")
	}

	return rename(fileNames, fileNamesBuff, currDir)
}

func renameFiles(args []string) error {
	fileNames := []string{}
	fileNamesBuff := bytes.Buffer{}
	for _, file := range args {
		_, err := os.Stat(file)
		if err != nil {
			return err
		}

		fileNames = append(fileNames, file)
		fileNamesBuff.WriteString(file + "\n")
	}
	return rename(fileNames, fileNamesBuff, "")
}
