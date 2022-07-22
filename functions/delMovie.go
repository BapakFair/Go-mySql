package functions

import (
	"context"
	"fmt"
	. "go_sql/query"
	"go_sql/utils"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func DelMovie(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var idMovie = ps.ByName("id")
	if err := DelData(ctx, idMovie); err != nil {
		errDelete := map[string]string{
			"error": fmt.Sprintf("%v", err),
		}
		utils.ResponseJSON(rw, errDelete, http.StatusInternalServerError)
		return
	}

	res := map[string]string{
		"status": "Success",
	}
	utils.ResponseJSON(rw, res, http.StatusOK)
}
