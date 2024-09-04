/*
MIT License

Copyright (c) 2024 nilroad

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
*/

package main

import (
	"context"
	"github.com/nilroad/kateb"
	"github.com/spf13/cobra"
	"os/signal"
	"sotoon/cmd/http"
	"sotoon/cmd/migration"
	"sotoon/internal/config"
	"syscall"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	cfg, err := config.Load()
	if err != nil {
		kateb.Fatal("failed to load config", map[string]any{
			"err": err.Error(),
		})
	}

	// registering commands
	root := registerCommands(ctx, cfg)

	err = root.Execute()
	if err != nil {
		return
	}
}

// registerCommands all commands register in this function
func registerCommands(ctx context.Context, cfg *config.Config) *cobra.Command {
	root := &cobra.Command{}

	HTTPCmd := new(http.Command)
	migrateCmd := new(migration.Command)

	root.AddCommand(HTTPCmd.Register(ctx, cfg), migrateCmd.Register(ctx, cfg))

	return root
}
