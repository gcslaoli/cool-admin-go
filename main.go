package main

import (
	_ "cool-admin-go-simple/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"

	"cool-admin-go-simple/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.New())
}
