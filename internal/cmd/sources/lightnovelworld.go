package sources

import (
	"github.com/roger-russel/novel-grabber/internal/cmd/normalizers"
	"github.com/roger-russel/novel-grabber/internal/output"
	"github.com/roger-russel/novel-grabber/internal/source/lightnovelworld"
	"github.com/roger-russel/novel-grabber/pkg/structs/cmd"
	"github.com/roger-russel/novel-grabber/pkg/structs/novel"
	"github.com/spf13/cobra"
)

//Lightnovelworld is the command that handle this source
func Lightnovelworld(flags *cmd.Flags) (lightnovelworldCmd *cobra.Command) {

	lightnovelworldCmd = &cobra.Command{
		Use:   "lightnovelworld",
		Short: "Lightnovelworld source",
		Long:  `Lightnovelworld source handler`,
		Run: func(cmd *cobra.Command, args []string) {
			normalizers.NormalizeFlags(flags)
			var n *novel.Novel = &novel.Novel{}
			lightnovelworld.New(n, flags.Novel)
			output.Writer(n, flags.FormatType, flags.Dir)
		},
	}

	lightnovelworldCmd.AddCommand(&cobra.Command{
		Use:   "info",
		Short: "Get info about the novel",
		Long:  `Get information about the novel: author, chapter numbers, etc`,
		Run: func(cmd *cobra.Command, args []string) {
			lightnovelworld.Info(flags.Novel)
		},
	})

	return lightnovelworldCmd
}
