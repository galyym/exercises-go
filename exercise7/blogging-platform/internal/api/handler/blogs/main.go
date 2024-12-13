package blogs

import (
	"github.com/talgat-ruby/exercises-go/exercise7/blogging-platform/internal/db"
	"log/slog"
)

type Blogs struct {
	logger *slog.Logger
	db     *db.ConfDB
}

func New(logger *slog.Logger, db *db.ConfDB) *Blogs {
	return &Blogs{logger: logger, db: db}
}