package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	file "github.com/CaptainKills/glangd/file"
	parser "github.com/CaptainKills/glangd/parser"
	regex "github.com/CaptainKills/glangd/regex"
)

var (
	projectWorkingDirectory string
	debugEnabled            bool
	inputPath               string
	outputPath              string
)

func main() {
	// Initial Configuration
	ProjectWorkingDirectory, err := os.Getwd()
	if err != nil {
		log.Fatalf("Could not get current working directory! %q\n", err)
	}

	regex.InitRegex()

	// Program Flags
	flag.StringVar(&inputPath, "f", "", "Specify input file path")
	flag.StringVar(&outputPath, "o", "", "Specify output file path")
	flag.BoolVar(&debugEnabled, "d", false, "Enable/Disable Debug Information")
	flag.Parse()

	fmt.Println(ProjectWorkingDirectory)
	fmt.Printf("Input Path: %s\n", inputPath)
	fmt.Printf("Output Path: %s\n", outputPath)
	fmt.Printf("Debug Enabled: %t\n", debugEnabled)

	// Handle -o
	if outputPath == "" {
		// Use $PWD/compile_commands.json as output
	} else {
		// Check if path is actually a path to a file/folder
		// if folder, output to folder/compile_commands.json
		// if file, output to folder/file
	}

	// Handle -f
	if inputPath == "" {
		// Use stdin as input
	} else {
		// Check if path is actually a path to a file
		// Run glangd for a specific file
	}

	// Main Program
	// entries, err := os.ReadDir(inputPath)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	//
	// for _, f := range entries {
	// 	inputFilePath := inputPath + f.Name()
	// 	outputFilePath := outputPath + strings.ReplaceAll(f.Name(), ".log", ".json")
	//
	// 	if strings.HasSuffix(inputFilePath, ".log") {
	// 		fmt.Printf("File Name: %s\n", inputFilePath)
	// 		compileCommands := parser.ParseFile(inputFilePath, debugEnabled)
	//
	// 		// Write to compile_commands.json
	// 		file.WriteFile(outputFilePath, compileCommands)
	// 	}
	// }
}
