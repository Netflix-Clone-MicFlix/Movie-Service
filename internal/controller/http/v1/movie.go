package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/Netflix-Clone-MicFlix/Movie-Service/internal"
	"github.com/Netflix-Clone-MicFlix/Movie-Service/internal/entity"
	"github.com/Netflix-Clone-MicFlix/Movie-Service/pkg/logger"
)

type MovieRoutes struct {
	t          internal.Movie
	l          logger.Interface
	corsConfig gin.HandlerFunc
}

func newMovieRoutes(handler *gin.RouterGroup, t internal.Movie, l logger.Interface, corsConfig gin.HandlerFunc) {
	r := &MovieRoutes{t, l, corsConfig}

	movie := handler.Group("/movie")
	{
		movie.Use(corsConfig)
		movie.GET("", r.GetAll)
		movie.POST("", r.Create)
		movie.GET("/:movie_id", r.GetById)
		genres := movie.Group("/genre")
		{
			genres.Use(corsConfig)
			genres.POST("", r.AddGenre)
			genres.GET("", r.GetAllGenre)
			genres.GET("/:genre_id", r.GetGenreById)
		}
	}

}

type movieCollectionResponse struct {
	Movies []entity.Movie `json:"movies"`
}
type CreateMovieRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

type GenreCollectionResponse struct {
	Genre []entity.Genre `json:"genre"`
}

type GenreRequest struct {
	Id   string `json:"id"    example:"6be244a7-25ac-34ce-31e3-04157d3d42e3"`
	Name string `json:"name"  example:"Horror"`
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
// @Router      /movie [post]
func (r *MovieRoutes) Create(c *gin.Context) {
	var request CreateMovieRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		r.l.Error(err, "http - v1 - Register user")
		errorResponse(c, http.StatusBadRequest, "invalid request body")

		return
	}

	err := r.t.Create(c.Request.Context(), request.Title, request.Description)
	if err != nil {
		r.l.Error(err, "http - v1 - doMovie")
		errorResponse(c, http.StatusInternalServerError, "movie service problems")

		return
	}

	c.JSON(http.StatusOK, nil)
}

// @Summary     creates genre
// @Description creates genre with discription and title
// @ID          genre
// @Tags  	    genre
// @Accept      json
// @Produce     json
// @Success     200 {object} CreateMovieRequest
// @Failure     500 {object} response
// @Router      /movie/genre [post]
func (r *MovieRoutes) AddGenre(c *gin.Context) {
	var request GenreRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		r.l.Error(err, "http - v1 - Register user")
		errorResponse(c, http.StatusBadRequest, "invalid request body")

		return
	}

	err := r.t.AddGenre(c.Request.Context(), request.Name)
	if err != nil {
		r.l.Error(err, "http - v1 - doMovie")
		errorResponse(c, http.StatusInternalServerError, "movie service problems")

		return
	}

	c.JSON(http.StatusOK, nil)
}

// @Summary     Show genre
// @Description Show all genre
// @ID          genre
// @Tags  	    genre
// @Accept      json
// @Produce     json
// @Success     200 {object} genreResponse
// @Failure     500 {object} response
// @Router      /movie/genre  [get]
func (r *MovieRoutes) GetAllGenre(c *gin.Context) {
	genres, err := r.t.GetAllGenre(c.Request.Context())
	if err != nil {
		r.l.Error(err, "http - v1 - GetAll")
		errorResponse(c, http.StatusInternalServerError, "database problems")

		return
	}

	c.JSON(http.StatusOK, GenreCollectionResponse{genres})
}

// @Summary     Show genre with id
// @Description Show genres with id
// @ID          genre
// @Tags  	    genre
// @Accept      json
// @Produce     json
// @Success     200 {object} genreResponse
// @Failure     500 {object} response
// @Router      /movie/genre  [get]
func (r *MovieRoutes) GetGenreById(c *gin.Context) {
	genreId := c.Param("genre_id")

	genre, err := r.t.GetGenreById(c.Request.Context(), genreId)
	if err != nil {
		r.l.Error(err, "http - v1 - doMovie")
		errorResponse(c, http.StatusInternalServerError, "movie service problems")
		return
	}

	c.JSON(http.StatusOK, genre)
}
