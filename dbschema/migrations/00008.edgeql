CREATE MIGRATION m13bnpp2gdzxcbu6szqpwl3mhd3gvi7jhdhkanvtl37bv6chrbit4a
    ONTO m1boh4n2tfvhivxbhcuwwyuvpaowl7bhytpbwfbntqp7mshmolb62q
{
  ALTER TYPE default::EngineManufacturer {
      ALTER PROPERTY code {
          DROP CONSTRAINT std::exclusive;
      };
  };
};
