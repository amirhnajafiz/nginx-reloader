package clone

// clone controller clones into a git repository to get the files
// from the given address.
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
