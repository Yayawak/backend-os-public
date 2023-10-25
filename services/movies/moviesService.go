package movies

import (
	"sort"
	"strings"

	// testModel "backend_proj_os/models/test.
	templateError "backend_proj_os/errors"
	movieModel "backend_proj_os/models/movies"
)

// var data = []movieModel.MovieBasicDetailResponse{

// 	movieModel.MovieBasicDetailResponse{
// 		TrailerLink: "",
// 	}
// 	// testModel.TestModelResponse{Code: "20000", Message: "Everything Alright"},
// 	// testModel.TestModelResponse{Code: apiErr.InternalServerError.Code(), Message: apiErr.InternalServerError.EnMessage()},
// }
// var data = movieModel.Sample_data{
// 	// movieModel.Sample_data
// 	for i := 0; i < count; i++ {

// 	}
// }

func GetMovie(id int) (*movieModel.Movie, error) {
	data := movieModel.Sample_data
	for _, datum := range data {
		if datum.ID == id {
			return &datum, nil
		}
	}
	return nil, templateError.DataNotFound
}

func GetMovies() (*[]movieModel.MovieBasicDetailResponse, error) {
	data := movieModel.Sample_data
	size := len(data)

	if size == 0 {
		return nil, templateError.DataNotFound
	}

	response := make([]movieModel.MovieBasicDetailResponse, size)
	for i, datum := range data {
		response[i] = movieModel.MovieBasicDetailResponse{
			ID:           datum.ID,
			Name:         datum.Name,
			Description:  datum.Description,
			Rating:       datum.Rating,
			DirectorName: datum.DirectorName,
			Category:     datum.Category,
			Views:        datum.Views,
			TrailerLink:  datum.TrailerLink,
			MovieImage:   datum.MovieImage,
		}
	}

	return &response, nil
}

func GetMoviesByTitle(name string) ([]movieModel.MovieBasicDetailResponse, error) {
	data := movieModel.Sample_data

	if len(data) == 0 {
		return nil, templateError.DataNotFound
	}

	var response []movieModel.MovieBasicDetailResponse
	for _, datum := range data {
		if datum.Name == name {
			response = append(response, movieModel.MovieBasicDetailResponse{
				ID:           datum.ID,
				Name:         datum.Name,
				Description:  datum.Description,
				Rating:       datum.Rating,
				DirectorName: datum.DirectorName,
				Category:     datum.Category,
				Views:        datum.Views,
				TrailerLink:  datum.TrailerLink,
				MovieImage:   datum.MovieImage,
			})
		}
	}
	return response, nil
}

func GetTopMovies(count int) ([]movieModel.MovieBasicDetailResponse, error) {
	// var tops [count]movieModel.MovieBasicDetailResponse
	tops := make([]movieModel.MovieBasicDetailResponse, count+1)
	var all = movieModel.Sample_data

	// * not good behavior to sort all data without permission
	// sort.SliceStable(all, func(i, j int) bool {
	// 	return all[i].Rating < all[j].Rating
	// })

	// for i := 0; i < count; i++ {

	// movieModel
	var maxRating float32 = -999
	var j int = 0
	for i := 0; i < len(all); i++ {
		// movieModel.Sample_data
		if all[i].Rating > maxRating {
			var cur = movieModel.MovieBasicDetailResponse{
				ID:           all[i].ID,
				Name:         all[i].Name,
				Description:  all[i].Description,
				Rating:       all[i].Rating,
				DirectorName: all[i].DirectorName,
				Category:     all[i].Category,
				Views:        all[i].Views,
				TrailerLink:  all[i].TrailerLink,
				MovieImage:   all[i].MovieImage,
			}

			tops[j] = cur

			j++
			j %= count + 1
		}

	}

	sort.SliceStable(tops, func(i, j int) bool {
		return tops[i].Rating < tops[j].Rating
	})

	// return tops, nill
	return tops, nil
}

func GetMoviesByCategory(categoryName string) ([]movieModel.MovieBasicDetailResponse, error) {
	var response []movieModel.MovieBasicDetailResponse
	var all = movieModel.Sample_data

	for i := 0; i < len(all); i++ {

		var foundCat bool = false
		for j := 0; j < len(all[i].Category); j++ {
			// if strings.ToLower(all[i].Category[j].Name) == strings.ToLower(categoryName) {
			if strings.EqualFold(strings.ToLower(all[i].Category[j].Name),
				strings.ToLower(categoryName)) {
				foundCat = true
			}
		}
		if foundCat {
			var cur = movieModel.MovieBasicDetailResponse{
				ID:           all[i].ID,
				Name:         all[i].Name,
				Description:  all[i].Description,
				Rating:       all[i].Rating,
				DirectorName: all[i].DirectorName,
				Category:     all[i].Category,
				Views:        all[i].Views,
				TrailerLink:  all[i].TrailerLink,
				MovieImage:   all[i].MovieImage,
			}
			response = append(response, cur)
		}

	}

	return response, nil
}

func GetMoviesByDirector(directorName string) ([]movieModel.MovieBasicDetailResponse, error) {
	var response []movieModel.MovieBasicDetailResponse
	var all = movieModel.Sample_data

	for i := 0; i < len(all); i++ {

		var found bool = false
		if strings.EqualFold(strings.ToLower(all[i].DirectorName),
			strings.ToLower(directorName)) {
			found = true
		}

		if found {
			var cur = movieModel.MovieBasicDetailResponse{
				ID:           all[i].ID,
				Name:         all[i].Name,
				Description:  all[i].Description,
				Rating:       all[i].Rating,
				DirectorName: all[i].DirectorName,
				Category:     all[i].Category,
				Views:        all[i].Views,
				TrailerLink:  all[i].TrailerLink,
				MovieImage:   all[i].MovieImage,
			}
			response = append(response, cur)
		}
	}

	return response, nil

}
