package router

import (
	"github.com/julienschmidt/httprouter"
	"log"
	"sync"
)

var instance JulienSchmidt
var once sync.Once

type JulienSchmidt struct {
	Router *httprouter.Router
}

func Build() {
	once.Do(func() {
		instance = JulienSchmidt{
			httprouter.New(),
		}
	})
}

func GetRouter() *httprouter.Router{
	if instance.Router == nil {
		log.Fatal("Router has not been created")
	}
	return instance.Router
}


