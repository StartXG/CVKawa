if [ "$1" = "linux" ]; then
    GOOS=linux GOARCH=amd64 go build -o ./jtp cmd/main.go
elif [ "$1" = "linux-arm" ]; then
    GOOS=linux GOARCH=arm64 go build -o ./jtp cmd/main.go
elif [ "$1" = "windows" ]; then
    GOOS=windows GOARCH=amd64 go build -o ./jtp.exe cmd/main.go
elif [ "$1" = "mac" ]; then
    GOOS=darwin GOARCH=amd64 go build -o ./jtp cmd/main.go
elif [ "$1" = "mac-m" ]; then
    GOOS=darwin GOARCH=arm64 go build -o ./jtp-aarch64-apple-darwin cmd/main.go
else
    echo "Invalid platform! Please use 'linux', 'windows' or 'mac'."
fi
