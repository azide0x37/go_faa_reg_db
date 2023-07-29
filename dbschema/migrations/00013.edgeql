CREATE MIGRATION m1vilurtpkw47w6qxgvohaqzo3zpjxrz5fo7fluhoyfo25pbbtucwa
    ONTO m1g6rjwta2nyuucg4q5funi4rz7ifsopxgxyua2uaf4iox6p7fljnq
{
  ALTER TYPE default::EngineModel {
      DROP CONSTRAINT std::exclusive ON ((.code, .name));
  };
  ALTER TYPE default::EngineModel {
      CREATE PROPERTY key {
          USING ((.code, .name));
          CREATE CONSTRAINT std::exclusive;
      };
  };
};
