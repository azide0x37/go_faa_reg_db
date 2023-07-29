CREATE MIGRATION m1w7dczpyoesmtxzwxd6csxhgi5enmlqhlnaearpqzmkmd7r7jcsfq
    ONTO m1ovp5mkaiy2erfd7c36tuoxilcydlti3ytctenfygqy4haadjsdea
{
  ALTER TYPE default::EngineManufacturer {
      CREATE CONSTRAINT std::exclusive ON ((.code, .name));
      DROP INDEX ON ((.code, .name));
  };
  ALTER TYPE default::EngineModel {
      CREATE CONSTRAINT std::exclusive ON ((.code, .name));
      DROP INDEX ON ((.code, .name));
  };
};
