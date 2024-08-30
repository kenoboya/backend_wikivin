package repo

import (
	"context"
	"fmt"
	"wikivin/internal/model"

	"github.com/jmoiron/sqlx"
)

type InfoBoxesRepository struct {
	db *sqlx.DB
}

func NewInfoBoxesRepository(db *sqlx.DB) *InfoBoxesRepository{
	return &InfoBoxesRepository{db}
}

func(r *InfoBoxesRepository)GetTypeAndObjectInfoBoxByArticleID(ctx context.Context, articleID int) (string, int, error){
	query := "SELECT type, object_info_box_id FROM info_box WHERE article_id = $1"
	var typeInfoBox string
	var objectInfoBoxID int
	rows, err:= r.db.Queryx(query,articleID)
	if err!= nil{
		return typeInfoBox, objectInfoBoxID, err
	}
	defer rows.Close()
	for rows.Next(){
		err:= rows.Scan(&typeInfoBox, &objectInfoBoxID)
		if err!= nil{
			return typeInfoBox, objectInfoBoxID, err
		}		
	}
	return typeInfoBox, objectInfoBoxID, nil
}
func(r *InfoBoxesRepository)GetInfoBoxByObjectInfoBoxIDAndType(ctx context.Context, objectInfoBoxID int, infoBoxType string) (model.InfoBox, error){
	
	factory, err:=model.GetInfoBoxFactory(infoBoxType)
	if err != nil{
		return nil, err
	}
	infoBox := factory()
	query:= fmt.Sprintf("SELECT * FROM %s_info_box WHERE %s_info_box_id = $1", infoBoxType, infoBoxType)
	err = r.db.Get(infoBox, query, objectInfoBoxID)
	if err!= nil{
		return nil, err
	}
	return infoBox, nil
}
