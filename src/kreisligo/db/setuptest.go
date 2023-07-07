package db

import (
	"testing"

	"github.com/ory/dockertest/v3"
	"github.com/ory/dockertest/v3/docker"
)

func SetupTestDB(t *testing.T) func() {
	pool, err := dockertest.NewPool("")
	if err != nil {
		t.Fatalf("Could not connect to docker: %s", err)
	}

	runDockerOpt := &dockertest.RunOptions{
		Repository: "mariadb",
		Tag:        "10.5",
		Env:        []string{"MYSQL_ROOT_PASSWORD=root", "MYSQL_DATABASE=myaktion"},
		PortBindings: map[docker.Port][]docker.PortBinding{
			"3306/tcp": {{HostIP: "localhost", HostPort: "3306/tcp"}},
		},
	}

	fnConfig := func(config *docker.HostConfig) {
		config.AutoRemove = true                     // set AutoRemove to true so that stopped container goes away by itself
		config.RestartPolicy = docker.NeverRestart() // don't restart container
	}

	resource, err := pool.RunWithOptions(runDockerOpt, fnConfig)
	if err != nil {
		t.Fatalf("Could not start test DB: %s", err)
	}

	// retry until db server is ready
	err = pool.Retry(func() error {
		return Connect("localhost:3306")
	})
	if err != nil {
		t.Fatalf("Could not connect to test DB: %s", err)
	}

	return func() {
		resource.Close()
	}
}

