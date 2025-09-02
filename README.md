# glangd
A compile_commands.json generator for clangd written in Go.

# Features

- **Make/Cmake compatible:** run `glangd` alongside you existing `make` or `cmake` workflow.
- **Prevent Duplicate Files:** `glangd` will prevent duplicate entries when a source file is processed multiple times by make.
- **JSON Output Format:** `glangd` will automatically output a `compile_commands.json`.

# Installation
Glangd can easily be installed via the Golang CLI:

```bash
go install github.com/CaptainKills/glangd@latest
```

Or simply clone the repository and build from source:
```bash
git clone https://github.com/CaptainKills/glangd.git
cd glangd
go install .
```

# Getting Started
Glangd can be used in a couple of different ways depending on your liking.
The first and easiest way is to pipe any Makefile output directly into glangd:

```bash
make | glangd
```

In this way glangd will create a `compile_commands.json` in the current working directory.
In case you want to read from a file, you can do this as follows:

```bash
make > make.log
glangd -f make.log
```

This will prompt glangd to read the output of the `make` command via the make.log file
and process the output directly. glangd also provides the option to output the processed contents to another file with the `-o` flag:
```bash
make | glangd -o path/to/output.json
```


Note that the specified output file needs to be a `.json` file.

In case you would like to see how glangd analyses your Makefile output, you can use the `-d` flag to enable debug mode:
```bash
make | glangd -d

# The following console output will be produced:
Input Path: examples/emulator.log
Output Path: compile_commands.json
Debug Enabled: true

gcc -c -o build/main.o src/main.c -Wall -Wextra -Wformat -O3 -I inc/
        (Compiler) gcc      # glangd found the compiler
        (Path) src/main.c   # glangd founf the file path
        (Directory) src/    # glangd found the file directory
        (File) main.c       # glangd found the source file
```

The generated output of glangd will look as follows as the `compile_commands.json`:
```json
[
	{
		"directory": "src/",
		"command": "gcc -c -o build/main.o src/main.c -Wall -Wextra -Wformat -O3 -I inc/",
		"file": "main.c"
	}
]
```

Assuming you placed the `compile_commands.json` file in the correct directory, your clangd tooling should automatically pick up the file.
In case of issues, please provide the input as a text/log file, and the corresponding debug output from glangd (and clangd).

