CREATE MIGRATION m1l6f4kk3efyas5yoq3fc6vuvictu7q2ukcjfrfdmxa4wtiuhb4jha
    ONTO m1vd4myg3wli7o33n7657fcxjbekgsth3wre4y7upijgblivqzxyaa
{
  ALTER TYPE default::NNumberRecord {
      ALTER PROPERTY registrant {
          SET REQUIRED USING ('Unknown');
      };
  };
};
