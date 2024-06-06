package main

import (
	"errors"
	"log"

	"github.com/amirhnajafiz/nginx-configmap-operator/internal/config"
	"github.com/amirhnajafiz/nginx-configmap-operator/internal/controller"
)

func main() {
	// load configs
	cfg, err := config.Load()
	if err != nil {
		panic(err)
	}

	// load controllers (first store them in tmp local dir, after that move them to nginx html dir)
	ctls := controller.LoadControllers(cfg.TmpLocalDir, cfg.NginxHTMLDir, cfg.Filename)

	// use the controller that is set in configs
	if ctl, ok := ctls[cfg.Type]; ok {
		// run the controller (input parameter is the source address)
		if err := ctl.GetFiles(cfg.Address); err != nil {
			log.Fatal(err)
		}
	} else {
		panic(errors.New("your input type is not supported by the system"))
	}
}
