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
	if articleHandler == nil {
		syncArticleHandler.Do(func() {
			articleHandler = handler.NewArticleHandler(GetRouter(), GetArticleFactory())
		})
	}
	return articleHandler
}
