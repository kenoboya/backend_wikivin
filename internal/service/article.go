package service

import (
	"context"
	"wikivin/internal/model"
	repo "wikivin/internal/repository/mysql"
)

type ArticlesService struct {
	articlesRepo repo.Articles
	chaptersRepo repo.Chapters
	infoBoxRepo repo.InfoBox
}

func NewArticlesService(articles repo.Articles, chapters repo.Chapters, infoBox repo.InfoBox) *ArticlesService{
	return &ArticlesService{articles, chapters, infoBox}
}

func (s *ArticlesService) CreateArticle(ctx context.Context, infoBoxDB model.InfoBoxDB,  article model.Article) error{
	articleID, err:= s.articlesRepo.Create(ctx, article); 
	if err != nil{
		return err;
	}
	infoBoxID, err:= s.infoBoxRepo.CreateInfoBoxByType(ctx, infoBoxDB)
	if err != nil{
		return err;
	}
	if err:= s.infoBoxRepo.Create(ctx,articleID, infoBoxID); err!= nil{
		return err;
	}

	// chapters
	return nil
}

func (s *ArticlesService) LoadArticles(ctx context.Context) ([]model.Article, error){
	return s.articlesRepo.GetArticles(ctx)
}
func (s *ArticlesService)LoadArticle(ctx context.Context, title string) (*model.ArticlePage, error){
	var article model.Article
	var chapters []model.Chapter
	var infoBox model.InfoBox

	article, err := s.articlesRepo.GetArticleByTitle(ctx, title)
	if err != nil{
		return nil, err
	}

	infoType, ObjectInfoBoxID, err := s.infoBoxRepo.GetTypeAndObjectInfoBoxByArticleID(ctx, article.ID)
    if err != nil {
        return nil, err
    }
	infoBox, err = s.infoBoxRepo.GetInfoBoxByObjectInfoBoxIDAndType(ctx, ObjectInfoBoxID, infoType)
	if err != nil{
		return nil, err
	}

	chapters, err = s.chaptersRepo.GetChaptersByArticleID(ctx, article.ID)
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
