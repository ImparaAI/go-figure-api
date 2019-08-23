package database

var Schema string = `
CREATE TABLE IF NOT EXISTS drawings (
  id int(11) unsigned NOT NULL AUTO_INCREMENT,
  featured tinyint NOT NULL DEFAULT 0,
  originalPoints text NOT NULL,
  drawVectors text NOT NULL,
  createdAt datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  lastDrawVectorCalculatedAt datetime NULL DEFAULT NULL,
  PRIMARY KEY (id),
  KEY creation_time (createdAt),
  KEY featured (featured, createdAt)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
`
