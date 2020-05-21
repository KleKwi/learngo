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
	"path"
)

func main() {
	// var dir, file string

	dir, file := path.Split("css/main.css")

	fmt.Println("dir :", dir)
	fmt.Println("file:", file)
}
