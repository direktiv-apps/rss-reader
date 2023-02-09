#!/bin/sh

docker build -t rss-reader . && docker run -p 9191:8080 rss-reader