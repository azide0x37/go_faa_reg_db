CREATE MIGRATION m1tqy2h4rbbe7sxxrxkuwbjpvwkgvcr3kuw3yem2t2ht6qmkobinka
    ONTO m1yra7jg6le7hbodhn4fwhd7qozz45ksqdeae6w6ruba4nmjv3ri7q
{
  ALTER TYPE default::EngineModel {
      ALTER PROPERTY name {
          CREATE CONSTRAINT std::exclusive;
      };
  };
};
