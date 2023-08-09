package cmd

import (
	"fmt"
	"os"
	"strconv"

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
			fmt.Printf("Error: %v", err)
			os.Exit(1)
		}

		if !ok {
			fmt.Println("IP Info token is not set or is empty. Falling back to a basic IP lookup method.", ipinfo_token)
			fmt.Printf("Current public IP Address is: %s\n", public["public_ip"])

			os.Exit(0)
		}

		if ok {
			data, err := api.GetIPInfoData(ipinfo_token, public["public_ip"])
			if err != nil {
				fmt.Printf("Error: %v", err)
			}

			fmt.Printf("You're currently using the public IP address %s (%s)\n", data.IP, data.Org)
			fmt.Printf("\tCountry: %s (%s, %s, %s, %s)\n", data.Country, data.City, data.CountryName, data.Continent.Name, data.Postal)
			fmt.Printf("\tBased in the EU: %s\n", strconv.FormatBool(data.IsEU))
			fmt.Printf("\tLocation: %s\n", data.Location)
			fmt.Printf("\tTimezone: %s\n", data.Timezone)
			fmt.Println()
			fmt.Println("Local Interfaces:")
			api.DeviceList()
			fmt.Println()
		}
	},
}
