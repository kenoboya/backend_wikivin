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
    FullName         *string    `json:"full_name" db:"full_name"`
    BirthDate        *time.Time `json:"birth_date" db:"birth_date"`
    PlaceOfBirth     *string    `json:"place_of_birth" db:"place_of_birth"`
    Nationality      *string    `json:"nationality" db:"nationality"`
    DeathDate        *time.Time `json:"death_date" db:"death_date"`
    Educations       *string    `json:"educations" db:"educations"`
    Occupations      *string    `json:"occupations" db:"occupations"`
    Parents          *string    `json:"parents" db:"parents"`
    Children         *string    `json:"children" db:"children"`
}

type BuildingInfoBox struct {
    FullName         *string    `json:"full_name" db:"full_name"`
    Location         *string    `json:"location" db:"location"`
    Founded          *time.Time `json:"founded" db:"founded"`
    Architects       *string    `json:"architects" db:"architects"`
    Height           *int       `json:"height" db:"height"`
    Floors           *int       `json:"floors" db:"floors"`
    Usage            *string    `json:"usage" db:"usage"`
}

type CompanyInfoBox struct {
    FullName         *string    `json:"full_name" db:"full_name"`
    Founded          *time.Time `json:"founded" db:"founded"`
    Industry         *string    `json:"industry" db:"industry"`
    Headquarters     *string    `json:"headquarters" db:"headquarters"`
    CEO              *string    `json:"ceo" db:"ceo"`
    Revenue          *string    `json:"revenue" db:"revenue"`
    Employees        *int       `json:"employees" db:"employees"`
}

type EventInfoBox struct {
    Title            *string    `json:"title" db:"title"`
    Date             *time.Time `json:"date" db:"date"`
    Locations        *string    `json:"locations" db:"locations"`
    Organizers       *string    `json:"organizers" db:"organizers"`
    Attendance       *int       `json:"attendance" db:"attendance"`
}

type FilmInfoBox struct {
    Title            *string       `json:"title" db:"title"`
    ReleaseDate      *time.Time    `json:"release_date" db:"release_date"`
    Director         *string       `json:"director" db:"director"`
    Genre            *string       `json:"genre" db:"genre"`
    Duration         *time.Duration `json:"duration" db:"duration"`
    Languages        *string       `json:"language" db:"language"`
    Rating           *float32      `json:"rating" db:"rating"`
}

type BookInfoBox struct {
    Title            *string    `json:"title" db:"title"`
    Authors          *string    `json:"authors" db:"authors"`
    PublishDate      *time.Time `json:"publish_date" db:"publish_date"`
    Genre            *string    `json:"genre" db:"genre"`
    ISBN             *string    `json:"isbn" db:"isbn"`
    Pages            *int       `json:"pages" db:"pages"`
    Publisher        *string    `json:"publisher" db:"publisher"`
}

type AlbumInfoBox struct {
    Title            *string    `json:"title" db:"title"`
    Artists          *string    `json:"artists" db:"artists"`
    ReleaseDate      *time.Time `json:"release_date" db:"release_date"`
    Genre            *string    `json:"genre" db:"genre"`
    Label            *string    `json:"label" db:"label"`
    Tracks           *string    `json:"tracks" db:"tracks"`
}

type AnimalInfoBox struct {
    FullName         *string `json:"full_name" db:"full_name"`
    Species          *string `json:"species" db:"species"`
    Habitat          *string `json:"habitat" db:"habitat"`
    ConservationStatus *bool  `json:"conservation_status" db:"conservation_status"`
    Lifespan         *string `json:"lifespan" db:"lifespan"`
}

type AwardInfoBox struct {
    FullName         *string    `json:"full_name" db:"full_name"`
    Year             *time.Time `json:"year" db:"year"`
    Category         *string    `json:"category" db:"category"`
    Recipients       *string    `json:"recipients" db:"recipients"`
    Organizations    *string    `json:"organizations" db:"organizations"`
}

type SongInfoBox struct {
    Title            *string       `json:"title" db:"title"`
    Artists          *string       `json:"artists" db:"artists"`
    Album            *string       `json:"album" db:"album"`
    ReleaseDate      *time.Time    `json:"release_date" db:"release_date"`
    Genre            *string       `json:"genre" db:"genre"`
    Duration         *time.Duration `json:"duration" db:"duration"`
    Label            *string       `json:"label" db:"label"`
}

type CountryInfoBox struct {
    FullName         *string    `json:"full_name" db:"full_name"`
    Capital          *string    `json:"capital" db:"capital"`
    Population       *int       `json:"population" db:"population"`
    Region           *string    `json:"region" db:"region"`
    Area             *float64   `json:"area" db:"area"`
    Currency         *string    `json:"currency" db:"currency"`
    OfficialLanguage *string    `json:"official_language" db:"official_language"`
}

type UniversityInfoBox struct {
    FullName         *string    `json:"full_name" db:"full_name"`
    Founded          *time.Time `json:"founded" db:"founded"`
    Location         *string    `json:"location" db:"location"`
    Programs         *string    `json:"programs" db:"programs"`
    Enrollment       *int       `json:"enrollment" db:"enrollment"`
    Affiliation      *string    `json:"affiliation" db:"affiliation"`
    Accreditation    *string    `json:"accreditation" db:"accreditation"`
}

type MuseumInfoBox struct {
    FullName         *string    `json:"full_name" db:"full_name"`
    Location         *string    `json:"location" db:"location"`
    Established      *time.Time `json:"established" db:"established"`
    Collections      *string    `json:"collections" db:"collections"`
    Director         *string    `json:"director" db:"director"`
    Exhibitions      *string    `json:"exhibitions" db:"exhibitions"`
}

type PoliticalPositionInfoBox struct {
    Position         *string    `json:"position" db:"position"`
    OfficeHolder     *string    `json:"office_holder" db:"office_holder"`
    TermStart        *time.Time `json:"term_start" db:"term_start"`
    TermEnd          *time.Time `json:"term_end" db:"term_end"`
    Party            *string    `json:"party" db:"party"`
    Constituency     *string    `json:"constituency" db:"constituency"`
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

type InfoBoxDB struct{
    InfoBoxType string
    InfoBox InfoBox
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

