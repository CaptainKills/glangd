package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

var (
	ProjectWorkingDirectory string
	DebugEnabled            bool
	inputPath               string
	outputPath              string
)

type CompileCommand struct {
	// JSON Information
	Directory string
	Command   string
	File      string

	// Additional Info
	Compiler string
	Path     string
}

func (c CompileCommand) ToJson() string {
	var builder strings.Builder

	builder.WriteString("\t{\n")

	builder.WriteString("\t\t\"directory\": \"")
	builder.WriteString(c.Directory)
	builder.WriteString("\",\n")

	builder.WriteString("\t\t\"command\": \"")
	builder.WriteString(c.Command)
	builder.WriteString("\",\n")

	builder.WriteString("\t\t\"file\": \"")
	builder.WriteString(c.File)
	builder.WriteString("\"\n")

	builder.WriteString("\t}")

	return builder.String()
}

func main() {
	// Initial Configuration
	ProjectWorkingDirectory, err := os.Getwd()
	if err != nil {
		log.Fatalf("Could not get current working directory! %q\n", err)
	}

	InitRegex()

	// Program Flags
	flag.StringVar(&inputPath, "f", "", "Specify input file path")
	flag.StringVar(&outputPath, "o", "", "Specify output file path")
	flag.BoolVar(&DebugEnabled, "d", false, "Enable/Disable Debug Information")
	flag.Parse()

	fmt.Println(ProjectWorkingDirectory)
	fmt.Printf("Input Path: %s\n", inputPath)
	fmt.Printf("Output Path: %s\n", outputPath)
	fmt.Printf("Debug Enabled: %t\n", DebugEnabled)

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
	// for _, file := range entries {
	// 	inputFilePath := inputPath + file.Name()
	// 	outputFilePath := outputPath + strings.ReplaceAll(file.Name(), ".log", ".json")
	//
	// 	if strings.HasSuffix(inputFilePath, ".log") {
	// 		fmt.Printf("File Name: %s\n", inputFilePath)
	// 		compileCommands := ParseFile(inputFilePath)
	//
	// 		// Write to compile_commands.json
	// 		writeFile(outputFilePath, compileCommands)
	// 	}
	// }
}
