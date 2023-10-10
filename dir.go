package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

type fileNameCng struct {
	changed   bool
	old, curr string
}

// type fileConfig struct {
// 	currDir string
// 	before, after []string
// }

// do i wanna test this? if yes then :)
func dir(currDir string, includeDir, includeHidden bool) ([]fileNameCng, error) {
	dirContents, err := os.ReadDir(currDir)
	if err != nil {
		return nil, err
	}

	fileNames := []string{}
	fileNamesBuff := bytes.Buffer{}
	// oh no
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

	changedFileNameFile, err := os.CreateTemp(os.TempDir(), "brn-")
	if err != nil {
		return nil, err
	}

	// cleaning up
	defer changedFileNameFile.Close()
	defer os.Remove(changedFileNameFile.Name())

	if _, err = changedFileNameFile.Write(fileNamesBuff.Bytes()); err != nil {
		return nil, err
	}

	editor := os.Getenv("EDITOR")
	if editor == "" {
		// TODO: handle err err
		log.Fatal("no editor name")
	}

	cmd := exec.Command(editor, changedFileNameFile.Name())
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout

	if err = cmd.Run(); err != nil {
		log.Fatal(err)
	}

	edited, _ := os.ReadFile(changedFileNameFile.Name())
	fmt.Println(string(edited))

	scanner := bufio.NewScanner(bytes.NewReader(edited))
	scanner.Split(bufio.ScanLines)

	fileCng := []fileNameCng{}

	for _, val := range fileNames {
		if !scanner.Scan() {
			return nil, fmt.Errorf("err: file names currupted")
		}

		changedName := string(scanner.Bytes())
		if val == changedName {
			fileCng = append(fileCng, fileNameCng{changed: false})
			continue
		}

		fileCng = append(fileCng, fileNameCng{
			changed: false,
			old:     filepath.Join(currDir, val),
			curr:    filepath.Join(currDir, changedName),
		})
	}

	return fileCng, nil
}
