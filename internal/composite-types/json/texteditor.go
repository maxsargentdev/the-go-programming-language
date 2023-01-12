package json

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/exec"
)

const tempFileLocation = "/tmp/test.txt"

// updateBodyParams recieves a bodyParams object and lets users edit the json in vim
func editBodyParams(bodyParams IssueBodyParams) IssueBodyParams {

	// Create the temp file filled with JSON
	tempFile := createFile(tempFileLocation)
	writeFile(tempFile, bodyParams)

	// Open vim and give a view onto the json object
	cmd := exec.Command("vim", tempFileLocation)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	err := cmd.Run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error launching vim: %v\n", err)
		os.Exit(1)
	}

	// Read the file back, its going to be our new boyParams
	tempBytes := readFile()

	// Now let's unmarshall the data into back into our struct
	var editedBodyParams IssueBodyParams
	err = json.Unmarshal(tempBytes, &editedBodyParams)
	if err != nil {
		log.Fatal("Error during Unmarshal(): ", err)
	}

	// Close handle on temp file so we can delete safely
	closeFile(tempFile)

	// Delete file
	deleteFile()

	return editedBodyParams
}

func createFile(p string) *os.File {
	f, err := os.Create(p)

	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}

	return f
}

func writeFile(f *os.File, params IssueBodyParams) {
	marshalledParams, err := json.Marshal(params)

	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
	fmt.Fprintln(f, string(marshalledParams))
}

func closeFile(f *os.File) {
	err := f.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}

func deleteFile() {
	err := os.Remove(tempFileLocation)

	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}

func readFile() []byte {
	data, err := os.ReadFile(tempFileLocation)

	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}

	return data
}
