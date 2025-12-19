package main

import (
	"failingversions"

	sensor "go.viam.com/rdk/components/sensor"
	"go.viam.com/rdk/module"
	"go.viam.com/rdk/resource"
)

func main() {
	panic("Im failing on version 0.0.3")

	// ModularMain can take multiple APIModel arguments, if your module implements multiple models.
	module.ModularMain(resource.APIModel{sensor.API, failingversions.TestModel})
}
