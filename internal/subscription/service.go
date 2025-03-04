package subscription

type Subscription struct {
	Name  string
	Users []int
}

func NewSubscription(name string, users []int) *Subscription {
	return &Subscription{
		Name:  name,
		Users: users,
	}
}
