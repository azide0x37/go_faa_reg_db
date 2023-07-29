CREATE MIGRATION m1vd4myg3wli7o33n7657fcxjbekgsth3wre4y7upijgblivqzxyaa
    ONTO initial
{
  CREATE FUTURE nonrecursive_access_policies;
  CREATE TYPE default::NNumberRecord {
      CREATE PROPERTY city -> std::str;
      CREATE PROPERTY expiration_date -> cal::local_date;
      CREATE REQUIRED PROPERTY n_number -> std::str;
      CREATE PROPERTY n_number_for_change -> std::str;
      CREATE PROPERTY purge_date -> cal::local_date;
      CREATE PROPERTY registrant -> std::str;
      CREATE PROPERTY reserve_date -> cal::local_date;
      CREATE PROPERTY state -> std::str;
      CREATE PROPERTY street2 -> std::str;
      CREATE PROPERTY street_address -> std::str;
      CREATE PROPERTY type_reservation -> std::str;
      CREATE PROPERTY unassigned -> std::bool;
      CREATE PROPERTY zip_code -> std::str;
  };
};
