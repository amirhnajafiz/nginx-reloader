package controller

// Controller interface represents a abstract structure
// for our controllers.
type Controller interface {
	GetFiles(address string) error
}

// LoadControllers returns a map of controllers to be used by main.
func LoadControllers(localDir, nginxDir string) map[string]Controller {
	return nil
}
