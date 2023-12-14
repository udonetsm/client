package main

import (
	"github.com/spf13/cobra"
	"github.com/udonetsm/client/http"
)

var (
	target, newname, newnumber string
	newnumlist                 []string
	rootCmd                    = &cobra.Command{
		Use:   "contactmgr",
		Short: "manage your contacts",
		Run:   func(cmd *cobra.Command, args []string) { cmd.Println("", target) },
	}
	delete = &cobra.Command{
		Use:   "delete",
		Short: "delete target contact",
		Run: func(cmd *cobra.Command, args []string) {
		},
	}
	//command update includes next subcommands ...
	update = &cobra.Command{
		Use:   "update",
		Short: "update target contact",
		Long: `
	For update number list you should create new contact with same name 
	and general number but with differrent list of additional numbers`,
	}
	//...changes general number ...
	number = &cobra.Command{
		Use:   "number",
		Short: "change number",
		Run: func(cmd *cobra.Command, args []string) {
			http.Upgrade(target, newnumber, newname, newnumlist)
		},
	}

	// ...changes contact name.
	name = &cobra.Command{
		Use:   "name",
		Short: "change contact name",
		Run: func(cmd *cobra.Command, args []string) {
			http.Upgrade(target, newnumber, newname, newnumlist)
		},
	}
	//show contact info
	info = &cobra.Command{
		Use:   "info",
		Short: "get info abount target contact",
		Run:   func(cmd *cobra.Command, args []string) { http.GetInfo(target) },
	}
	create = &cobra.Command{
		Use:   "create",
		Short: "creates new contact",
		Run: func(cmd *cobra.Command, args []string) {
			http.Create(target, newname, newnumlist)
			/* use.Create(target, newname, newnumlist) */
		},
	}
)

func addTarget(c ...*cobra.Command) {
	for _, item := range c {
		item.Flags().StringVarP(&target, "target", "t", "", "-t <phone number>")
		item.MarkFlagRequired("target")
	}
}

func addSubcommandInCommand(command *cobra.Command, subcommands ...*cobra.Command) {
	for _, subcommand := range subcommands {
		command.AddCommand(subcommand)
	}
}

func init() {
	addSubcommandInCommand(rootCmd, delete, info, update, create)
	addSubcommandInCommand(update, number, name)
	addTarget(number, delete, info, name, create)

	name.Flags().StringVarP(&newname, "new", "n", "", "-n <new contact name>")
	name.MarkFlagRequired("new")
	number.Flags().StringVarP(&newnumber, "new", "n", target, "-n <new general number>")

	create.Flags().StringVarP(&newname, "fname", "f", "", "-f <full name of contact>")
	create.Flags().StringSliceVarP(&newnumlist, "list", "l", nil, "-l <addnumber1,addnumber2,...addnumber5>")
	create.MarkFlagRequired("fname")

	rootCmd.Execute()
}
