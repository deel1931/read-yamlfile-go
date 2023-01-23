package main

import (
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

func write() {
	data := map[string]string{
		"version": "2",
		"key":     "adjkfaj",
	}

	f, err := os.OpenFile("local.yaml", os.O_WRONLY|os.O_CREATE, 0664)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	d := yaml.NewEncoder(f)

	if err := d.Encode(&data); err != nil {
		log.Fatal(err)
	}

	d.Close()
}
