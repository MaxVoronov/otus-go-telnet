package cmd

import (
	"errors"
	"fmt"
	"github.com/MaxVoronov/otus-go-telnet/internal/tcpclient"
	"github.com/spf13/cobra"
	"os"
	"time"
)

var (
	timeout int
	rootCmd = &cobra.Command{
		Use: "go-telnet host port",
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) < 2 {
				return errors.New("requires host and port")
			}
			return nil
		},
		Run: func(cmd *cobra.Command, args []string) {
			client, err := tcpclient.NewClient(&tcpclient.ConnectOptions{
				Host:    args[0],
				Port:    args[1],
				Timeout: time.Duration(timeout) * time.Second,
			})
			if err != nil {
				fmt.Printf("Error: %s\n", err)
			}
			defer client.Close()

			client.Run()
		},
	}
)

func init() {
	rootCmd.PersistentFlags().IntVar(&timeout, "timeout", 10, "Connection timeout in seconds")
}

// Execute Default command
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
