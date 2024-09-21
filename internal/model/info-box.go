package model

import (
	"encoding/json"
	"fmt"
	"time"
)

const (
    PersonInfoBoxType            = "person"
    BuildingInfoBoxType           = "building"
    CompanyInfoBoxType           = "company"
    EventInfoBoxType              = "event"
    FilmInfoBoxType               = "film"
    BookInfoBoxType               = "book"
    AlbumInfoBoxType              = "album"
    AnimalInfoBoxType             = "animal"
    AwardInfoBoxType              = "award"
    SongInfoBoxType               = "song"
    CountryInfoBoxType            = "country"
    UniversityInfoBoxType         = "university"
    MuseumInfoBoxType             = "museum"
    PoliticalPositionInfoBoxType  = "political"
)

var infoBoxFactories map[string]func() InfoBox


type PersonInfoBox struct {
    ID               int        `json:"person_info_box_id" db:"person_info_box_id"`
    FullName         *string    `json:"full_name" db:"full_name"`
    BirthDate        *string `json:"birth_date" db:"birth_date"`
    PlaceOfBirth     *string    `json:"place_of_birth" db:"place_of_birth"`
    Nationality      *string    `json:"nationality" db:"nationality"`
    DeathDate        *string`json:"death_date" db:"death_date"`
    Educations       *string    `json:"educations" db:"educations"`
    Occupations      *string    `json:"occupations" db:"occupations"`
    Parents          *string    `json:"parents" db:"parents"`
    Children         *string    `json:"children" db:"children"`
}

type BuildingInfoBox struct {
    ID               int        `json:"building_info_box_id" db:"building_info_box_id"`
    FullName         *string    `json:"full_name" db:"full_name"`
    Location         *string    `json:"location" db:"location"`
    Founded          *string`json:"founded" db:"founded"`
    Architects       *string    `json:"architects" db:"architects"`
    Height           *int       `json:"height" db:"height"`
    Floors           *int       `json:"floors" db:"floors"`
    Usage            *string    `json:"usage" db:"usage"`
}

type CompanyInfoBox struct {
    ID               int        `json:"company_info_box_id" db:"company_info_box_id"`
    FullName         *string    `json:"full_name" db:"full_name"`
    Founded          *string    `json:"founded" db:"founded"`
    Industry         *string    `json:"industry" db:"industry"`
    Headquarters     *string    `json:"headquarters" db:"headquarters"`
    CEO              *string    `json:"ceo" db:"ceo"`
    Revenue          *string    `json:"revenue" db:"revenue"`
    Employees        *int       `json:"employees" db:"employees"`
}

type EventInfoBox struct {
    ID               int        `json:"event_info_box_id" db:"event_info_box_id"`
    Title            *string    `json:"title" db:"title"`
    Date             *string    `json:"date" db:"date"`
    Locations        *string    `json:"locations" db:"locations"`
    Organizers       *string    `json:"organizers" db:"organizers"`
    Attendance       *int       `json:"attendance" db:"attendance"`
}

type FilmInfoBox struct {
    ID               int        `json:"film_info_box_id" db:"film_info_box_id"`
    Title            *string       `json:"title" db:"title"`
    ReleaseDate      *string    `json:"release_date" db:"release_date"`
    Director         *string       `json:"director" db:"director"`
    Genre            *string       `json:"genre" db:"genre"`
    Duration         *time.Duration `json:"duration" db:"duration"`
    Languages        *string       `json:"language" db:"language"`
    Rating           *float32      `json:"rating" db:"rating"`
}

type BookInfoBox struct {
    ID               int        `json:"book_info_box_id" db:"book_info_box_id"`
    Title            *string    `json:"title" db:"title"`
    Authors          *string    `json:"authors" db:"authors"`
    PublishDate      *string `json:"publish_date" db:"publish_date"`
    Genre            *string    `json:"genre" db:"genre"`
    ISBN             *string    `json:"isbn" db:"isbn"`
    Pages            *int       `json:"pages" db:"pages"`
    Publisher        *string    `json:"publisher" db:"publisher"`
}

type AlbumInfoBox struct {
    ID               int        `json:"album_info_box_id" db:"album_info_box_id"`
    Title            *string    `json:"title" db:"title"`
    Artists          *string    `json:"artists" db:"artists"`
    ReleaseDate      *string `json:"release_date" db:"release_date"`
    Genre            *string    `json:"genre" db:"genre"`
    Label            *string    `json:"label" db:"label"`
    Tracks           *string    `json:"tracks" db:"tracks"`
}

type AnimalInfoBox struct {
    ID               int     `json:"animal_info_box_id" db:"animal_info_box_id"`
    FullName         *string `json:"full_name" db:"full_name"`
    Species          *string `json:"species" db:"species"`
    Habitat          *string `json:"habitat" db:"habitat"`
    ConservationStatus *bool  `json:"conservation_status" db:"conservation_status"`
    Lifespan         *string `json:"lifespan" db:"lifespan"`
}

type AwardInfoBox struct {
    ID               int        `json:"award_info_box_id" db:"award_info_box_id"`
    FullName         *string    `json:"full_name" db:"full_name"`
    Year             *string `json:"year" db:"year"`
    Category         *string    `json:"category" db:"category"`
    Recipients       *string    `json:"recipients" db:"recipients"`
    Organizations    *string    `json:"organizations" db:"organizations"`
}

type SongInfoBox struct {
    ID               int           `json:"song_info_box_id" db:"song_info_box_id"`
    Title            *string       `json:"title" db:"title"`
    Artists          *string       `json:"artists" db:"artists"`
    Album            *string       `json:"album" db:"album"`
    ReleaseDate      *string   `json:"release_date" db:"release_date"`
    Genre            *string       `json:"genre" db:"genre"`
    Duration         *time.Duration `json:"duration" db:"duration"`
    Label            *string       `json:"label" db:"label"`
}

type CountryInfoBox struct {
    ID               int        `json:"country_info_box_id" db:"country_info_box_id"`
    FullName         *string    `json:"full_name" db:"full_name"`
    Capital          *string    `json:"capital" db:"capital"`
    Population       *int       `json:"population" db:"population"`
    Region           *string    `json:"region" db:"region"`
    Area             *float64   `json:"area" db:"area"`
    Currency         *string    `json:"currency" db:"currency"`
    OfficialLanguage *string    `json:"official_language" db:"official_language"`
}

type UniversityInfoBox struct {
    ID               int        `json:"university_info_box_id" db:"university_info_box_id"`
    FullName         *string    `json:"full_name" db:"full_name"`
    Founded          *string `json:"founded" db:"founded"`
    Location         *string    `json:"location" db:"location"`
    Programs         *string    `json:"programs" db:"programs"`
    Enrollment       *int       `json:"enrollment" db:"enrollment"`
    Affiliation      *string    `json:"affiliation" db:"affiliation"`
    Accreditation    *string    `json:"accreditation" db:"accreditation"`
}

type MuseumInfoBox struct {
    ID               int        `json:"museum_info_box_id" db:"museum_info_box_id"`
    FullName         *string    `json:"full_name" db:"full_name"`
    Location         *string    `json:"location" db:"location"`
    Established      *string    `json:"established" db:"established"`
    Collections      *string    `json:"collections" db:"collections"`
    Director         *string    `json:"director" db:"director"`
    Exhibitions      *string    `json:"exhibitions" db:"exhibitions"`
}

type PoliticalPositionInfoBox struct {
    ID               int        `json:"political_info_box_id" db:"political_info_box_id"`
    Position         *string    `json:"position" db:"position"`
    OfficeHolder     *string    `json:"office_holder" db:"office_holder"`
    TermStart        *string    `json:"term_start" db:"term_start"`
    TermEnd          *string `json:"term_end" db:"term_end"`
    Party            *string    `json:"party" db:"party"`
    Constituency     *string    `json:"constituency" db:"constituency"`
}


func init(){
    if infoBoxFactories == nil{
        infoBoxFactories = map[string]func() InfoBox{
            PersonInfoBoxType :             func() InfoBox { return &PersonInfoBox{} },
            BuildingInfoBoxType :           func() InfoBox { return &BuildingInfoBox{} },
            CompanyInfoBoxType :            func() InfoBox { return &CompanyInfoBox{} },
            EventInfoBoxType :              func() InfoBox { return &EventInfoBox{} },
            FilmInfoBoxType :               func() InfoBox { return &FilmInfoBox{} },
            BookInfoBoxType :               func() InfoBox { return &BookInfoBox{} },
            AlbumInfoBoxType :              func() InfoBox { return &AlbumInfoBox{} },
            AnimalInfoBoxType :             func() InfoBox { return &AnimalInfoBox{} },
            AwardInfoBoxType :              func() InfoBox { return &AwardInfoBox{} },
            SongInfoBoxType :               func() InfoBox { return &SongInfoBox{} },
            CountryInfoBoxType :            func() InfoBox { return &CountryInfoBox{} },
            UniversityInfoBoxType :         func() InfoBox { return &UniversityInfoBox{} },
            MuseumInfoBoxType :             func() InfoBox { return &MuseumInfoBox{} },
            PoliticalPositionInfoBoxType :  func() InfoBox { return &PoliticalPositionInfoBox{} },
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
    return PersonInfoBoxType 
}

func (p *PersonInfoBox) ToJSON() (string, error) {
    data, err := json.Marshal(p)
    if err != nil {
        return "", err
    }
    return string(data), nil
}


func (b *BuildingInfoBox) GetType() string {
    return BuildingInfoBoxType 
}

func (b *BuildingInfoBox) ToJSON() (string, error) {
    data, err := json.Marshal(b)
    if err != nil {
        return "", err
    }
    return string(data), nil
}


func (c *CompanyInfoBox) GetType() string {
    return CompanyInfoBoxType 
}

func (c *CompanyInfoBox) ToJSON() (string, error) {
    data, err := json.Marshal(c)
    if err != nil {
        return "", err
    }
    return string(data), nil
}


func (e *EventInfoBox) GetType() string {
    return EventInfoBoxType 
}

func (e *EventInfoBox) ToJSON() (string, error) {
    data, err := json.Marshal(e)
    if err != nil {
        return "", err
    }
    return string(data), nil
}


func (f *FilmInfoBox) GetType() string {
    return FilmInfoBoxType 
}

func (f *FilmInfoBox) ToJSON() (string, error) {
    data, err := json.Marshal(f)
    if err != nil {
        return "", err
    }
    return string(data), nil
}


func (b *BookInfoBox) GetType() string {
    return BookInfoBoxType 
}

func (b *BookInfoBox) ToJSON() (string, error) {
    data, err := json.Marshal(b)
    if err != nil {
        return "", err
    }
    return string(data), nil
}


func (a *AlbumInfoBox) GetType() string {
    return AlbumInfoBoxType 
}

func (a *AlbumInfoBox) ToJSON() (string, error) {
    data, err := json.Marshal(a)
    if err != nil {
        return "", err
    }
    return string(data), nil
}


func (a *AnimalInfoBox) GetType() string {
    return AnimalInfoBoxType 
}

func (a *AnimalInfoBox) ToJSON() (string, error) {
    data, err := json.Marshal(a)
    if err != nil {
        return "", err
    }
    return string(data), nil
}


func (a *AwardInfoBox) GetType() string {
    return AwardInfoBoxType 
}

func (a *AwardInfoBox) ToJSON() (string, error) {
    data, err := json.Marshal(a)
    if err != nil {
        return "", err
    }
    return string(data), nil
}


func (s *SongInfoBox) GetType() string {
    return SongInfoBoxType 
}

func (s *SongInfoBox) ToJSON() (string, error) {
    data, err := json.Marshal(s)
    if err != nil {
        return "", err
    }
    return string(data), nil
}


func (c *CountryInfoBox) GetType() string {
    return CountryInfoBoxType 
}

func (c *CountryInfoBox) ToJSON() (string, error) {
    data, err := json.Marshal(c)
    if err != nil {
        return "", err
    }
    return string(data), nil
}


func (u *UniversityInfoBox) GetType() string {
    return UniversityInfoBoxType 
}

func (u *UniversityInfoBox) ToJSON() (string, error) {
    data, err := json.Marshal(u)
    if err != nil {
        return "", err
    }
    return string(data), nil
}


func (m *MuseumInfoBox) GetType() string {
    return MuseumInfoBoxType 
}

func (m *MuseumInfoBox) ToJSON() (string, error) {
    data, err := json.Marshal(m)
    if err != nil {
        return "", err
    }
    return string(data), nil
}

func (p *PoliticalPositionInfoBox) GetType() string {
    return PoliticalPositionInfoBoxType 
}

func (p *PoliticalPositionInfoBox) ToJSON() (string, error) {
    data, err := json.Marshal(p)
    if err != nil {
        return "", err
    }
    return string(data), nil
}

