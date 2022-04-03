package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/Netflix-Clone-MicFlix/Movie-Service/internal"
	"github.com/Netflix-Clone-MicFlix/Movie-Service/internal/entity"
	"github.com/Netflix-Clone-MicFlix/Movie-Service/pkg/logger"
)

type MovieRoutes struct {
	t internal.Movie
	l logger.Interface
}

func newMovieRoutes(handler *gin.RouterGroup, t internal.Movie, l logger.Interface) {
	r := &MovieRoutes{t, l}

	h := handler.Group("/movie")
	{
		h.GET("", r.GetAll)
		h.POST("", r.Create)
		h.GET("/:movie_id", r.GetById)

	}
}

type movieCollectionResponse struct {
	Movies []entity.Movie `json:"movies"`
}
type MovieRequest struct {
	Movie entity.Movie `json:"movies"`
}
type CreateMovieRequest struct {
	Title       string `json:"title"`
	Discription string `json:"discription"`
}

// @Summary     Show movies
// @Description Show all movies
// @ID          movie
// @Tags  	    movie
// @Accept      json
// @Produce     json
// @Success     200 {object} movieResponse
// @Failure     500 {object} response
// @Router      /movie [get]
func (r *MovieRoutes) GetAll(c *gin.Context) {
	movies, err := r.t.GetAll(c.Request.Context())
	if err != nil {
		r.l.Error(err, "http - v1 - GetAll")
		errorResponse(c, http.StatusInternalServerError, "database problems")

		return
	}

	c.JSON(http.StatusOK, movieCollectionResponse{movies})
}

// @Summary     Show movie with id
// @Description Show movies with id
// @ID          movie
// @Tags  	    movie
// @Accept      json
// @Produce     json
// @Success     200 {object} movieResponse
// @Failure     500 {object} response
// @Router      /movie [get]
func (r *MovieRoutes) GetById(c *gin.Context) {
	movieId := c.Param("movie_id")

	movie, err := r.t.GetById(c.Request.Context(), movieId)
	if err != nil {
		r.l.Error(err, "http - v1 - doMovie")
		errorResponse(c, http.StatusInternalServerError, "movie service problems")

		return
	}

	c.JSON(http.StatusOK, movie)
}

// @Summary     creates movie
// @Description creates movie with discription and title
// @ID          movie
// @Tags  	    movie
// @Accept      json
// @Produce     json
// @Success     200 {object} CreateMovieRequest
// @Failure     500 {object} response
// @Router      /movie [get]
func (r *MovieRoutes) Create(c *gin.Context) {
	var request CreateMovieRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		r.l.Error(err, "http - v1 - Register user")
		errorResponse(c, http.StatusBadRequest, "invalid request body")

		return
	}

	err := r.t.Create(c.Request.Context(), request.Title, request.Discription)
	if err != nil {
		r.l.Error(err, "http - v1 - doMovie")
		errorResponse(c, http.StatusInternalServerError, "movie service problems")

		return
	}

	c.JSON(http.StatusOK, nil)
}
