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

## Important

1. When you build the app for msvc[windows brach] you need the lib file for sqlite, in the project we support the file, but it build for `x64` system
2. On mac or linux build windows,please read the [tauri doc](https://tauri.app/v1/guides/building/cross-platform/#%E5%A6%82%E4%BD%95%E8%A7%A6%E5%8F%91) first.