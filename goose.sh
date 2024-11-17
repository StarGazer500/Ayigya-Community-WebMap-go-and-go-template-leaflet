#!/bin/bash



export GOOSE_DRIVER=postgres
export GOOSE_DBSTRING=postgresql://goose:password@127.0.0.1:8092/go_migrations?sslmode=disable