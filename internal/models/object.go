package models

// GameObject is an interface that all game objects must implement to be part of the game.
type GameObject interface {
	// Update is called to update the state of the game object.
	Update()

	// Render is called to render the game object.
	Render()

	// HandleEvent is called to handle events that affect the game object.
	HandleEvent(event Event)
}
