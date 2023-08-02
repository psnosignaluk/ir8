package cmd

import (
	"fmt"
	"os"

	"github.com/psnosignaluk/ir8/api"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(selfCmd)
}

var selfCmd = &cobra.Command{
	Use:   "self",
	Short: "Dig up your local network settings",
	Long:  "Use various tools to determine your local network configuration.",
	Run: func(cmd *cobra.Command, args []string) {
		ipinfo_token, ok := os.LookupEnv("IR8_IPINFO_TOKEN")

		fmt.Println("Running diagnostics on self.")
		public, err := api.SelfGetPublicIP()
		if err != nil {
			fmt.Errorf("Error: %v", err)
		}

		if !ok {
			fmt.Printf("%s is not set or is empty. Falling back to another IP lookup method.\n", ipinfo_token)

			fmt.Printf("Current public IP Address is: %s\n", public["public_ip"])

			os.Exit(0)
		}

		if ok {
			data, err := api.GetIPInfoData(ipinfo_token, public["public_ip"])
			if err != nil {
				fmt.Errorf("%v", err)
			}

			fmt.Printf("You're currently using the public IP address %s\n", data["ip"])
			fmt.Printf("\tCountry: %s (%s, %s)\n", data["country"], data["city"], data["countryName"])
			fmt.Printf("\tBased in the EU: %s\n", data["isEU"])
			fmt.Println()
		}
	},
}
