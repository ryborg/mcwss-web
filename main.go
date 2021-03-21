
package main

import (
	"github.com/ryborg/mcwss"
)

func main() {

	server := mcwss.NewServer(nil)

	server.OnConnection(func(player *mcwss.Player){
player.Tell("Hello %s!", player.Name())
})
	server.OnDisconnection(func(player *mcwss.Player){
	})
	server.Run()
}
