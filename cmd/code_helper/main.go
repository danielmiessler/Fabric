package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Command line flags
	maxDepth := flag.Int("depth", 3, "Maximum directory depth to scan")
	ignorePatterns := flag.String("ignore", ".git,node_modules,vendor", "Comma-separated patterns to ignore")
	outputFile := flag.String("out", "", "Output file (default: stdout)")
	flag.Usage = printUsage
	flag.Parse()

	// Check if stdin has data (is a pipe)
	stdinInfo, _ := os.Stdin.Stat()
	hasStdin := (stdinInfo.Mode() & os.ModeCharDevice) == 0

	var jsonData []byte
	var err error

	if hasStdin {
		// Stdin mode: read file list from stdin, instructions from argument
		if flag.NArg() != 1 {
			fmt.Fprintf(os.Stderr, "Error: When piping file list via stdin, provide exactly 1 argument: <instructions>\n")
			fmt.Fprintf(os.Stderr, "Usage: find . -name '*.go' | code_helper \"instructions\"\n")
			os.Exit(1)
		}

		instructions := flag.Arg(0)

		// Read file paths from stdin
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

		// Validate directory
		if info, err := os.Stat(directory); err != nil || !info.IsDir() {
			fmt.Fprintf(os.Stderr, "Error: Directory '%s' does not exist or is not a directory\n", directory)
			os.Exit(1)
		}

		// Parse ignore patterns and scan directory
		jsonData, err = ScanDirectory(directory, *maxDepth, instructions, strings.Split(*ignorePatterns, ","))
	}

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error scanning: %v\n", err)
		os.Exit(1)
	}

	// Output result
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
	fmt.Fprintf(os.Stderr, `code_helper - Code project scanner for use with Fabric AI

Usage:
  code_helper [options] <directory> <instructions>
  <file_list> | code_helper [options] <instructions>

Examples:
  code_helper . "Add input validation to all user inputs"
  code_helper -depth 4 ./my-project "Implement error handling"
  code_helper -out project.json ./src "Fix security issues"
  find . -name '*.go' | code_helper "Refactor error handling"
  git ls-files '*.py' | code_helper "Add type hints"

Options:
`)
	flag.PrintDefaults()
}
