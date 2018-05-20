package bot

import (
	"fmt"
	"os"
	"os/signal"
	"regexp"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/mgerb/spam-filter-bot/config"
)

const urlRegexp = `https?:\/\/(www\.)?[-a-zA-Z0-9@:%._\+~#=]{2,256}\.[a-z]{2,6}\b([-a-zA-Z0-9@:%_\+.~#?&//=]*)`

func Init() {

	// Create a new Discord session using the provided bot token.
	dg, err := discordgo.New("Bot " + config.Config.Token)
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	// Register the messageCreate func as a callback for MessageCreate events.
	dg.AddHandler(messageCreate)

	// Open a websocket connection to Discord and begin listening.
	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Cleanly close down the Discord session.
	dg.Close()
}

// This function will be called (due to AddHandler above) every time a new
// message is created on any channel that the autenticated bot has access to.
func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	guild, err := getGuild(s, m)

	if err != nil {
		return
	}

	member, err := s.GuildMember(guild.ID, m.Author.ID)

	if err != nil {
		return
	}

	match := false
	for _, v := range member.Roles {
		if v == config.Config.FilterRoleID {
			match = true
			break
		}
	}

	isURL, _ := regexp.Match(urlRegexp, []byte(m.Content))

	if !match || !isURL {
		return
	}

	message := m.Author.Username + ":\n" + m.Content
	s.ChannelMessageSend(config.Config.FilterChannelID, message)
	s.ChannelMessageDelete(m.ChannelID, m.ID)
}

func getGuild(s *discordgo.Session, m *discordgo.MessageCreate) (*discordgo.Guild, error) {
	var guild *discordgo.Guild

	channel, err := s.State.Channel(m.ChannelID)
	if err != nil {
		channel, err = s.Channel(m.ChannelID)
		if err != nil {
			return guild, err
		}
	}

	// Attempt to get the guild from the state,
	// If there is an error, fall back to the restapi.
	guild, err = s.State.Guild(channel.GuildID)
	if err != nil {
		guild, err = s.Guild(channel.GuildID)
		if err != nil {
			return guild, err
		}
	}

	return guild, nil
}
