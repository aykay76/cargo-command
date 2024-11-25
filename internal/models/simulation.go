package models

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Simulation represents the simulation environment
type Simulation struct {
	ActualStartTime  time.Time
	ActualFinishTime time.Time
	StartTime        time.Time
	CurrentTime      time.Time
	TimeStep         time.Duration
	AutoStep         bool
	DoneChan         chan struct{}
	Mutex            sync.Mutex
	RNG              *rand.Rand
	StepsTaken       int
	EndSteps         int
	ActorCount       int
	CurrentActors    int
}

// NewSimulation creates a new simulation instance
func NewSimulation() *Simulation {
	return &Simulation{
		AutoStep:        true,
		ActualStartTime: time.Now(),
		StartTime:       time.Date(2000, 1, 1, 0, 0, 0, 0, time.Now().UTC().Location()),
		CurrentTime:     time.Now(),
		TimeStep:        time.Hour,
		DoneChan:        make(chan struct{}),
		RNG:             rand.New(rand.NewSource(99)),
		EndSteps:        100,
		ActorCount:      1,
		CurrentActors:   0,
	}
}

func (sim *Simulation) Start(ctx context.Context) (err error) {
	// do some initialisation
	sim.StepsTaken = 0
	sim.ActualStartTime = time.Now()

	// TODO: for each actor type I need to define how many of each actor there will be and possibly define initial state

	// TODO: call Initialise on all actor interfaces

	// main sim loop (if auto-stepping, otherwise done and wait for step invocation)
	if sim.AutoStep {
		// TODO: raise step event - channels or redis streams?
	} else {
		fmt.Println("Simulation started and paused, invoke next step by calling http://localhost:6001/step")
	}

	sim.ActualFinishTime = time.Now()
	close(sim.DoneChan)
	fmt.Println("Simulation finished, simulation end time: ", sim.CurrentTime, "; actual finish time: ", sim.ActualFinishTime, "; run time: ", sim.ActualFinishTime.Sub(sim.ActualStartTime))

	return
}

func (sim *Simulation) Step(ctx context.Context) (err error) {
	sim.StepsTaken++

	if sim.StepsTaken >= sim.EndSteps {
		fmt.Println("Simulation complete, get stats by calling http://localhost:6001/stats")
	} else {
		// TODO: call step on all actor interfaces

		if sim.AutoStep {
			// TODO: raise step event - channels or redis streams?
		} else {
			fmt.Printf("%d steps taken of %d - call http://localhost:6001/step\n", sim.StepsTaken, sim.EndSteps)
		}
	}

	return
}
