package gate

type CategoryCreator struct{}

func (c *CategoryCreator) Allow(u User) bool {
	return true
}
