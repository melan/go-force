package cmd

import (
	"fmt"
	"os"

	"github.com/melan/go-force/sobjects"
	"github.com/spf13/cobra"
)

var (
	sobject, id string
)

var getCmd = &cobra.Command{
	Use: "get",
	Run: func(cmd *cobra.Command, args []string) {
		if sobject == "" || id == "" {
			panic("Both sobject and id parameters must be defined")
		}

		client, err := initClient()
		if err != nil {
			fmt.Println(fmt.Errorf("can't initialize client for Force.com: %v", err))
			os.Exit(1)
		}

		if val, ok := sobjects.SObjectsImplementations[sobject]; !ok {
			fmt.Printf("SObject %s is unsupported at this moment. You can try to run %s gen %s to "+
				"generate representation for the SObject, rebuild the client and try again\n",
				sobject, os.Args[0], sobject)
			os.Exit(1)
		} else {
			if err := client.GetSObject(id, make([]string, 0), val); err != nil {
				fmt.Printf("can't get SObject %s with ID %s: %v\n", sobject, id, err)
				os.Exit(1)
			}

			fmt.Printf("Found SObject %s with ID %s:\n%v\n", sobject, id, val)
		}
	},
}

func init() {
	getCmd.Flags().StringVar(&sobject, "sobject", "", "Name of an SObject to query, like User or Lead")
	getCmd.Flags().StringVar(&id, "id", "", "Id of an SObject to query")
	getCmd.MarkFlagRequired("sobject")
	getCmd.MarkFlagRequired("id")
	rootCmd.AddCommand(getCmd)
}
