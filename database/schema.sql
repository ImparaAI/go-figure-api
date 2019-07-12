BEGIN;
CREATE TABLE IF NOT EXISTS submissions (
  id INTEGER PRIMARY KEY,
  requestedDrawVectorCount int unsigned NOT NULL,
  featured tinyint NOT NULL DEFAULT 0,
  originalPoints text NOT NULL,
  drawVectors text NULL DEFAULT NULL,
  calculatedDrawVectorCount int unsigned NOT NULL DEFAULT 0,
  createdAt datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  lastDrawVectorCalculatedAt datetime NULL DEFAULT NULL
);
CREATE INDEX IF NOT EXISTS creation_time ON submissions (createdAt);
CREATE INDEX IF NOT EXISTS featured ON submissions (featured, createdAt);

COMMIT;