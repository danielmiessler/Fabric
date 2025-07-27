package main

import (
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

	// Require exactly two positional arguments: directory and instructions
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
	jsonData, err := ScanDirectory(directory, *maxDepth, instructions, strings.Split(*ignorePatterns, ","))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error scanning directory: %v\n", err)
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

Examples:
  code_helper . "Add input validation to all user inputs"
  code_helper -depth 4 ./my-project "Implement error handling"
  code_helper -out project.json ./src "Fix security issues"

Options:
`)
	flag.PrintDefaults()
}
