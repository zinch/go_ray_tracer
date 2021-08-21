package main

import (
	"fmt"

	"example.com/ray-tracer/core"
)

type Projectile struct {
	position core.Tuple
	velocity core.Tuple
}

type Environment struct {
	gravity core.Tuple
	wind    core.Tuple
}

func tick(env Environment, p Projectile) Projectile {
	position := p.position.Plus(p.velocity)
	velocity := p.velocity.Plus(env.gravity).Plus(env.wind)
	return Projectile{position: position, velocity: velocity}
}

func main() {
	projectile := Projectile{
		position: core.Point(0, 1, 0),
		velocity: core.Vector(1, 1, 0).Normalize(),
	}
	env := Environment{
		gravity: core.Vector(0, -0.1, 0),
		wind:    core.Vector(-0.01, 0, 0),
	}

	fmt.Println("Launch!")

	t := 0
	for i := 0; i < 20; i++ {
		t += 1
		projectile = tick(env, projectile)
		fmt.Printf("Position: %v\n", projectile.position)
		if projectile.position.Y < 1e-6 {
			fmt.Printf("Hit the ground after %d ticks\n", t)
			break
		}
	}
}
