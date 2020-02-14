package global

type commands struct {
	Name, Seasonal, URL, Genre              string
	Score, Rank, Episodes, Info, Aired, Top bool // One of these bools must be true
}

type results struct {
	URL                                string
	Seasonal, Genre, Top               []string
	Score, Rank, Episodes, Info, Aired string
}

//type Map2func func(string, string) (string, bool)
type Website interface {
	Search(string, string) ([]string, bool)
	GetScore() string
	GetTop() (string, bool)
	GetInfo() string
	GetAired() string
	GetRank() string
	GetEpisodes() string
	GetSeasonal(string) ([]string, bool)
}

var CommandMap = *new(commands)
var ResultMap = *new(results)
