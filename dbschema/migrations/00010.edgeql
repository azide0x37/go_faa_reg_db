CREATE MIGRATION m1ovp5mkaiy2erfd7c36tuoxilcydlti3ytctenfygqy4haadjsdea
    ONTO m1litw644ay6q3jsi5knkp5k4fyw2ct53rzrda3qon2akp7hbxpfha
{
  ALTER TYPE default::EngineManufacturer {
      CREATE INDEX ON ((.code, .name));
  };
  ALTER TYPE default::EngineModel {
      CREATE INDEX ON ((.code, .name));
  };
};
