CREATE MIGRATION m1yra7jg6le7hbodhn4fwhd7qozz45ksqdeae6w6ruba4nmjv3ri7q
    ONTO m1wbecwwj2rkn47o56pyofyo5wcvwsasitozkcwn5mvxe6ylivul3q
{
  ALTER TYPE default::EngineModel {
      ALTER PROPERTY type {
          RENAME TO engine_type;
      };
  };
};
