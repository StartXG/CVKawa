cd ./JsonToPdf  || exit
sh Build.sh $1
cd ..
cp ./JsonToPdf/jtp* ./src-tauri/assets/
pnpm run taurib