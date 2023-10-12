package main

import (
	"flag"
	"fmt"
	"os"
)

const vertion = "1.1"

func main() {
	incudeDirs := flag.Bool("d", false, "incule dirs")
	inclueHiddenFiles := flag.Bool("h", false, "incule hidden files")
	flag.Usage = usages
	flag.Parse()
	args := flag.Args()

	var err error

	switch len(args) {
	case 0:
		usages()
		os.Exit(1)
	case 1:
		var stat os.FileInfo
		if stat, err = os.Stat(args[0]); err != nil {
		} else if !stat.IsDir() {
			fmt.Fprintf(os.Stderr, "err: %s is not a dir\nYou have provided 1 argumet so the programm will assume you want to remane the contens in that directory\n", args[0])
			return
		}
		err = renameDir(args[0], *incudeDirs, *inclueHiddenFiles)
	default:
		err = renameFiles(args)
	}

	if err != nil {
		fmt.Fprintf(os.Stderr, "unexpected: %s\n", err)
		os.Exit(1)
	}
}

func usages() {
	fmt.Printf(`brn %s
Bulk renamer

DESCRIPTION
	brn takes the specified files and puts them into a text buffer and open your editor "EDITOR" env var
	or "nvim", "vim", "vi", "nano" and lets you edit the the file names. if the file any name was changed only
	those fils are renamed.

	Note: Deleting fils are not supported. This will resut in "err: file names currupted"

OPTIONS:
    -d
	inclue directories while renaming. By default diretories are exclued.

    -h
	inclue hidden files or directories while renaming.By default hidden files are are exclued.

EXAMPLES:
	$ export EDITOR=nvim # set the env var

	$ brn .			# rename only the files in current dir "."
	$ brn . -d		# name files and directoris in current dir "."
	$ brn . -d -h	# name files and directoris including hidden files in current dir "."
	$ brn fo		# rename only the files in the dir "fo"
	$ brn f f2 f3	# rename only the files "f f2 f3"
	$ brn *.mp4		# rename only the files ending with "mp4" // bash magic!

`, vertion)
}
