package main

import (
	"math/rand"
	"time"

	"github.com/docker/docker/pkg/namesgenerator"
)

func main() {
	// seed the random number generator with the current time
	rand.Seed(time.Now().UnixNano()) // #nosec G401: this is not a security risk
	names := namesgenerator.GetRandomName(0)
	println(names)
}
