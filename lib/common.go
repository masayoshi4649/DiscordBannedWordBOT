package lib

import (
	"log"
	"strings"

	"github.com/BurntSushi/toml"
)

// Config ... config.tml
type Config struct {
	Token string
}

// ReadConfig ... 設定ファイル読み込み
func ReadConfig(filename string) Config {
	var c Config
	_, err := toml.DecodeFile(filename, &c)
	if err != nil {
		log.Fatal(err)
	}
	return c
}

// UserIDEncode ... ID => <@!ID>
func UserIDEncode(ID string) string {
	return "<@!" + ID + ">"
}

// UserIDDecode ... <@!ID> => ID
func UserIDDecode(IDcode string) string {
	IDcode = strings.Replace(IDcode, "<@!", "", 1)
	IDcode = strings.Replace(IDcode, ">", "", 1)
	return IDcode
}

// RoleIDEncode ... ID => <@&ID>
func RoleIDEncode(ID string) string {
	return "<@&" + ID + ">"
}

// RoleIDDecode ... <@&ID> => ID
func RoleIDDecode(IDcode string) string {
	IDcode = strings.Replace(IDcode, "<@&", "", 1)
	IDcode = strings.Replace(IDcode, ">", "", 1)
	return IDcode
}
