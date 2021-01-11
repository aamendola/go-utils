#!/bin/bash
git add .
git commit -m "."
git tag -a $1 -m "."
git push origin main
git push origin --tags
