package model

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type Recipe struct {
	Name        string
	Ingredients []any
	Seasons     []Season `json:",omitempty"`
	Diets       []Diet   `json:",omitempty"`
	Tags        []Tag    `json:",omitempty"`
	Picture     []byte   `json:",omitempty"`
}

func (r Recipe) Validate() error {
	if r.Name == "" {
		return errors.New("empty recipe name")
	}

	if len(r.Ingredients) == 0 {
		return errors.New("empty ingredient list")
	}

	for i, season := range r.Seasons {
		if !season.IsASeason() {
			return fmt.Errorf("invalid season at index %d: '%d'", i, season)
		}
	}

	for i, diet := range r.Diets {
		if !diet.IsADiet() {
			return fmt.Errorf("invalid diet at index %d: '%d'", i, diet)
		}
	}

	for i, tag := range r.Tags {
		if !tag.IsATag() {
			return fmt.Errorf("invalid tag at index %d: '%d'", i, tag)
		}
	}

	//TODO image validation

	return nil
}

func (r *Recipe) loadPicture(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}

	r.Picture = data
	return nil
}

func (r *Recipe) loadWebPicture(url string) error {
	res, err := http.Get(url)
	if err != nil || res.StatusCode != 200 {
		return err
	}
	defer res.Body.Close()

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	r.Picture = data
	return nil
}
