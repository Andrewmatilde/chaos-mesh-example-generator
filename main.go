//go:generate go get -u github.com/valyala/quicktemplate/qtc
//go:generate qtc -dir=./template/

package main

import (
	"chaos-mesh-example-generator/template/chaos-mesh/api/v1alpha1"
	"chaos-mesh-example-generator/template/chaos-mesh/cmd/controllerManager"
	"chaos-mesh-example-generator/template/chaos-mesh/config/crd"
	"chaos-mesh-example-generator/template/chaos-mesh/config/crd/bases"
	"chaos-mesh-example-generator/template/chaos-mesh/controllers"
	helm "chaos-mesh-example-generator/template/chaos-mesh/helm/chaos-mesh/templates"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/urfave/cli/v2"
)

func writeTemplate(path,s string) {
	fmt.Println(path)
	_, err := os.Stat(path)
	if err == nil || os.IsExist(err) {
		err = os.Rename(path, path+".bak")
		if err != nil {
			panic(err)
		}
	} else {
		if !os.IsNotExist(err) {
			panic(err)
		}
	}

	file,err := os.Create(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	_, err = file.WriteString(s)
	if err != nil {
		panic(err)
	}
}

func main() {
	err := (&cli.App{
		Name:  "chaos-mesh-example-generator",
		Usage: "Generate an chaos mesh example automatically!",
		Commands: []*cli.Command{
			&cli.Command{
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:    "chaos-mesh-dir",
						Aliases: []string{"d"},
						Usage:   "chaos mesh dir position, just like /chaos-mesh",
					},
					&cli.StringFlag{
						Name:    "chaos-name",
						Aliases: []string{"n"},
						Usage:   "chaos name",
					},
				},
				Name:    "generate",
				Aliases: []string{"gen"},
				Usage:   "Generate the example",
				Action: func(context *cli.Context) error {
					chaosName := context.String("chaos-name")
					chaosDirPath := context.String("chaos-mesh-dir")
					chaosType := v1alpha1.ChaosType(chaosName)
					controllerManager := controllerManager.ControllerManager(chaosName)
					crdYaml := bases.CRDYaml(chaosName)
					kustomizationYaml := crd.KustomizationYaml(chaosName)
					chaosContoller := controllers.ChaosContoller(chaosName)
					controllerManagerRbacYaml := helm.ControllerManagerRbacYaml(chaosName)
					writeTemplate(filepath.Join(chaosDirPath,"api/v1alpha1/" + chaosName +"chaos_types.go"),chaosType)
					writeTemplate(filepath.Join(chaosDirPath,"cmd/controllerManager/main.go"),controllerManager)
					writeTemplate(filepath.Join(chaosDirPath,"config/crd/bases/chaos-mesh.org_"+chaosName +"_chaos.yaml"),crdYaml)
					writeTemplate(filepath.Join(chaosDirPath,"config/crd/kustomization.yaml"),kustomizationYaml)
					writeTemplate(filepath.Join(chaosDirPath,"controllers/" + chaosName +"_controller.go"),chaosContoller)
					writeTemplate(filepath.Join(chaosDirPath,"helm/chaos-mesh/templates/controller-manager-rbac.yaml"),controllerManagerRbacYaml)
					return nil
				},
			},
			&cli.Command{
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:    "chaos-mesh-dir",
						Aliases: []string{"d"},
						Usage:   "chaos mesh dir position, just like /chaos-mesh",
					},
					&cli.StringFlag{
						Name:    "chaos-name",
						Aliases: []string{"n"},
						Usage:   "chaos name",
					},
				},
				Name:  "clear",
				Usage: "clear the example and recover",
				Action: func(context *cli.Context) error {
					return nil
				},
			},
		},
	}).Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
