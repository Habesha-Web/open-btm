package manager

import (
	"open-btm.com/messages"

	"github.com/spf13/cobra"
)

var (
	startconsumercli = &cobra.Command{
		Use:   "start",
		Short: "start rabbit consumer",
		Long:  "Start rabbit app consumer",
		Run: func(cmd *cobra.Command, args []string) {
			startconsumer()
		},
	}
)

func startconsumer() {
	messages.RabbitConsumer("esb")
}

func init() {
	goFrame.AddCommand(startconsumercli)

}
