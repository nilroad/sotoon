package hello

import (
	"context"
	"fmt"
	"github.com/spf13/cobra"
)

type Command struct {
	name string
}

func (r *Command) Register(ctx context.Context) *cobra.Command {
	return &cobra.Command{
		Use:   "hello [name]",
		Short: "Say hello",
		Args:  cobra.MaximumNArgs(1),

		Run: func(_ *cobra.Command, args []string) {
			if len(args) == 0 {
				r.name = "sotoon"
			} else {
				r.name = args[0]
			}

			r.run(ctx)
		},
	}
}

func (r *Command) run(_ context.Context) {
	fmt.Println("Hello " + r.name)
}
