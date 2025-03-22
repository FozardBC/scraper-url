package index

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"scraper-url/internal/index"
)

type IndexGetter interface {
	GetIndex() *index.Index
}

func New(log *slog.Logger, urlGetter IndexGetter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "/internal/netsrv/http-server/handlers/index/New"

		log.With(
			slog.String("op:%s", op),
		)

		i := urlGetter.GetIndex()

		err := json.Unmarshal(data)

	}

}
