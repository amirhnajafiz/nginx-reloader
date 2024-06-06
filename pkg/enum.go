package pkg

// Types represent how reloader should get files from the given address.
// It can clone into a repository or download them from the give source.
const (
	TypeClone string = "clone"
	TypeFetch string = "download"
)
