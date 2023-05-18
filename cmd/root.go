package cmd

import (
	"bufio"
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"ccwc/pkg/wc"
)

var version = "0.1"
var countBytesFlag bool
var countLinesFlag bool
var countWordsFlag bool
var countCharsFlag bool

var rootCmd = &cobra.Command{
	Use:     "ccwc",
	Version: version,
	Short:   "print newline, word, and byte counts for each file",
	Run:     RunWC,
}

func RunWC(cmd *cobra.Command, args []string) {
	if len(args) > 0 {
		filePath := args[0]
		bytes, err := os.ReadFile(filePath)
		if err != nil {
			panic(err)
		}
		processFlags(bytes, filePath)
		return
	}
	var data []byte
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		data = append(data, scanner.Bytes()...)
		data = append(data, '\n')
	}
	processFlags(data, "")
}

func processFlags(bytes []byte, filePath string) {
	if countBytesFlag {
		fmt.Println(wc.GetByteCount(bytes), filePath)
		return
	}
	if countLinesFlag {
		fmt.Println(wc.GetLineCount(bytes), filePath)
		return
	}
	if countWordsFlag {
		fmt.Println(wc.GetWordCount(bytes), filePath)
		return
	}
	if countCharsFlag {
		fmt.Println(wc.GetCharCount(bytes), filePath)
		return
	}
	fmt.Println(" ", wc.GetLineCount(bytes), "", wc.GetWordCount(bytes), wc.GetByteCount(bytes), filePath)
	return
}

func init() {
	rootCmd.Flags().BoolVarP(&countBytesFlag, "bytes", "c", false, "print the byte counts")
	rootCmd.Flags().BoolVarP(&countLinesFlag, "lines", "l", false, "print the newline counts")
	rootCmd.Flags().BoolVarP(&countWordsFlag, "words", "w", false, "print the word counts")
	rootCmd.Flags().BoolVarP(&countCharsFlag, "chars", "m", false, "print the character counts")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
