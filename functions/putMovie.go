package functions

import (
	"context"
	"encoding/json"
	"go_sql/models"
	. "go_sql/query"
	"go_sql/utils"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func PutMovie(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(rw, "Please Use Content-Type application/json", http.StatusBadRequest)
		return
	}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var mov models.Movie
	if err := json.NewDecoder(r.Body).Decode(&mov); err != nil {
		utils.ResponseJSON(rw, err, http.StatusBadRequest)
		return
	}

	var idMovie = ps.ByName("id")
	if err := PutData(ctx, mov, idMovie); err != nil {
		utils.ResponseJSON(rw, err, http.StatusInternalServerError)
		return
	}

	res := map[string]string{
		"status": "Success",
	}
	utils.ResponseJSON(rw, res, http.StatusOK)
}
