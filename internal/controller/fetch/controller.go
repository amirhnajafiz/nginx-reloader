package fetch

// fetch controller downloads the files from the given address.
// it uses wget command.
type controller struct {
	localDir string
	callback func() error
}

func New(ld string, cb func() error) *controller {
	return &controller{
		localDir: ld,
		callback: cb,
	}
}

func (c controller) GetFiles(address string) error {
	return c.callback()
}
