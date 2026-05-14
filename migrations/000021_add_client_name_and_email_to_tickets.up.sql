ALTER TABLE tickets
ADD COLUMN client_name VARCHAR(255) DEFAULT '' AFTER client_dni,
ADD COLUMN client_email VARCHAR(255) DEFAULT '' AFTER client_name;
