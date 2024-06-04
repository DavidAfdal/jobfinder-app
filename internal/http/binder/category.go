package binder


type CreateCategoryRequest struct {
	Title string `json:"title"`
	Icon string `json:"icon"`
}

type UpdateCategoryRequest struct {
	ID   string `param:"id"`
	Title string `json:"title"`
	Icon string `json:"icon"`
}

type DeleteCategoryRequest struct {
	ID   string `param:"id"`
}
type FindCategoryByIDRequest struct {
	ID   string `param:"id"`
}
