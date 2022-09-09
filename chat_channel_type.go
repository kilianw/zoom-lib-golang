package zoom // types for /chat endpoints

// CreateChatChannelAction specifies how to create a new chat channel
type CreateChatChannelAction string

// DeleteChatChannelAction specifies how to delete a new chat channel
type DeleteChatChannelAction string

// ChatChannelType is one of a fixed number of possible chat channel types
type ChatChannelType int

const (

	// PrivateChannel channel type
	PrivateChannel ChatChannelType = 1

	// PrivateChannelWithAccountMembers channel type
	PrivateChannelWithAccountMembers ChatChannelType = 2

	// PublicChannel channel type
	PublicChannel ChatChannelType = 3

	// InstantChannel channel type
	InstantChannel ChatChannelType = 4
)

// String provides a string representation of user types
func (t ChatChannelType) String() (str string) {
	switch t {
	case PrivateChannel:
		str = "private"
	case PrivateChannelWithAccountMembers:
		str = "privateWithAccountMembers"
	case PublicChannel:
		str = "public"
	case InstantChannel:
		str = "instant"
	}
	return
}

// ChatChannel represents an account user
type ChatChannel struct {
	ID   string          `json:"id"`
	JID  string          `json:"jid"`
	Name Time            `json:"name,string"`
	Type ChatChannelType `json:"type"`
}
