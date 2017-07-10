// Gerbo - Rodent and data extractor
// https://github.com/luk4z7/gerbo for the canonical source repository
// Copyright Lucas Alves 2017

package operation

import (
	"gerbo/services/movie"
	"gerbo/lib/validation"
	"gerbo/lib/logs"
	"gerbo/services/models"
	"database/sql"
	"log"
	"sync"
	"strconv"
)

func CheckSync(db *sql.DB, mu *sync.Mutex) {
	 mu.Lock()
	 defer mu.Unlock()

	rows, err := movie.CheckAndGet(db)
	if err != nil {
		logs.INFO.Println(err.Error())
		return
	}
	err = Sync(rows)
	if err != nil {
		panic(err)
	}
}

func Sync(rows *sql.Rows) error {

	var arrayMovie  = []int{}
	var arrayRating = []string{}
	var arrayGenre  = []string{}

	var lastMovie int
	var counter   int =  0
	var total     int

	var moRes  = models.MoviesResponse{}
	var genre  = models.Genre{}
	var rating = models.Rating{}

	moviesReq := new(models.MovieRequest)
	done      := make(chan struct{})

	for rows.Next() {
		if counter == 0 {
			logs.INFO.Println("Synchronizing...")
		}
		err := rows.Scan(
			&moviesReq.ID,
			&moviesReq.Filme,
			&moviesReq.Ano,
			&moviesReq.GeneroID,
			&moviesReq.Genero,
			&moviesReq.AvaliacaoID,
			&moviesReq.Avaliacao,
			&moviesReq.AvaliacaoNota,
			&moviesReq.TwitterID,
			&total,
		)
		if err != nil {
			log.Fatalln("Could not scan row:", err)
			return err
		}
		idMovie := strconv.Itoa(moviesReq.ID)
		idRating := strconv.Itoa(moviesReq.TwitterID)
		idGenre := strconv.Itoa(moviesReq.GeneroID)
		idMovieRating := idMovie + idRating
		idMovieGenre := idMovie + idGenre

		existMovie, _ := validation.InArrayInteger(moviesReq.ID, arrayMovie)
		if !existMovie {
			if lastMovie != 0 && lastMovie != moviesReq.ID {
				go func() {
					logs.INFO.Println("movie -> ", moviesReq.ID)
					write(moRes)
					done <- struct{}{}
				}()
				<-done
			}
			lastMovie = moviesReq.ID
			// add the movie in the slice
			arrayMovie = append(arrayMovie, moviesReq.ID)
			// set the variable movie for the struct
			moRes.ID = moviesReq.ID
			moRes.Title = moviesReq.Filme
			moRes.Year = moviesReq.Ano
			// clear the Genre and Rating slices
			moRes.Genre = nil
			moRes.Rating = nil
		}
		existRating, _ := validation.InArrayString(idMovieRating, arrayRating)
		if !existRating {
			arrayRating = append(arrayRating, idMovieRating)
			rating.ID = moviesReq.AvaliacaoID
			rating.Feeling = moviesReq.Avaliacao
			rating.Score = moviesReq.AvaliacaoNota
			rating.User.TwitterID = moviesReq.TwitterID
			moRes.Rating = append(moRes.Rating, rating)
		}
		existGenre, _ := validation.InArrayString(idMovieGenre, arrayGenre)
		if !existGenre {
			arrayGenre = append(arrayGenre, idMovieGenre)
			genre.ID = moviesReq.GeneroID
			genre.Title = moviesReq.Genero
			moRes.Genre = append(moRes.Genre, genre)
		}

		counter++
		if total == counter {
			go func() {
				logs.INFO.Println("movie -> ", moviesReq.ID)
				write(moRes)
				done <- struct{}{}
			}()
			<-done
		}
	}

	if counter == 0 {
		logs.INFO.Println("Nothing to sync!")
	}
	rows.Close()

	return nil
}

func write(movies models.MoviesResponse) error {
	err := movie.Insert(movies)
	if err != nil {
		return err
	}
	return nil
}