package ts3

// CHANNEL CODECS VALUES
const CODEC_SPEEX_NARROWBAND = 0x00
const CODEC_SPEEX_WIDEBAND = 0x01
const CODEC_SPEEX_ULTRAWIDEBAND = 0x02
const CODEC_CELT_MONO = 0x03
const CODEC_OPUS_VOICE = 0x04
const CODEC_OPUS_MUSIC = 0x05

// CHANNEL STRUCTURE 
type Channel struct {
	ChannelName string
	ChannelPassword string
	ChannelCodec int
	ChannelCodecQuality int
	ChannelFlagPermanent int
	ChannelFlagSemiPermanent int

}

// TEAMSPEAK 3 TARGET 
const TARGET_SERVER = 0x03
const TARGET_CHANNEL = 0x02
const TARGET_CLIENT = 0x01