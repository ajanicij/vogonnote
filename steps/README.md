# steps directory contains steps that I took to develop vogonnote

## s001
Read entries of directory ./test by os.ReadDir.

## s002
Read directory tree recursively by filepath.Walk.

## s003
Read directory tree ./test recursively and list only files.

## s004
Read note file. Every note has a line with date and then one
or more lines with text. One note file can have more than one
note.

## s005
Parse date: from 2024-11-16 to time.Time to "2024 November 16".

## s006
Use bleve to search for text in two documents.

## s007
Parsing TOML file using go-toml.

## s008
Creating a temporary directory and deleting it at the end.

## s009
Regular expressions: check if a file name matches a pattern.

## s010
Getting current user's home directory.

## s011
OS-independent way to join path. Joining home directory and
file name .vogonnote.cfg.

## s012
Reading configuration, reading files from the note directory, indexing
files. Missing: process key from command line.

## s013
Reading value of command line argument -key.

## s014
Like s012; added processing key from command line.
Missing: parsing the result and printing the found note.

## s015
Split string that has fields separated by "^" into fiels
and display all the fields and the number of fields found.

## s016
Reading a note file and generating a slice with all the notes from one
file.

## s017
Same as s016, but moved code for reading a notes file to note.go.

## s018
vogonnote initial version.
