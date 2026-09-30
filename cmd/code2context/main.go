package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	maxDepth := flag.Int("depth", 3, "Maximum directory depth to scan")
	ignorePatterns := flag.String("ignore", ".git,node_modules,vendor", "Comma-separated patterns to ignore")
	outputFile := flag.String("out", "", "Output file (default: stdout)")
	flag.Usage = printUsage
	flag.Parse()

	// hasStdin is true when stdin is not a terminal.
	stdinInfo, _ := os.Stdin.Stat()
	hasStdin := (stdinInfo.Mode() & os.ModeCharDevice) == 0

	var jsonData []byte
	var err error

	if hasStdin {
		// Stdin mode: read file list from stdin, instructions from argument
		if flag.NArg() != 1 {
			fmt.Fprintf(os.Stderr, "Error: When piping file list via stdin, provide exactly 1 argument: <instructions>\n")
			fmt.Fprintf(os.Stderr, "Usage: find . -name '*.go' | code2context \"instructions\"\n")
			os.Exit(1)
		}

		instructions := flag.Arg(0)

		var files []string
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			line := strings.TrimSpace(scanner.Text())
			if line != "" {
				files = append(files, line)
			}
		}
		if err := scanner.Err(); err != nil {
			fmt.Fprintf(os.Stderr, "Error reading stdin: %v\n", err)
			os.Exit(1)
		}

		if len(files) == 0 {
			fmt.Fprintf(os.Stderr, "Error: No files provided via stdin\n")
			os.Exit(1)
		}

		jsonData, err = ScanFiles(files, instructions)
	} else {
		// Directory mode: require directory and instructions arguments
		if flag.NArg() != 2 {
			printUsage()
			os.Exit(1)
		}

		directory := flag.Arg(0)
		instructions := flag.Arg(1)

		if info, err := os.Stat(directory); err != nil || !info.IsDir() {
			fmt.Fprintf(os.Stderr, "Error: Directory '%s' does not exist or is not a directory\n", directory)
			os.Exit(1)
		}

		jsonData, err = ScanDirectory(directory, *maxDepth, instructions, strings.Split(*ignorePatterns, ","))
	}

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error scanning: %v\n", err)
		os.Exit(1)
	}

	if *outputFile != "" {
		if err := os.WriteFile(*outputFile, jsonData, 0644); err != nil {
			fmt.Fprintf(os.Stderr, "Error writing file: %v\n", err)
			os.Exit(1)
		}
	} else {
		fmt.Print(string(jsonData))
	}
}

func printUsage() {
	fmt.Fprintf(os.Stderr, `code2context - Code project scanner for use with Fabric AI

Usage:
  code2context [options] <directory> <instructions>
  <file_list> | code2context [options] <instructions>

Examples:
  code2context . "Add input validation to all user inputs"
  code2context -depth 4 ./my-project "Implement error handling"
  code2context -out project.json ./src "Fix security issues"
  find . -name '*.go' | code2context "Refactor error handling"
  git ls-files '*.py' | code2context "Add type hints"

Options:
`)
	flag.PrintDefaults()
}
