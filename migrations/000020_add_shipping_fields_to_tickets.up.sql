ALTER TABLE tickets 
ADD COLUMN shipping_address VARCHAR(255) NULL,
ADD COLUMN shipping_zip_code VARCHAR(20) NULL,
ADD COLUMN shipping_phone VARCHAR(50) NULL,
ADD COLUMN shipping_message TEXT NULL;
