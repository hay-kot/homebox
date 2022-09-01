package v1

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/hay-kot/content/backend/internal/services"
	"github.com/hay-kot/content/backend/internal/types"
	"github.com/hay-kot/content/backend/pkgs/logger"
	"github.com/hay-kot/content/backend/pkgs/server"
)

/*
This is where we put partial snippets/functions for actions that are commonly
used within the controller class. This _hopefully_ helps with code duplication
and makes it a little more consistent when error handling and logging.
*/

// partialParseIdAndUser parses the ID from the requests URL and pulls the user
// from the context. If either of these fail, it will return an error. When an error
// occurs it will also write the error to the response. As such, if an error is returned
// from this function you can return immediately without writing to the response.
func (ctrl *V1Controller) partialParseIdAndUser(w http.ResponseWriter, r *http.Request) (uuid.UUID, *types.UserOut, error) {
	uid, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		ctrl.log.Debug(err.Error(), logger.Props{
			"details": "failed to convert id to valid UUID",
		})
		server.RespondError(w, http.StatusBadRequest, err)
		return uuid.Nil, nil, err
	}

	user := services.UseUserCtx(r.Context())
	return uid, user, nil
}
