package movie

import (
	"gerbo/core/drive/mongo"
	"gerbo/services/models"
	"gopkg.in/mgo.v2/bson"
)

var paginationLimit int64 = 20

var session = mongo.GetSession("movies")

type DistributionPerDecade struct {
	Decades []DataYears `json:"decades"`
}

type DataYears struct {
	Years []int `json:"years"`
	Count int   `json:"total_distribution"`
}

// 	db.movies.aggregate([
// 	{
// 		"$addFields": {
//			"rating_average": {
// 				"$avg": "$rating.score"
// 			}
// 		}},
// 		{ "$sort": { "id": -1 } },
// 		{ "$limit": 41 },
// 		{ "$skip": 20 }
//	]);
func GetBestMovies(pagination int64) ([]models.MoviesResponse, error) {
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

//	db.movies.aggregate([
// 	{ "$unwind": "$genre" },
// 	{
// 		"$addFields": {
// 			"rating_average": {
// 				"$avg": "$rating.score"
// 			}
// 		}},
// 		{ "$sort": { "id": -1 } },
// 		{
// 			"$project": {
//				"genre": "$genre"
// 			}},
// 			{ $group : { _id : { "genre" : "$genre" }
// 		}}
//	]);
func GetBestMoviesGenres() ([]bson.M, error) {
	resp := []bson.M{}
	query := []bson.M{
		{ "$unwind": "$genre" },
		{ "$addFields": bson.M{
				"rating_average": bson.M{ "$avg": "$rating.score" },
		}},
		{ "$project": bson.M{
				"genre": "$genre",
		}},
		{ "$group": bson.M{ "_id": bson.M{ "genre": "$genre" } } },
	}
	pipe := session.Pipe(query)
	err := pipe.All(&resp)
	return resp, err
}

//	db.movies.aggregate([
//		{"$unwind": "$genre" },
//		{
// 			$group : {
//				_id: { genge: "$genre" },
//				occurrence: { $sum: 1 }
//			}
//		}
//	]);
func GetMoreMoviesGenre() ([]bson.M, error) {
	resp := []bson.M{}
	query := []bson.M{
		{ "$unwind": "$genre" },
		{ "$group": bson.M{ "_id": bson.M{ "genre": "$genre" }, "occurrence": bson.M{ "$sum": 1 } } },
	}
	pipe := session.Pipe(query)
	err := pipe.All(&resp)
	return resp, err
}

//	db.movies.aggregate([
//	     { "$unwind": "$rating" },
//	     {
//	         "$group": {
//	             "_id": {
//	                 "_id": "$genre"
//	             },
//	             "rating_average": {
//	                 "$avg": "$rating.score"
//	             }
//	         }
//	     },
//	     {
//	         "$project": {
//	             "_id": "$_id._id",
//	             "genre": "$genre",
//	             "rating_average": 1,
//	             "rating": 1
//	         }
//	     }
//	 ]);
func GetMoviesRatingPerGenre() ([]bson.M, error) {
	resp := []bson.M{}
	query := []bson.M{
		{ "$unwind": "$rating" },
		{
			"$group": bson.M{
				"_id": bson.M{
					"_id": "$genre",
				},
				"rating_average": bson.M{
					"$avg": "$rating.score",
				},
			},
		},
		{
			"$project": bson.M{
				"_id": "$_id._id",
				"genre": "$genre",
				"rating_average": 1,
				"rating": 1,
			},
		},
	}
	pipe := session.Pipe(query)
	err := pipe.All(&resp)
	return resp, err
}

//	db.movies.aggregate([
//	     { "$unwind": "$rating" },
//	     {
//	         "$group": {
//	             "_id": {
//	                 "year": "$year"
//	             },
//	             "rating_average": {
//	                 "$avg": "$rating.score"
//	             }
//	         }
//	     },
//	     {
//	         "$project": {
//	             "_id": "$_id",
//	             "genre": "$genre",
//	             "rating_average": 1,
//	             "rating": 1
//	         }
//	     }
//	 ]);
func GetMoviesRatingPerYear() ([]bson.M, error) {
	resp := []bson.M{}
	query := []bson.M{
		{ "$unwind": "$rating" },
		{
			"$group": bson.M{
				"_id": bson.M{
					"year": "$year",
				},
				"rating_average": bson.M{
					"$avg": "$rating.score",
				},
			},
		},
		{
			"$project": bson.M{
				"_id": "$_id",
				"genre": "$genre",
				"rating_average": 1,
				"rating": 1,
			},
		},
	}
	pipe := session.Pipe(query)
	err := pipe.All(&resp)
	return resp, err
}

//	db.movies.aggregate(
//	    {
// 			$group : {
// 				_id : {
// 					"year": '$year'
// 			},
// 			count : {
// 				$sum : 1
// 			}
// 		}
// 	},
//	{
//		"$project": {
//	    	"count": 1
//	    }
//	});
func GetMoviesDistributionPerYear() ([]bson.M, error) {
	resp := []bson.M{}
	query := []bson.M{
		{
			"$group": bson.M{
				"_id": bson.M{
					"year": "$year",
				},
				"count": bson.M{
					"$sum": 1,
				},
			},
		},
		{
			"$project": bson.M{
				"count": 1,
		},
	}}
	pipe := session.Pipe(query)
	err := pipe.All(&resp)
	return resp, err
}

//	db.movies.aggregate([
//	     {
//	         "$group": {
//	             "_id": "$year",
//	             "count": { "$sum": 1 }
//	          }
//	     },
//	     {
//	         "$project": {
//	             "Year": "$_id",
//	             _id: 0,
//	             "count": 1
//	         }
//	     },
//	     { $sort: {count: 1} }
//	]);
func GetMoviesDistributionPerDecade() (DistributionPerDecade, error) {
	resp := []bson.M{}
	query := []bson.M{
		{
			"$group": bson.M{
				"_id": "$year",
				"count": bson.M{
					"$sum": 1,
				},
			},
		},
		{
			"$project": bson.M{
				"year": "$_id",
				"_id": 0,
				"count": 1,
			},
		},
	}
	pipe := session.Pipe(query)
	err := pipe.All(&resp)

	yearPerDecade := DistributionPerDecade{}
	dataYears := DataYears{}
  	years := []int{}

	contador := 0
	totalPerDecade := 0
	for i := range resp {
			totalPerDecade += resp[i]["count"].(int)
			years = append(years, resp[i]["year"].(int))
		contador++
		if contador == 10 {
			dataYears.Years = years
			dataYears.Count = totalPerDecade
			yearPerDecade.Decades = append(yearPerDecade.Decades, dataYears)

			dataYears.Count = 0
			totalPerDecade = 0
			dataYears.Years = nil
			years = nil
			contador = 0
		}
	}
	return yearPerDecade, err
}