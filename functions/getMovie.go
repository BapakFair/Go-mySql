package functions

import (
	"context"
	. "go_sql/query"
	"go_sql/utils"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func GetMovie(rw http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	movies, err := GetAll(ctx)
	if err != nil {
		panic(err)
	}
	utils.ResponseJSON(rw, movies, http.StatusOK)
}
