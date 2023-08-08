package repositories

import (
	"biFebriansyah/gogin/config"
	"biFebriansyah/gogin/internal/models"
	"fmt"
	"log"
	"math"

	"github.com/jmoiron/sqlx"
)

type RepoMovie struct {
	*sqlx.DB
}

func NewMovie(db *sqlx.DB) *RepoMovie {
	return &RepoMovie{db}
}

func (r *RepoMovie) CreateMovie(data *models.Movie) (string, error) {
	q := `INSERT INTO tiketz.movie(
		movie_name,
		movie_banner,
		release_date
	) VALUES(
		:movie_name, 
		:movie_banner, 
		:release_date
	)
	`

	_, err := r.NamedExec(q, data)
	if err != nil {
		log.Println(err)
		return "", err
	}

	return "1 data movie created", nil
}

func (r *RepoMovie) UpdateMovie(data *models.Movie) (string, error) {

	q := `UPDATE tiketz.movie SET
		movie_name = COALESCE(NULLIF(:movie_name, ''), movie_name),
		movie_banner = COALESCE(NULLIF(:movie_banner, ''), movie_banner),
		release_date = COALESCE(NULLIF(:release_date, CURRENT_DATE), release_date),
		updated_at = now()
	WHERE movie_id = :movie_id
	`

	_, err := r.NamedExec(q, data)
	if err != nil {
		log.Println(err)
		return "", err
	}

	return "1 data movie updated", nil
}

func (r *RepoMovie) RemoveMovie(idMovie string) (string, error) {
	q := `DELETE FROM tiketz.movie WHERE movie_id = ?`

	if _, err := r.Query(r.Rebind(q), idMovie); err != nil {
		return "", err
	}

	return "1 data movie deleted", nil
}

func (r *RepoMovie) GetAllMovie() (*config.Result, error) {
	q := `SELECT * FROM tiketz.movie ORDER BY created_at DESC`
	data := models.Movies{}

	if err := r.Select(&data, q); err != nil {
		return nil, err
	}

	return &config.Result{Data: data}, nil
}

func (r *RepoMovie) GetMovie(params models.Meta) (*config.Result, error) {
	var data models.Movies
	var metas config.Metas
	var filterQuery string
	var metaQuery string
	var count int
	var args []interface{}
	var filter []interface{}

	if params.Name != "" {
		filterQuery = "AND movie_name = ?"
		args = append(args, params.Name)
		filter = append(filter, params.Name)
	}

	offset := (params.Page - 1) * params.Limit
	metaQuery = "LIMIT ? OFFSET ? "
	args = append(args, params.Limit, offset)

	m := fmt.Sprintf(`SELECT COUNT(movie_id) as count FROM tiketz.movie WHERE true %s`, filterQuery)
	err := r.Get(&count, r.Rebind(m), filter...)
	if err != nil {
		return nil, err
	}

	q := fmt.Sprintf(`
	SELECT 
		movie_id, 
		movie_name, 
		movie_banner, 
		release_date, 
		created_at, 
		updated_at
	FROM tiketz.movie
	WHERE true %s
	%s
	`, filterQuery, metaQuery)

	err = r.Select(&data, r.Rebind(q), args...)
	if err != nil {
		return nil, err
	}

	check := math.Ceil(float64(count) / float64(params.Limit))
	metas.Total = count
	if count > 0 && params.Page != int(check) {
		metas.Next = params.Page + 1
	}

	if params.Page != 1 {
		metas.Prev = params.Page - 1
	}

	return &config.Result{Data: data, Meta: metas}, nil
}
