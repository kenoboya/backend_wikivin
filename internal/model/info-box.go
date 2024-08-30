package model

import (
	"encoding/json"
	"fmt"
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
    PoliticalPosition = "politicalPosition"
)

var infoBoxFactories map[string]func() InfoBox

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

type PersonInfoBox struct {
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

type BuildingInfoBox struct {
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

type CompanyInfoBox struct {
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

type EventInfoBox struct {
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

type FilmInfoBox struct {
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

type BookInfoBox struct {
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

type AlbumInfoBox struct {
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

type AnimalInfoBox struct {
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

type AwardInfoBox struct {
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

type SongInfoBox struct {
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

type CountryInfoBox struct {
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

type UniversityInfoBox struct {
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

type MuseumInfoBox struct {
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

type PoliticalPositionInfoBox struct {
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

