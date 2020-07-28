package cmd

import (
	v "novel-grabber/internal/cmd/version"
	"novel-grabber/internal/source/lightnovelworld"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd *cobra.Command

//Root Command
func Root(vf v.FullVersion) {
	checkDefaultCommand()

	rootCmd = &cobra.Command{
		Use:   "novel-grabber",
		Short: "novel-grabber",
		Run: func(cmd *cobra.Command, args []string) {

			var n lightnovelworld.Novel

			n.New("")

		},
	}

	rootCmd.AddCommand(version(vf))

}

func checkDefaultCommand() {
	if len(os.Args) < 2 {
		os.Args = append([]string{os.Args[0], "--help"}, os.Args[1:]...)
	}
}
