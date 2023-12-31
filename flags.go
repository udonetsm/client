package main

import (
	"log"

	"github.com/spf13/cobra"
	"github.com/udonetsm/client/http"
	"github.com/udonetsm/client/use"
)

var (
	// initialize commands
	target, newname, newnumber string
	newnumlist                 []string
	rootCmd                    = &cobra.Command{
		Use:   "cmngr",
		Short: "manage your contacts",
		Run: func(cmd *cobra.Command, args []string) {
			log.Println("See help using \"cmngr -h\"")
		},
	}
	delete = &cobra.Command{
		Use:   "delete",
		Short: "delete target contact",
		Run: func(cmd *cobra.Command, args []string) {
			http.Delete(target)
		},
	}
	//command update includes next subcommands:...
	update = &cobra.Command{
		Use:   "update",
		Short: "update target contact",
		Long: `
	For update number list you should create new contact with same name 
	and general number but with differrent list of additional numbers`,
	}
	//...changed general number ...
	number = &cobra.Command{
		Use:   "number",
		Short: "change number",
		Run: func(cmd *cobra.Command, args []string) {
			http.Upgrade(target, "number", newnumber, newname, newnumlist)
		},
	}

	// ...changed contact name...
	name = &cobra.Command{
		Use:   "name",
		Short: "change contact name",
		Run: func(cmd *cobra.Command, args []string) {
			http.Upgrade(target, "name", newnumber, newname, newnumlist)
		},
	}

	// ...canged additional number list for target command.
	numlist = &cobra.Command{
		Use:   "list",
		Short: "change additional number list (less or equal 3)",
		Run: func(cmd *cobra.Command, args []string) {
			use.LimitNumList(newnumlist)
			http.Upgrade(target, "listnumber", newnumber, newname, newnumlist)
		},
	}
	//show contact info
	info = &cobra.Command{
		Use:   "info",
		Short: "get info abount target contact",
		Run: func(cmd *cobra.Command, args []string) {
			http.Info(target)
		},
	}
	// creates new contact.
	// Number and name are required flag
	// Additional number list is optional flag
	create = &cobra.Command{
		Use:   "create",
		Short: "creates new contact",
		Run: func(cmd *cobra.Command, args []string) {
			// additional number list can't be more than 3
			// otherwise program craches
			use.LimitNumList(newnumlist)
			http.Create(target, newname, newnumlist)
		},
	}

	search = &cobra.Command{
		Use:   "search",
		Short: "find contact by target name",
		Run: func(cmd *cobra.Command, args []string) {
			http.Search(newname)
		},
	}
)

// several same flags in several different commands are required
// addTarget is a local function for set it
func addTarget(c ...*cobra.Command) {
	for _, item := range c {
		item.Flags().StringVarP(&target, "target", "t", "", "-t <general phone number>")
		item.MarkFlagRequired("target")
	}
}

// some commands includes other commands
// this function set subcommands in commands
func addSubcommandInCommand(command *cobra.Command, subcommands ...*cobra.Command) {
	for _, subcommand := range subcommands {
		command.AddCommand(subcommand)
	}
}

// Load flags
func init() {
	addSubcommandInCommand(rootCmd, delete, info, update, create, search)
	addSubcommandInCommand(update, number, name, numlist)
	addTarget(number, numlist, delete, info, name, create)

	name.Flags().StringVarP(&newname, "new", "n", "", "-n <new contact name>")
	name.MarkFlagRequired("new")

	number.Flags().StringVarP(&newnumber, "new", "n", target, "-n <new general number>")

	numlist.Flags().StringSliceVarP(&newnumlist, "new", "n", nil, "-n <addnumber1,...,addnumber3>")
	numlist.MarkFlagRequired("new")

	create.Flags().StringVarP(&newname, "fname", "f", "", "-f <full name of contact>")
	create.Flags().StringSliceVarP(&newnumlist, "list", "l", nil, "-l <addnumber1,...,addnumber3>")
	create.MarkFlagRequired("fname")

	search.Flags().StringVarP(&newname, "name", "n", "", "-n <target name>")
	search.MarkFlagRequired("name")
	rootCmd.Execute()
}
