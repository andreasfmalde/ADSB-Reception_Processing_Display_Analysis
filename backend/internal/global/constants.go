package global

import "strings"

// Database variables
var (
	DbUser     string
	DbPassword string
	DbName     = "adsb_db"
	DbHost     = "localhost"
	DbPort     = 5432
)

// API constants
const (
	DefaultPort         = "8080"
	VERSION             = "1.0.3"
	DefaultPath         = "/"
	AircraftCurrentPath = "/aircraft/current/"
	AircraftHistoryPath = "/aircraft/history/"
)

var (
	CurrentPathMaxLength = len(strings.Split(AircraftCurrentPath, "/")) - 1
	HistoryPathMaxLength = len(strings.Split(AircraftHistoryPath, "/"))
)

// SBS processing constants
var (
	SbsSource       string
	WaitingTime     = 4
	UpdatingPeriod  = 10
	MaxDaysHistory  = 1
	CleanupSchedule = "0 0 * * *" // once a day
)
