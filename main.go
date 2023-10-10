package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	incudeDirs := flag.Bool("d", false, "incule dirs")
	inclueHiddenFiles := flag.Bool("p", false, "incule hidden files")
	flag.Usage = usages
	flag.Parse()
	args := flag.Args()

	var err error

	switch len(args) {
	case 0:
		err = renameDir(".", *incudeDirs, *inclueHiddenFiles)
	case 1:
		var stat os.FileInfo
		if stat, err = os.Stat(args[0]); err != nil {
		} else if !stat.IsDir() {
			fmt.Fprintf(os.Stderr, "err: %s is not a dir\nYou have provided 1 argumet so the programm will assume you want to remane the contens in that directory\n", args[0])
		}
		err = renameDir(args[0], *incudeDirs, *inclueHiddenFiles)
	default:
		err = renameFiles(args)
	}

	if err != nil {
		fmt.Fprintf(os.Stderr, "err: %s is not a dir\nYou have provided 1 argumet so the programm will assume you want to remane the contens in that directory\n", args[0])
	}
}

func usages() {
	fmt.Printf(`
brn
Bulk renamer

USAGE:
    brn [OPTIONS [Directory name]] [FILE]...

OPTIONS:
    -d
	inclue directories while renaming. (default "false")

    -p
	inclue hidden files or directories while renaming. (default "false")

EXAMPLES:
	1. By defaut if directory name is provided it will use "." current dir.
		$ brn media # will rename the contents of the media dir

	2. Remaning selected files
		$ brn *.mp4 # to rename the mp4 only
`)

}
