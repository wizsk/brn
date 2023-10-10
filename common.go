package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

func rename(fileNames []string, fileNamesBuff bytes.Buffer, currDir string) error {
	changedFileNameFile, err := os.CreateTemp(os.TempDir(), "brn-")
	if err != nil {
		return err
	}

	// cleaning up
	defer changedFileNameFile.Close()
	defer os.Remove(changedFileNameFile.Name())

	if _, err = changedFileNameFile.Write(fileNamesBuff.Bytes()); err != nil {
		return err
	}

	editor := os.Getenv("EDITOR")
	if editor == "" {
		editor = getPrefferedEditor()
		// TODO: handle err err
		if editor == "" {
			return fmt.Errorf("err: could not find an editor")
		}
	}

	cmd := exec.Command(editor, changedFileNameFile.Name())
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout

	if err = cmd.Run(); err != nil {
		log.Fatal(err)
	}

	edited, err := os.ReadFile(changedFileNameFile.Name())
	if err != nil {
		return err
	}
	scanner := bufio.NewScanner(bytes.NewReader(edited))
	scanner.Split(bufio.ScanLines)

	changedNames := []string{}
	for scanner.Scan() {
		changedNames = append(changedNames, string(scanner.Bytes()))
	}

	if len(changedNames) != len(fileNames) {
		return fmt.Errorf("err: file names currupted")

	}

	// rename the dam thing
	renameCount := 0
	fmt.Println("Files renamed:")
	for i, val := range fileNames {
		if val == changedNames[i] {
			continue
		}

		old := filepath.Join(currDir, val)
		curr := filepath.Join(currDir, changedNames[i])
		if err = os.Rename(old, curr); err != nil {
			return err
		}
		renameCount++
		fmt.Printf("\t%s -> %s\n", old, curr)
	}
	fmt.Printf("\ntotal %d files remaned out of %d files\n", renameCount, len(fileNames))

	return nil
}

func getPrefferedEditor() string {
	// i like vim
	editors := []string{"nvim", "vim", "vi", "nano"}
	for _, editor := range editors {
		cmd := exec.Command(editor, "--version")
		err := cmd.Run()
		if err == nil {
			return editor
		}
	}
	return ""
}
