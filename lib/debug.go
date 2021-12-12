package lib

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

// Debug ... デバッグ
func Debug(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.Bot == true {
		return
	}

	fmt.Println("GuildID=>", m.GuildID)
	fmt.Println("ChannelID=>", m.ChannelID)
	fmt.Println("Author=>", m.Author)
	fmt.Println("Roles=>", m.Member.Roles)
	fmt.Println("ID=>", m.Author.ID)
	fmt.Println("MentionRoles=>", append(m.MentionRoles))
	fmt.Println("MentionChannels=>", append(m.MentionChannels))

	guild, _ := s.Guild(m.GuildID)
	// 所有者
	fmt.Println("OwnerID=>", guild.OwnerID)

	fmt.Println(append(m.Member.Roles))

	fmt.Print("GuildRoles=>")
	fmt.Println(s.GuildRoles(m.GuildID))

}
