package controller

import (
	"github.com/amirhnajafiz/nginx-configmap-operator/internal/controller/clone"
	"github.com/amirhnajafiz/nginx-configmap-operator/internal/controller/fetch"
	"github.com/amirhnajafiz/nginx-configmap-operator/internal/utils"
	"github.com/amirhnajafiz/nginx-configmap-operator/pkg"
)

// Controller interface represents a abstract structure
// for our controllers. Controllers get files from the given source, however their
// logic is different based on their map key.
type Controller interface {
	GetFiles(address string) error
}

// LoadControllers returns a map of controllers to be used by main.
func LoadControllers(localDir, nginxDir, filename string) map[string]Controller {
	list := make(map[string]Controller)

	list[pkg.TypeClone] = clone.New(localDir, controllersCallbackFunc(nginxDir))
	list[pkg.TypeFetch] = fetch.New(localDir, filename, controllersCallbackFunc(nginxDir))

	return list
}

// controllersCallbackFunc is function that each controller calls after it finishes
// the getting files process successfully.
func controllersCallbackFunc(nginxDir string) func(string) error {
	return func(path string) error {
		return utils.MoveDir(path, nginxDir)
	}
}
