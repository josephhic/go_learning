package main

type Projectile struct {
	position, speed, acceleration Vector
}

func (projectile *Projectile) getPosition() Vector {
	return projectile.position
}

func (projectile *Projectile) getSpeed() Vector {
	return projectile.speed
}

func (projectile *Projectile) getAcceleration() Vector {
	return projectile.acceleration
}
