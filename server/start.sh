#!/bin/bash

cp -r ../network/crypto-config/ ./
mv crypto-config organizations

go mod init application
go mod tidy

go run main.go
