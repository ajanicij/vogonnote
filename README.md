# vogonnote
A simple note searching CLI utility

## Introduction

vogonnote is a simple note searching utility. You as a user maintain
a directory with note files, where each note file is a text file with
one or more notes. You use vogonnote to search for a phrase in the
notes.

Currently vogonnote only works on Linux. I developed and tested it
on Ubuntu 22.04.

## Configuration

The configuration is in ~/.vogonnote.cfg. It is in
[TOML format](https://toml.io) and has the following structure:

```
rootdir = "/home/someuser/vogonnote.dir"
notefilepattern = ".*\\.note"
```

- rootdir is the root directory of the directory tree in which you
  keep your notes.
- notefilepattern is the file pattern that vogonnote uses to find
  all note files.

In the above configuration, all notes are in ~/vogonnote.dir. Note
files are all files with the extension .note.

## Notes

I wrote vogonnote with the idea that it will be used for a long time
(years, really). So you use it by creating a root directory and
add any subdirectories to structure the notes in any way meaningful
to you. One simple way is to have one directory for each year, e.g
2024, 2025 and so on. vogonnote will look for note files in the
root directory of the note tree and all subdirectories.

Every note in a note file has a very simple structure: one line
for date and then one or more lines for the text. For example:

```
2024-11-16
Today I created this wonderful utility that will conquer the world!
It will show everyone how smart... 
Never mind!

2024-11-17
It's no good. I put great hopes on it, but now I
am having my doubts.
```

The date line is in the format year-month-day, so 2024-11-17 means
"November 17, 20024".

Assuming that the configuration file is as in the example above and
you created this note file within the note directory tree and named
it with .note extension, vogonnote will find it.

Then if you search for "utility" like this:

    vogonnote -key utility

it will report something like this:

```
Temporary directory for index: /tmp/vogonnote.bleve3090128185
Search results: 2 matches, took 88.939µs

Note 0 in file /home/someuser/Documents/vogonnote.dir/tests/2024-11-17.note
Found note: 
-------------------
Date: 2024-11-16 00:00:00 +0000 UTC
Text:
Today I created this wonderful utility that will conquer the world!
It will show everyone how smart... 
Never mind!
-------------------

```

## How vogonnote works

It will never change the notes in any way. It just reads them, creates
a search index and then searches for the user's key phrase in the index.

vogonnote uses [Bleve](http://blevesearch.com/) for search. Every time
it is run, it creates an index in a temporary directory and before
exiting deletes the index directory, so it doesn't leave any traces
behind.

This utility is something I wrote for myself. It is deliberately a small
program that does one thing - searches for my notes - and hopefully
does it well. For detailed steps of how I developed it, look at the
directory [steps](https://github.com/ajanicij/vogonnote/tree/main/steps).

## Installing the code

    go install github.com/ajanicij/vogonnote/

## Building from source

In the root of the source tree, type

    go build .

## Usage

Type

    vogonnote -h

and it displays the following helpful message:

```
vogonnote tool.
Copyright Aleksandar Janicijevic ajanicij@yahoo.com 2024 
Usage:
  -key string
    	Search pattern
  -verbose
    	Run in verbose mode
```
