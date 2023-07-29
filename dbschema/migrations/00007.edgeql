CREATE MIGRATION m1boh4n2tfvhivxbhcuwwyuvpaowl7bhytpbwfbntqp7mshmolb62q
    ONTO m1tqy2h4rbbe7sxxrxkuwbjpvwkgvcr3kuw3yem2t2ht6qmkobinka
{
  ALTER TYPE default::EngineModel {
      ALTER PROPERTY code {
          DROP CONSTRAINT std::exclusive;
      };
  };
};
