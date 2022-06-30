package main

import (
	"fmt"
	//. "internal/tui"
	"log"

	"github.com/spf13/cobra"
	//. "internal/tui"
)

/*
	This var sets up the root command and then all other commands. The root command, according to Cobra's structure, is the first thing we hit when we run the program.
	Imagine it as an automatic constructor that's allowing us to run an instance of this program.
*/
var rootCmd = &cobra.Command{
	Use: "put usage example here",
	//TraverseChildren: true,
	Short: "This program shows disk usage",
	Long:  "Put longer version of Short here",
	Run: func(cmd *cobra.Command, args []string) {
		v, _ := cmd.Flags().GetBool("version")
		if v {
			version()
		} else {
			log.Fatal("version unavailable")
		}
	},
}

func version() {
	fmt.Println("Version goes here")
}

func file() {

}

func init() {
	rootCmd.PersistentFlags().BoolP("file", "f", true, "-f [FILE] Loads the given file, which has earlier been created with the -o option. If [FILE] is equivalent to -, the file is read from standard input.")
	rootCmd.PersistentFlags().BoolP("version", "v", true, "-v or --version: will output the version of godu currently in use")

}

func main() {
	//commented out to avoid going into tui every time I tested code.
	/*p := tea.NewProgram(InitialModel(), tea.WithAltScreen())
	if err := p.Start(); err != nil {
		log.Fatal(err)
	}*/
	//p := tea.NewProgram(InitialModel(), tea.WithAltScreen())

	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}

}
