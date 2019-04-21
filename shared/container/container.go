package container

import (
	"github.com/fgrosse/goldi"
	AWS "github.com/mywishes/go-ses/shared/aws"
	Config "github.com/mywishes/go-ses/shared/config"
)

// GoldiDefaultContainer return default dependecy injection
func GoldiDefaultContainer() *goldi.Container {

	// create new container when application loads
	registry := goldi.NewTypeRegistry()
	config := make(map[string]interface{})
	container := goldi.NewContainer(registry, config)

	// defines the type using container
	container.RegisterType("shared.config", Config.NewImmutableConfig)
	container.RegisterType("shared.aws.session", AWS.NewAWS, "@shared.config")
	return container
}
