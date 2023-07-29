package main

import (
	"time"
)

const (
	DEFAULT_DB                   = "edgedb"
	DEFAULT_DB_USER              = "edgedb"
	DEFAULT_FAA_MASTER_FILE_URI  = "https://registry.faa.gov/database/ReleasableAircraft.zip"
	DEFAULT_FAA_MASTER_FILE_PATH = "ReleasableAircraft.zip"
	DEFAULT_FAA_MASTER_FILE_DIR  = "ReleasableAircraft"
	DEFAULT_DB_CONNECT_TIMEOUT   = 30 * time.Second
	DEFAULT_DB_CONCURRENCY       = 4
)

func main() {

	// Process ENGINE.txt
	processEnginesFile()

	// Process RESERVED.txt
	processReservedNNumbersFile()
}
