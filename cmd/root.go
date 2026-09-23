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
	"sort"
	"strings"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "pls [job]",
	Short: "Run named command jobs from a plsfile",
	Long: `pls runs named groups of shell commands defined in a local plsfile.

	For example, use "pls build" to run the job named "build". If no file is
	specified, pls uses ./plsfile. Run "pls" without a job name to list the
	available jobs.`,
	SilenceUsage: true,

	// Uncomment the following line if your bare application
	// has an action associated with it:
	RunE: func(cmd *cobra.Command, args []string) error {
		file, err := cmd.Flags().GetString("file")
		if err != nil {
			return err
		}

		plsF, err := plsfile.ReadFile(file)
		if err != nil {
			if errors.Is(err, os.ErrNotExist) {
				return fmt.Errorf("plsfile %q not found", file)
			}
			return fmt.Errorf("could not read plsfile %q: %w", file, err)
		}

		if len(args) == 0 {
			fmt.Fprintf(cmd.OutOrStdout(), "No job specified. Available jobs: %s\nRun 'pls <job>' to execute one.\n", getJobsList(plsF))
			return nil
		}

		job, found := plsF.Jobs[args[0]]
		if !found {
			return fmt.Errorf("job %q not found. Available jobs: %s", args[0], getJobsList(plsF))
		}

		err = runCommands(job.Commands)
		if err != nil {
			return fmt.Errorf("job %q failed: %w", args[0], err)
		}

		return nil
	},
}

func getJobsList(plsF plsfile.PlsFile) string {
	names := make([]string, 0, len(plsF.Jobs))
	for name := range plsF.Jobs {
		names = append(names, name)
	}
	sort.Strings(names)

	var builder strings.Builder

	for i, name := range names {
		if i > 0 {
			builder.WriteString(", ")
		}
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
	rootCmd.Flags().StringP("file", "f", "plsfile", "path to the plsfile to use")
}
