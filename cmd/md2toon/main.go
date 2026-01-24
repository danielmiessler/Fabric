package main

import (
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {
	output := flag.String("o", "", "output file")
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: md2toon [file.md] [-o output.toon]\n")
		fmt.Fprintf(os.Stderr, "Converts Markdown prompts to TOON format.\n\n")
		flag.PrintDefaults()
	}
	flag.Parse()

	var content []byte
	var err error
	if flag.NArg() > 0 {
		content, err = os.ReadFile(flag.Arg(0))
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
