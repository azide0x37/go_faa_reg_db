package main

import (
	"bufio"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func processReservedNNumbersFile() {
	// download and unzip FAA data
	err := downloadFile(DEFAULT_FAA_MASTER_FILE_URI, DEFAULT_FAA_MASTER_FILE_PATH)
	if err != nil {
		log.Fatal(err)
	}

	err = unzip(DEFAULT_FAA_MASTER_FILE_PATH, DEFAULT_FAA_MASTER_FILE_DIR)
	if err != nil {
		log.Fatal(err)
	}

	defer os.RemoveAll(DEFAULT_FAA_MASTER_FILE_DIR) // clean up
	defer os.Remove(DEFAULT_FAA_MASTER_FILE_PATH)   // clean up

	f, err := os.Open(filepath.Join(
		DEFAULT_FAA_MASTER_FILE_DIR,
		"RESERVED.txt"))
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	var records []NNumberRecord

	for scanner.Scan() {
		line := scanner.Text()

		// strip BOM if present
		line = strings.TrimPrefix(line, "\xef\xbb\xbf")

		if strings.HasPrefix(line, "N-NUMBER") {
			continue
		}

		reserveDateStr := safeSubstring(line, 158, 166)
		expirationDateStr := safeSubstring(line, 170, 178)
		purgeDateStr := safeSubstring(line, 185, 194)

		var reserveDate, expirationDate, purgeDate time.Time
		var err error
		if reserveDateStr != "" {
			reserveDate, err = time.Parse("20060102", reserveDateStr)
			if err != nil {
				log.Fatal(err)
			}
		}
		if expirationDateStr != "" {
			expirationDate, err = time.Parse("20060102", expirationDateStr)
			if err != nil {
				log.Fatal(err)
			}
		}
		if purgeDateStr != "" {
			purgeDate, err = time.Parse("20060102", purgeDateStr)
			if err != nil {
				log.Fatal(err)
			}
		}

		unassigned := false
		registrant := safeSubstring(line, 6, 56)
		if strings.HasPrefix(registrant, "CANCELLED/NOT ASSIGNED") {
			unassigned = true
		}

		record := NNumberRecord{
			NNumber:          safeSubstring(line, 0, 5),
			Registrant:       registrant,
			StreetAddress:    safeSubstring(line, 57, 90),
			Street2:          safeSubstring(line, 91, 124),
			City:             safeSubstring(line, 125, 143),
			State:            safeSubstring(line, 144, 146),
			ZipCode:          safeSubstring(line, 147, 157),
			ReserveDate:      reserveDate,
			TypeReservation:  safeSubstring(line, 167, 169),
			ExpirationDate:   expirationDate,
			NNumberForChange: safeSubstring(line, 179, 184),
			PurgeDate:        purgeDate,
			Unassigned:       unassigned,
		}
		records = append(records, record)
		//DEBUG: fmt.Printf("Record parsed: %+v\n", record)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// insert reserved N-Number records into the database
	insertIntoEdgeDB(records)
}
