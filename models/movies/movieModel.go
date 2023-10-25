package movies

type Movie struct {
	ID           int             `json:"id"`
	Name         string          `json:"name"`
	Description  string          `json:"description"`
	Rating       float32         `json:"rating"`
	DirectorName string          `json:"directorName"`
	TrailerLink  string          `json:"trailerLink"`
	Category     []MovieCategory `json:"category"`
	Views        int             `json:"views"`
	MovieSeasons []MovieSeasons  `json:"movieSeasons"`
	CreatedBy    string          `json:"createdBy"`
	CreatedAt    string          `json:"createdAt"`
	UpdatedBy    *string         `json:"updatedBy"`
	UpdatedAt    *string         `json:"updatedAt"`
}

type MovieCategory struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type MovieSeasons struct {
	SeasonID     int           `json:"id"`
	MovieDetails []MovieDetail `json:"movieDetails"`
}

type MovieDetail struct {
	EpisodeID   int    `json:"epdisodeId"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Video_link  string `json:"video_link"`
}

type MovieBasicDetailResponse struct {
	ID           int             `json:"id"`
	Name         string          `json:"name"`
	Description  string          `json:"description"`
	Rating       float32         `json:"rating"`
	DirectorName string          `json:"directorName"`
	Category     []MovieCategory `json:"Category"`
	Views        int             `json:"views"`
	TrailerLink  string          `json:"trailerLink"`
}

var Sample_data = []Movie{
	Movie{
		ID:           1,
		Name:         "The Batman",
		Description:  "When a sadistic serial killer begins murdering key political figures in Gotham, Batman is forced to investigate the city's hidden corruption and question his family's involvement.",
		Rating:       7.8,
		DirectorName: "santana",
		Category: []MovieCategory{
			MovieCategory{
				Name:        "Action",
				Description: "",
			},
			MovieCategory{
				Name:        "Crime",
				Description: "",
			},
			MovieCategory{
				Name:        "Drama",
				Description: "",
			},
		},
		TrailerLink: "https://www.imdb.com/video/vi3215114777/?playlistId=tt1877830&ref_=tt_ov_vi",
		Views:       9,
		// MovieSeasons: []MovieSeasons{},
		CreatedBy: "Matt Reeves",
		CreatedAt: "02/15/14",
		UpdatedBy: nil,
		UpdatedAt: nil,
	},
	Movie{
		ID:           2,
		Name:         "Joker",
		Description:  "During the 1980s, a failed stand-up comedian is driven insane and turns to a life of crime and chaos in Gotham City while becoming an infamous psychopathic crime figure.",
		Rating:       8.4,
		DirectorName: "Todd Phillips",
		Category: []MovieCategory{
			MovieCategory{
				Name:        "Action",
				Description: "",
			},
			MovieCategory{
				Name:        "Crime",
				Description: "",
			},
			MovieCategory{
				Name:        "Drama",
				Description: "",
			},
		},
		TrailerLink: "https://www.imdb.com/video/vi1723318041/?playlistId=tt7286456&ref_=tt_pr_ov_vi",
		Views:       20,
		// MovieSeasons: []MovieSeasons{
		// 	MovieSeasons{
		// 		SeasonID: 1,
		// 		MovieDetails: []MovieDetail{
		// 			MovieDetail{
		// 				EpisodeID:   1,
		// 				Title:       "",
		// 				Description: "",
		// 				Video_link:  "",
		// 			},
		// 		},
		// 	},
		// },
		CreatedBy: "Todd Phillips",
		CreatedAt: "",
		UpdatedBy: nil,
		UpdatedAt: nil,
	},
	Movie{
		ID:           3,
		Name:         "Game of Thrones",
		Description:  "Nine noble families fight for control over the lands of Westeros, while an ancient enemy returns after being dormant for a millennia.",
		Rating:       9.8,
		DirectorName: "apex",
		Category: []MovieCategory{
			MovieCategory{
				Name:        "Action",
				Description: "",
			},
			MovieCategory{
				Name:        "Crime",
				Description: "",
			},
			MovieCategory{
				Name:        "Drama",
				Description: "",
			},
		},
		TrailerLink: "https://www.imdb.com/video/vi59490329/?playlistId=tt0944947&ref_=tt_ov_vi",
		Views:       25,
		MovieSeasons: []MovieSeasons{
			MovieSeasons{
				SeasonID: 1,
				MovieDetails: []MovieDetail{
					MovieDetail{
						EpisodeID:   1,
						Title:       "S1.E1 ∙ Winter Is Coming",
						Description: "Eddard Stark is torn between his family and an old friend when asked to serve at the side of King Robert Baratheon; Viserys plans to wed his sister to a nomadic warlord in exchange for an army.",
						Video_link:  "https://www.youtube.com/watch?v=TZE9gVF1QbA&pp=ygUOZ2FtZSBvZiB0cmhvbmU%3D",
					},
					MovieDetail{
						EpisodeID:   2,
						Title:       "S1.E2 ∙ The Kingsroad",
						Description: "While Bran recovers from his fall, Ned takes only his daughters to King's Landing. Jon Snow goes with his uncle Benjen to the Wall. Tyrion joins them.",
						Video_link:  "https://www.youtube.com/watch?v=TZE9gVF1QbA&pp=ygUOZ2FtZSBvZiB0cmhvbmU%3D",
					},
				},
			},
			MovieSeasons{
				SeasonID: 2,
				MovieDetails: []MovieDetail{
					MovieDetail{
						EpisodeID:   1,
						Title:       "S2.E1 ∙ The North Remembers",
						Description: "Tyrion arrives at King's Landing to take his father's place as Hand of the King. Stannis Baratheon plans to take the Iron Throne for his own. Robb tries to decide his next move in the war. The Night's Watch arrive at the house of Craster.",
						Video_link:  "https://www.youtube.com/watch?v=TZE9gVF1QbA&pp=ygUOZ2FtZSBvZiB0cmhvbmU%3D",
					},
				},
			},
		},
		CreatedBy: "David Benioff & D.B. Weiss",
		CreatedAt: "",
		UpdatedBy: nil,
		UpdatedAt: nil,
	},

	Movie{
		ID:           4,
		Name:         "Your Name",
		Description:  "Two teenagers share a profound, magical connection upon discovering they are swapping bodies. Things manage to become even more complicated when the boy and girl decide to meet in person.",
		Rating:       3,
		DirectorName: "Makoto Shinkai",
		Category: []MovieCategory{
			MovieCategory{
				Name:        "Animation",
				Description: "",
			},
			MovieCategory{
				Name:        "Drama",
				Description: "",
			},
			MovieCategory{
				Name:        "Fantasy",
				Description: "",
			},
		},
		TrailerLink: "https://www.imdb.com/video/vi1705097753/?playlistId=tt5311514&ref_=tt_ov_vi",
		Views:       9,
		CreatedBy:   "Makoto Shinkai",
		CreatedAt:   "",
		UpdatedBy:   nil,
		UpdatedAt:   nil,
	},
	Movie{
		ID:           5,
		Name:         "The Continental",
		Description:  "Set in 1970s New York City, The Continental explores the origin of the iconic hotel-for-assassins centerpiece of the John Wick universe seen through the eyes and action of a young Winston Scott.",
		Rating:       7.3,
		DirectorName: "Kirk Ward",
		Category: []MovieCategory{
			MovieCategory{
				Name:        "Action",
				Description: "",
			},
			MovieCategory{
				Name:        "Crime",
				Description: "",
			},
			MovieCategory{
				Name:        "Triller",
				Description: "",
			},
		},
		TrailerLink: "https://www.imdb.com/video/vi4002006809/?playlistId=tt6486762&ref_=tt_ov_vi",
		Views:       4,
		CreatedBy:   "Kirk Ward",
		CreatedAt:   "",
		UpdatedBy:   nil,
		UpdatedAt:   nil,
	},
}
