/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"errors"
	"fmt"
	"local/plsfile/internal/plsfile"
	"os"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "pls",
	Short: "pls help me run a set of commands",
	Long: `pls is a cli tool that helps you run a set of commands in your console.

For example for a set of commands that's necessary to build your application.`,

	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		file, err := cmd.Flags().GetString("file")
		if err != nil {
			panic(err)
		}

		plsF, err := plsfile.ReadFile(file)
		if err != nil {
			if errors.Is(err, os.ErrNotExist) {
				fmt.Printf("The file '%s' was not found. Use --help for more information.\n", file)
				return
			}
			panic(err)
		}

		if len(args) == 0 {
			fmt.Printf("No job name provided. Available are: %s\n", getJobsList(plsF))
			return
		}

		job, found := plsF.Jobs[args[0]]
		if !found {
			fmt.Printf("Job %s not found. Available jobs are: %s\n", args[0], getJobsList(plsF))
			return
		}

		err = runCommands(job.Commands)
		if err != nil {
			panic(err)
		}

	},
}

func getJobsList(plsF plsfile.PlsFile) string {
	var builder strings.Builder

	first := true
	for name, _ := range plsF.Jobs {
		if !first {
			builder.WriteString(", ")
		}
		first = false
		builder.WriteString(name)
	}

	return builder.String()
}

func runCommands(commands []string) error {
	script := "set -e\n" + strings.Join(commands, "\n")
	cmd := exec.Command("sh", "-c", script)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	return cmd.Run()
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
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.plsfile.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().StringP("file", "f", "plsfile", "which file/path should be used. Default is ./plsfile")
}
