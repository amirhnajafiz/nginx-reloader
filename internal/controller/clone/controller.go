package clone

type controller struct{}

func New() *controller {
	return &controller{}
}

func (c controller) GetFiles(address string) error {
	return nil
}
