package main

import (
	"bufio"
	"bytes"
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/edgedb/edgedb-go"
)

func processEnginesFile() {
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

	// read engine records
	engineRecords, err := os.ReadFile(filepath.Join(DEFAULT_FAA_MASTER_FILE_DIR, "ENGINE.txt"))
	if err != nil {
		log.Fatal(err)
	}

	opts := edgedb.Options{
		Database:       DEFAULT_DB,
		User:           DEFAULT_DB_USER,
		ConnectTimeout: DEFAULT_DB_CONNECT_TIMEOUT,
		Concurrency:    DEFAULT_DB_CONCURRENCY,
	}

	conn, err := edgedb.CreateClient(context.Background(), opts)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	scanner := bufio.NewScanner(bytes.NewReader(engineRecords))
	for scanner.Scan() {
		line := scanner.Text()

		// strip BOM if present
		line = strings.TrimPrefix(line, "\xef\xbb\xbf")

		// skip header
		if strings.HasPrefix(line, "CODE") {
			continue
		}

		// split line by commas
		fields := strings.Split(line, ",")

		// check for correct number of fields
		if len(fields) != 7 {
			log.Printf("Warning: Incorrect number of fields in line: %s: expected 7, got %d\n", line, len(fields))
			continue
		}

		// trim extra spaces from each field
		for i := range fields {
			fields[i] = strings.TrimSpace(fields[i])
		}

		// parse engine type
		engineType, err := strconv.Atoi(fields[3])
		if err != nil {
			log.Printf("Warning: Invalid engine type in line: %s\n", line)
			engineType = 0
			continue
		}

		// parse thrust and horsepower based on engine type
		horsepower, err := strconv.ParseInt(fields[4], 10, 64)
		if err != nil {
			log.Printf("Warning: Invalid horsepower in line: %s\n", line)
			continue
		}

		thrust, err := strconv.ParseInt(fields[5], 10, 64)
		if err != nil {
			log.Printf("Warning: Invalid thrust in line: %s\n", line)
			continue
		}

		if EngineType(engineType) == 4 || engineType == 5 || engineType == 6 {
			horsepower = 0
		} else {
			thrust = 0
		}
		// construct an EngineModel object
		engineModel := EngineModel{
			code: fields[0][3:],
			manufacturer: EngineManufacturer{
				name: strings.TrimSpace(fields[1]),
				code: fields[0][:3],
			},
			name:        strings.TrimSpace(fields[2]),
			engine_type: EngineType(engineType),
			horsepower:  horsepower,
			thrust:      thrust,
		}

		// confirm that the engine manufacturer exists
		var manufacturerExists bool
		err = conn.QuerySingle(
			context.Background(),
			engineManufacturerSelectQuery,
			&manufacturerExists,
			map[string]interface{}{
				"manufacturer_name": engineModel.manufacturer.name,
			})

		// check for errors and insert the engine manufacturer if it doesn't exist
		if !manufacturerExists {
			log.Printf("Manufacturer %s does not exist; inserting\n", engineModel.manufacturer.name)
			// insert the engine manufacturer
			err = conn.Execute(
				context.Background(),
				engineManufacturerInsertQuery,
				map[string]interface{}{
					"name": engineModel.manufacturer.name,
					"code": engineModel.manufacturer.code,
				})
			if err != nil {
				log.Printf("Error saving engine manufacturer: %v\n", err)
				fmt.Printf("manufacturer name: %s\n", engineModel.manufacturer.name)
				fmt.Printf("manufacturer code: %s\n", engineModel.manufacturer.code)
				continue
			}
		}

		if err != nil {
			// DEBUG: print the engine manufacturer that failed to check
			log.Printf("Error checking for engine manufacturer: %v\n", err)
			continue
		}

		// Save or process the engineModel
		err = conn.Execute(
			context.Background(),
			engineModelInsertQuery,
			map[string]interface{}{
				"code":         engineModel.code,
				"manufacturer": engineModel.manufacturer.name,
				"name":         engineModel.name,
				"engine_type":  engineModel.engine_type.String(),
				"horsepower":   engineModel.horsepower,
				"thrust":       engineModel.thrust,
			})

		if err != nil {
			// DEBUG: print the engine model that failed to save
			log.Printf("Error saving engine model: %v\n", err)
			fmt.Printf("name: %s\n", engineModel.name)
			fmt.Printf("manufacturer name %s\n", engineModel.manufacturer.name)
			continue
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
