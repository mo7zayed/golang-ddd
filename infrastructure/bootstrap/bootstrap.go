package bootstrap

import "golang-ddd/infrastructure/routes"

// StartApplication : Starts New Instance From The Application
func StartApplication(port string) {
	routes.RegisterRoutes(port)
}
