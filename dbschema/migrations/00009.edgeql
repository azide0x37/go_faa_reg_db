CREATE MIGRATION m1litw644ay6q3jsi5knkp5k4fyw2ct53rzrda3qon2akp7hbxpfha
    ONTO m13bnpp2gdzxcbu6szqpwl3mhd3gvi7jhdhkanvtl37bv6chrbit4a
{
  ALTER TYPE default::EngineModel {
      ALTER PROPERTY name {
          DROP CONSTRAINT std::exclusive;
      };
  };
};
