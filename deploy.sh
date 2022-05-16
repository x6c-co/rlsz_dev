#!/bin/sh

go run cmd/client/main.go
npm run build
npx wrangler pages publish build --commit-dirty=true
