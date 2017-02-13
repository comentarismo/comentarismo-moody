package server

import (
	//util "comentarismo-moody/util"
	"fmt"
	r "github.com/dancannon/gorethink"
)

// UUID creates a new unique ID, which can be used as database access ID.
func UUID(targetUUID string) (string, error) {
	// Create a new unique ID.
	r, err := r.UUID(targetUUID).Run(Session)
	if err != nil {
		return "", fmt.Errorf("failed to obtain a new unique ID: %v", err)
	}

	// Get the value.
	var id string
	err = r.One(&id)
	if err != nil {
		return "", fmt.Errorf("failed to obtain a new unique ID: %v", err)
	}

	if len(id) == 0 {
		return "", fmt.Errorf("failed to obtain a new unique ID: %v", err)
	}

	return id, nil
}
