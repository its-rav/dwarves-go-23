package cmd

import (
	"fmt"

	"github.com/its-rav/dwarves-go-23/ass-d2/sortarray"
	"github.com/spf13/cobra"
)

func NewCommand() *cobra.Command {
	var inputArr []string
	var sortType string

	command := &cobra.Command{
		Use: "sort",
		Run: func(cmd *cobra.Command, args []string) {
			rs := sortarray.Sort(sortType, inputArr)
			fmt.Println(rs)
		},
	}

	command.Flags().StringSliceVarP(&inputArr, "input array", "i", []string{}, "")
	command.Flags().StringVarP(&sortType, "sort type", "t", "int", "")

	return command
}
