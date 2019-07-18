BEGIN;
CREATE TABLE IF NOT EXISTS drawings (
  id INTEGER PRIMARY KEY,
  featured tinyint NOT NULL DEFAULT 0,
  originalPoints text NOT NULL,
  drawVectors text NULL DEFAULT NULL,
  calculatedDrawVectorCount int unsigned NOT NULL DEFAULT 0,
  createdAt datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  lastDrawVectorCalculatedAt datetime NULL DEFAULT NULL
);
CREATE INDEX IF NOT EXISTS creation_time ON drawings (createdAt);
CREATE INDEX IF NOT EXISTS featured ON drawings (featured, createdAt);

COMMIT;