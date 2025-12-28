package srp

import (
	"fmt"
	"time"
)

type (
	Player struct {
		X, Y   float64
		Health int
		Sprite string
	}

	PlayerMovement struct{}

	PlayerRenderer struct{}

	PlayerSaver struct{}

	PlayerUpdater struct {
		Movement *PlayerMovement
		Renderer *PlayerRenderer
		Saver    *PlayerSaver
	}
)

func (m *PlayerMovement) Move(p *Player, delta float64) {
	p.X += 10 * delta
}

func (r *PlayerRenderer) Draw(p *Player) {
	fmt.Printf("Desenhando %s em (%.2f, %.2f)\n", p.Sprite, p.X, p.Y)
}

func (s *PlayerSaver) SaveIfNeeded(p *Player) {
	if p.X > 100 {
		fmt.Println("Salvando jogo...")
	}
}

func (u *PlayerUpdater) Update(p *Player, delta float64) {
	u.Movement.Move(p, delta)
	u.Renderer.Draw(p)
	u.Saver.SaveIfNeeded(p)
}

func main() {
	player := &Player{
		X:      0,
		Y:      0,
		Health: 100,
		Sprite: "Hero",
	}

	movement := &PlayerMovement{}
	renderer := &PlayerRenderer{}
	saver := &PlayerSaver{}

	updater := &PlayerUpdater{
		Movement: movement,
		Renderer: renderer,
		Saver:    saver,
	}

	lastTime := time.Now()

	for {
		now := time.Now()
		delta := now.Sub(lastTime).Seconds()
		lastTime = now

		updater.Update(player, delta)

		time.Sleep(16 * time.Millisecond)
	}
}
