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

func (s *ArticlesService) CreateArticle(ctx context.Context, infoBoxDB model.InfoBoxDB,  article model.Article, chapters []model.Chapter) error{
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

	chapters, err = unbuildHierarchy(chapters)
	if err!= nil{
		return err
	}
	for _, chapter:= range chapters{
		if err:= s.chaptersRepo.Create(ctx, chapter); err!= nil{
			return err
		}
	}
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

func unbuildHierarchy(roots []model.Chapter) ([]model.Chapter, error) {
    var chapters []model.Chapter

    convertToPointers := func(chaptersList []model.Chapter) []*model.Chapter {
        var pointers []*model.Chapter
        for i := range chaptersList {
            pointers = append(pointers, &chaptersList[i])
        }
        return pointers
    }

    pointers := convertToPointers(roots)

    var collectChapters func([]*model.Chapter) error
    collectChapters = func(chaptersList []*model.Chapter) error {
        for _, ch := range chaptersList {
            chapters = append(chapters, *ch)
            if ch.Child != nil {
                if err := collectChapters(ch.Child); err != nil {
                    return err
                }
            }
        }
        return nil
    }

    if err := collectChapters(pointers); err != nil {
        return nil, err
    }
    return chapters, nil
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
