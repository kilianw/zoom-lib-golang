package zoom

import "fmt"

// GetChatChannelPath - v2 path for getting a specific chat channel
const GetChatChannelPath = "/chat/channels/%s"

// GetChatChannelOpts contains options for GetUser
type GetChatChannelOpts struct {
	ChannelID string `url:"channelId"`
}

// GetChatChannel calls /chat/channels/{channelId}, searching for a chat channel by channelID, using the default client
func GetChatChannel(opts GetChatChannelOpts) (User, error) {
	return defaultClient.GetChatChannel(opts)
}

// GetChatChannel calls /chat/channels/{channelId, searching for a chat channel by channelID, using a specific client
func (c *Client) GetChatChannel(opts GetChatChannelOpts) (User, error) {
	var ret = User{}
	return ret, c.requestV2(requestV2Opts{
		Method:        Get,
		Path:          fmt.Sprintf(GetChatChannelPath, opts.ChannelID),
		URLParameters: opts,
		Ret:           &ret,
	})
}
