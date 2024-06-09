# CVKawa core module

This module will be rewritten to Rust in the future, and the new version will be embedded to CVKawa.

## For Use

```shell
jtp-aarch64-apple-darwin -f ./test.json -l cn
jtp-aarch64-apple-darwin --file ./test.json --lang cn
```

## For Development

```shell
git clone <the project>
cd JsonToPdf
go mod tidy
go run cmd/main.go
./main -f <json file path>
```
## About json

you can use the file which is in the `assets` folder as a template, and you can add your own json file to the `assets` folder, then you can use the command to generate the pdf file.

# important

Don't delete these files

- `assets/`: For Demo, you can add your own json file to this folder.
- `lang/`: For language translation, you can add your own language file to this folder.
- `JTPAction/fontDir/`: For font, you can add your own font file to this folder.