package main

import (
	"context"
	"log"
	"os"
	"path/filepath"

	boss "github.com/hashload/boss/pkg/cli"
)

func main() {
	baseName := filepath.Base(os.Args[0])

	err := boss.InitializeBossCommandLine(baseName).Execute()
	if err != nil {
		if err != context.Canceled {
			log.Fatalf("An error occurred: %s\n", err)
		}
	}
}
