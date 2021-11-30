package gen

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/2637309949/dolphin/packages/logrus"
	"github.com/2637309949/dolphin/packages/swagger"
	"github.com/ghodss/yaml"
)

// Gen presents a generate tool for swagger.
type Gen struct {
	jsonIndent func(data interface{}) ([]byte, error)
	jsonToYAML func(data []byte) ([]byte, error)
}

// New creates a new Gen.
func New() *Gen {
	return &Gen{
		jsonIndent: func(data interface{}) ([]byte, error) {
			return json.MarshalIndent(data, "", "    ")
		},
		jsonToYAML: yaml.JSONToYAML,
	}
}

// Config presents Gen configurations.
type Config struct {
	// SearchDir the swag would be parse
	SearchDir string

	// excludes dirs and files in SearchDir,comma separated
	Excludes string

	// OutputDir represents the output directory for all the generated files
	OutputDir string

	// MainAPIFile the Go file path in which 'swagger general API Info' is written
	MainAPIFile string

	// PropNamingStrategy represents property naming strategy like snakecase,camelcase,pascalcase
	PropNamingStrategy string

	// ParseVendor whether swag should be parse vendor folder
	ParseVendor bool

	// ParseDependencies whether swag should be parse outside dependency folder
	ParseDependency bool

	// ParseInternal whether swag should parse internal packages
	ParseInternal bool

	// MarkdownFilesDir used to find markdownfiles, which can be used for tag descriptions
	MarkdownFilesDir string

	// GeneratedTime whether swag should generate the timestamp at the top of docs.go
	GeneratedTime bool

	// CodeExampleFilesDir used to find code example files, which can be used for x-codeSamples
	CodeExampleFilesDir string

	// ParseDepth dependency parse depth
	ParseDepth int
}

// Build builds swagger json file  for given searchDir and mainAPIFile. Returns json
func (g *Gen) Build(config *Config) error {
	if _, err := os.Stat(config.SearchDir); os.IsNotExist(err) {
		return fmt.Errorf("dir: %s is not exist", config.SearchDir)
	}

	logrus.Println(context.TODO(), "Generate swagger docs....")
	p := swagger.New(swagger.SetMarkdownFileDirectory(config.MarkdownFilesDir),
		swagger.SetExcludedDirsAndFiles(config.Excludes),
		swagger.SetCodeExamplesDirectory(config.CodeExampleFilesDir))
	p.PropNamingStrategy = config.PropNamingStrategy
	p.ParseVendor = config.ParseVendor
	p.ParseDependency = config.ParseDependency
	p.ParseInternal = config.ParseInternal

	if err := p.ParseAPI(config.SearchDir, config.MainAPIFile, config.ParseDepth); err != nil {
		return err
	}
	swagger := p.GetSwagger()

	b, err := g.jsonIndent(swagger)
	if err != nil {
		return err
	}

	if err := os.MkdirAll(config.OutputDir, os.ModePerm); err != nil {
		return err
	}

	// absOutputDir, err := filepath.Abs(config.OutputDir)
	// if err != nil {
	// 	return err
	// }
	// packageName := filepath.Base(absOutputDir)
	// docFileName := filepath.Join(config.OutputDir, "docs.go")
	// jsonFileName := filepath.Join(config.OutputDir, "swagger.json")
	yamlFileName := filepath.Join(config.OutputDir, "swagger.yaml")

	// docs, err := os.Create(docFileName)
	// if err != nil {
	// 	return err
	// }
	// defer docs.Close()

	// err = g.writeFile(b, jsonFileName)
	// if err != nil {
	// 	return err
	// }

	y, err := g.jsonToYAML(b)
	if err != nil {
		return fmt.Errorf("cannot convert json to yaml error: %s", err)
	}

	err = g.writeFile(y, yamlFileName)
	if err != nil {
		return err
	}

	// Write doc
	// err = g.writeGoDoc(packageName, docs, swagger, config)
	// if err != nil {
	// 	return err
	// }

	// log.Printf("create docs.go at %+v", docFileName)
	// logrus.Printf("create swagger.json at %+v", jsonFileName)
	logrus.Printf(context.TODO(), "create swagger.yaml at %+v", yamlFileName)

	return nil
}

func (g *Gen) writeFile(b []byte, file string) error {
	f, err := os.Create(file)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.Write(b)
	return err
}
