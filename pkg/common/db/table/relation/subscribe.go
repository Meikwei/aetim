package relation

import "context"

// SubscribeUser collection constant.
const (
	SubscribeUser = "subscribe_user"
)

// SubscribeUserModel collection structure.
type SubscribeUserModel struct {
	UserID     string   `bson:"user_id"      json:"userID"`
	UserIDList []string `bson:"user_id_list" json:"userIDList"`
}

func (SubscribeUserModel) TableName() string {
	return SubscribeUser
}

// SubscribeUserModelInterface Operation interface of user mongodb.
type SubscribeUserModelInterface interface {
	// AddSubscriptionList Subscriber's handling of thresholds.
	AddSubscriptionList(ctx context.Context, userID string, userIDList []string) error
	// UnsubscriptionList Handling of unsubscribe.
	UnsubscriptionList(ctx context.Context, userID string, userIDList []string) error
	// RemoveSubscribedListFromUser Among the unsubscribed users, delete the user from the subscribed list.
	RemoveSubscribedListFromUser(ctx context.Context, userID string, userIDList []string) error
	// GetAllSubscribeList Get all users subscribed by this user
	GetAllSubscribeList(ctx context.Context, id string) (userIDList []string, err error)
	// GetSubscribedList Get the user subscribed by those users
	GetSubscribedList(ctx context.Context, id string) (userIDList []string, err error)
}
