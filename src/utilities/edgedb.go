package utilities

import (
	"context"
	"log"
	"runtime"
	"strings"
	"time"

	"github.com/edgedb/edgedb-go"
)

func insertIntoEdgeDB(records []NNumberRecord) {
	opts := edgedb.Options{
		Database:       "edgedb",
		User:           "edgedb",
		ConnectTimeout: 35 * time.Second,
		Concurrency:    uint(runtime.NumCPU()), // use all available cores
	}

	ctx := context.Background()
	client, err := edgedb.CreateClient(ctx, opts)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	// insert n number records
	for _, record := range records {
		var inserted struct{ id edgedb.UUID }
		err = client.QuerySingle(ctx, `
			WITH ns := (SELECT NNumberRecord FILTER .n_number = <str>$n_number)

			INSERT NNumberRecord {
				n_number := <str>$n_number,
				registrant := <str>$registrant,
				street_address := <str>$street_address,
				street2 := <str>$street2,
				city := <str>$city,
				state := <str>$state,
				zip_code := <str>$zip_code,
				reserve_date := <cal::local_date>$reserve_date,
				type_reservation := <str>$type_reservation,
				expiration_date := <cal::local_date>$expiration_date,
				n_number_for_change := <str>$n_number_for_change,
				purge_date := <cal::local_date>$purge_date,
				unassigned := <bool>$unassigned,
			} UNLESS CONFLICT ON .n_number ELSE (
				UPDATE ns
				SET {
					registrant := <str>$registrant,
					street_address := <str>$street_address,
					street2 := <str>$street2,
					city := <str>$city,
					state := <str>$state,
					zip_code := <str>$zip_code,
					reserve_date := <cal::local_date>$reserve_date,
					type_reservation := <str>$type_reservation,
					expiration_date := <cal::local_date>$expiration_date,
					n_number_for_change := <str>$n_number_for_change,
					purge_date := <cal::local_date>$purge_date,
					unassigned := <bool>$unassigned,
				}
			)			
		`, &inserted,
			map[string]interface{}{
				"n_number":            strings.TrimSpace(record.NNumber),
				"registrant":          strings.TrimSpace(record.Registrant),
				"street_address":      strings.TrimSpace(record.StreetAddress),
				"street2":             strings.TrimSpace(record.Street2),
				"city":                strings.TrimSpace(record.City),
				"state":               strings.TrimSpace(record.State),
				"zip_code":            strings.TrimSpace(record.ZipCode),
				"reserve_date":        edgedb.NewLocalDate(record.ReserveDate.Date()),
				"type_reservation":    strings.TrimSpace(record.TypeReservation),
				"expiration_date":     edgedb.NewLocalDate(record.ExpirationDate.Date()),
				"n_number_for_change": strings.TrimSpace(record.NNumberForChange),
				"purge_date":          edgedb.NewLocalDate(record.PurgeDate.Date()),
				"unassigned":          record.Unassigned,
			})

		if err != nil {
			log.Fatal(err)
		}
	}
}
