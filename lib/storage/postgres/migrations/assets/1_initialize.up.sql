CREATE TABLE flats (
  id bigint not null primary key,
  street text,
  home smallint,
  structure smallint,
  flat_number smallint,
  state text,
  floor smallint,
  is_corner boolean,
  flat_type text,
  description text,
  owner text
);
