
echo "Building for macOS"
echo "-----Starting Build-----"
echo "Installing dependencies"
pnpm install
if [ $? -ne 0 ]; then
    echo "Failed to install dependencies"
    exit 1
fi

ARCH=$(rustc -Vv | grep host | cut -f2 -d' ')


echo "-----Building JTP -----"
cd ./JsonToPdf  || exit
go build -o "jtp-$ARCH" ./cmd/main.go
cd ..
mv "./JsonToPdf/jtp-$ARCH" "./src-tauri/assets/jtp-$ARCH"

echo "-----Building Tauri-----"
pnpm run tauri build

SIG=$(cat ./src-tauri/target/release/bundle/macos/cv-kawa.app.tar.gz.sig)
echo "$SIG"

