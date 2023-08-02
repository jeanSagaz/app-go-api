package entity

type Album struct {
	Id     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

func AlbumList() []Album {
	return []Album{
		{Id: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
		{Id: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
		{Id: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
	}
}
