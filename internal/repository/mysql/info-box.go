package repo

import (
	"context"
	"fmt"
	"log"
	"reflect"
	"strings"
	"wikivin/internal/model"
	"wikivin/pkg/logger"

	"github.com/jmoiron/sqlx"
)

type InfoBoxesRepository struct {
	db *sqlx.DB
}

func NewInfoBoxesRepository(db *sqlx.DB) *InfoBoxesRepository{
	return &InfoBoxesRepository{db}
}



func (r *InfoBoxesRepository) Create(ctx context.Context, articleID int, infoboxType string, infoBoxID int) error {
    query := "INSERT INTO info_box (article_id, type, object_info_box_id) VALUES (?, ?, ?)"
    if _, err := r.db.ExecContext(ctx, query, articleID, infoboxType, infoBoxID); err != nil {
        return err
    }
    return nil
}


func (r *InfoBoxesRepository) CreateInfoBoxByType(ctx context.Context, infoBoxDB model.InfoBoxDB) (int, error) {
    columns, placeholders, values := getReflectionFieldForInfoBox(infoBoxDB.InfoBox)

    query := fmt.Sprintf(
        "INSERT INTO %s_info_box (%s) VALUES (%s)",
        infoBoxDB.InfoBoxType,
        strings.Join(columns, ", "),
        strings.Join(placeholders, ", "),
    )

    result, err := r.db.ExecContext(ctx, query, values...)
    if err != nil {
        return -1, err
    }

    id, err := result.LastInsertId()
    if err != nil {
        return -1, err
    }

    return int(id), nil
}



func(r *InfoBoxesRepository)GetTypeAndObjectInfoBoxByArticleID(ctx context.Context, articleID int) (string, int, error){
	query := "SELECT type, object_info_box_id FROM info_box WHERE article_id = ?"
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
	query:= fmt.Sprintf("SELECT * FROM %s_info_box WHERE %s_info_box_id = ?", infoBoxType, infoBoxType)
	err = r.db.Get(infoBox, query, objectInfoBoxID)
	if err!= nil{
		return nil, err
	}
	return infoBox, nil
}

func getReflectionFieldForInfoBox(infoBox interface{}) ([]string, []string, []interface{}) {
    val := reflect.ValueOf(infoBox)
    typ := val.Type()

    if val.Kind() == reflect.Ptr {
        if val.IsNil() {
            logger.Error(model.ErrNilPointerFromReflection)
            return nil, nil, nil
        }
        val = val.Elem()
        typ = val.Type()
    }

    if val.Kind() != reflect.Struct {
        log.Fatal("Expected struct, got:", val.Kind())
        return nil, nil, nil
    }

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
        placeholders = append(placeholders, "?")
        values = append(values, val.Field(i).Interface())
    }

    return columns, placeholders, values
}
