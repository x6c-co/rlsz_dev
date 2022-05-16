#!/bin/sh

go run cmd/client/main.go
npx parcel build ./web/scss/app.scss --dist-dir ./build/assets/css
npx wrangler pages publish build --commit-dirty=true
