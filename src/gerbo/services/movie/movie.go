// Gerbo - Rodent and data extractor
// https://github.com/luk4z7/gerbo for the canonical source repository
// Copyright Lucas Alves 2017

package movie

import (
	liberr "gerbo/lib/error"
	"gerbo/core/drive/mongo"
	"gerbo/services/models"
	"database/sql"
	"gopkg.in/mgo.v2/bson"
	"strconv"
)

var session = mongo.GetSession("movies")

type MaxMin struct {
	Max int `bson:"max"`
	Min int `bson:"min"`
}

func Insert(movie models.MoviesResponse) error {
	err := session.Insert(movie)
	if err != nil {
		return err
	}
	return nil
}

func CheckAndGet(db *sql.DB) (*sql.Rows, error) {
	result, err := getMaxValue()
	if err != nil {
		panic(err)
	}
	if len(result) == 0 {
		return &sql.Rows{}, &liberr.Err{Name: "Nothing max value found!"}
	}
	rows, err := Get(db, result[0].Max, ">")
	return rows, err
}

func getMaxValue() ([]MaxMin, error) {
	result := []MaxMin{}
	query := []bson.M{{
		"$group": bson.M{
			"_id": bson.M{},
			"max": bson.M{
				"$max": "$id",
			},
		},
	}}
	pipe := session.Pipe(query)
	err := pipe.All(&result)
	return result, err
}

func Get(db *sql.DB, id int, param string) (*sql.Rows, error) {
	query := `
		SELECT
			mo.id         AS id,
			mo.title      AS filme,
			mo.year       AS ano,
			ge.id         AS genero_id,
			ge.title      AS genero,
			ra.id         AS avaliacao_id,
			ra.feeling    AS avaliacao,
			ra.id         AS avaliacao_nota,
			us.twitter_id AS twitter,
			(
				SELECT COUNT(*) FROM (
					SELECT
						mo.id,
						ge.id,
						ra.id,
						ra.id,
						us.twitter_id
					FROM genres AS ge
					INNER JOIN movie_genre AS mg ON ge.id = mg.genre_id
					INNER JOIN movies AS mo ON mg.movie_id = mo.id
					INNER JOIN movie_ratings AS mr ON mo.id = mr.movie_id
					INNER JOIN users AS us ON mr.user_id = us.id
					INNER JOIN ratings AS ra ON ra.id = mr.rating_id
					WHERE mo.id ` + param + strconv.Itoa(id) + `
					GROUP BY
						mo.id,
						ge.id,
						ra.id,
						ra.id,
						us.twitter_id
					ORDER BY mo.title
				) AS x
			) AS total
		FROM genres AS ge
		INNER JOIN movie_genre AS mg ON ge.id = mg.genre_id
		INNER JOIN movies AS mo ON mg.movie_id = mo.id
		INNER JOIN movie_ratings AS mr ON mo.id = mr.movie_id
		INNER JOIN users AS us ON mr.user_id = us.id
		INNER JOIN ratings AS ra ON ra.id = mr.rating_id
		WHERE mo.id ` + param + strconv.Itoa(id) + `
		GROUP BY
			mo.id,
			mo.title,
			mo.year,
			ge.id,
			ge.title,
			ra.id,
			ra.feeling,
			ra.id,
			us.twitter_id
		ORDER BY
			mo.title
		DESC;
	`
	rows, err := db.Query(query)
	return rows, err
}