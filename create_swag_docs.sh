#!/bin/bash

# Ejecutar swag init
# Se asume que 'swag' está instalado en el PATH
swag init -g cmd/api/main.go -o docs

# Verificar si el comando fue exitoso
if [ $? -eq 0 ]; then
    echo "Swagger documentation generated successfully"
    
    # Ruta al archivo docs.go
    DOCS_FILE="docs/docs.go"
    
    # Verificar si el archivo existe
    if [ -f "$DOCS_FILE" ]; then
        echo "Cleaning up docs.go..."
        
        # Eliminar las líneas LeftDelim y RightDelim que a veces causan problemas en Echo
        # Usamos sed para una edición simple
        sed -i '/LeftDelim:.*"{{",/d' "$DOCS_FILE"
        sed -i '/RightDelim:.*"}}",/d' "$DOCS_FILE"
        
        echo "docs.go cleaned successfully"
    else
        echo "Warning: docs.go not found at $DOCS_FILE"
    fi
else
    echo "Error: swag init failed"
    exit 1
fi
