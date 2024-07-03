#!/bin/zsh

go build cmd/main.go && ./main -f ./assets/cv_data.json -o ~/Desktop/cv.pdf  -l zh-CN