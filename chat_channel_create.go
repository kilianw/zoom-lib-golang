package zoom

import "fmt"

// CreateChatChannelPath - v2 path for creating a chat channel
const CreateChatChannelPath = " /chat/users/{userId}/channels"

// CreateChatChannelInfo are details about a chat channel to create
type CreateChatChannelInfo struct {
	Name    string   `json:"name"`
	Type    int      `json:"type"`
	Members []Member `json:"members,omitempty"`
}

// CreateChatChannelOptions are the options to create a user with
type CreateChatChannelOptions struct {
	UserID          string                `json:"userId"`
	ChatChannelInfo CreateChatChannelInfo `json:"user_info"`
}

// CreateChatChannel calls POST /users/{userId}/meetings
func CreateChatChannel(opts CreateChatChannelOptions) (ChatChannel, error) {
	return defaultClient.CreateChatChannel(opts)
}

// CreateChatChannel calls POST /chat/users/{userId}/channels
// https://marketplace.zoom.us/docs/api-reference/chat/methods/#operation/createChannel
func (c *Client) CreateChatChannel(opts CreateChatChannelOptions) (ChatChannel, error) {
	var ret = ChatChannel{}
	return ret, c.requestV2(requestV2Opts{
		Method:         Post,
		Path:           fmt.Sprintf(CreateChatChannelPath, opts.UserID),
		DataParameters: &opts.ChatChannelInfo,
		Ret:            &ret,
	})
}
