# brn

Bulk rename files in a sweep with your favourite text [editor](#editor). (I know it's nvim).

## Install

See [releases](https://github.com/wizsk/brn/releases) for pre-built binaries.

Building from the source requires [Go](https://go.dev/).

On Unix:

```bash
env CGO_ENABLED=0 go install -ldflags="-s -w" github.com/wizsk/brn@latest
```

On Windows `cmd`:
<!-- guess people does use windows -->

```cmd
set CGO_ENABLED=0
go install -ldflags="-s -w" github.com/wizsk/brn@latest
```

On Windows `powershell`:

```powershell
$env:CGO_ENABLED = '0'
go install -ldflags="-s -w" github.com/wizsk/brn@latest
```
or see release.


## Demo

![Demo gif](/imgs/demo.gif)

## Usages

```
brn --help
brn v1.0
Bulk renamer

DESCRIPTION
	brn takes the specified files and puts them into a text buffer and open your editor "EDITOR" env var
	or "nvim", "vim", "vi", "nano" and lets you edit the the file names. if the file name was changed only
	those fils are renamed.

	Note: Deleting fils are not supported.

OPTIONS:
    -d
	inclue directories while renaming. By default diretories are exclued.

    -p
	inclue hidden files or directories while renaming.By default hidden files are are exclued.

EXAMPLES:
	$ export EDITOR=nvim # set the env var
	$ brn			# rename only the files in current dir "."
	$ brn -d		# name files and directoris in current dir "."
	$ brn -d -h		# name files and directoris including hidden files in current dir "."
	$ brn fo		# rename only the files in the dir "fo"
	$ brn f f2 f3	# rename only the files "f f2 f3"
	$ brn *.mp4		# rename only the files ending with "mp4" // bash magic!
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


