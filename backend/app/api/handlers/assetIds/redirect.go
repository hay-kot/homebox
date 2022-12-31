package assetIds

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/rs/zerolog/log"

	"github.com/go-chi/chi/v5"
	"github.com/hay-kot/homebox/backend/internal/data/repo"
	"github.com/hay-kot/homebox/backend/pkgs/server"
)

func HandleAssetRedirect(repos *repo.AllRepos) server.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) error {
		// Get the asset ID from the URL
		assetIdParam := chi.URLParam(r, "id")
		assetIdParam = strings.ReplaceAll(assetIdParam, "-", "") // Remove dashes
		// Convert the asset ID to an int64
		assetId, err := strconv.ParseInt(assetIdParam, 10, 64)
		if err != nil {
			return err
		}

		// Get the asset from the database
		itemIds, err := repos.Items.GetIDsByAssetID(r.Context(), repo.AssetID(assetId));
		if err != nil {
			return err
		}
		// check if we got more than one item
		if len(itemIds) > 1 {
			log.Err(err).Msg("More than one item found for asset ID")
			return server.Respond(w, http.StatusInternalServerError, "More than one item found for asset ID")
		}
		// check if we got any items
		if len(itemIds) == 0 {
			log.Err(err).Msg("No items found for asset ID")
			return server.Respond(w, http.StatusNotFound, "No items found for asset ID")
		}
	
		http.Redirect(w, r, "/item/" + itemIds[0].String(), http.StatusSeeOther)
		return nil
	}
}