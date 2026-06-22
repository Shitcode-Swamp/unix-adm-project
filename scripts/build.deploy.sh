#!/bin/sh
CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o /out/unix-adm-backend ./source/app
