CREATE MIGRATION m1wbecwwj2rkn47o56pyofyo5wcvwsasitozkcwn5mvxe6ylivul3q
    ONTO m1k6q57blgf6ymapf36zxlrbufrkmeblfl2pehswcjcbf42r324wjq
{
  CREATE TYPE default::EngineManufacturer {
      CREATE REQUIRED PROPERTY code -> std::str {
          CREATE CONSTRAINT std::exclusive;
          CREATE CONSTRAINT std::max_len_value(3);
          CREATE CONSTRAINT std::min_len_value(3);
      };
      CREATE REQUIRED PROPERTY name -> std::str {
          CREATE CONSTRAINT std::max_len_value(10);
      };
  };
  CREATE SCALAR TYPE default::EngineType EXTENDING enum<None, Reciprocating, TurboProp, TurboShaft, TurboJet, TurboFan, Ramjet, TwoCycle, FourCycle, Unknown, Electric, Rotary>;
  CREATE TYPE default::EngineModel {
      CREATE REQUIRED LINK manufacturer -> default::EngineManufacturer;
      CREATE REQUIRED PROPERTY code -> std::str {
          CREATE CONSTRAINT std::exclusive;
          CREATE CONSTRAINT std::max_len_value(2);
          CREATE CONSTRAINT std::min_len_value(2);
      };
      CREATE PROPERTY horsepower -> std::int64 {
          CREATE CONSTRAINT std::max_value(99999);
      };
      CREATE REQUIRED PROPERTY name -> std::str {
          CREATE CONSTRAINT std::max_len_value(13);
      };
      CREATE PROPERTY thrust -> std::int64 {
          CREATE CONSTRAINT std::max_value(999999);
      };
      CREATE REQUIRED PROPERTY type -> default::EngineType;
  };
  ALTER TYPE default::EngineManufacturer {
      CREATE MULTI LINK models -> default::EngineModel;
  };
};
