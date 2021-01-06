package main

import (
	"context"
	"log"
	"os"
	"path/filepath"

	"github.com/hashload/boss/pkg/boss"
)

func main() {
	baseName := filepath.Base(os.Args[0])

	err := boss.NewBossCommand(baseName).Execute()
	if err != nil {
		if err != context.Canceled {
			log.Fatalf("An error occurred: %s\n", err)
		}
	}
}
