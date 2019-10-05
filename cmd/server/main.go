package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"

	"github.com/felts94/al-go/storage"
	"github.com/pkg/errors"
)

type server struct {
	db     storage.Client
	router *httprouter.Router
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stdout, "%s\n", err.Error())
		os.Exit(1)
	}

}

func run() error {
	db, dbCleanup, err := setupDatabase("file")
	if err != nil {
		return errors.Wrap(err, "setup db")
	}
	defer dbCleanup()
	svr := &server{
		db:     db,
		router: httprouter.New(),
	}

	svr.setRoutes()

	err = http.ListenAndServe(":8080", svr.router)

	return err
}

func setupDatabase(t string) (storage.Client, func() error, error) {
	switch t {
	case "file":
		os.Mkdir("./storage", os.ModePerm)
		f, err := ioutil.TempFile("./storage/", "storage*.json")
		if err != nil {
			return nil, nil, errors.Wrap(err, "Creating temp file")
		}
		_, err = f.Write([]byte(`{"statusCheckCount": 0}`))
		if err != nil {
			return nil, nil, errors.Wrap(err, "Writing initial data")
		}
		fc := storage.FileClient{
			File: f,
		}
		// pass the pointer of fileclient, that implements the storage.Client interface
		// https://stackoverflow.com/questions/44370277/type-is-pointer-to-interface-not-interface-confusion
		return &fc, fc.Cleanup, nil
	}
	return nil, nil, errors.New("Invalid DB type")
}

func (s *server) Logf(msg string, v ...interface{}) {
	log.Printf(msg, v...)
}
