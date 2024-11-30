package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	ignoreCase    bool
	recursive     bool
	afterContext  int
	beforeContext int
	context       int
)

func main() {
	var rootCmd = &cobra.Command{
		Use:   "mygrep [flags] PATTERN [FILE...]",
		Short: "A simplified grep-like search utility",
		Long: `mygrep is a simplified version of grep that supports basic search functionality
with options for case-insensitive search (-i), recursive search (-r),
and context lines (-A, -B, -C).`,
		Args: cobra.MinimumNArgs(1),
		Run:  runGrep,
	}

	// Define flags
	rootCmd.Flags().BoolVarP(&ignoreCase, "ignore-case", "i", false, "Perform case insensitive matching")
	rootCmd.Flags().BoolVarP(&recursive, "recursive", "r", false, "Recursively search subdirectories")
	rootCmd.Flags().IntVarP(&afterContext, "after-context", "A", 0, "Print NUM lines of trailing context")
	rootCmd.Flags().IntVarP(&beforeContext, "before-context", "B", 0, "Print NUM lines of leading context")
	rootCmd.Flags().IntVarP(&context, "context", "C", 0, "Print NUM lines of output context")

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func runGrep(cmd *cobra.Command, args []string) {
	pattern := args[0]
	files := args[1:]

	if len(files) == 0 {
		files = []string{"-"}
	}

	if context > 0 {
		afterContext = context
		beforeContext = context
	}

	fmt.Printf("Searching for: %s\n", pattern)
	fmt.Printf("Options:\n")
	fmt.Printf("  Ignore case: %v\n", ignoreCase)
	fmt.Printf("  Recursive: %v\n", recursive)
	fmt.Printf("  After context: %d\n", afterContext)
	fmt.Printf("  Before context: %d\n", beforeContext)
	fmt.Printf("  Files: %v\n", files)
}
