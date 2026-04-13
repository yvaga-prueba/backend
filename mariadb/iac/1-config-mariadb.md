1. Configurar el Bind Address de MariaDB
Por defecto, MariaDB suele escuchar solo en 127.0.0.1 (localhost). Desde el punto de vista de un Pod, el "localhost" es el propio contenedor, no tu servidor físico.

    1. Edita el archivo de configuración (normalmente en /etc/mysql/mariadb.conf.d/50-server.cnf o /etc/mysql/my.cnf).

    2. Busca la línea bind-address.

    3. Cámbiala a 0.0.0.0 para que escuche en todas las interfaces, o a la IP de la interfaz de red de K3s (usualmente la de la interfaz cni0 o la IP privada del nodo).

    Ini, TOML
    `bind-address = 0.0.0.0`

    4. Reinicia el servicio: sudo systemctl restart mariadb.