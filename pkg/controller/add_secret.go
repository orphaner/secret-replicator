package controller

import (
	"github.com/kiwigrid/secret-replicator/pkg/controller/secret"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, secret.Add)
}
