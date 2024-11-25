package main

import (
	"context"
	"fmt"

	"github.com/aykay76/cargo-command/internal/models"
)

func main() {
	ctx := context.Background()

	// TODO: add endpoint to create simulation and accept definition (which I need to create also)

	// TODO: obviously this needs to move to be triggered on some event, either when the game is started
	// or when a number of players is reached or something
	sim := models.NewSimulation()
	sim.Start(ctx)

	fmt.Println("You can start a simulation by calling http://localhost:6001/start")

	// TODO: Start API server

	fmt.Println("Simulation closed nicely")
}
