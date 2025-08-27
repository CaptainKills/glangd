package main

import (
	"flag"
	"fmt"
	"log"
	"strings"

	file "github.com/CaptainKills/glangd/file"
	parser "github.com/CaptainKills/glangd/parser"
	regex "github.com/CaptainKills/glangd/regex"
)

var (
	debugEnabled bool
	inputPath    string
	outputPath   string
)

func main() {
	regex.InitRegex()

	// Program Flags
	flag.StringVar(&inputPath, "f", "", "Specify input file path")
	flag.StringVar(&outputPath, "o", "compile_commands.json", "Specify output file path")
	flag.BoolVar(&debugEnabled, "d", false, "Enable/Disable Debug Information")
	flag.Parse()

	fmt.Printf("Input Path: %s\n", inputPath)
	fmt.Printf("Output Path: %s\n", outputPath)
	fmt.Printf("Debug Enabled: %t\n", debugEnabled)

	// Handle -o
	if !strings.HasSuffix(outputPath, ".json") {
		log.Fatalf("Specified output file is not a .json file!")
	}

	// Handle -f
	if inputPath == "" {
		compileCommands := parser.ParseStdin(debugEnabled)
		file.WriteFile(outputPath, compileCommands)
	} else {
		compileCommands := parser.ParseFile(inputPath, debugEnabled)
		file.WriteFile(outputPath, compileCommands)
	}
}
