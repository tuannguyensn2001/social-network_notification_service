package cmd

import (
	"github.com/spf13/cobra"
	"social-work_notification_service/src/config"
)

type command = func(i config.IConfig) *cobra.Command

func GetRoot(config config.IConfig) *cobra.Command {
	cmds := []command{
		server,
	}

	root := &cobra.Command{}

	for _, item := range cmds {
		root.AddCommand(item(config))
	}

	return root
}
