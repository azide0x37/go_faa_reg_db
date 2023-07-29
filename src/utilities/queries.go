package main

var engineModelInsertQuery = `
						WITH ns := (
							SELECT EngineManufacturer
							FILTER .name = <str>$manufacturer
							LIMIT 1
						)
		
						INSERT EngineModel {
							code := <str>$code,
							manufacturer := (SELECT ns),
							name := <str>$name,
							engine_type := <EngineType>$engine_type,
							horsepower := <int64>$horsepower,
							thrust := <int64>$thrust,
						}
				`

var engineManufacturerInsertQuery = `
				WITH ns := (
					SELECT EngineManufacturer
					FILTER .name = <str>$name
				)

				INSERT EngineManufacturer {
					name := <str>$name,
					code := <str>$code,
					}`

var engineManufacturerSelectQuery = `
			SELECT EXISTS (
				SELECT EngineManufacturer
				FILTER .name = <str>$manufacturer_name
			)
		`
