package main

import (
	"flag"
	"log"
	"os"
)

func main() {
	incudeDirs := flag.Bool("d", false, "incule dirs")
	incudeDirs = flag.Bool("dirs", false, "incule dirs")
	inclueHiddenFiles := flag.Bool("h", false, "incule hidden files")
	// inclueHiddenFiles = flag.Bool("hidde", false, "incule dirs")
	flag.Parse()
	args := flag.Args()
	switch len(args) {
	case 0:
		dir(".", *incudeDirs, *inclueHiddenFiles)
	case 1:
		if stat, err := os.Stat(args[0]); err != nil {
			// TODO: handle err err
			log.Fatal(err)
		} else if !stat.IsDir() {
			log.Fatalf("err: %s is not a dir\n", args[0])
		}
		dir(args[0], *incudeDirs, *inclueHiddenFiles)
		// the provided arg has to be a dir
		// other wise no idea
	default:
		log.Fatal(args)
		// iplemnet is
	}
}
