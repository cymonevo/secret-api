package provider

import (
	"sync"

	"github.com/cymon1997/go-backend/internal/render"
	"github.com/cymon1997/go-backend/internal/router"
)

var (
	appRouter     router.Router
	syncAppRouter sync.Once

	renderEngine     render.Client
	syncRenderEngine sync.Once
)

func GetRouter() router.Router {
	syncAppRouter.Do(func() {
		appRouter = router.New(GetRenderEngine())
	})
	return appRouter
}

func GetRenderEngine() render.Client {
	syncRenderEngine.Do(func() {
		renderEngine = render.New(render.Config{
			TemplatePath: "files/",
		})
	})
	return renderEngine
}
