package main

import (
	"encoding/json"
	"errors"
	"log"

	"github.com/amirhnajafiz/nginx-configmap-operator/internal/config"
	"github.com/amirhnajafiz/nginx-configmap-operator/internal/controller"

	"github.com/tidwall/pretty"
)

func main() {
	// load configs
	cfg, err := config.Load()
	if err != nil {
		panic(err)
	}

	indent, err := json.MarshalIndent(cfg, "", "\t")
	if err != nil {
		log.Fatalf("error marshaling config to json: %s", err)
	}

	indent = pretty.Color(indent, nil)
	tmpl := `
	================ Loaded Configuration ================
	%s
	======================================================
	`
	log.Printf(tmpl, string(indent))

	// load controllers (first store them in tmp local dir, after that move them to nginx html dir)
	ctls := controller.LoadControllers(cfg.TmpLocalDir, cfg.NginxHTMLDir, cfg.Filename)

	// use the controller that is set in configs
	if ctl, ok := ctls[cfg.Type]; ok {
		log.Printf("controller %s is loaded.\n", cfg.Type)

		// run the controller (input parameter is the source address)
		if err := ctl.GetFiles(cfg.Address); err != nil {
			log.Fatal(err)
		}

		log.Println("nginx-reloader done.")
	} else {
		panic(errors.New("your input type is not supported by the system"))
	}
}
