#!/bin/sh

FILE=./main
if [ ! -f "$FILE" ]; then
    echo "Compiling..."
    go build cmd/client/main.go
fi

./main
npm run build
# wrangler pages deploy build --commit-dirty=true
# scp -r ./build/* levi@x6c.us:/var/www/html/rlsz_dev/
