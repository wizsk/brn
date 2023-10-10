# brn

Bulk rename files in a sweep with your favourite text [editor](#editor). (I know it's nvim).

## Install

```bash
go install https://github.com/wizsk/brn
```
or see release.


## Demo

![Demo gif](/imgs/demo.gif)

## Usages

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
## Editor

You can use whatever editor you want by setting `EDITOR` env var. Preferably the editor should be terminal based.

```bash
export EDITOR=nvim # nvim user btw
```

But for some wired reason you wanna use vscode use this flag too. Use [vscodium](https://vscodium.com/) or maybe you love to give Microsoft your data ;). Or If you are using windows may God have mercy on you.

```bash
export EDITOR='code -w' # --wait
```


