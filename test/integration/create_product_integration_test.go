package integration

import (
	"fmt"
	"os"
	"testing"

	"github.com/archel/product_store/test/containers"
)

func TestCreatesAProduct(t *testing.T) {
	container, err := containers.NewPostgresContainer()

	if err != nil {
		fmt.Fprint(os.Stderr, err)
		panic(-1)
	}

	defer func() {
		if err := container.Terminate(); err != nil {
			t.Fatalf("failed to terminate container: %s", err.Error())
		}
	}()
}
