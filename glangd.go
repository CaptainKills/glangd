package main

import (
	"fmt"
	"log"
	"regexp"
	"strings"
)

const (
	// Files
	inputFileName  = "make.log"
	outputFileName = "compile_commands.json"

	// Regex
	sourceFileRegex = "([^\\s]+[.]c)"
	fileNameRegex   = "\\w+[.]c"
	directoryRegex  = "'.+'"
)

type CompileCommand struct {
	Directory  string
	Command    string
	SourceFile string
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
	builder.WriteString(c.SourceFile)
	builder.WriteString("\"\n")

	builder.WriteString("\t}")

	return builder.String()
}

func main() {
	// osArgs := os.Args
	// for index, element := range osArgs {
	// 	fmt.Printf("Argument %d: %s\n", index, element)
	// }

	// Read File
	lines := readFile(inputFileName)

	// Regex
	directoryRegex, err := regexp.Compile(directoryRegex)
	if err != nil {
		log.Fatalf("Could not compile directory regex! %q\n", err)
	}

	sourceFileRegex, err := regexp.Compile(sourceFileRegex)
	if err != nil {
		log.Fatalf("Could not compile directory regex! %q\n", err)
	}

	fileRegex, err := regexp.Compile(fileNameRegex)
	if err != nil {
		log.Fatalf("Could not compile directory regex! %q\n", err)
	}

	for _, line := range lines {
		if strings.Contains(line, "directory") {
			directory := strings.ReplaceAll(directoryRegex.FindString(line), "'", "")
			fmt.Printf("(Directory) %s\n", directory)
		}

		if strings.Contains(line, ".c") {
			sourceFile := sourceFileRegex.FindString(line)
			fmt.Printf("(Source File) %s\n", sourceFile)
			file := fileRegex.FindString(line)
			fmt.Printf("(File) %s\n", file)
		}

	}

	var compileCommands []CompileCommand
	command := CompileCommand{Directory: "/home/danick/", Command: "gcc -o main main.c", SourceFile: "main.c"}
	compileCommands = append(compileCommands, command)
	command = CompileCommand{Directory: "/home/senno/", Command: "gcc -o main /home/senno/src/main.c", SourceFile: "src/tree.c"}
	compileCommands = append(compileCommands, command)
	command = CompileCommand{Directory: "/home/tonnie/", Command: "gcc -o main main.c", SourceFile: "file.c"}
	compileCommands = append(compileCommands, command)

	// Write to compile_commands.json
	writeFile(outputFileName, compileCommands)
}
