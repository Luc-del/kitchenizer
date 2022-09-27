package service

import "Kitchenizer/model"

type recipeRepository interface {
	Save(recipe model.Recipe)
	GetAll() []model.Recipe
}
