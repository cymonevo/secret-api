package handler

import "github.com/cymonevo/secret-api/internal/router"

type BaseHandler interface {
	Register() router.Router
}
