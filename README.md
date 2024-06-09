# CVKawa

A simple, localized resume creation and management tool that helps job seekers highlight their strengths and simplifies resume writing

# To Developer

## base env

- Go: 1.22+
- Rustc: 1.78+

# Build by yourself

> Eg with M1 mac

```shell
cd  <project path>/JsonToPdf
go build -o ./jtp-aarch64-apple-darwin cmd/main.go
cp jtp-aarch64-apple-darwin <project path>/src-tauri/jtp-aarch64-apple-darwin
cd  <project path>
pnpm/yarn run taurib
```