package parser

import (
	"fmt"
	"strings"

	cmd "github.com/CaptainKills/glangd/cmd"
	file "github.com/CaptainKills/glangd/file"
	regex "github.com/CaptainKills/glangd/regex"
)

func ParseStdin(pwd string, debug bool) []cmd.CompileCommand {
	var commands []cmd.CompileCommand
	var foundFiles []string

	stdin := file.ReadStdin()

	for _, line := range stdin {
		cmd := parseLine(pwd, line)

		if strings.Compare(cmd.Compiler, "") != 0 && strings.Compare(cmd.File, "") != 0 {
			duplicateFile := false

			for _, foundFile := range foundFiles {
				if strings.Compare(foundFile, cmd.File) == 0 {
					duplicateFile = true
					break
				}
			}

			if duplicateFile {
				continue
			}

			// Debug Output
			if debug {
				fmt.Println(line)
				fmt.Printf("\t(Compiler) %s\n", cmd.Compiler)
				fmt.Printf("\t(Directory) %s\n", cmd.Directory)
				fmt.Printf("\t(File) %s\n", cmd.File)
			}

			// Register Command & Record Path
			commands = append(commands, cmd)
			foundFiles = append(foundFiles, cmd.File)
		}
	}

	return commands
}

func ParseFile(pwd string, filePath string, debug bool) []cmd.CompileCommand {
	var commands []cmd.CompileCommand
	var foundFiles []string

	lines := file.ReadFile(filePath)

	for _, line := range lines {
		cmd := parseLine(pwd, line)

		if strings.Compare(cmd.Compiler, "") != 0 && strings.Compare(cmd.File, "") != 0 {
			duplicateFile := false

			for _, foundFile := range foundFiles {
				if strings.Compare(foundFile, cmd.File) == 0 {
					duplicateFile = true
					break
				}
			}

			if duplicateFile {
				continue
			}

			// Debug Output
			if debug {
				fmt.Println(line)
				fmt.Printf("\t(Compiler) %s\n", cmd.Compiler)
				fmt.Printf("\t(Directory) %s\n", cmd.Directory)
				fmt.Printf("\t(File) %s\n", cmd.File)
			}

			// Register Command & Record Path
			commands = append(commands, cmd)
			foundFiles = append(foundFiles, cmd.File)
		}
	}

	return commands
}

func parseLine(pwd string, line string) (c cmd.CompileCommand) {
	var cmd cmd.CompileCommand

	compiler := regex.CompilerRegex.FindString(line)
	path := regex.PathRegex.FindString(line)

	cmd.Directory = pwd
	cmd.Command = line
	cmd.File = path
	cmd.Compiler = compiler

	return cmd
}
