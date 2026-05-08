package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	cmd "github.com/CaptainKills/glangd/cmd"
	file "github.com/CaptainKills/glangd/file"
	parser "github.com/CaptainKills/glangd/parser"
	regex "github.com/CaptainKills/glangd/regex"
)

var (
	debugEnabled bool
	singleMode   bool

	usedCompiler string
	inputPath    string
	outputPath   string

	compileCommands []cmd.CompileCommand
)

func main() {
	// Project Working Directory
	workingDirectory, err := os.Getwd()
	if err != nil {
		log.Fatalf("Could not retrieve Project Working Directory!")
	}

	// Program Flags
	flag.StringVar(&usedCompiler, "c", "", "Specify compiler program")
	flag.StringVar(&workingDirectory, "w", workingDirectory, "Overwrite current working directory")
	flag.StringVar(&inputPath, "f", "", "Specify input file path")
	flag.StringVar(&outputPath, "o", "", "Specify output file path")

	flag.BoolVar(&debugEnabled, "d", false, "Enable/Disable Debug Information")
	flag.BoolVar(&singleMode, "s", false, "Format the output as .txt instead of .json")
	flag.Parse()

	// Handle -c
	regex.InitRegex(usedCompiler)

	// Handle -o
	if outputPath == "" {
		if !singleMode {
			outputPath = "compile_commands.json"
		} else {
			outputPath = "compile_flags.txt"
		}
	}

	// Handle -s
	if !singleMode {
		if !strings.HasSuffix(outputPath, ".json") {
			log.Fatalf("Specified output file is not a .json file!")
		}
	} else if singleMode {
		if !strings.HasSuffix(outputPath, ".txt") {
			log.Fatalf("Specified output file is not a .txt file!")
		}
	}

	// Handle -d
	if debugEnabled {
		fmt.Printf("Specified Compiler: %s\n", usedCompiler)
		fmt.Printf("Working Directory: %s\n", workingDirectory)
		fmt.Printf("Input Path: %s\n", inputPath)
		fmt.Printf("Output Path: %s\n", outputPath)
		fmt.Printf("Debug Enabled: %t\n", debugEnabled)
		fmt.Printf("Single Mode: %t\n", singleMode)
		fmt.Println()
	}

	// Handle -f
	if inputPath == "" {
		compileCommands = parser.ParseStdin(workingDirectory, debugEnabled)
	} else {
		compileCommands = parser.ParseFile(workingDirectory, inputPath, debugEnabled)
	}

	file.WriteFile(outputPath, compileCommands, singleMode)
}
