# go-bimenu

go-bimenu is a simple tool to parse golang into a json model. This tool is a backend to the emacs plugin
https://github.com/gdrte/emacs-settings/blob/master/emacs.d/lisp/go-bimenu.el

## Installation

[Go][] version 1.1 or higher is required. Install or update go-bimenu using the
`go get` command:

	go get -u github.com/gdrte/go-bimenu

## Usage

	go-bimenu [options] file(s)

	-L="": source file names are read from the specified file. If file is "-", input is read from standard in.
	-R=false: recurse into directories in the file list.
	-f="": write output to specified file. If file is "-", output is written to standard out.
	-silent=false: do not produce any output on error.
	-format   Supported formats json(json-compact)/ctags. (default "json")
	-sort=true: sort tags.
	-tag-relative=false: file paths should be relative to the directory containing the tag file.
	-v=false: print version.

Also watch the demo screencast![demo screencast](./go-bimenu.gif?raw=true).
