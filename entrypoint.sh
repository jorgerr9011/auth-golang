#!/bin/bash
set -e

echo "Ejecutando migraciones..."
just migrate-production

echo "Iniciando la aplicación..."
exec ./auth-app
