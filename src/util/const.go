package util

const (
	WantOneGotMany = "want 1, got %d"
	WantQGotQ      = "want %q, got %q"
	WantFGotF      = "want %f, got %f"
	WantVGotV      = "want %v, got %v"
	CtrlC          = "ctrl+c"

	MockDBShema = `
CREATE TABLE IF NOT EXISTS rating (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	name TEXT,
	art REAL,
	support REAL,
	plot REAL,
	bias REAL,
	rating TEXT,
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
	`
)

var (
	DefaultParams = []string{"Art/Animation", "Character/Cast", "Plot", "Bias"}
)
