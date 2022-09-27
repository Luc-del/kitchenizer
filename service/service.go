package service

import "Kitchenizer/model"

type Service struct {
	Repo recipeRepository
}

func (s Service) InsertRecipe(recipe model.Recipe) {
	s.Repo.Save(recipe)
}

func (s Service) LoadRecipe() {

}
