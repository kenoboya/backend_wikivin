package model

import "errors"

var (
	ErrArticleData = errors.New("invalid article data")
	ErrChapterData = errors.New("invalid chapter data")
	ErrNotFoundParentChapter = errors.New("the parent chapter wasn't found")
	ErrSpecifiedIdInParam = errors.New("an incorrect ID was specified in the parameter, which led to an error during conversion")
	ErrInfoBoxType = errors.New("infoBoxType is required")
	ErrNilPointerFromReflection = errors.New("received nil pointer")
)