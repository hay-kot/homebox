package reporting

import (
	"encoding/csv"
	"io"

	"github.com/gocarina/gocsv"
	"github.com/hay-kot/homebox/backend/internal/data/repo"
	"github.com/rs/zerolog"
)

type ReportingService struct {
	repos *repo.AllRepos
	l     *zerolog.Logger
}

func NewReportingService(repos *repo.AllRepos, l *zerolog.Logger) *ReportingService {
	gocsv.SetCSVWriter(func(out io.Writer) *gocsv.SafeCSVWriter {
		writer := csv.NewWriter(out)
		writer.Comma = '\t'
		return gocsv.NewSafeCSVWriter(writer)
	})

	return &ReportingService{
		repos: repos,
		l:     l,
	}
}
