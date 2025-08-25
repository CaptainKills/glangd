package main

import (
	"regexp"
	"testing"
)

func TestCompilerRegex(t *testing.T) {
	t.Run("Compiler: cc", func(t *testing.T) {
		command := "cc -Wall -Wextra -pedantic -O0 -g3 -std=c99 -I. -I../vep_0/libbsp/include  -L. -L../vep_0/libchannel/lib  channel_init.c  -lchannel-arm -lm -o channel_init"

		regex, err := regexp.Compile(compilerRegexExpr)
		if err != nil {
			t.Fatalf("Could not compile 'compiler' regex! %q\n", err)
		}

		got := regex.FindString(command)
		want := "cc"

		if got != want {
			t.Fatalf("got %q but wanted %q\n", got, want)
		}
	})

	t.Run("Compiler: gcc", func(t *testing.T) {
		command := "gcc -c -o build/main.o src/main.c -Wall -Wextra -Wformat -O3 -I inc/"

		regex, err := regexp.Compile(compilerRegexExpr)
		if err != nil {
			t.Fatalf("Could not compile 'compiler' regex! %q\n", err)
		}

		got := regex.FindString(command)
		want := "gcc"

		if got != want {
			t.Fatalf("got %q but wanted %q\n", got, want)
		}
	})

	t.Run("Compiler: clang", func(t *testing.T) {
		command := "clang /path/test.cc -DNDEBUG -fsyntax-only -resource-dir=/usr/lib/clang/lib/12/"

		regex, err := regexp.Compile(compilerRegexExpr)
		if err != nil {
			t.Fatalf("Could not compile 'compiler' regex! %q\n", err)
		}

		got := regex.FindString(command)
		want := "clang"

		if got != want {
			t.Fatalf("got %q but wanted %q\n", got, want)
		}
	})

	t.Run("Compiler: riscv32-unknown-elf-gcc", func(t *testing.T) {
		command := "/opt/riscv/bin/riscv32-unknown-elf-gcc -march=rv32imc -mabi=ilp32 -O2 -g3 -static --specs=nano.specs -mcmodel=medany -ffunction-sections -fdata-sections -fvisibility=hidden -nostartfiles   -I../../vep_0/libbsp/include -I../../vep_0/libfifo-riscv/include -I../../vep_0/tiles -I../../vep_0/libmutex-riscv/include -I../../vep_0/tiles -Iinclude -c -o libsrc/tile0_mb_greyscale.o  libsrc/greyscale.c"

		regex, err := regexp.Compile(compilerRegexExpr)
		if err != nil {
			t.Fatalf("Could not compile 'compiler' regex! %q\n", err)
		}

		got := regex.FindString(command)
		want := "riscv32-unknown-elf-gcc"

		if got != want {
			t.Fatalf("got %q but wanted %q\n", got, want)
		}
	})
}

func TestPathRegex(t *testing.T) {
	command := "gcc -c -o build/main.o src/main.c -Wall -Wextra -Wformat -O3 -I inc/"

	regex, err := regexp.Compile(pathRegexExpr)
	if err != nil {
		t.Fatalf("Could not compile 'path' regex! %q\n", err)
	}

	got := regex.FindString(command)
	want := "src/main.c"

	if got != want {
		t.Fatalf("got %q but wanted %q\n", got, want)
	}
}

func TestFileRegex(t *testing.T) {
	t.Run("File: main.c", func(t *testing.T) {
		command := "gcc -c -o build/main.o src/main.c -Wall -Wextra -Wformat -O3 -I inc/"

		regex, err := regexp.Compile(fileRegexExpr)
		if err != nil {
			t.Fatalf("Could not compile 'file' regex! %q\n", err)
		}

		got := regex.FindString(command)
		want := "main.c"

		if got != want {
			t.Fatalf("got %q but wanted %q\n", got, want)
		}
	})

	t.Run("File: channel_init.c", func(t *testing.T) {
		command := "cc -Wall -Wextra -pedantic -O0 -g3 -std=c99 -I. -I../vep_0/libbsp/include  -L. -L../vep_0/libchannel/lib  channel_init.c  -lchannel-arm -lm -o channel_init"

		regex, err := regexp.Compile(fileRegexExpr)
		if err != nil {
			t.Fatalf("Could not compile 'file' regex! %q\n", err)
		}

		got := regex.FindString(command)
		want := "channel_init.c"

		if got != want {
			t.Fatalf("got %q but wanted %q\n", got, want)
		}
	})

	t.Run("File: generate-json.c", func(t *testing.T) {
		command := "gcc -Wall -Wextra -pedantic -O0 -g3 -std=c99 -I. -I../vep_0/libbsp/include -L. -L../vep_0/libchannel/lib -o generate-json generate-json.c ../vep_0/libbsp/libsrc/platform.c -lchannel-arm -lm"

		regex, err := regexp.Compile(fileRegexExpr)
		if err != nil {
			t.Fatalf("Could not compile 'file' regex! %q\n", err)
		}

		got := regex.FindString(command)
		want := "generate-json.c"

		if got != want {
			t.Fatalf("got %q but wanted %q\n", got, want)
		}
	})
}
