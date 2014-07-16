#!/usr/bin/env bash
cd ./blog
go build
./blog --content="../content" --template="../template" --static="../static"
