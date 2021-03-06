package list

import (
	"fmt"
	"net/http"

	"github.com/fatih/color"
	"github.com/gojektech/proctor/daemon"
	"github.com/gojektech/proctor/io"
	"github.com/gojektech/proctor/proctord/utility"
	"github.com/spf13/cobra"
)

func NewCmd(printer io.Printer, proctorEngineClient daemon.Client) *cobra.Command {
	return &cobra.Command{
		Use:     "list",
		Short:   "List procs available with proctor for execution",
		Long:    "List procs available with proctor for execution",
		Example: "proctor list",
		Run: func(cmd *cobra.Command, args []string) {
			procList, err := proctorEngineClient.ListProcs()
			if err != nil {
				if err.Error() == http.StatusText(http.StatusUnauthorized) {
					printer.Println(utility.UnauthorizedError, color.FgRed)
					return
				}

				printer.Println(utility.GenericListCmdError, color.FgRed)
				return
			}

			printer.Println("List of Procs:\n", color.FgGreen)

			for _, proc := range procList {
				printer.Println(fmt.Sprintf("%-40s %-100s", proc.Name, proc.Description), color.Reset)
			}

			printer.Println("\nFor detailed information of procs, run:\nproctor describe <proc_name>", color.FgGreen)
		},
	}
}
