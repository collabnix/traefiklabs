package cmd

import (
	"fmt"
	"os"

	"github.com/nithin-bose/pekka/pkg"
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(updateCmd)
}

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Updates a wordpress deployment",
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

		fmt.Println("\nPulling updates...")

		os.Chdir(deploymentPath)

		dockerComposePath := pkg.GetDockerComposePath()
		dockerComposeArgs := []string{
			"pull",
		}
		err = pkg.Execute(dockerComposePath, dockerComposeArgs)
		if err != nil {
			pkg.FatalF("An error occurred while pulling updates:\n %s \n", err.Error())
		}

		fmt.Println("Updates pulled")

		fmt.Println("\nStopping current deployment...")
		dockerComposeArgs = []string{
			"stop",
		}
		err = pkg.Execute(dockerComposePath, dockerComposeArgs)
		if err != nil {
			pkg.FatalF("An error occurred while stopping services:\n %s \n", err.Error())
		}

		dockerComposeArgs = []string{
			"rm",
			"-f",
			"-v",
		}
		err = pkg.Execute(dockerComposePath, dockerComposeArgs)
		if err != nil {
			pkg.FatalF("An error occurred while removing services:\n %s \n", err.Error())
		}

		fmt.Println("\nRestarting deployment with updates...")
		dockerComposeArgs = []string{
			"up",
			"-d",
		}
		err = pkg.Execute(dockerComposePath, dockerComposeArgs)
		if err != nil {
			pkg.FatalF("An error occurred while restarting services:\n %s \n", err.Error())
		}
		fmt.Println("\nDeployment updated")
	},
}
