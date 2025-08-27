package parser

import (
	"fmt"
	"strings"

	cmd "github.com/CaptainKills/glangd/cmd"
	file "github.com/CaptainKills/glangd/file"
	regex "github.com/CaptainKills/glangd/regex"
)

func ParseStdin(stdin []string, debug bool) []cmd.CompileCommand {
	var commands []cmd.CompileCommand
	var foundPaths []string

	for _, line := range stdin {
		cmd := parseLine(line)

		if strings.Compare(cmd.Compiler, "") != 0 && strings.Compare(cmd.Path, "") != 0 {
			duplicatePath := false

			for _, foundPath := range foundPaths {
				if strings.Compare(foundPath, cmd.Path) == 0 {
					duplicatePath = true
					break
				}
			}

			if duplicatePath {
				continue
			}

			// Debug Output
			if debug {
				fmt.Println(line)
				fmt.Printf("\t(Compiler) %s\n", cmd.Compiler)
				fmt.Printf("\t(Path) %s\n", cmd.Path)
				fmt.Printf("\t(Directory) %s\n", cmd.Directory)
				fmt.Printf("\t(File) %s\n", cmd.File)
			}

			// Register Command & Record Path
			commands = append(commands, cmd)
			foundPaths = append(foundPaths, cmd.Path)
		}
	}

	return commands
}

func ParseFile(f string, debug bool) []cmd.CompileCommand {
	var commands []cmd.CompileCommand
	var foundPaths []string

	lines := file.ReadFile(f)

	for _, line := range lines {
		cmd := parseLine(line)

		if strings.Compare(cmd.Compiler, "") != 0 && strings.Compare(cmd.Path, "") != 0 {
			duplicatePath := false

			for _, foundPath := range foundPaths {
				if strings.Compare(foundPath, cmd.Path) == 0 {
					duplicatePath = true
					break
				}
			}

			if duplicatePath {
				continue
			}

			// Debug Output
			if debug {
				fmt.Println(line)
				fmt.Printf("\t(Compiler) %s\n", cmd.Compiler)
				fmt.Printf("\t(Path) %s\n", cmd.Path)
				fmt.Printf("\t(Directory) %s\n", cmd.Directory)
				fmt.Printf("\t(File) %s\n", cmd.File)
			}

			// Register Command & Record Path
			commands = append(commands, cmd)
			foundPaths = append(foundPaths, cmd.Path)
		}
	}

	return commands
}

func parseLine(line string) (c cmd.CompileCommand) {
	var cmd cmd.CompileCommand

	compiler := regex.CompilerRegex.FindString(line)
	path := regex.PathRegex.FindString(line)
	file := regex.FileRegex.FindString(line)
	directory, _, _ := strings.Cut(path, file)

	cmd.Directory = directory
	cmd.Command = line
	cmd.File = file
	cmd.Compiler = compiler
	cmd.Path = path

	return cmd
}
