package event

import (
	"trashy-ecs/pkg/entity"
)

type CollisionEvent struct {
	firstEntity  entity.Entity
	secondEntity entity.Entity
}

type KeypressEvent struct {
	key int
}

type JumpEvent struct {
	Entity entity.Entity
}
