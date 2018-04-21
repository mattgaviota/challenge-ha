-- Destroy and Create table locations and copy the data from csvFile
DROP TABLE IF EXISTS locations;

CREATE TABLE locations(
  id INT PRIMARY KEY,
  lat FLOAT,
  lng FLOAT
);

\copy locations(id, lat, lng) FROM '/var/lib/postgresql/testdata/geoDataLarge.csv' DELIMITER ',' CSV HEADER;
