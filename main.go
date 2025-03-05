package main

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"nez-server/internal/logic/process/scanner"
	_ "nez-server/internal/packed"

	_ "nez-server/internal/logic"

	"github.com/gogf/gf/v2/os/gctx"

	"nez-server/internal/cmd"
)

func main() {
	if err := gtime.SetTimeZone("UTC"); err != nil {
		g.Log().Error(context.Background(), err)
		return
	}
	go func() {
		//task.RunReward()
		scanner.Scanner()
	}()
	cmd.Main.Run(gctx.GetInitCtx())
}
