package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

const (
	// Files
	inputDirectory  = "test/"
	outputDirectory = "test/"

	// Regex
	compilerRegexExpr = "cc|gcc|clang|riscv32-unknown-elf-gcc"
	pathRegexExpr     = "([^\\s]+[.]c)"
	fileRegexExpr     = "[a-zA-Z0-9_-]+[.]c"
)

type CompileCommand struct {
	Directory string
	Command   string
	File      string
}

func (c CompileCommand) ToString() string {
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
	projectWorkingDirectory, err := os.Getwd()
	if err != nil {
		log.Fatalf("Could not get current working directory! %q\n", err)
	}

	// Regex
	compilerRegex, err := regexp.Compile(compilerRegexExpr)
	if err != nil {
		log.Fatalf("Could not compile 'compiler' regex! %q\n", err)
	}

	pathRegex, err := regexp.Compile(pathRegexExpr)
	if err != nil {
		log.Fatalf("Could not compile 'path' regex! %q\n", err)
	}

	fileRegex, err := regexp.Compile(fileRegexExpr)
	if err != nil {
		log.Fatalf("Could not compile 'file' regex! %q\n", err)
	}

	entries, err := os.ReadDir(inputDirectory)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range entries {
		inputFileName := inputDirectory + file.Name()
		outputFileName := outputDirectory + strings.ReplaceAll(file.Name(), ".log", ".json")

		fmt.Printf("File Name: %s\n", file.Name())

		// Read File
		lines := readFile(inputFileName)

		var compileCommands []CompileCommand
		var command CompileCommand
		var foundPaths []string

		for _, line := range lines {
			// Extract Data
			compiler := compilerRegex.FindString(line)
			path := pathRegex.FindString(line)
			file := fileRegex.FindString(line)
			directory, _, _ := strings.Cut(path, file)

			if strings.Compare(compiler, "") != 0 && strings.Compare(path, "") != 0 {
				duplicatePath := false
				for _, foundPath := range foundPaths {
					if strings.Compare(foundPath, path) == 0 {
						duplicatePath = true
						break
					}
				}

				if duplicatePath {
					continue
				}

				// Debug Output
				fmt.Println(line)
				fmt.Printf("\t(Compiler) %s\n", compiler)
				fmt.Printf("\t(Path) %s\n", path)
				fmt.Printf("\t(Directory) %s\n", directory)
				fmt.Printf("\t(File) %s\n", file)

				// Create Command
				command.Directory = projectWorkingDirectory + directory
				command.Command = line
				command.File = file

				// Register Command & Record Path
				compileCommands = append(compileCommands, command)
				foundPaths = append(foundPaths, path)
			}
		}

		// Write to compile_commands.json
		writeFile(outputFileName, compileCommands)
	}
}
