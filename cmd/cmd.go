package cmd

import "strings"

type CompileCommand struct {
	// JSON Information
	Directory string
	Command   string
	File      string

	// Additional Info
	Compiler string
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
