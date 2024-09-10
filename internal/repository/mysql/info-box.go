package repo

import (
	"context"
	"fmt"
	"reflect"
	"strings"
	"wikivin/internal/model"

	"github.com/jmoiron/sqlx"
)

type InfoBoxesRepository struct {
	db *sqlx.DB
}

func NewInfoBoxesRepository(db *sqlx.DB) *InfoBoxesRepository{
	return &InfoBoxesRepository{db}
}



func (r *InfoBoxesRepository) Create(ctx context.Context, articleID int, infoBoxID int) error{
	query:="INSERT INTO info_box (article_id, object_info_box_id) VALUES($1,$2)"
	if _, err:= r.db.Exec(query, articleID, infoBoxID); err != nil{
		return err
	}
	return nil
}

func (r *InfoBoxesRepository) CreateInfoBoxByType(ctx context.Context, infoBoxDB model.InfoBoxDB) (int, error) {
    columns, placeholders, values := getReflectionFieldForInfoBox(infoBoxDB.InfoBox)

    query := fmt.Sprintf(
        "INSERT INTO %s_info_box (%s) VALUES (%s) RETURNING %s_info_box_id",
        infoBoxDB.InfoBoxType,
        strings.Join(columns, ", "),
        strings.Join(placeholders, ", "),
        infoBoxDB.InfoBoxType,
    )

    var id int
    row := r.db.QueryRowContext(ctx, query, values...) 

    if err := row.Scan(&id); err != nil {
        return -1, err
    }

    return id, nil
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

func getReflectionFieldForInfoBox(infoBox model.InfoBox) ([]string, []string, []interface{}){
	val := reflect.ValueOf(infoBox).Elem()
    typ := val.Type()

    var columns []string
    var placeholders []string
    var values []interface{}

    for i := 0; i < val.NumField(); i++ {
        field := typ.Field(i)
        columnName := field.Tag.Get("db") 
        if columnName == "" {
            continue
        }

        columns = append(columns, columnName)
        placeholders = append(placeholders, fmt.Sprintf(":%s", columnName))
        values = append(values, val.Field(i).Interface())
    }
	return columns, placeholders, values
}