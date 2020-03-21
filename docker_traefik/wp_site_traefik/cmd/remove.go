package cmd

import (
	"fmt"
	"os"

	"github.com/nithin-bose/pekka/pkg"
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(removeCmd)
}

var removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "Removes a wordpress deployment",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			pkg.FatalF("Command requires exactly 1 argumant, deployment name \n")
		}

		deploymentName := args[0]
		deploymentPath := pkg.GetDeploymentPath(deploymentName)
		_, err := os.Stat(deploymentPath)
		if os.IsNotExist(err) {
			pkg.FatalF("Deployment does not exist \n")
		}

		os.Chdir(deploymentPath)

		dockerComposePath := pkg.GetDockerComposePath()
		dockerComposeArgs := []string{
			"stop",
		}
		err = pkg.Execute(dockerComposePath, dockerComposeArgs)
		if err != nil {
			pkg.FatalF("An error occurred while stopping services:\n %s \n", err.Error())
		}

		fmt.Println("Deployment stopped")

		dockerComposeArgs = []string{
			"rm",
			"-f",
			"-v",
		}
		err = pkg.Execute(dockerComposePath, dockerComposeArgs)
		if err != nil {
			pkg.FatalF("An error occurred while removing services:\n %s \n", err.Error())
		}

		fmt.Println("Removing files...")

		err = os.RemoveAll(deploymentPath)
		if err != nil {
			pkg.FatalF("An error occurred:\n %s \n", err.Error())
		}
		fmt.Println("Deployment removed")
	},
}
