package main

import (
	"adsb-api/internal/db"
	"adsb-api/internal/global"
	"adsb-api/internal/handler/currentAircraftHandler"
	"adsb-api/internal/handler/defaultHandler"
	"adsb-api/internal/logger"
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"os"
)

// main method for the RESTFUL API
func main() {
	logger.InitLogger()
	global.InitEnvironment()
	dbConn, err := db.InitDatabase()
	if err != nil {
		logger.Error.Fatalf("Error opening database: %q", err)
	} else {
		logger.Info.Println("Successfully connected to database.")
	}

	defer func(conn *sql.DB) {
		err := db.CloseDatabase(conn)
		if err != nil {
			logger.Error.Fatalf("Could not close database connection: %q", err)
		}
	}(dbConn)

	http.HandleFunc(global.DefaultPath, defaultHandler.DefaultHandler)
	http.HandleFunc(global.CurrentAircraftPath, currentAircraftHandler.CurrentAircraftHandler(dbConn))

	port := os.Getenv("PORT")

	if port == "" {
		log.Println("$PORT has not been set. Default: " + global.DefaultPort)
		port = global.DefaultPort
	}

	logger.Info.Println("Listening on port: " + port)
	logger.Info.Fatal(http.ListenAndServe(":"+port, nil))

}
