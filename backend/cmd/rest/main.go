package main

import (
	"adsb-api/internal/db"
	"adsb-api/internal/global"
	"adsb-api/internal/global/errorMsg"
	"adsb-api/internal/handler/aircraftCurrentHandler"
	"adsb-api/internal/handler/aircraftHistory"
	"adsb-api/internal/handler/defaultHandler"
	"adsb-api/internal/service/restService"
	"adsb-api/internal/utility/logger"
	"net/http"
	"os"
)

// main method for the RESTFUL API
func main() {
	// Initialize logger
	logger.InitLogger()
	// Initialize environment variables
	global.InitEnvironment()
	// Initialize the database
	database, err := db.InitDB()
	if err != nil {
		logger.Error.Fatalf("error opening database: %q", err)
	}

	defer func() {
		err = database.Close()
		if err != nil {
			logger.Error.Fatalf(errorMsg.ErrorClosingDatabase+": %q", err)
		}
	}()

	logger.Info.Printf("REST API successfully connected to database with database user: %s name: %s host: %s port: %d",
		global.DbUser, global.DbName, global.DbHost, global.DbPort)

	restSvc := restService.InitRestService(database)

	http.HandleFunc(global.DefaultPath, defaultHandler.DefaultHandler)
	http.HandleFunc(global.AircraftCurrentPath, aircraftCurrentHandler.CurrentAircraftHandler(restSvc))
	http.HandleFunc(global.AircraftHistoryPath, aircraftHistory.HistoryAircraftHandler(restSvc))

	port := os.Getenv("PORT")
	if port == "" {
		port = global.DefaultPort
		logger.Info.Printf("PORT has not been set. Using default port: %s", port)
	}

	logger.Info.Println("Listening on port: " + port)
	logger.Info.Fatal(http.ListenAndServe(":"+port, nil))

}
