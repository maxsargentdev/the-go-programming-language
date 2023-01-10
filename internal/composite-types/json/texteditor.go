package json

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/exec"
)

// SWITCH TO USING JSON.MARSHAL FOR THIS
// STILL WIP

// updateBodyParams recieves a bodyParams object, creates a temp file, opens a text editor on that file that renders that object in a nice format
func editBodyParams(bodyParams IssueBodyParams) IssueBodyParams {

	tempFile := createFile("/tmp/test.txt")
	writeFile(tempFile, bodyParams)

	cmd := exec.Command("vim", "/tmp/test.txt")
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	err := cmd.Run()

	if err != nil {
		fmt.Fprintf(os.Stderr, "error updating issue: %v\n", err)
		os.Exit(1)
	}
	testString := readFile(tempFile)
	fmt.Println(testString)

	// Now let's unmarshall the data into `payload`
	var payload IssueBodyParams
	err = json.Unmarshal(testString, &payload)
	if err != nil {
		log.Fatal("Error during Unmarshal(): ", err)
	}

	closeFile(tempFile)
	deleteFile(tempFile)

	return payload
}

func createFile(p string) *os.File {
	fmt.Println("creating")
	f, err := os.Create(p)
	if err != nil {
		panic(err)
	}
	return f
}

func writeFile(f *os.File, params IssueBodyParams) {
	fmt.Println("writing")
	fmt.Fprintln(f, fmt.Sprintf("%#v", params))

}

func closeFile(f *os.File) {
	fmt.Println("closing")
	err := f.Close()

	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}

func deleteFile(f *os.File) {
	os.Remove("/tmp/test.txt")
}

func readFile(f *os.File) []byte {
	data, _ := os.ReadFile("/tmp/test.txt")
	return data
}

//This part of the code will do a few things:
//
//	- create temp files with a setup of input prefilled, like git CLI does
//	- open a text editor & temp file, prompting user for input
//	- save file to disk and load to memory for validation
//	- if valid, use as input for some REST request
//	- if invalid, save file back to disk and reload, give hints in comments for user
//
//The file format should NOT be json, it should be something commentable and more user friendly.
