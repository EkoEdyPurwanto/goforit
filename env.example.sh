#!/bin/bash

# Database Configuration
export APP_DB_DRIVER=postgres
export APP_DB_HOST=localhost
export APP_DB_PORT=5432
export APP_DB_USER=eepsql
export APP_DB_PASSWORD=1903
export APP_DB_NAME=goforit
export APP_SSL_MODE=disable
export APP_TIME_ZONE=Asia/Jakarta

# Api Configuration
export APP_API_HOST=localhost
export APP_API_PORT=8080

# File Log Configuration
export APP_FILE_PATH=logger.log

# Token Configuration
export APP_ISSUER=GoForIt
export APP_SIGNATURE_KEY=verySecret
export APP_EXPIRES_IN_MINUTES=60