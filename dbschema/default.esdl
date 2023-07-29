module default {
    type NNumberRecord {
        required property n_number -> str {
            constraint exclusive;
        };
        required property registrant -> str;
        property street_address -> str;
        property street2 -> str;
        property city -> str;
        property state -> str;
        property zip_code -> str;
        property reserve_date -> cal::local_date;
        property type_reservation -> str;
        property expiration_date -> cal::local_date;
        property n_number_for_change -> str;
        property purge_date -> cal::local_date;
        property unassigned -> bool;
    }

    type EngineManufacturer {
        required property code -> str {
            constraint max_len_value(3);
            constraint min_len_value(3);
        }
        required property name -> str {
            constraint max_len_value(10);
        }
        constraint exclusive on ((.code, .name));
        multi link models := .<manufacturer[IS EngineModel];
    }

    scalar type EngineType extending enum<None, Reciprocating, TurboProp, TurboShaft, TurboJet, TurboFan, Ramjet, TwoCycle, FourCycle, Unknown, Electric, Rotary>;

    type EngineModel {
        required property code -> str {
            constraint max_len_value(2);
            constraint min_len_value(2);
        }
        required property name -> str {
            constraint max_len_value(13);
        }
        required property engine_type -> EngineType;
        property horsepower -> int64 {
            constraint max_value(99999);
        }
        property thrust -> int64 {
            constraint max_value(999999);
        }
        property key {                                                        
            using ((.code, .name));                                       
            constraint exclusive;                                             
        };
        required link manufacturer -> EngineManufacturer;   
    }
}
