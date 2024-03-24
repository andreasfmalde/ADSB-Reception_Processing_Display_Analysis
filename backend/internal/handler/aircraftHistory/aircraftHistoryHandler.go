package aircraftHistory

import (
	"adsb-api/internal/db"
	"adsb-api/internal/global"
	"adsb-api/internal/global/errors"
	"adsb-api/internal/global/geoJSON"
	"adsb-api/internal/logger"
	"adsb-api/internal/utility/apiUtility"
	"fmt"
	"net/http"
)

var params = []string{"icao"}

// HistoryAircraftHandler handles HTTP requests for /aircraft/history endpoint.
func HistoryAircraftHandler(db db.Database) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := apiUtility.ValidateURL(r.URL.Path, r.URL.Query(), len(global.AircraftHistoryPath), params)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		switch r.Method {
		case http.MethodGet:
			handleHistoryAircraftGetRequest(w, r, db)
		default:
			http.Error(w, fmt.Sprintf(errors.MethodNotSupported, r.Method), http.StatusMethodNotAllowed)
		}
	}
}

// handleHistoryAircraftGetRequest handles GET requests for the /aircraft/history endpoint.
// Sends history data for aircraft given by the icao query parameter.
// A valid icao: "ABC123"
func handleHistoryAircraftGetRequest(w http.ResponseWriter, r *http.Request, db db.Database) {
	var search = r.URL.Query().Get("icao")
	res, err := db.GetHistoryByIcao(search)
	if err != nil {
		http.Error(w, errors.ErrorRetrievingAircraftWithIcao+search, http.StatusInternalServerError)
		logger.Error.Printf(errors.ErrorRetrievingAircraftWithIcao+search+": %q URL: %q", err, r.URL)
		return
	}
	if len(res) == 0 {
		http.Error(w, errors.NoAircraftFound, http.StatusNoContent)
		return
	}

	aircraft, err := geoJSON.ConvertHistoryModelToGeoJson(res)
	if err != nil {
		http.Error(w, errors.ErrorConvertingDataToGeoJson, http.StatusInternalServerError)
		logger.Error.Printf(errors.ErrorConvertingDataToGeoJson+": %q", err)
		return
	}

	apiUtility.EncodeJsonData(w, aircraft)
}
