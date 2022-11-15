package eacstoken

import (
	"fmt"
	"time"
)

type Entity struct {
	username    string
	generatedAt time.Time
}

func Generate(username string) *Entity {
	return &Entity{
		username:    username,
		generatedAt: time.Now(),
	}
}

func (e Entity) String() string {
	return fmt.Sprintf("{%s, %s}", e.username, e.generatedAt.String())
}
