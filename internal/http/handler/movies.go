package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/wannn28/TASK-MIKTI/internal/http/dto"
	"github.com/wannn28/TASK-MIKTI/internal/service"
	"github.com/wannn28/TASK-MIKTI/pkg/response"
)

type MovieHandler struct {
	movieService service.MovieService
}

func NewMovieHandler(movieService service.MovieService) MovieHandler {
	return MovieHandler{movieService}
}

func (h *MovieHandler) GetMovies(ctx echo.Context) error {
	var req dto.GetAllMovieRequest

	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, err.Error()))
	}
	users, err := h.movieService.GetAll(ctx.Request().Context(), req)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, response.ErrorResponse(http.StatusInternalServerError, err.Error()))
	}

	return ctx.JSON(http.StatusOK, response.SuccessResponse("successfully showing all movies", users))
}

func (h *MovieHandler) GetMovie(ctx echo.Context) error {
	var req dto.GetMovieByIDRequest

	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, err.Error()))
	}

	movie, err := h.movieService.GetByID(ctx.Request().Context(), req.ID)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, response.ErrorResponse(http.StatusInternalServerError, err.Error()))
	}

	return ctx.JSON(http.StatusOK, response.SuccessResponse("successfully showing movie", movie))
}

func (h *MovieHandler) CreateMovie(ctx echo.Context) error {
	var req dto.CreateMovieRequest

	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, err.Error()))
	}

	err := h.movieService.Create(ctx.Request().Context(), req)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, response.ErrorResponse(http.StatusInternalServerError, err.Error()))
	}

	return ctx.JSON(http.StatusOK, response.SuccessResponse("successfully creating movie", nil))
}

func (h *MovieHandler) UpdateMovie(ctx echo.Context) error {
	var req dto.UpdateMovieRequest

	if err := ctx.Bind(&req); err != nil {

		return ctx.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, err.Error()))
	}

	err := h.movieService.Update(ctx.Request().Context(), req)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, response.ErrorResponse(http.StatusInternalServerError, err.Error()))
	}

	return ctx.JSON(http.StatusOK, response.SuccessResponse("successfully update movie", nil))
}

func (h *MovieHandler) DeleteMovie(ctx echo.Context) error {
	var req dto.DeleteMovieRequest

	if err := ctx.Bind(&req); err != nil {

		return ctx.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, err.Error()))
	}

	movie, err := h.movieService.GetByID(ctx.Request().Context(), req.ID)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, response.ErrorResponse(http.StatusInternalServerError, err.Error()))
	}

	err = h.movieService.Delete(ctx.Request().Context(), movie)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, response.ErrorResponse(http.StatusInternalServerError, err.Error()))
	}

	return ctx.JSON(http.StatusOK, response.SuccessResponse("successfully delete movie", nil))
}
