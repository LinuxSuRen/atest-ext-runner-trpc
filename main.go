package main

import (
	"os"

	"github.com/linuxsuren/atest-ext-runner-trpc/cmd"
)

func main() {
	if err := cmd.NewRootCommand().Execute(); err != nil {
		os.Exit(1)
	}
}
