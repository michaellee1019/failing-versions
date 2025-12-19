package main

import (
	"failingversions"
	"go.viam.com/rdk/module"
	"go.viam.com/rdk/resource"
	sensor "go.viam.com/rdk/components/sensor"
)

func main() {
	// ModularMain can take multiple APIModel arguments, if your module implements multiple models.
	module.ModularMain(resource.APIModel{ sensor.API, failingversions.TestModel})

	panic("Im failing on version 0.0.1")
}
