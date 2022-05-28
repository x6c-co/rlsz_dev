#!/bin/sh

FILE=./main
if [ ! -f "$FILE" ]; then
    echo "Compiling..."
    go build cmd/client/main.go
fi

./main
npm run build
npx wrangler pages publish build --commit-dirty=true
