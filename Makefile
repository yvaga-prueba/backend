.PHONY: help run build test fmt lint sec check tidy clean swagger migrate-up migrate-down migrate-status migrate-new tree consolidate

# Detectar sistema operativo
ifeq ($(OS),Windows_NT)
    DETECTED_OS := Windows
else
    DETECTED_OS := $(shell uname -s)
endif

# Variables
APP_NAME := core
CMD_API := ./cmd/api
CMD_MIGRATE := ./cmd/migrate
DOCS_DIR := ./internal/docs
DOCS_FILE := $(DOCS_DIR)/docs.go
BIN_DIR := bin

help: ## Mostrar ayuda
	@echo "Comandos disponibles:"
	@echo "  make run              - Ejecutar la API en modo desarrollo (requiere air)"
	@echo "  make build            - Compilar binarios en bin/"
	@echo "  make test             - Ejecutar tests con cobertura"
	@echo "  make fmt              - Verificar formato (go fmt)"
	@echo "  make lint             - Ejecutar go vet + staticcheck"
	@echo "  make sec              - Ejecutar análisis de seguridad (gosec)"
	@echo "  make check            - fmt + lint + sec + test (pre-push)"
	@echo "  make tidy             - Limpiar y verificar go.mod"
	@echo "  make swagger          - Generar documentación Swagger"
	@echo "  make migrate-up       - Aplicar migraciones pendientes"
	@echo "  make migrate-down     - Revertir última migración"
	@echo "  make migrate-status   - Ver estado de migraciones"
	@echo "  make migrate-new      - Crear nueva migración (name=nombre)"
	@echo "  make tree             - Mostrar estructura del proyecto"
	@echo "  make gdrive-auth      - Obtener nuevo Refresh Token para Google Drive"
	@echo "  make clean            - Limpiar archivos generados"

run:
	@echo "Starting API server with AIR..."
	air

build:
	@echo "Building binarios..."
	@mkdir -p $(BIN_DIR)
	go build -ldflags="-w -s" -o $(BIN_DIR)/api ./cmd/api
	go build -ldflags="-w -s" -o $(BIN_DIR)/migrate ./cmd/migrate
	@echo "Build complete: $(BIN_DIR)/"

test:
	@echo "Running tests..."
	go test ./... -cover -coverprofile=coverage.out
	go tool cover -html=coverage.out -o coverage.html
	@echo "Coverage report: coverage.html"

fmt:
	@echo "Checking format..."
ifeq ($(DETECTED_OS),Windows)
	@powershell -Command "$$out = (go fmt ./...); if ($$out) { Write-Host 'ERROR: archivos sin formatear:'; Write-Host $$out; exit 1 }"
else
	@test -z "$$(go fmt ./...)" || (echo "ERROR: archivos sin formatear, ejecuta 'go fmt ./...'" && exit 1)
endif
	@echo "Format OK"

lint:
	@echo "Running go vet..."
	go vet ./...
	@echo "Running staticcheck..."
ifeq ($(DETECTED_OS),Windows)
	@where staticcheck >NUL 2>&1 || go install honnef.co/go/tools/cmd/staticcheck@latest
else
	@which staticcheck > /dev/null 2>&1 || go install honnef.co/go/tools/cmd/staticcheck@latest
endif
	staticcheck ./...
	@echo "Lint OK"

sec:
	@echo "Running gosec..."
ifeq ($(DETECTED_OS),Windows)
	@where gosec >NUL 2>&1 || go install github.com/securego/gosec/v2/cmd/gosec@latest
else
	@which gosec > /dev/null 2>&1 || go install github.com/securego/gosec/v2/cmd/gosec@latest
endif
	gosec ./...
	@echo "Security OK"

check: fmt lint sec test
	@echo "All checks passed. Listo para subir."

tidy:
	go mod tidy
	go mod verify

swagger:
	@echo "Generating Swagger documentation..."
ifeq ($(DETECTED_OS),Windows)
	@powershell -ExecutionPolicy Bypass -File create_swag_docs.ps1
else
	@chmod +x create_swag_docs.sh
	@./create_swag_docs.sh
endif
	@echo "Swagger documentation ready at /swagger/index.html"

migrate-up:
	@echo "Running migrations up..."
	go run $(CMD_MIGRATE) up

migrate-down:
	@echo "Running migrations down..."
	go run $(CMD_MIGRATE) down

migrate-status:
	@echo "Checking migration status..."
	go run $(CMD_MIGRATE) status

migrate-new:
ifndef name
	@echo "Error: Debes especificar name=nombre_migracion"
	@echo "Ejemplo: make migrate-new name=add_roles"
	@exit 1
endif
ifeq ($(DETECTED_OS),Windows)
	@powershell -Command "$$timestamp = Get-Date -Format 'yyyyMMddHHmmss'; \
		$$version = \"$${timestamp}_$(name)\"; \
		$$upfile = \"cmd\\migrate\\migrations\\$${version}.up.sql\"; \
		$$downfile = \"cmd\\migrate\\migrations\\$${version}.down.sql\"; \
		New-Item -ItemType Directory -Path 'cmd\\migrate\\migrations' -Force | Out-Null; \
		New-Item -ItemType File -Path $$upfile -Force | Out-Null; \
		New-Item -ItemType File -Path $$downfile -Force | Out-Null; \
		Write-Host 'Created migration files:' -ForegroundColor Green; \
		Write-Host \"  $$upfile\"; \
		Write-Host \"  $$downfile\""
else
	@timestamp=$$(date +%Y%m%d%H%M%S); \
	version="$${timestamp}_$(name)"; \
	upfile="migrations/$${version}.up.sql"; \
	downfile="migrations/$${version}.down.sql"; \
	mkdir -p migrations; \
	touch "$$upfile" "$$downfile"; \
	echo "Created migration files:"; \
	echo "  $$upfile"; \
	echo "  $$downfile"
endif

tree:
	@go run ./internal/pkg/tree.go

gdrive-auth: ## Ejecutar asistente de autenticación para Google Drive
	@go run ./cmd/gdrive_auth/main.go

consolidate:
	@go run ./internal/pkg/consolidate.go

clean:
	@echo "Cleaning..."
	@rm -rf $(BIN_DIR) coverage.out coverage.html
	@echo "Clean complete"

.DEFAULT_GOAL := help