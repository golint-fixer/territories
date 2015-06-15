package main

import (
	"flag"
	"log"
	"os"
	"runtime"
	"runtime/pprof"

	"github.com/codegangsta/cli"

	"github.com/Quorumsco/contact/components/commands"
)

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	flag.Parse()
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}

	app := cli.NewApp()
	app.Name = "Quorums"
	app.Usage = "Quorums"
	app.Version = "0.0.1"
	app.Commands = []cli.Command{
		commands.CmdServe,
		commands.CmdMigrate,
	}
	app.Flags = append(app.Flags, []cli.Flag{cli.StringFlag{"cpu, cpuprofile", "", "cpu profiling", ""}}...)
	app.RunAndExitOnError()
}
