#!/bin/zsh

go build cmd/main.go && ./main -f ./assets/cv_data.json -o ./cv.pdf -l jp