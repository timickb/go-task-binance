package cli

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/timickb/task-17apr/internal/cli/api"
	"strings"
)

// rateCmd represents the rate command
var rateCmd = &cobra.Command{
	Use:   "rate",
	Short: "Get pair from Binance API",
	Run: func(cmd *cobra.Command, args []string) {
		pFlag := cmd.Flag("pair")

		split := strings.Split(pFlag.Value.String(), "-")
		if len(split) != 2 {
			fmt.Println("Please specify pair in format TOKEN1-TOKEN2")
			return
		}

		srv := api.New("localhost:3001", false)
		price, err := srv.GetPairPrice(pFlag.Value.String())
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		fmt.Println(price)
	},
}

func init() {
	rootCmd.AddCommand(rateCmd)
	rateCmd.PersistentFlags().String("pair", "", "")
}
