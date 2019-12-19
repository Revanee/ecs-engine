package manager

import "trashy-ecs/pkg/entity"

// EntityManager manages instantiation of entities
type EntityManager interface {
	NewEntity() entity.Entity
	Exists(entity.Entity) bool
	All() []entity.Entity
}

// IEntityManager is an implementation of EntityManager
type IEntityManager struct {
	entities []entity.Entity
	nextID   int64
}

var _ EntityManager = (*IEntityManager)(nil)

// NewEntityManager instantiates an IEntityManager
func NewEntityManager() *IEntityManager {
	return &IEntityManager{
		entities: make([]entity.Entity, 0),
		nextID:   0,
	}
}

// NewEntity creates a new entity without conflicts with other entities
func (em *IEntityManager) NewEntity() entity.Entity {
	e := entity.Entity(em.nextID)
	em.entities = append(em.entities, e)
	em.nextID++
	return e
}

// Exists returns ture if an entity is already instantiated
func (em IEntityManager) Exists(entity entity.Entity) bool {
	for _, e := range em.entities {
		if e == entity {
			return true
		}
	}
	return false
}

// All returns a slice of all instantiated entities
func (em IEntityManager) All() []entity.Entity {
	return em.entities
}
