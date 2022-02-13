package main

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"text/template"

	"gopkg.in/yaml.v2"

	"github.com/lemenendez/gogo/pkg"

	"github.com/spf13/cobra"
)

var (
	cfgFileName string
	cfg         *pkg.Config
)

const defCfgFileName string = "gogo-config"

func initConfig() {
	var realCfgFileName = cfgFileName
	if cfgFileName == "" {
		// fall back to default filename and path
		wd, err := os.Getwd()
		if err != nil {
			panic(err)
		}
		realCfgFileName = fmt.Sprintf("%v/%v.%v", wd, defCfgFileName, "yaml")
		err = saveDefaultConfig(realCfgFileName)
		if err != nil {
			panic(err)
		}
	}
	var err error
	cfg, err = newConfigFromFile(realCfgFileName)
	if err != nil {
		panic(err)
	}
}

var rootCmd = &cobra.Command{
	Use:   "gogo input.yml generator-name template-name",
	Short: "A code generator",
	Long:  `A code generator written in go. It uses tpl to generate code.`,
	Args:  cobra.MinimumNArgs(3),
	Run: func(cmd *cobra.Command, args []string) {
		ent, err := newEntityFromFile(args[0])
		if err != nil {
			panic(err)
		}
		dirPath := filepath.Dir(args[0])
		// fmt.Println(dirPath)
		if gen, found := cfg.Gens[args[1]]; found {
			tpl, err := template.New("examples").Funcs(pkg.FuncMap).ParseGlob(fmt.Sprintf("%v/%v/*.tpl", dirPath, args[1]))
			// tpl, err := template.New("examples").Funcs(pkg.FuncMap).ParseGlob("pkg/examples/mysql/*.tpl")
			if err != nil {
				fmt.Println(err)
				panic(err)
			}
			gen.Tpl = tpl
			ent.Map(gen.FieldMap)
			if res, err := gen.Exec(ent, args[2]); err != nil {
				fmt.Println(err)
				panic(err)
			} else {
				fmt.Print(res)
			}
			// fmt.Println(gen.Exec(ent, args[2]))
		} else {
			fmt.Println("generator not found")
			os.Exit(1)
		}
	},
}

func Setup() {
	rootCmd.PersistentFlags().StringVarP(&cfgFileName, "config", "c", "", fmt.Sprintf("config file name (default is PWD/%v.yml)", defCfgFileName))
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

func saveDefaultConfig(filename string) error {
	exists, err := fileExists(filename)
	if err != nil {
		return err
	}
	if !exists {
		cfg := getDefaultConfig()
		bdata, err := yaml.Marshal(cfg)
		if err != nil {
			return err
		}
		err = os.WriteFile(filename, bdata, 0644)
		if err != nil {
			return err
		}
		return nil
	}
	return nil
}
func newConfigFromFile(filename string) (*pkg.Config, error) {
	bData, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	cfg := &pkg.Config{}
	err = cfg.UnmarshalYAML(bData)
	if err != nil {
		panic(err)
	}
	return cfg, nil
}

func newEntityFromFile(filename string) (*pkg.Entity, error) {
	bData, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	ent := &pkg.Entity{}
	err = ent.UnmarshalYAML(bData)
	if err != nil {
		panic(err)
	}
	return ent, nil
}

func fileExists(name string) (bool, error) {
	_, err := os.Stat(name)
	if err == nil {
		return true, nil
	}
	if errors.Is(err, os.ErrNotExist) {
		return false, nil
	}
	return false, err
}

func getDefaultConfig() *pkg.Config {
	gens := make(map[string]pkg.Generator)
	mysql := NewMySQLDDLDefault()
	gens["mysql"] = mysql
	c := &pkg.Config{
		Name: "gogo default config file",
		Gens: gens,
	}
	return c
}

func NewMySQLDDLDefault() pkg.Generator {

	fm := make(pkg.FieldMap)
	fm["int"] = "int"
	fm["big_int"] = "bigint"
	fm["string"] = "varchar"
	fm["date"] = "date"
	fm["date_time"] = "datetime"
	fm["small_int"] = "smallint"
	fm["uuid"] = "binary(16)"
	fm["medium_text"] = "mediumtext"
	myddl := pkg.Generator{
		FieldMap:    fm,
		Description: "MySQL DDL Generator",
	}
	return myddl
}
