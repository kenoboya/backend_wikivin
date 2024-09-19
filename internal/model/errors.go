package model

import "errors"

var (
	ErrArticleData = errors.New("invalid article data")
	ErrChapterData = errors.New("invalid chapter data")
	ErrNotFoundParentChapter = errors.New("the parent chapter wasn't found")
	ErrSpecifiedIdInParam = errors.New("an incorrect ID was specified in the parameter, which led to an error during conversion")
	ErrInfoBoxType = errors.New("infoBoxType is required")
	ErrNilPointerFromReflection = errors.New("received nil pointer")
	ErrUserNotFound = errors.New("user doesn't exist")
	ErrUserBlocked = errors.New("user is blocked")
	ErrInvalidPassword = errors.New("invalid password")
	ErrInvalidLogin = errors.New("invalid login")
	ErrLoginEmpty = errors.New("login cannot be empty")
	ErrNotImageUploaded = errors.New("no image uploaded")
	ErrSaveFile = errors.New("could not save file")
	ErrJSONData = errors.New("no JSON data found")
)