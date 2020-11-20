package gunplas

import (
	"log"

	"github.com/slowmanchan/grapqlGunpla/internal/pkg/db/postgres"
)

type Gunpla struct {
	ID              string `json:"id"`
	BoxArtImageLink string `json:"boxArtImageLink" db:"box_art_image_link"`
	Title           string `json:"title"`
	Subtitle        string `json:"subtitle"`
	Classification  string `json:"classification"`
	LineupNo        string `json:"lineupNo" db:"lineup_no"`
	Scale           string `json:"scale"`
	Franchise       string `json:"franchise"`
	ReleaseDate     string `json:"releaseDate" db:"release_date"`
	JanIsbn         string `json:"janIsbn" db:"jan_isbn"`
	Run             string `json:"run"`
	Price           string `json:"price"`
	Includes        string `json:"includes"`
	Features        string `json:"features"`
}

func GetAll() []Gunpla {
	var gunplas []Gunpla
	if err := postgres.Db.Select(&gunplas, "SELECT * FROM gunpla"); err != nil {
		log.Fatal(err)
	}
	return gunplas
}
