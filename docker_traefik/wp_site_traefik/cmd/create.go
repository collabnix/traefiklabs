package cmd

import (
	"os"
	"path/filepath"

	prompt "gopkg.in/distillerytech/prompt.v1"

	"github.com/nithin-bose/pekka/pkg"
	"github.com/nithin-bose/pekka/templates"
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(createCmd)
}

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Creates a wordpress deployment",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			pkg.FatalF("Command requires exactly 1 argumant, deployment name \n")
		}

		pekkaTraefikPath := pkg.GetPekkaTraefikPath()
		_, err := os.Stat(pekkaTraefikPath)
		if os.IsNotExist(err) {
			pkg.FatalF("Please initialize Pekka before creating a deployment \n")
		}

		deploymentName := args[0]
		deploymentPath := pkg.GetDeploymentPath(deploymentName)
		_, err = os.Stat(deploymentPath)
		if !os.IsNotExist(err) {
			pkg.FatalF("Deployment already exists \n")
		}

		wordPressFilesData := templates.WordPressFilesData{}
		wordPressFilesData.MySQLRootPassword = pkg.RandString(8)
		wordPressFilesData.WPTablePrefix = pkg.RandString(4)
		wordPressFilesData.TraefikBackend = deploymentName
		wordPressFilesData.WPHosts = prompt.AskStringRequired("Enter domain name:")
		addWWW := prompt.Confirm("Add entry for www." + wordPressFilesData.WPHosts + "?")
		if addWWW {
			wordPressFilesData.WPHosts += ",www." + wordPressFilesData.WPHosts
		}

		err = os.Mkdir(deploymentPath, os.ModePerm)
		if err != nil {
			pkg.FatalF("An error occurred:\n %s \n", err.Error())
		}

		for fileName, template := range templates.WordPressFiles {
			fPath := filepath.Join(deploymentPath, fileName)
			pkg.CreateFile(fPath, template, wordPressFilesData)
		}

		os.Chdir(deploymentPath)

		dockerComposePath := pkg.GetDockerComposePath()
		dockerComposeArgs := []string{
			"up",
			"-d",
		}
		return pkg.Execute(dockerComposePath, dockerComposeArgs)

	},
}
