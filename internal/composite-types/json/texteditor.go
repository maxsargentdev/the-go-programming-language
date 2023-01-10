package json

import "fmt"

func EditFile() {
	fmt.Println("Edit a file")
}

This part of the code will do a few things:

	- create temp files with a setup of input prefilled, like git CLI does
	- open a text editor & temp file, prompting user for input
	- save file to disk and load to memory for validation
	- if valid, use as input for some REST request
	- if invalid, save file back to disk and reload, give hints in comments for user

The file format should NOT be json, it should be something commentable and more user friendly.