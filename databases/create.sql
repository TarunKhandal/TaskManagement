-- Create the Authors table
CREATE TABLE IF NOT EXISTS Authors (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  first_name TEXT NOT NULL,
  last_name TEXT NOT NULL,
  username TEXT NOT NULL UNIQUE,
  added_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Create trigger to update updated_at on Authors table
CREATE TRIGGER update_authors_updated_at
AFTER UPDATE ON Authors
FOR EACH ROW
BEGIN
  UPDATE Authors SET updated_at = CURRENT_TIMESTAMP WHERE id = NEW.id;
END;

-- Create the Tasks table
CREATE TABLE IF NOT EXISTS Tasks (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  name TEXT NOT NULL,
  description TEXT,
  author_id INTEGER NOT NULL,
  added_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  FOREIGN KEY (author_id) REFERENCES Authors(id) ON DELETE CASCADE
);

-- Create trigger to update updated_at on Tasks table
CREATE TRIGGER update_tasks_updated_at
AFTER UPDATE ON Tasks
FOR EACH ROW
BEGIN
  UPDATE Tasks SET updated_at = CURRENT_TIMESTAMP WHERE id = NEW.id;
END;

-- Create the Steps table
CREATE TABLE IF NOT EXISTS Steps (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  task_id INTEGER NOT NULL,
  step TEXT NOT NULL,
  completed  NOT NULL DEFAULT 0 CHECK(completed IN (0, 1)),
  added_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  FOREIGN KEY (task_id) REFERENCES Tasks(id) ON DELETE CASCADE
);

-- Create trigger to update updated_at on Steps table
CREATE TRIGGER update_steps_updated_at
AFTER UPDATE ON Steps
FOR EACH ROW
BEGIN
  UPDATE Steps SET updated_at = CURRENT_TIMESTAMP WHERE id = NEW.id;
END;

