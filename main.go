package main

import (
	"runtime"

	"github.com/Quorumsco/contact/components/application"
	"github.com/Quorumsco/contact/components/database"
	"github.com/Quorumsco/contact/components/settings"
	"github.com/Quorumsco/contact/models"
	"github.com/Quorumsco/contact/routers"
	"github.com/codegangsta/cli"
)

//var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")

const HelpTemplate = `NAME:
   {{.Name}} - {{.Usage}}
USAGE:
   {{.Name}} {{if .Flags}}[global options] {{end}}command{{if .Flags}} [command options]{{end}} [arguments...]
VERSION:
   {{.Version}}{{if len .Authors}}
AUTHOR(S): 
   {{range .Authors}}{{ . }}{{end}}{{end}}
{{if .Commands}}
COMMANDS:
   {{range .Commands}}{{join .Names ", "}}{{ "\t" }}{{.Usage}}
   {{end}}{{end}}{{if .Flags}}OPTIONS:
   {{range .Flags}}{{.}}
   {{end}}{{end}}
`

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	//flag.Parse()
	//if *cpuprofile != "" {
	//f, err := os.Create(*cpuprofile)
	//if err != nil {
	//log.Fatal(err)
	//}
	//pprof.StartCPUProfile(f)
	//defer pprof.StopCPUProfile()
	//}

	app := cli.NewApp()
	app.Name = "contact"
	app.Usage = "Quorums contact backend"
	app.Version = "0.0.1"
	app.Before = serve
	app.Action = func(ctx *cli.Context) {}
	app.HideHelp = true
	cli.AppHelpTemplate = HelpTemplate
	app.Flags = append(app.Flags, []cli.Flag{
		cli.StringFlag{"cpu, cpuprofile", "", "cpu profiling", ""},
		cli.BoolFlag{"m, migrate", "migrate the database", ""},
		cli.StringFlag{"port, p", "8080", "server listening port", ""},
		cli.HelpFlag,
	}...)
	app.RunAndExitOnError()
}

func migrate(ctx *cli.Context) error {
	return database.Migrate(models.GetModels())
}

func serve(ctx *cli.Context) error {
	if ctx.Bool("migrate") {
		if err := migrate(ctx); err != nil {
			return err
		}
	}

	var app = &application.Application{}
	err := app.Init()
	if err != nil {
		return err
	}

	settings.Port = ctx.String("port")

	err = routers.Init(app)
	if err != nil {
		return err
	}

	return nil
}
