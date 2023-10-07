/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bufio"
	"cc-json-parser/pkg"
	"fmt"
	"os"
	"slices"
	"strings"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "cc-json-parser",
	Short: "A code challenge JSON Parser",
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		validCharacters := []string{"{", "}"}
		fileNameFlag, _ := cmd.Flags().GetString("filename")

		if fileNameFlag != "" {
			fileContent, err := os.ReadFile(fileNameFlag)
			if err != nil {
				fmt.Fprintf(os.Stderr, "%v\n", err)
				os.Exit(1)
			} else if len(fileContent) == 0 {
				fmt.Fprintf(os.Stderr, "%v\n", "JSON Invalid: Empty File")
				os.Exit(1)
			}

			s := bufio.NewScanner(strings.NewReader(string(fileContent)))
			s.Split(bufio.ScanRunes)
			for s.Scan() {
				if !slices.Contains(validCharacters, s.Text()) {
					fmt.Fprintf(os.Stderr, "%v\n", "JSON Invalid: Characters not allowed in JSON")
					os.Exit(1)
				} else {
					fmt.Fprintf(os.Stdout, "%v\n", "Valid JSON")
					os.Exit(0)
				}
			}
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().StringVarP(&pkg.FileName, "filename", "f", "", "Defines the file path of the json to be parsed")
}
