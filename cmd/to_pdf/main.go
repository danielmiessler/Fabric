// to_pdf
//
// Usage:
//   [no args]             Read from stdin, write to output.pdf
//   <file.tex>            Read from .tex file, write to <file>.pdf
//   <output.pdf>          Read stdin, write to specified PDF
//   <output>              Read stdin, write to <output>.pdf
//   <input> <output>      Read input (.tex appended if needed), write to output.pdf
//
// Examples:
//   to_pdf                  # stdin -> output.pdf
//   to_pdf doc.tex          # doc.tex -> doc.pdf
//   to_pdf report           # stdin -> report.pdf
//   to_pdf chap.tex out/    # Creates out/chap.pdf
//
// Error handling:
// - Validates pdflatex installation
// - Creates missing directories
// - Cleans temp files on exit

package main

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

// hasSuffix checks if a string ends with the given suffix, case-insensitive.
func hasSuffix(s, suffix string) bool {
	return strings.HasSuffix(strings.ToLower(s), strings.ToLower(suffix))
}

// resolveInputFile attempts to open the input file.
// If tryAppendTex is true and the initial attempt fails, it appends ".tex" and retries.
func resolveInputFile(filename string, tryAppendTex bool) (io.ReadCloser, string) {
	file, err := os.Open(filename)
	if err == nil {
		return file, filename
	}
	if tryAppendTex {
		newFilename := filename + ".tex"
		file, err = os.Open(newFilename)
		if err == nil {
			return file, newFilename
		}
	}
	return nil, ""
}

func main() {
	var input io.Reader
	var outputFile string

	args := os.Args
	argCount := len(args) - 1 // excluding the program name

	switch argCount {
	case 0:
		// Case 1: No arguments
		input = os.Stdin
		outputFile = "output.pdf"

	case 1:
		// Case 2: One argument
		arg := args[1]
		if hasSuffix(arg, ".tex") {
			// Case 2a: Argument ends with .tex
			file, actualName := resolveInputFile(arg, false)
			if file == nil {
				fmt.Fprintf(os.Stderr, "Error opening file: %s\n", arg)
				os.Exit(1)
			}
			defer file.Close()

			input = file

			// Derive output file name by replacing .tex with .pdf
			ext := filepath.Ext(actualName)
			outputFile = strings.TrimSuffix(actualName, ext) + ".pdf"
		} else if hasSuffix(arg, ".pdf") {
			// Case 2b: Argument ends with .pdf
			input = os.Stdin
			outputFile = arg
		} else {
			// Case 2c: Argument without .pdf
			input = os.Stdin
			outputFile = arg + ".pdf"
		}

	case 2:
		// Case 3: Two arguments
		inputArg := args[1]
		outputArg := args[2]

		// Resolve input file, ignore actualName
		file, _ := resolveInputFile(inputArg, true)
		if file == nil {
			fmt.Fprintf(os.Stderr, "Error: Input file '%s' not found, even after appending '.tex'.\n", inputArg)
			os.Exit(1)
		}
		defer file.Close()

		input = file

		// Resolve output file
		if hasSuffix(outputArg, ".pdf") {
			outputFile = outputArg
		} else {
			outputFile = outputArg + ".pdf"
		}

	default:
		fmt.Fprintf(os.Stderr, "Usage:\n")
		fmt.Fprintf(os.Stderr, "  %s                 # Read from stdin, output to 'output.pdf'\n", args[0])
		fmt.Fprintf(os.Stderr, "  %s <file.tex>      # Read from 'file.tex', output to 'file.pdf'\n", args[0])
		fmt.Fprintf(os.Stderr, "  %s <output.pdf>    # Read from stdin, output to 'output.pdf'\n", args[0])
		fmt.Fprintf(os.Stderr, "  %s <output>        # Read from stdin, output to '<output>.pdf'\n", args[0])
		fmt.Fprintf(os.Stderr, "  %s <input> <output># Read from 'input' (tries 'input.tex'), output to 'output.pdf'\n", args[0])
		os.Exit(1)
	}

	// Check if pdflatex is installed
	if _, err := exec.LookPath("pdflatex"); err != nil {
		fmt.Fprintln(os.Stderr, "Error: pdflatex is not installed or not in your PATH.")
		fmt.Fprintln(os.Stderr, "Please install a LaTeX distribution (e.g., TeX Live or MiKTeX) and ensure pdflatex is in your PATH.")
		os.Exit(1)
	}

	// Create a temporary directory
	tmpDir, err := os.MkdirTemp("", "latex_")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error creating temporary directory: %v\n", err)
		os.Exit(1)
	}
	defer os.RemoveAll(tmpDir)

	// Create a temporary .tex file
	tmpFilePath := filepath.Join(tmpDir, "input.tex")
	tmpFile, err := os.Create(tmpFilePath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error creating temporary file: %v\n", err)
		os.Exit(1)
	}

	// Copy input to the temporary file
	_, err = io.Copy(tmpFile, input)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error writing to temporary file: %v\n", err)
		tmpFile.Close()
		os.Exit(1)
	}
	tmpFile.Close()

	// Run pdflatex with nonstopmode
	cmd := exec.Command("pdflatex", "-interaction=nonstopmode", "-output-directory", tmpDir, "input.tex")
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error running pdflatex: %v\n", err)
		fmt.Fprintf(os.Stderr, "pdflatex output:\n%s\n", output)
		os.Exit(1)
	}

	// Check if PDF was actually created
	pdfPath := filepath.Join(tmpDir, "input.pdf")
	if _, err := os.Stat(pdfPath); os.IsNotExist(err) {
		fmt.Fprintln(os.Stderr, "Error: PDF file was not created. There might be an issue with your LaTeX source.")
		fmt.Fprintf(os.Stderr, "pdflatex output:\n%s\n", output)
		os.Exit(1)
	}

	// Move the output PDF to the desired location
	err = copyFile(pdfPath, outputFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error moving output file: %v\n", err)
		os.Exit(1)
	}

	// Remove the generated PDF from the temporary directory
	err = os.Remove(pdfPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error cleaning up temporary file: %v\n", err)
		// Not exiting as the main process succeeded
	}

	fmt.Printf("PDF created: %s\n", outputFile)
}

// copyFile copies a file from src to dst.
// If dst exists, it will be overwritten.
func copyFile(src, dst string) error {
	sourceFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer sourceFile.Close()

	// Ensure the destination directory exists
	dstDir := filepath.Dir(dst)
	err = os.MkdirAll(dstDir, 0755)
	if err != nil {
		return err
	}

	destFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer destFile.Close()

	_, err = io.Copy(destFile, sourceFile)
	if err != nil {
		return err
	}

	return destFile.Sync()
}
