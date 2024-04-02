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
	editorWithArgs := []string{}
	if editor == "" {
		editor = getPrefferedEditor()
		// TODO: handle err err
		if editor == "" {
			return fmt.Errorf("err: could not find an editor")
		}
	} else {
		editorWithArgs = strings.Split(strings.TrimSpace(editor), " ")
		if len(editorWithArgs) == 0 {
			return fmt.Errorf("err: invalid editor name")
		}
	}

	editorWithArgs = append(editorWithArgs, changedFileNameFile.Name())
	cmd := exec.Command(editorWithArgs[0], editorWithArgs[1:]...)
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
	renameTracker := strings.Builder{}
	renameCount := 0
	renameTracker.WriteString("Files renamed:\n")
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
		renameTracker.WriteString(fmt.Sprintf("\t'%s' -> '%s'\n", old, curr))
	}

	fmt.Print(renameTracker.String())
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
