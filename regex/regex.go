package regex

import (
	"log"
	"regexp"
)

const (
	// Regex
	compilerRegexExpr = "cc|gcc|clang|riscv32-unknown-elf-gcc"
	pathRegexExpr     = "([^\\s]+[.]c)"
	fileRegexExpr     = "[a-zA-Z0-9_-]+[.]c"
)

var (
	CompilerRegex *regexp.Regexp
	PathRegex     *regexp.Regexp
	FileRegex     *regexp.Regexp
)

func InitRegex() {
	var err error

	CompilerRegex, err = regexp.Compile(compilerRegexExpr)
	if err != nil {
		log.Fatalf("Could not compile 'compiler' regex! %q\n", err)
	}

	PathRegex, err = regexp.Compile(pathRegexExpr)
	if err != nil {
		log.Fatalf("Could not compile 'path' regex! %q\n", err)
	}

	FileRegex, err = regexp.Compile(fileRegexExpr)
	if err != nil {
		log.Fatalf("Could not compile 'file' regex! %q\n", err)
	}
}
