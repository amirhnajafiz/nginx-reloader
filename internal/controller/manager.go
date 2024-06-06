package controller

import (
	"github.com/amirhnajafiz/nginx-configmap-operator/internal/controller/clone"
	"github.com/amirhnajafiz/nginx-configmap-operator/internal/controller/fetch"
	"github.com/amirhnajafiz/nginx-configmap-operator/pkg"
)

// Controller interface represents a abstract structure
// for our controllers. Controllers get files from the given source, however their
// logic is different based on their map key.
type Controller interface {
	GetFiles(address string) error
}

// LoadControllers returns a map of controllers to be used by main.
func LoadControllers(localDir, nginxDir string) map[string]Controller {
	list := make(map[string]Controller)

	list[pkg.TypeClone] = clone.New()
	list[pkg.TypeFetch] = fetch.New()

	return list
}
