package sources

import (
	"github.com/roger-russel/novel-grabber/internal/cmd/normalizers"
	"github.com/roger-russel/novel-grabber/internal/output"
	"github.com/roger-russel/novel-grabber/internal/source/wuxiaworld"
	"github.com/roger-russel/novel-grabber/pkg/structs/cmd"
	"github.com/roger-russel/novel-grabber/pkg/structs/novel"
	"github.com/spf13/cobra"
)

//Wuxiaworld is the command that handle this source
func Wuxiaworld(flags *cmd.Flags) (wuxiaworldCmd *cobra.Command) {

	wuxiaworldCmd = &cobra.Command{
		Use:   "wuxiaworld",
		Short: "wuxiaworld source",
		Long:  `wuxiaworld source handler`,
		Run: func(cmd *cobra.Command, args []string) {
			normalizers.NormalizeFlags(flags)
			var n *novel.Novel = &novel.Novel{}
			wuxiaworld.New(n, flags.Novel)
			output.Writer(n, flags.FormatType, flags.Dir)
		},
	}

	return wuxiaworldCmd
}
