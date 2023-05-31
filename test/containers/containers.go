package containers

import (
	"context"

	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

type Container interface {
	Terminate() error
}

type PostgresContainer struct {
	container testcontainers.Container
	ctx       context.Context
}

func (pc *PostgresContainer) Terminate() error {
	return pc.container.Terminate(pc.ctx)
}

func NewPostgresContainer() (Container, error) {
	ctx := context.Background()

	req := testcontainers.ContainerRequest{
		Image:        "postgres:15.3-alpine",
		Name:         "products-postgres",
		ExposedPorts: []string{"5432/tcp"},
		WaitingFor:   wait.ForLog("database system is ready to accept connections"),
		Env: map[string]string{
			"POSTGRES_PASSWORD": "s3cr3t",
			"POSTGRES_USER":     "postgres",
			"POSTGRES_DB":       "products",
		},
	}

	container, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
		Reuse:            true,
	})

	if err != nil {
		return nil, err
	}

	return &PostgresContainer{
		container: container,
		ctx:       ctx,
	}, nil
}
