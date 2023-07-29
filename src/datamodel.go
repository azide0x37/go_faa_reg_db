package main

import "time"

// NNumberRecord represents a record from the FAA's N-Number Registry
type NNumberRecord struct {
	NNumber          string
	Registrant       string
	StreetAddress    string
	Street2          string
	City             string
	State            string
	ZipCode          string
	ReserveDate      time.Time
	TypeReservation  string
	ExpirationDate   time.Time
	NNumberForChange string
	PurgeDate        time.Time
	Unassigned       bool
}

// EngineManufacturer represents an engine manufacturer
type EngineManufacturer struct {
	name string
	code string
}

// EngineType represents an engine type
type EngineType int

const (
	None EngineType = iota
	Reciprocating
	TurboProp
	TurboShaft
	TurboJet
	TurboFan
	Ramjet
	TwoCycle
	FourCycle
	Unknown
	Electric
	Rotary
)

// String method to return the string representation of the EngineType
func (e EngineType) String() string {
	names := [...]string{
		"None",
		"Reciprocating",
		"TurboProp",
		"TurboShaft",
		"TurboJet",
		"TurboFan",
		"Ramjet",
		"TwoCycle",
		"FourCycle",
		"Unknown",
		"Electric",
		"Rotary",
	}

	if e < None || e > Rotary {
		return "Unknown"
	}

	return names[e]
}

// EngineModel represents an engine model
type EngineModel struct {
	code         string
	manufacturer EngineManufacturer
	name         string
	engine_type  EngineType
	horsepower   int64
	thrust       int64
}
