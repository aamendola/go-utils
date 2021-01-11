#!/bin/bash
git add .
git commit -m "."
git tag -a $1 -m "."
git push origin feature/start-up 
git push origin --tags
