package utils

import (
	"github.com/Succo/emoji"
	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/snowflake/v2"
	"math/rand"
	"regexp"
)

var decodeDiscordEmojiSource = regexp.MustCompile("^<(a?):(.+?):(\\d{18,19})>")

func DecodeDiscordEmoji(a string) (string, bool, int) {
	decode := decodeDiscordEmojiSource.FindString(a)
	if decode == "" {
		return emoji.DecodeString(a)
	}
	return decode, true, len(decode)
}

func DecodeAllDiscordEmoji(a string) (emojiStr []string) {
	for {
		g, ok, n := DecodeDiscordEmoji(a)
		if n == 0 {
			break
		}
		if ok {
			emojiStr = append(emojiStr, g)
		}
		a = a[n:]
	}
	return
}

func ConvertToComponentEmoji(a string) discord.ComponentEmoji {
	sub := decodeDiscordEmojiSource.FindStringSubmatch(a)
	if len(sub) == 4 {
		if sId, err := snowflake.Parse(sub[3]); err == nil {
			return discord.ComponentEmoji{
				ID:       sId,
				Name:     sub[2],
				Animated: sub[1] == "a",
			}
		}
	}
	decode, ok, _ := emoji.DecodeString(a)
	if ok {
		return discord.ComponentEmoji{Name: decode}
	}
	return discord.ComponentEmoji{}
}

func RandomEmoji(a string) string {
	emojis := DecodeAllDiscordEmoji(a)
	if len(emojis) <= 0 {
		return ""
	}
	n := rand.Intn(len(emojis))
	return emojis[n]
}
