package cmd

import (
	"github.com/roger-russel/novel-grabber/internal/cmd/sources"
	v "github.com/roger-russel/novel-grabber/internal/cmd/version"
	"github.com/roger-russel/novel-grabber/pkg/structs/cmd"

	"github.com/spf13/cobra"
)

var rootCmd *cobra.Command

//Root Command
func Root(vf v.FullVersion) {

	var flags cmd.Flags

	rootCmd = &cobra.Command{
		Use:   "novel-grabber",
		Short: "novel-grabber",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	rootCmd.PersistentFlags().StringVarP(
		&flags.Dir, "dir", "d", "./",
		"The output dir: -d ~/Docs",
	)

	rootCmd.PersistentFlags().StringVarP(
		&flags.FormatType, "format-type", "t", "epub",
		"The output format: -t epub",
	)

	rootCmd.PersistentFlags().StringVarP(
		&flags.Novel, "novel", "n", "",
		"The novel wich will be taken: -n i-alone-level-up-solo-leveling-web-novel",
	)

	rootCmd.AddCommand(version(vf))
	addSources(rootCmd, &flags)

	rootCmd.Execute()

}

func addSources(rootCmd *cobra.Command, flags *cmd.Flags) {
	rootCmd.AddCommand(sources.Lightnovelworld(flags))
	rootCmd.AddCommand(sources.Wuxiaworld(flags))
}
