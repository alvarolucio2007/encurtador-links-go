package main

import (
	"github.com/alvarolucio2007/encurtador-links-go/internal/cache"
	"github.com/alvarolucio2007/encurtador-links-go/internal/database"
	routes "github.com/alvarolucio2007/encurtador-links-go/internal/routesAPI"
)

func main() {
	database.ConectarDataBase()
	cache.ConectarRedis()
	routes.SetupAPI()
}
