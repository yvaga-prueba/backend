CREATE DATABASE core_t;

-- 1. Crear el usuario con acceso global ('%') y una contraseña segura
CREATE USER 'usr_yvaga_t'@'%' IDENTIFIED BY 'TuPasswordSeguro123!';

-- 2. Otorgar privilegios específicos sobre la base de datos core_t
-- Puedes ajustar 'ALL PRIVILEGES' a permisos específicos (SELECT, INSERT, etc.) según sea necesario
GRANT ALL PRIVILEGES ON core_t.* TO 'usr_yvaga_t'@'%';

-- 3. Refrescar la tabla de privilegios
FLUSH PRIVILEGES;