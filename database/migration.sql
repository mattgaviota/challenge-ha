-- Destroy and Create table locations and copy the data from csvFile
DROP TABLE IF EXISTS locations;

CREATE TABLE locations(
  id INT PRIMARY KEY,
  lat FLOAT,
  lng FLOAT
);

\copy locations(id, lat, lng) FROM './files/geoData.csv' DELIMITER ',' CSV HEADER;
