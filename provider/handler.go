package provider

import (
	"sync"

	"github.com/cymon1997/go-backend/handler"
)

var (
	articleHandler     handler.BaseHandler
	syncArticleHandler sync.Once
)

func GetArticleHandler() handler.BaseHandler {
	syncArticleHandler.Do(func() {
		//articleHandler = handler.NewArticleHandler(GetRouter(), GetArticleFactory())
		articleHandler = handler.NewArticleHandler(GetRouter(), nil)
	})
	return articleHandler
}
