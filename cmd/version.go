package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	version  = "3.0.2"
	codename = "Xray-Server"
	intro    = "Xray-Server is a simple and easy-to-use server application"
)

func init() {
	rootCmd.AddCommand(&cobra.Command{
		Use:   "version",
		Short: "Print current version of XrayR",
		Run: func(cmd *cobra.Command, args []string) {
			showVersion()
		},
	})
}

func showVersion() {
	fmt.Printf("%s %s (%s) \n", codename, version, intro)
}
