package model

import (
	"encoding/json"
	"fmt"
	"time"
)

const (
    Person            = "person"
    Building          = "building"
    Company           = "company"
    Event             = "event"
    Film              = "film"
    Book              = "book"
    Album             = "album"
    Animal            = "animal"
    Award             = "award"
    Song              = "song"
    Country           = "country"
    University        = "university"
    Museum            = "museum"
    PoliticalPosition = "political"
)

var infoBoxFactories map[string]func() InfoBox


type PersonInfoBox struct {
    FullName         *string `json:"full_name"`
    BirthDate        *time.Time `json:"birth_date"`
    PlaceOfBirth     *string `json:"place_of_birth"`
    Nationality      *string `json:"nationality"`
    DeathDate        *time.Time `json:"death_date"`
    Educations        *string `json:"educations"`
    Occupations       *string `json:"occupations"`
    Parents          *string `json:"parents"`
    Children        *string `json:"children"`
}

type BuildingInfoBox struct {
    FullName         *string `json:"full_name"`
    Location         *string `json:"location"`
    Founded          *time.Time `json:"founded"`
    Architects        *string `json:"architects"`
    Height           *int    `json:"height"`
    Floors           *int    `json:"floors"`
    Usage            *string `json:"usage"`
}

type CompanyInfoBox struct {
    FullName         *string `json:"full_name"`
    Founded          *time.Time `json:"founded"`
    Industry         *string `json:"industry"`
    Headquarters     *string `json:"headquarters"`
    CEO              *string `json:"ceo"`
    Revenue          *string `json:"revenue"`
    Employees        *int    `json:"employees"`
}

type EventInfoBox struct {
    Title            *string `json:"title"`
    Date             *time.Time `json:"date"`
    Locations         *string `json:"locations"`
    // Description      *string `json:"description"`
    Organizers        *string `json:"organizers"`
    Attendance       *int    `json:"attendance"`
    // Type             *string `json:"type"`
}

type FilmInfoBox struct {
    Title            *string `json:"title"`
    ReleaseDate      *time.Time `json:"release_date"`
    Director         *string `json:"director"`
    Genre            *string `json:"genre"`
        Duration         *time.Duration    `json:"duration"`
    Languages         *string `json:"language"`
    Rating           *float32 `json:"rating"` // It will be seperate class for rating (Rotten, imdb, ...)
}

type BookInfoBox struct {
    Title            *string `json:"title"`
    Authors           *string `json:"authors"`
    PublishDate      *time.Time `json:"publish_date"`
    Genre            *string `json:"genre"`
    ISBN             *string `json:"isbn"`
    Pages            *int    `json:"pages"`
    Publisher        *string `json:"publisher"`
}

type AlbumInfoBox struct {
    Title            *string   `json:"title"`
    Artists           *string   `json:"artists"`
    ReleaseDate      *time.Time   `json:"release_date"`
    Genre            *string   `json:"genre"`
    Label            *string   `json:"label"`
    Tracks           *string `json:"tracks"`
}

type AnimalInfoBox struct {
    FullName         *string `json:"full_name"`
    Species          *string `json:"species"`
    Habitat          *string `json:"habitat"`
    ConservationStatus *bool `json:"conservation_status"`
    Lifespan         *string `json:"lifespan"`
}

type AwardInfoBox struct {
    FullName         *string `json:"full_name"`
    Year             *time.Time    `json:"year"`
    Category         *string `json:"category"`
    Recipients        *string `json:"recipients"`
    Organizations     *string `json:"organizations"`
}

type SongInfoBox struct {
    Title            *string `json:"title"`
    // In the future, i can added classes: artist, album, etc...
    Artists           *string `json:"artists"`
    Album            *string `json:"album"`
    ReleaseDate      *time.Time `json:"release_date"`
    Genre            *string `json:"genre"`
    Duration         *time.Duration    `json:"duration"`
    Label            *string `json:"label"`
}

type CountryInfoBox struct {
    FullName         *string `json:"full_name"`
    Capital          *string  `json:"capital"`
    Population       *int     `json:"population"`
    Region           *string  `json:"region"`
    Area             *float64 `json:"area"`
    Currency         *string  `json:"currency"`
    OfficialLanguage *string  `json:"official_language"`
}

type UniversityInfoBox struct {
    FullName         *string `json:"full_name"`
    Founded          *time.Time   `json:"founded"`
    Location         *string   `json:"location"`
    Programs         *string `json:"programs"`
    Enrollment       *int      `json:"enrollment"`
    Affiliation      *string   `json:"affiliation"`
    Accreditation    *string   `json:"accreditation"`
}

type MuseumInfoBox struct {
    FullName         *string `json:"full_name"`
    Location         *string   `json:"location"`
    Established      *time.Time   `json:"established"`
    Collections      *string `json:"collections"`
    Director         *string   `json:"director"`
    Exhibitions      *string `json:"exhibitions"`
}

type PoliticalPositionInfoBox struct {
    Position         *string `json:"position"`
    OfficeHolder     *string `json:"office_holder"`
    TermStart        *time.Time `json:"term_start"`
    TermEnd          *time.Time `json:"term_end"`
    Party            *string `json:"party"`
    Constituency     *string `json:"constituency"`
}

func init(){
    if infoBoxFactories == nil{
        infoBoxFactories = map[string]func() InfoBox{
            Person:             func() InfoBox { return &PersonInfoBox{} },
            Building:           func() InfoBox { return &BuildingInfoBox{} },
            Company:            func() InfoBox { return &CompanyInfoBox{} },
            Event:              func() InfoBox { return &EventInfoBox{} },
            Film:               func() InfoBox { return &FilmInfoBox{} },
            Book:               func() InfoBox { return &BookInfoBox{} },
            Album:              func() InfoBox { return &AlbumInfoBox{} },
            Animal:             func() InfoBox { return &AnimalInfoBox{} },
            Award:              func() InfoBox { return &AwardInfoBox{} },
            Song:               func() InfoBox { return &SongInfoBox{} },
            Country:            func() InfoBox { return &CountryInfoBox{} },
            University:         func() InfoBox { return &UniversityInfoBox{} },
            Museum:             func() InfoBox { return &MuseumInfoBox{} },
            PoliticalPosition:  func() InfoBox { return &PoliticalPositionInfoBox{} },
        }
    }
}

func GetInfoBoxFactory(infoBoxType string) (func() InfoBox, error) {
    factory, exists := infoBoxFactories[infoBoxType]
    if !exists {
        return nil, fmt.Errorf("unknown infoBoxType: %s", infoBoxType)
    }
    return factory, nil
}

type InfoBox interface {
    GetType() string
    ToJSON() (string, error)
}


func (p *PersonInfoBox) GetType() string {
    return Person
}

func (p *PersonInfoBox) ToJSON() (string, error) {
    data, err := json.Marshal(p)
    if err != nil {
        return "", err
    }
    return string(data), nil
}


func (b *BuildingInfoBox) GetType() string {
    return Building
}

func (b *BuildingInfoBox) ToJSON() (string, error) {
    data, err := json.Marshal(b)
    if err != nil {
        return "", err
    }
    return string(data), nil
}


func (c *CompanyInfoBox) GetType() string {
    return Company
}

func (c *CompanyInfoBox) ToJSON() (string, error) {
    data, err := json.Marshal(c)
    if err != nil {
        return "", err
    }
    return string(data), nil
}


func (e *EventInfoBox) GetType() string {
    return Event
}

func (e *EventInfoBox) ToJSON() (string, error) {
    data, err := json.Marshal(e)
    if err != nil {
        return "", err
    }
    return string(data), nil
}


func (f *FilmInfoBox) GetType() string {
    return Film
}

func (f *FilmInfoBox) ToJSON() (string, error) {
    data, err := json.Marshal(f)
    if err != nil {
        return "", err
    }
    return string(data), nil
}


func (b *BookInfoBox) GetType() string {
    return Book
}

func (b *BookInfoBox) ToJSON() (string, error) {
    data, err := json.Marshal(b)
    if err != nil {
        return "", err
    }
    return string(data), nil
}


func (a *AlbumInfoBox) GetType() string {
    return Album
}

func (a *AlbumInfoBox) ToJSON() (string, error) {
    data, err := json.Marshal(a)
    if err != nil {
        return "", err
    }
    return string(data), nil
}


func (a *AnimalInfoBox) GetType() string {
    return Animal
}

func (a *AnimalInfoBox) ToJSON() (string, error) {
    data, err := json.Marshal(a)
    if err != nil {
        return "", err
    }
    return string(data), nil
}


func (a *AwardInfoBox) GetType() string {
    return Award
}

func (a *AwardInfoBox) ToJSON() (string, error) {
    data, err := json.Marshal(a)
    if err != nil {
        return "", err
    }
    return string(data), nil
}


func (s *SongInfoBox) GetType() string {
    return Song
}

func (s *SongInfoBox) ToJSON() (string, error) {
    data, err := json.Marshal(s)
    if err != nil {
        return "", err
    }
    return string(data), nil
}


func (c *CountryInfoBox) GetType() string {
    return Country
}

func (c *CountryInfoBox) ToJSON() (string, error) {
    data, err := json.Marshal(c)
    if err != nil {
        return "", err
    }
    return string(data), nil
}


func (u *UniversityInfoBox) GetType() string {
    return University
}

func (u *UniversityInfoBox) ToJSON() (string, error) {
    data, err := json.Marshal(u)
    if err != nil {
        return "", err
    }
    return string(data), nil
}


func (m *MuseumInfoBox) GetType() string {
    return Museum
}

func (m *MuseumInfoBox) ToJSON() (string, error) {
    data, err := json.Marshal(m)
    if err != nil {
        return "", err
    }
    return string(data), nil
}

func (p *PoliticalPositionInfoBox) GetType() string {
    return PoliticalPosition
}

func (p *PoliticalPositionInfoBox) ToJSON() (string, error) {
    data, err := json.Marshal(p)
    if err != nil {
        return "", err
    }
    return string(data), nil
}

