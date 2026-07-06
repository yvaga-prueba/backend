CREATE TABLE IF NOT EXISTS contact_messages (
    id INT AUTO_INCREMENT PRIMARY KEY,
    user_id INT NULL, -- Puede ser null si no está logueado
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    order_number VARCHAR(50) NULL,
    message TEXT NOT NULL,
    status ENUM('Pendiente', 'En revisión', 'Resuelto') DEFAULT 'Pendiente',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_contact_user FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE SET NULL
);