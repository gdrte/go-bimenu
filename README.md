# gotags

go-emacsoutline is a [ctags][]-compatible tag generator for [Go][].

## Installation

[Go][] version 1.1 or higher is required. Install or update gotags using the
`go get` command:

	go get -u github.com/gdrte/go-emacsoutline

## Usage

	go-emacsoutline [options] file(s)

	-L="": source file names are read from the specified file. If file is "-", input is read from standard in.
	-R=false: recurse into directories in the file list.
	-f="": write output to specified file. If file is "-", output is written to standard out.
	-silent=false: do not produce any output on error.
	-format   Supported formats json(json-compact)/ctags. (default "json")
	-sort=true: sort tags.
	-tag-relative=false: file paths should be relative to the directory containing the tag file.
	-v=false: print version.

