package zoom

import "fmt"

// DeleteChatChannelOptions are the options to delete a meeting
type DeleteChatChannelOptions struct {
	ChatChannelID string
}

// DeleteChatChannelPath - v2 delete a chat channel
const DeleteChatChannelPath = "/chat/channels/%s"

// DeleteChatChannel calls DELETE /chat/channels/{channelId}
func DeleteChatChannel(opts DeleteChatChannelOptions) error {
	return defaultClient.DeleteChatChannel(opts)
}

// DeleteChatChannel calls DELETE /chat/channels/{channelId}
func (c *Client) DeleteChatChannel(opts DeleteChatChannelOptions) error {
	return c.requestV2(requestV2Opts{
		Method:        Delete,
		Path:          fmt.Sprintf(DeleteChatChannelPath, opts.ChatChannelID),
		URLParameters: &opts,
		HeadResponse:  true,
	})
}
