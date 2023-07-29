CREATE MIGRATION m1g6rjwta2nyuucg4q5funi4rz7ifsopxgxyua2uaf4iox6p7fljnq
    ONTO m1w7dczpyoesmtxzwxd6csxhgi5enmlqhlnaearpqzmkmd7r7jcsfq
{
  ALTER TYPE default::EngineManufacturer {
      ALTER LINK models {
          USING (.<manufacturer[IS default::EngineModel]);
      };
  };
};
