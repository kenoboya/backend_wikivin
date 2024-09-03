package model

import "errors"

var (
	ErrNotFoundParentChapter = errors.New("the parent chapter wasn't found")
	ErrSpecifiedIdInParam = errors.New("an incorrect ID was specified in the parameter, which led to an error during conversion")
)