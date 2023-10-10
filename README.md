# brn

Bulk rename. Renger stle rename

# Demo

![Demo gif](/imgs/demo.gif)

# Usages

```
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
```

# Insall

```bash
go install https://github.com/wizsk/brn
```
or see releases