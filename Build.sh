
# Install frontend
pnpm install

ARCH=$(rustc -Vv | grep host | cut -f2 -d' ')


# build jtp
cd ./JsonToPdf  || exit
go build -o "jtp-${ARCH}" cmd/main.go

cp "./JsonToPdf/jtp-${ARCH}" ./src-tauri/assets/
pnpm run taurib

