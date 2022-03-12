CREATE TABLE item (
		id INT PRIMARY KEY AUTO_INCREMENT NOT NULL,
		title TEXT NOT NULL,
		created_at DATETIME DEFAULT current_timestamp,
		updated_at DATETIME DEFAULT current_timestamp on update current_timestamp,
		url TEXT,
		memo TEXT,
		tag TEXT
)