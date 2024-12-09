#!/usr/bin/env bash

echo "Getting started to compile spec"

# Bundle docs into zero-dependency HTML file
npx redoc-cli bundle docs/spec.yaml && \
mv redoc-static.html docs/spec.html && \
echo "Changed name from redoc-static.html to spec.html and moved to docs folder" && \
echo -e "\nDone!"
