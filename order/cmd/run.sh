#!/bin/bash

# Be sure you have a running db (../scripts/db-up.sh)

DATA_SOURCE_URL='root:secret-password@tcp(127.0.0.1:3306)/order' \
APPLICATION_PORT=3000 \
ENV=development \
go run main.go
