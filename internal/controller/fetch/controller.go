package fetch

// fetch controller downloads the files from the given address.
// it uses wget command.
type controller struct{}

func New() *controller {
	return &controller{}
}

func (c controller) GetFiles(address string) error {
	return nil
}
