package flags

import (
	"github.com/git-town/git-town/v22/internal/config/configdomain"
	. "github.com/git-town/git-town/v22/pkg/prelude"
	"github.com/spf13/cobra"
)

const proposeNoWebLong = "no-web"

// type-safe access to the CLI arguments of type configdomain.ProposeNoWeb
func ProposeNoWeb() (AddFunc, ReadProposeNoWebFlagFunc) {
	addFlag := func(cmd *cobra.Command) {
		cmd.Flags().BoolP(proposeNoWebLong, "", false, "do not open a browser after creating the proposal")
	}
	readFlag := func(cmd *cobra.Command) (Option[configdomain.ProposeNoWeb], error) {
		return readBoolOptFlag[configdomain.ProposeNoWeb](cmd.Flags(), proposeNoWebLong)
	}
	return addFlag, readFlag
}

// ReadProposeNoWebFlagFunc is the type signature for the function that reads the "no-web" flag from the args to the given Cobra command.
type ReadProposeNoWebFlagFunc func(*cobra.Command) (Option[configdomain.ProposeNoWeb], error)
