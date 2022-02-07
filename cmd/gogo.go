package main

import (
	"fmt"
	"os"

	"github.com/lemenendez/gogo/pkg"
	"github.com/lemenendez/gogo/pkg/in"

	"github.com/spf13/cobra"
)

var (
	cfgFileName string
	cfg         *in.Config
)

const defCfgFileName string = "gogo"

func initConfig() {
	var realCfgFileName = cfgFileName
	if cfgFileName == "" {
		// fall back to default filename and path
		home, err := os.UserHomeDir()
		if err != nil {
			panic(err)
		}
		realCfgFileName = fmt.Sprintf("%v/%v.%v", home, defCfgFileName, "yaml")
		err = in.SaveDefaultConfig(realCfgFileName)
		if err != nil {
			panic(err)
		}
	}
	var err error
	cfg, err = in.NewConfigFromFile(realCfgFileName)
	if err != nil {
		panic(err)
	}

}

var rootCmd = &cobra.Command{
	Use:   "gogo input.yml generator template",
	Short: "A code generator",
	Long:  `A code generator written in go. It uses tpl to generate code.`,
	Args:  cobra.MinimumNArgs(3),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Print("Code Generator Start\n")
		iEntity, err := in.NewEntityFromFile(args[0])
		if err != nil {
			panic(err)
		}
		entity, err := pkg.ToEntity(*iEntity)
		if err != nil {
			panic(err)
		}
		if gen, found := cfg.Gens[args[1]]; found {
			codeGenerator, err := pkg.NewGen(gen.TplPath)
			if err != nil {
				panic(err)
			}
			entity.Map(gen.FieldMap)
			fmt.Println(codeGenerator.Generate(entity, args[2]))
		}
	},
}

func Setup() {
	rootCmd.PersistentFlags().StringVarP(&cfgFileName, "config", "c", "", fmt.Sprintf("config file name (default is $HOME/%v.yml)", defCfgFileName))
	cobra.OnInitialize(initConfig)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func main() {
	Setup()
	Execute()
}
