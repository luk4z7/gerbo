package movie

import (
	"gerbo/core/drive/mongo"
	"gerbo/services/models"
	"gopkg.in/mgo.v2/bson"
)

var paginationLimit int64 = 20

var session = mongo.GetSession("movies")

// Example using aggregate
// db.movies.aggregate([
// { "$addFields": {
//		"rating_average": { "$avg": "$rating.score" }
// }},
// { "$sort": { "id": -1 } },
// { "$limit": 41 },
// { "$skip": 20 }
// ]);
func GetBestMovies(pagination int64)  ([]models.MoviesResponse, error) {
	result := []models.MoviesResponse{}
	query := []bson.M{{
		"$addFields": bson.M{
			"rating_average": bson.M{
				"$avg": "$rating.score",
			},
		},
	},{
		"$sort": bson.M{
			"id": -1,
		},
	},{
		"$limit": pagination + paginationLimit+1,
	},{
		"$skip": pagination,
	}}
	pipe := session.Pipe(query)
	err := pipe.All(&result)

	return result, err
}