package cmd

import (
	"fmt"

	"github.com/rjdailey/labcheck/config"
	labssh "github.com/rjdailey/labcheck/ssh"
	"github.com/spf13/cobra"
)

var pingCmd = &cobra.Command{
	Use:   "ping",
	Short: "Test SSH connection to your homelab",
	RunE: func(cmd *cobra.Command, args []string) error {
		cfg, err := config.Load()
		if err != nil {
			return fmt.Errorf("could not load config, run 'labcheck init' first: %w", err)
		}

		client, err := labssh.Connect(cfg.Host, cfg.Port, cfg.User, cfg.KeyPath)
		if err != nil {
			return err
		}
		defer client.Close()

		out, err := client.Run("echo pong")
		if err != nil {
			return err
		}

		fmt.Println(out)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(pingCmd)
}