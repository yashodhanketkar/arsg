PRAGMA foreign_keys = ON;

CREATE TABLE rating (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  name TEXT NOT NULL,
  art REAL NOT NULL,
  support REAL NOT NULL,
  plot REAL NOT NULL,
  bias REAL NOT NULL,
  rating TEXT NOT NULL,
  comments TEXT
);

CREATE TABLE IF NOT EXISTS anime (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  rating_id integer NOT NULL,
  FOREIGN KEY (rating_id) REFERENCES rating (id)
);

CREATE TABLE IF NOT EXISTS manga (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  rating_id integer NOT NULL,
  FOREIGN KEY (rating_id) REFERENCES rating (id)
);

CREATE TABLE IF NOT EXISTS lightnovel (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  rating_id integer NOT NULL,
  FOREIGN KEY (rating_id) REFERENCES rating (id)
);
