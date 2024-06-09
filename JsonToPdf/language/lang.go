package language

import (
	"embed"
	_ "embed"
	"github.com/BurntSushi/toml"
)

var LangStruct map[string]map[string]string

//go:embed lang.toml
var LangFs embed.FS

func GetLangFile(env string, key string) string {
	tomlFile, err := LangFs.ReadFile("lang.toml")
	if err != nil {
		panic(err)
	}
	err = toml.Unmarshal(tomlFile, &LangStruct)
	if err != nil {
		return ""
	}
	return LangStruct[env][key]
}
