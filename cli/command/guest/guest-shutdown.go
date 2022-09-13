package guest

import (
	"github.com/Telmate/proxmox-api-go/cli"
	"github.com/Telmate/proxmox-api-go/proxmox"
	"github.com/spf13/cobra"
)

var guest_shutdownCmd = &cobra.Command{
	Use:   "shutdown GUESTID",
	Short: "Shuts the speciefid guest down",
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		vmr := proxmox.NewVmRef(cli.ValidateIntIDset(args, "GuestID"))
		c := cli.NewClient()
		_, err = c.ShutdownVm(vmr)
		if err == nil {
			cli.PrintGuestStatus(GuestCmd.OutOrStdout(), vmr.VmId(), "stopped")
		}
		return
	},
}

func init() {
	GuestCmd.AddCommand(guest_shutdownCmd)
}
