package main

import (
	"log"
	"net"
	"net/http"
	"net/rpc"
	"os"
	"runtime"

	"github.com/codegangsta/cli"
	"github.com/jinzhu/gorm"
	"github.com/quorumsco/cmd"
	"github.com/quorumsco/contacts/controllers"
	"github.com/quorumsco/contacts/models"
	"github.com/quorumsco/databases"
	"github.com/quorumsco/logs"
	"github.com/quorumsco/settings"

	"gopkg.in/olivere/elastic.v2"
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	cmd := cmd.New()
	cmd.Name = "contacts"
	cmd.Usage = "quorums contacts backend"
	cmd.Version = "0.0.1"
	cmd.Before = serve
	cmd.Flags = append(cmd.Flags, []cli.Flag{
		cli.StringFlag{Name: "config, c", Usage: "configuration file", EnvVar: "CONFIG"},
		cli.HelpFlag,
	}...)
	cmd.RunAndExitOnError()
}

func serve(ctx *cli.Context) error {
	var (
		config settings.Config
		err    error
	)

	if ctx.String("config") != "" {
		config, err = settings.Parse(ctx.String("config"))
		if err != nil {
			logs.Error(err)
		}
	}

	if config.Debug() {
		logs.Level(logs.DebugLevel)
	}

	dialect, args, err := config.SqlDB()
	if err != nil {
		logs.Critical(err)
		os.Exit(1)
	}
	logs.Debug("database type: %s", dialect)

	var db *gorm.DB
	if db, err = databases.InitGORM(dialect, args); err != nil {
		logs.Critical(err)
		os.Exit(1)
	}

	logs.Debug("connected to %s", args)

	if config.Migrate() {
		db.AutoMigrate(models.Models()...)
		logs.Debug("database migrated successfully")
	}

	if config.Debug() {
		db.LogMode(true)
	}

	var server settings.Server
	server, err = config.Server()
	if err != nil {
		logs.Critical(err)
		os.Exit(1)
	}

	ElasticSettings, err := config.Elasticsearch()
	var client *elastic.Client
	client, err = elastic.NewClient(elastic.SetURL(ElasticSettings.String()))
	if err != nil {
		logs.Critical(err)
		os.Exit(1)
	}

	// // Use the IndexExists service to check if a specified index exists.
	// exists, err := client.IndexExists("contacts").Do()
	// if err != nil {
	// 	logs.Critical(err)
	// 	os.Exit(1)
	// }
	// if !exists {
	// 	createIndex, err := client.CreateIndex("contacts").Do()
	// 	if err != nil {
	// 		logs.Critical(err)
	// 	}
	// 	if !createIndex.Acknowledged {
	// 		logs.Critical("Index creation wasn't aknowledged")
	// 	}
	// }

	rpc.Register(&controllers.Search{Client: client})
	rpc.Register(&controllers.Contact{DB: db})
	rpc.Register(&controllers.Note{DB: db})
	rpc.Register(&controllers.Tag{DB: db})
	rpc.Register(&controllers.Mission{DB: db})
	rpc.HandleHTTP()

	l, e := net.Listen("tcp", server.String())
	if e != nil {
		log.Fatal("listen error:", e)
	}
	logs.Info("Listening on " + server.String())
	return http.Serve(l, nil)
}
