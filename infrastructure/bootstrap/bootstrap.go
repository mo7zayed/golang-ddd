package bootstrap

import "ddd/infrastructure/routes"

// StartApplication : Starts New Instance From The Application
func StartApplication(port string) {
	routes.RegisterRoutes(port)
}
