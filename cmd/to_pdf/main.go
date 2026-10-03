// to_pdf
//
// Usage:
//   [no args]             Read from stdin, write to output.pdf
//   <file.tex>            Read from .tex file, write to <file>.pdf
//   <output.pdf>          Read stdin, write to specified PDF
//   <output>              Read stdin, write to <output>.pdf
//   <input> <output>      Read <input> or <input>.tex, write to <output>.pdf

package main

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

// hasSuffix is strings.HasSuffix with a case-insensitive comparison.
func hasSuffix(s, suffix string) bool {
	return strings.HasSuffix(strings.ToLower(s), strings.ToLower(suffix))
}

// resolveInputFile opens filename and returns the file and the name it opened.
// If tryAppendTex is true and the first open fails, it retries with ".tex" appended.
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
	argCount := len(args) - 1 // without the program name

	switch argCount {
	case 0:
		input = os.Stdin
		outputFile = "output.pdf"

	case 1:
		arg := args[1]
		if hasSuffix(arg, ".tex") {
			file, actualName := resolveInputFile(arg, false)
			if file == nil {
				fmt.Fprintf(os.Stderr, "Error opening file: %s\n", arg)
				os.Exit(1)
			}
			defer file.Close()

			input = file

			ext := filepath.Ext(actualName)
			outputFile = strings.TrimSuffix(actualName, ext) + ".pdf"
		} else if hasSuffix(arg, ".pdf") {
			input = os.Stdin
			outputFile = arg
		} else {
			input = os.Stdin
			outputFile = arg + ".pdf"
		}

	case 2:
		inputArg := args[1]
		outputArg := args[2]

		file, _ := resolveInputFile(inputArg, true)
		if file == nil {
			fmt.Fprintf(os.Stderr, "Error: Input file '%s' not found, even after appending '.tex'.\n", inputArg)
			os.Exit(1)
		}
		defer file.Close()

		input = file

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

	if _, err := exec.LookPath("pdflatex"); err != nil {
		fmt.Fprintln(os.Stderr, "Error: pdflatex is not installed or not in your PATH.")
		fmt.Fprintln(os.Stderr, "Please install a LaTeX distribution (e.g., TeX Live or MiKTeX) and ensure pdflatex is in your PATH.")
		os.Exit(1)
	}

	tmpDir, err := os.MkdirTemp("", "latex_")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error creating temporary directory: %v\n", err)
		os.Exit(1)
	}
	defer os.RemoveAll(tmpDir)

	tmpFilePath := filepath.Join(tmpDir, "input.tex")
	tmpFile, err := os.Create(tmpFilePath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error creating temporary file: %v\n", err)
		os.Exit(1)
	}

	_, err = io.Copy(tmpFile, input)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error writing to temporary file: %v\n", err)
		tmpFile.Close()
		os.Exit(1)
	}
	tmpFile.Close()

	cmd := exec.Command("pdflatex", "-interaction=nonstopmode", "-output-directory", tmpDir, "input.tex")
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error running pdflatex: %v\n", err)
		fmt.Fprintf(os.Stderr, "pdflatex output:\n%s\n", output)
		os.Exit(1)
	}

	pdfPath := filepath.Join(tmpDir, "input.pdf")
	if _, err := os.Stat(pdfPath); os.IsNotExist(err) {
		fmt.Fprintln(os.Stderr, "Error: PDF file was not created. There might be an issue with your LaTeX source.")
		fmt.Fprintf(os.Stderr, "pdflatex output:\n%s\n", output)
		os.Exit(1)
	}

	err = copyFile(pdfPath, outputFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error moving output file: %v\n", err)
		os.Exit(1)
	}

	err = os.Remove(pdfPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error cleaning up temporary file: %v\n", err)
		// Do not exit. The output PDF is already written.
	}

	fmt.Printf("PDF created: %s\n", outputFile)
}

// copyFile copies src to dst. It creates missing parent directories and overwrites dst.
func copyFile(src, dst string) error {
	sourceFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer sourceFile.Close()

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
