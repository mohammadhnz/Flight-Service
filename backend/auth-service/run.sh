#!/bin/bash

sleep 3
echo "Running auth service..."

exec go run ./grpc_server/server.go