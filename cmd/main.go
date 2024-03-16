package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/Mlstermass/task2/api/controller"
	"github.com/Mlstermass/task2/api/router"
	"github.com/Mlstermass/task2/pkg/env"
	"github.com/Mlstermass/task2/storage"
	"github.com/Mlstermass/task2/storage/immudb"
)

func main() {
	// load env variables to the Config struct
	var conf env.Config
	env.LoadConfig(&conf)

	listeningAddress := flag.String("listening-address",
		conf.AppHost, "Address which server handle")
	flag.Parse()

	immuDBClient, err := immudb.NewImmuDBConn(conf)
	if err != nil {
		log.Fatalf("Failed to initialize immudb client: %v", err)
	}
	log.Printf("Connected to immudb at %s:%d", conf.ImmuDBAdress, conf.ImmuBDPort)
	
	immuDriver := immudb.NewImmu(immuDBClient, conf)
	ctl := newControllers(conf, immuDriver)

	r := router.New(ctl, conf)

	srv := &http.Server{
		Handler: r,
		Addr:    *listeningAddress,
		// Enforce timeouts for servers you create!
		WriteTimeout: conf.WriteTimeout,
		ReadTimeout:  conf.ReadTimeout,
	}
	log.Println("Listening to", *listeningAddress)
	err = srv.ListenAndServe()
	if err != nil {
		log.Fatalf(err.Error())
	}
}

func newControllers(
	config env.Config,
	immuDriver storage.DocumentActions,
) controller.LogService {
	return controller.NewLogService(
		config, immuDriver)
}
