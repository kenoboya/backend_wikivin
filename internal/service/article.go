package service

import (
	"context"
	"wikivin/internal/model"
	repo "wikivin/internal/repository/mysql"
)

type ArticlesService struct {
	articlesRepo *repo.ArticlesRepository
	chaptersRepo *repo.ChaptersRepository
	infoBoxRepo *repo.InfoBoxesRepository
}

func NewArticlesService(articlesRepo *repo.ArticlesRepository, chaptersRepo *repo.ChaptersRepository, infoBoxRepo *repo.InfoBoxesRepository) *ArticlesService{
	return &ArticlesService{articlesRepo, chaptersRepo, infoBoxRepo}
}
func (s *ArticlesService) LoadArticlesBriefInfo(ctx context.Context) ([]model.ArticleBriefInfo, error){
	return s.articlesRepo.GetArticlesBriefInfo(ctx)
}
func (s *ArticlesService)LoadArticle(ctx context.Context, articleID int) (*model.ArticlePage, error){
	var article model.Article
	var chapters []model.Chapter
	var infoBox model.InfoBox

	infoType, ObjectInfoBoxID, err := s.infoBoxRepo.GetTypeAndObjectInfoBoxByArticleID(ctx, articleID)
    if err != nil {
        return nil, err
    }
	infoBox, err = s.infoBoxRepo.GetInfoBoxByObjectInfoBoxIDAndType(ctx, ObjectInfoBoxID,infoType)
	if err != nil{
		return nil, err
	}

	article, err = s.articlesRepo.GetArticleByID(ctx, articleID)
	if err != nil{
		return nil, err
	}

	chapters, err = s.chaptersRepo.GetChaptersByArticleID(ctx, articleID)
	if err != nil{
		return nil, err
	}
	chapters, err = buildHierarchy(chapters)
	if err != nil{
		return nil, err
	}
	
	return &model.ArticlePage{
		Article: article,
		Chapters: chapters,
		InfoBox: infoBox,
		}, nil
}


func buildHierarchy(chapters []model.Chapter) ([]model.Chapter, error){
	chapterMap := make(map[int]*model.Chapter)
	for i:= range chapters{
		chapterMap[chapters[i].ID] = &chapters[i]
	}
	var roots []model.Chapter
	for _, ch:= range chapterMap{
		if ch.ParentID != nil{
			parent, exists:= chapterMap[*ch.ParentID]
			if exists{
				parent.Child = append(parent.Child, ch)
			}else{
				return roots, model.ErrNotFoundParentChapter
			}
		}else{
			roots = append(roots, *ch)
		}
	}
	return roots, nil
}
