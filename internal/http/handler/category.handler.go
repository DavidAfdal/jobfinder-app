package handler

import (
	"net/http"

	"github.com/DavidAfdal/workfinder/internal/entity"
	"github.com/DavidAfdal/workfinder/internal/http/binder"
	"github.com/DavidAfdal/workfinder/internal/service"
	"github.com/DavidAfdal/workfinder/pkg/response"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)


type CategoryHandler interface {
	FindAllCategory(ctx echo.Context) error
	FindCategoryByID(ctx echo.Context) error
	CreateCategory(ctx echo.Context) error
	UpdateCategory(ctx echo.Context) error
	DeleteCategory(ctx echo.Context) error
}

type categoryHandler struct {
	categoryService service.CategoryService
}


func NewCategoryHandler(categoryService service.CategoryService) CategoryHandler {
	return &categoryHandler{categoryService}
}

// FindAllCategory godoc
// @Summary Get all categories
// @Description Get all categories
// @Tags categories
// @Produce json
// @Success 200 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /categories [get]
func (h *categoryHandler) FindAllCategory(ctx echo.Context) error {
	categories, err := h.categoryService.FindAllCategory()

	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, response.ErrorResponse(http.StatusInternalServerError, err.Error()))
	}

	return ctx.JSON(http.StatusOK, response.SuccessResponse(http.StatusOK, "success get all categories", categories))

}


// FindCategoryByID godoc
// @Summary Get a category by ID
// @Description Get a category by ID
// @Tags categories
// @Produce json
// @Param id path string true "Category ID"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /categories/{id} [get]
func (c *categoryHandler) FindCategoryByID(ctx echo.Context)  error {
    var input binder.FindCategoryByIDRequest

	if err := ctx.Bind(&input); err != nil {
		return ctx.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, err.Error()))
	}

    id := uuid.MustParse(input.ID)
	category, err := c.categoryService.FindCategoryByID(id)

	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, response.ErrorResponse(http.StatusInternalServerError, err.Error()))
	}

	return ctx.JSON(http.StatusOK, response.SuccessResponse(http.StatusOK, "success get category", category))
}

// CreateCategory godoc
// @Summary Create a new category
// @Description Create a new category
// @Tags categories
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param createCategory body binder.CreateCategoryRequest true "Create Category Request"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /categories [post]
func (c *categoryHandler) CreateCategory(ctx echo.Context) error {
	var input binder.CreateCategoryRequest

	if err := ctx.Bind(&input); err != nil {
		return ctx.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, err.Error()))
	}

	newCategory := entity.NewCategory(input.Title, input.Icon)

	category, err := c.categoryService.CreateCategory(newCategory)

	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, response.ErrorResponse(http.StatusInternalServerError, err.Error()))
	}

	return ctx.JSON(http.StatusOK, response.SuccessResponse(http.StatusOK, "success create category", category))
}

// DeleteCategory godoc
// @Summary Delete a category
// @Description Delete a category
// @Tags categories
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "Category ID"
// @Param deleteCategory body binder.DeleteCategoryRequest true "Delete Category Request"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /categories [delete]
func (c *categoryHandler) DeleteCategory(ctx echo.Context) error {
	var input binder.DeleteCategoryRequest

	if err := ctx.Bind(&input); err != nil {
		return ctx.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, err.Error()))
	}

	id := uuid.MustParse(input.ID)

	isDeleted, err := c.categoryService.DeleteCategory(id)

	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, response.ErrorResponse(http.StatusInternalServerError, err.Error()))
	}

	return ctx.JSON(http.StatusOK, response.SuccessResponse(http.StatusOK, "success delete category", isDeleted))
}


// UpdateCategory godoc
// @Summary Update a category
// @Description Update a category
// @Tags categories
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "Category ID"
// @Param updateCategory body binder.UpdateCategoryRequest true "Update Category Request"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /categories/{id} [put]
func (c *categoryHandler) UpdateCategory(ctx echo.Context) error {
	var input binder.UpdateCategoryRequest

	if err := ctx.Bind(&input); err != nil {
		return ctx.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, err.Error()))
	}

	id := uuid.MustParse(input.ID)

	updateCategory := entity.UpdateCategory(id, input.Title, input.Icon)

	updatedCategory, err := c.categoryService.UpdateCategory(updateCategory)

	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, response.ErrorResponse(http.StatusInternalServerError, err.Error()))
	}

	return ctx.JSON(http.StatusOK, response.SuccessResponse(http.StatusOK, "success update category", updatedCategory))
}
