// md2toon - Markdown to TOON transpiler
//
// Converts Fabric Markdown prompts to TOON (Token-Oriented Object Notation)
// format for significant token savings when sending prompts to LLMs.
//
// Usage:
//
//	md2toon file.md              # file → stdout
//	cat file.md | md2toon        # stdin → stdout
//	md2toon file.md -o out.toon  # file → file
//	md2toon -o out.toon file.md  # file → file (alternate order)
package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	output := flag.String("o", "", "output file")
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: md2toon [-o output.toon] [file.md]\n")
		fmt.Fprintf(os.Stderr, "       md2toon [file.md] -o output.toon\n")
		fmt.Fprintf(os.Stderr, "Converts Markdown prompts to TOON format.\n\n")
		flag.PrintDefaults()
	}
	// Parse flags allowing them after positional args
	args := os.Args[1:]
	var inputFile string
	for i := 0; i < len(args); i++ {
		if args[i] == "-o" && i+1 < len(args) {
			*output = args[i+1]
			i++
		} else if args[i] == "-h" || args[i] == "--help" {
			flag.Usage()
			os.Exit(0)
		} else if !strings.HasPrefix(args[i], "-") && inputFile == "" {
			inputFile = args[i]
		}
	}

	var content []byte
	var err error
	if inputFile != "" {
		content, err = os.ReadFile(inputFile)
	} else {
		content, err = io.ReadAll(os.Stdin)
	}
	if err != nil {
		fmt.Fprintf(os.Stderr, "md2toon: %v\n", err)
		os.Exit(1)
	}
	if len(content) == 0 {
		fmt.Fprintf(os.Stderr, "md2toon: empty input\n")
		os.Exit(1)
	}

	toon := PromptToTOON(ParseMarkdownPrompt(string(content)))

	if *output != "" {
		if err := os.WriteFile(*output, []byte(toon+"\n"), 0644); err != nil {
			fmt.Fprintf(os.Stderr, "md2toon: %v\n", err)
			os.Exit(1)
		}
	} else {
		fmt.Println(toon)
	}
}
