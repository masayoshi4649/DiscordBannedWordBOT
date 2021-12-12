package lib

import "github.com/bwmarrin/discordgo"

// MessageSend ... メッセージ応答
func MessageSend(s *discordgo.Session, m *discordgo.MessageCreate, content string) {
	s.ChannelMessageSend(m.ChannelID, content)
}

// MessageSendWithMention ... メッセージ応答(メンション付き)
func MessageSendWithMention(s *discordgo.Session, m *discordgo.MessageCreate, content string) {
	s.ChannelMessageSend(m.ChannelID, m.Author.Mention()+content)
}

// MessageReply ... メッセージ返信
func MessageReply(s *discordgo.Session, content string, m *discordgo.MessageCreate) {
	s.ChannelMessageSendReply(m.ChannelID, content, m.Reference())
}

// MessageDelete ... メッセージ削除
func MessageDelete(s *discordgo.Session, m *discordgo.MessageCreate) {
	s.ChannelMessageDelete(m.ChannelID, m.Reference().MessageID)
}
