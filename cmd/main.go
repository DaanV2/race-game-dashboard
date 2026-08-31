package cmd

import (
	"embed"
	"fmt"
	"os"
)

// Its a hack, but this way we get the wails assets in
var assets embed.FS

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute(wailsAssets embed.FS) {
	assets = wailsAssets

	err := rootCmd.Execute()
	if err != nil {
		fmt.Println(fmt.Errorf("fatal: %w", err))
		os.Exit(1)
	}
}
