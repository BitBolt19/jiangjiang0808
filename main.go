package main

import (
	_ "nez-server/internal/packed"

	_ "nez-server/internal/logic"

	"github.com/gogf/gf/v2/os/gctx"

	"nez-server/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
