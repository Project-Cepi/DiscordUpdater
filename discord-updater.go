package main

import (
	"os"
	"strings"

	"github.com/andersfylling/disgord"
)

func main() {
	args := os.Args[1:]
	repo_name := strings.TrimPrefix(args[0], "Project-Cepi/")
	commit_hash := args[1][0:9]

	discord := disgord.New(disgord.Config{
		BotToken: os.Getenv("DISCORD_KEY"),
	})

	discord.Gateway().Connect()

	messageBuilder := discord.Channel(872557055934890095).Message(872557344029020190)
    message, err := messageBuilder.Get()
    if err != nil { 
        panic(err) 
    }


    embeds := message.Embeds

    field := new(disgord.EmbedField)
    field.Name = repo_name
    field.Value = "Latest commit: `" + commit_hash + "`"

    var fields []*disgord.EmbedField
    if len(message.Embeds) > 0 {
        oldEmbed := embeds[0]
        fields = oldEmbed.Fields

        var i int
        for range fields {
            if fields[i].Name == repo_name { 
                break
            }
            i++
        }

        if i < len(fields) {
            fields[i] = field
        } else {
            fields = append(fields, field)
        }
    } else {
        fields = append(fields, field)
    }

	newEmbed := new(disgord.Embed)
    newEmbed.Fields = append(newEmbed.Fields, fields...)
    newEmbed.Color = 0xFF0000

    messageBuilder.UpdateBuilder().SetEmbed(newEmbed).Execute()
}