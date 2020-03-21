package pkg

import (
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

//GetDockerComposePath returns the docker-compose path else prints an error and exits
func GetDockerComposePath() string {
	dockerComposePath, err := exec.LookPath("docker-compose")
	if err != nil {
		log.Fatalln("docker-compose could not be found")
	}
	return dockerComposePath
}

//GetPekkaTraefikPath returns the path to the pekka traefik folder
func GetPekkaTraefikPath() string {
	cwd, err := os.Getwd()
	if err != nil {
		FatalF("An error occurred:\n %s \n", err.Error())
	}

	return filepath.Join(cwd, "pekka-traefik")
}

//GetDeploymentPath returns the path to the given deployment
func GetDeploymentPath(deploymentName string) string {
	cwd, err := os.Getwd()
	if err != nil {
		FatalF("An error occurred:\n %s \n", err.Error())
	}

	return filepath.Join(cwd, deploymentName)
}
