package repository

import (
	"Kitchenizer/model"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
	"time"

	"github.com/rs/zerolog/log"
)

const (
	dumpFolder = "./.ignore/dumps"
	filePrefix = "kitchenizer"
)

func init() {
	if err := os.MkdirAll(dumpFolder, 0777); err != nil {
		panic(err)
	}
}

type Fake struct {
	recipes []model.Recipe
}

func NewFake() *Fake {
	f := Fake{}

	// Load last dump if available
	if err := f.LoadLast(); err != nil {
		panic(err)
	}

	// Store db content in a file at regular interval
	go func() {
		ticker := time.NewTicker(time.Minute)
		defer ticker.Stop()

		for range ticker.C {
			log.Info().Str("package", "fake_repository").Msg("dumping db content")
			if err := f.Dump(); err != nil {
				log.Error().Str("package", "fake_repository").Err(err).Msg("couldn't dump db content")
			}
		}
	}()

	return &f
}

func (f *Fake) Save(recipe model.Recipe) {
	f.recipes = append(f.recipes, recipe)
}

func (f Fake) GetAll() []model.Recipe {
	return f.recipes
}

func (f Fake) Dump() error {
	data, err := json.MarshalIndent(f.recipes, "", "	")
	if err != nil {
		return err
	}

	file := filepath.Join(dumpFolder, fmt.Sprintf("%s.%s.json", filePrefix, time.Now().Format("2006.01.02.15.04.05")))
	return ioutil.WriteFile(file, data, 0644)
}

func (f Fake) LoadLast() error {
	files, err := ioutil.ReadDir(dumpFolder)
	if err != nil {
		return err
	}

	if len(files) == 0 {
		log.Warn().Str("package", "fake_repository").Msg("no file to load")
		return nil
	}

	sort.Slice(files, func(i, j int) bool {
		return files[i].Name() > files[j].Name()
	})

	file, err := os.Open(filepath.Join(dumpFolder, files[0].Name()))
	if err != nil {
		return err
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}

	return json.Unmarshal(data, &f.recipes)
}
