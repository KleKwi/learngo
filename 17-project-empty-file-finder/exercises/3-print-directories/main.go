// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

// ---------------------------------------------------------
// EXERCISE: Find and write the names of subdirectories to a file
//
//  Create a program that can get multiple directory paths from
//  the command-line, and prints only their subdirectories into a
//  file named: dirs.txt
//
//
//  1. Get the directory paths from command-line
//
//  2. Append the names of subdirectories inside each directory
//     to a byte slice
//
//  3. Write that byte slice to dirs.txt file
//
//
// EXPECTED OUTPUT
//
//   go run main.go
//     Please provide directory paths
//
//   go run main.go dir/ dir2/
//
//   cat dirs.txt
//
//     dir/
//             subdir1/
//             subdir2/
//
//     dir2/
//             subdir1/
//             subdir2/
//             subdir3/
//
//
// HINTS
//
//   ONLY READ THIS IF YOU GET STUCK
//
//   + Get all the files in a directory using ioutil.ReadDir
//     (A directory is also a file)
//
//   + You can use IsDir method of a FileInfo value to detect
//     whether a file is a directory or not.
//
//     Check out its documentation:
//
//     go doc os.FileInfo.IsDir
//
//     # or using godocc
//     godocc os.FileInfo.IsDir
//
//   + You can use '\t' escape sequence for indenting the subdirs.
//
//   + You can find a sample directory structure under:
//     solution/ directory
//
// ---------------------------------------------------------

func main() {
	paths := os.Args[1:]

	if len(paths) == 0 {
		fmt.Println("Give me the dirs.")
		return
	}

	var data []byte

	for _, path := range paths {
		files, err := ioutil.ReadDir(path)
		if err != nil {
			fmt.Println(err)
			continue
		}

		data = append(data, path...)
		data = append(data, '\n')

		for _, file := range files {
			if file.IsDir() {
				name := file.Name()
				data = append(data, '\t')
				data = append(data, name...)
				data = append(data, '/', '\n')
			}
		}

		data = append(data, '\n')
	}

	err := ioutil.WriteFile("out.txt", data, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
}
