package main

import (
	"flag"
	"fmt"
	"os/user"
	"path/filepath"

	"github.com/LamkasDev/alpaca/cmd/alpaca_ai/ai"
	"github.com/LamkasDev/alpaca/cmd/alpaca_ai/run"
	"github.com/LamkasDev/alpaca/cmd/common/arch"
	"github.com/LamkasDev/alpaca/cmd/common/flags"
	"github.com/LamkasDev/alpaca/cmd/common/logger"
	"github.com/jwalton/gchalk"
)

func main() {
	u, _ := user.Current()
	arch.DataPath = filepath.Join(u.HomeDir, "Desktop", "alpaca")

	logger.Log(fmt.Sprintf("hi from %s :3\n", gchalk.Red(arch.AlpacaPlatform)))

	var quiet bool
	flag.BoolVar(&quiet, "q", false, "set to disable log")
	flag.Parse()
	logger.AlpacaLoggerEnabled = !quiet
	flags.ResolveColor()

	cai := ai.SetupAI()
	run.LoadAI(&cai)
	run.RunAI(&cai)

	logger.Log("bay! :33\n")
}
