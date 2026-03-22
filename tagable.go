package ludum

// Tagable implies an entity carries a numeric classification mark, useful for identification across various systems.
type Tagable interface {
	// GetTag retrieves the entity's designated integer affiliation id.
	GetTag() int
}
