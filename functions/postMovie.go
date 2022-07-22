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

func PostMovie(rw http.ResponseWriter, r *http.Request, _ httprouter.Params) {
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

	if err := PostData(ctx, mov); err != nil {
		utils.ResponseJSON(rw, err, http.StatusInternalServerError)
		return
	}

	res := map[string]string{
		"status": "Success",
	}
	utils.ResponseJSON(rw, res, http.StatusOK)
}
