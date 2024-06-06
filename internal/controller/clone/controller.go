package clone

// clone controller clones into a git repository to get the files
// from the given address.
type controller struct{}

func New() *controller {
	return &controller{}
}

func (c controller) GetFiles(address string) error {
	return nil
}
