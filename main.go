package main

import (
	"boyter/portfold/data"
	"boyter/portfold/handlers"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
	"os"
)

func main() {
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	// Connect to the database
	db, err := openDb("root:@tcp(localhost:3306)/portfold?autocommit=true")
	if err != nil {
		errorLog.Fatal(err)
	}
	defer db.Close()

	// Initialize a new instance of application containing dependencies.
	app := handlers.Application{
		ErrorLog:     errorLog,
		InfoLog:      infoLog,
		ProjectModel: &data.ProjectModel{DB: db},
	}

	srv := &http.Server{
		Addr:     ":8080",
		ErrorLog: errorLog,
		Handler:  app.Routes(),
	}

	// Use the http.ListenAndServe() function to start a new web server. We pass in
	// two parameters: the TCP network address to listen on (in this case ":8080")
	// and the servemux we just created. If http.ListenAndServe() returns an error
	// we use the log.Fatal() function to log the error message and exit.
	infoLog.Println("Starting server on :8080")
	err = srv.ListenAndServe()
	errorLog.Fatal(err)
}

func openDb(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)

	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
