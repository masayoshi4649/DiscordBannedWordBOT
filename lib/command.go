package lib

import (
	"fmt"
	"log"
	"strings"

	"github.com/bwmarrin/discordgo"
)

// ParseCommand ... @メンション 命令 命令対象 文字列やロール
// 返り値はコマンドか否か
func ParseCommand(s *discordgo.Session, m *discordgo.MessageCreate) bool {
	// 管理者ではない場合無視(将来的にロール付与も)
	guild, err := s.Guild(m.GuildID)
	if err != nil {
		log.Fatal(err)
	}

	if m.Author.ID != guild.OwnerID {
		return false
	}

	botID := s.State.User.ID
	cmd := strings.Fields(m.Content)

	// 1つめのコマンドがメンションでない場合は無視
	if UserIDEncode(botID) != cmd[0] {
		return false
	}

	if len(cmd) > 1 {
		switch strings.ToUpper(cmd[1]) {
		case "ADD":
			fmt.Println("ADD")

		case "DELETE":
			fmt.Println("DELETE")

		case "LIST":
			fmt.Println("LIST")

		case "HELP":
			fmt.Println("HELP")
			return false

		default:
			fmt.Println("default")
			return false
		}
	} else {
		fmt.Println("aug0 HELP")
		return false
	}

	// DEV
	return true
}
