package ts3

import (
	"strconv"
)

// Provide's the Channel Parameters
type Channel struct {
	ChannelName              string
	ChannelPassword          string
	ChannelCodec             int
	ChannelCodecQuality      int
	ChannelFlagPermanent     int
	ChannelFlagSemiPermanent int
}

func ChannelList() Command {
	return Command{
		Command: "channellist",
		Flags:   []string{"flags"},
	}
}

func (c Channel) ChannelCreate() Command {

	return Command{
		Command: "channelcreate",
		Params: map[string][]string{
			"channel_name":                []string{c.ChannelName},
			"channel_password":            []string{c.ChannelPassword},
			"channel_codec":               []string{strconv.Itoa(c.ChannelCodec)},
			"channel_flag_permanent":      []string{strconv.Itoa(c.ChannelFlagPermanent)},
			"channel_flag_semi_permanent": []string{strconv.Itoa(c.ChannelFlagSemiPermanent)},
			"channel_codec_quality":       []string{strconv.Itoa(c.ChannelCodecQuality)},
		},
	}

}

func ChannelEdit(m map[string][]string) Command {

	return Command{
		Command: "channeledit",
		Params:  m,
	}

}
