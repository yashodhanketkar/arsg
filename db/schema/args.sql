CREATE TABLE IF NOT EXISTS rating (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  name TEXT NOT NULL,
  art REAL NOT NULL,
  support REAL NOT NULL,
  plot REAL NOT NULL,
  bias REAL NOT NULL,
  rating TEXT NOT NULL,
  comments TEXT
)
