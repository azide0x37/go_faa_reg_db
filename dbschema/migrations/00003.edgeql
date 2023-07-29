CREATE MIGRATION m1k6q57blgf6ymapf36zxlrbufrkmeblfl2pehswcjcbf42r324wjq
    ONTO m1l6f4kk3efyas5yoq3fc6vuvictu7q2ukcjfrfdmxa4wtiuhb4jha
{
  ALTER TYPE default::NNumberRecord {
      ALTER PROPERTY n_number {
          CREATE CONSTRAINT std::exclusive;
      };
  };
};
