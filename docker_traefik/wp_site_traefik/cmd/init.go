package cmd

import (
	"os"
	"path/filepath"

	prompt "gopkg.in/distillerytech/prompt.v1"

	"github.com/nithin-bose/pekka/pkg"
	"github.com/nithin-bose/pekka/templates"
	"github.com/spf13/cobra"
)

var force bool

func init() {
	RootCmd.AddCommand(initCmd)
	initCmd.Flags().BoolVarP(&force, "force", "f", false, "forces an init")
}

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initializes infra required for pekka to run",
	RunE: func(cmd *cobra.Command, args []string) error {
		pekkaTraefikPath := pkg.GetPekkaTraefikPath()
		_, err := os.Stat(pekkaTraefikPath)
		if !os.IsNotExist(err) {
			if !force {
				pkg.FatalF("Pekka is already initialized \n")
			} else {
				err = os.RemoveAll(pekkaTraefikPath)
				if err != nil {
					pkg.FatalF("An error occurred:\n %s \n", err.Error())
				}
			}
		}

		traefikFilesData := templates.TraefikFilesData{}
		traefikFilesData.TraefikDashboardURL = prompt.AskStringRequired("Enter traefik dashboard URL:")
		traefikFilesData.AcmeEmail = prompt.AskStringRequired("Enter let's encrypt email:")

		err = os.Mkdir(pekkaTraefikPath, os.ModePerm)
		if err != nil {
			pkg.FatalF("An error occurred:\n %s \n", err.Error())
		}

		for fileName, template := range templates.TraefikFiles {
			fPath := filepath.Join(pekkaTraefikPath, fileName)
			pkg.CreateFile(fPath, template, traefikFilesData)
		}

		os.Chdir(pekkaTraefikPath)

		dockerComposePath := pkg.GetDockerComposePath()
		dockerComposeArgs := []string{
			"up",
			"-d",
		}
		return pkg.Execute(dockerComposePath, dockerComposeArgs)

	},
}
