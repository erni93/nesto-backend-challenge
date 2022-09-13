package router

import (
	"backend-api/repository"
	response "backend-api/router/response"
	"log"
	"net/http"
	"strconv"
	"strings"
)

var (
	ErrInvalidFilter = "Some filters are not valid"
)

type BooksRouter struct {
	Repository *repository.BookRepository
}

func (router *BooksRouter) GetBooksHandler(w http.ResponseWriter, r *http.Request) {
	filters, err := getFilters(r)
	if err != nil {
		log.Printf("Error reading GetBooks filters: %s", err)
		response.WriteError(w, ErrInvalidFilter)
	}

	books, err := router.Repository.GetBooks(*filters)
	if err != nil {
		log.Printf("Error retrieving books: %s", err)
		response.WriteGeneralError(w)
	} else {
		response.WriteJsonObject(w, books)
	}
}

func getFilters(r *http.Request) (*repository.GetBooksFilters, error) {
	filters := &repository.GetBooksFilters{}

	authorsParam := r.FormValue("authors")
	if len(authorsParam) > 0 {
		ids, err := splitIdList(authorsParam)
		if err != nil {
			return nil, err
		}
		filters.Authors = ids
	}

	genresParam := r.FormValue("genres")
	if len(genresParam) > 0 {
		ids, err := splitIdList(genresParam)
		if err != nil {
			return nil, err
		}
		filters.Genres = ids
	}

	minPagesParam := r.FormValue("min-pages")
	if len(minPagesParam) > 0 {
		minPages, err := strconv.ParseInt(minPagesParam, 10, 64)
		if err != nil {
			return nil, err
		}
		minPagesInt32 := int32(minPages)
		filters.MinPages = &minPagesInt32
	}

	maxPagesParam := r.FormValue("max-pages")
	if len(maxPagesParam) > 0 {
		maxPages, err := strconv.ParseInt(maxPagesParam, 10, 64)
		if err != nil {
			return nil, err
		}
		maxPagesInt32 := int32(maxPages)
		filters.MaxPages = &maxPagesInt32
	}

	minYearParam := r.FormValue("min-year")
	if len(minYearParam) > 0 {
		minYear, err := strconv.ParseInt(minYearParam, 10, 64)
		if err != nil {
			return nil, err
		}
		minYearInt32 := int32(minYear)
		filters.MinYear = &minYearInt32
	}

	maxYearParam := r.FormValue("max-year")
	if len(maxYearParam) > 0 {
		maxYear, err := strconv.ParseInt(maxYearParam, 10, 64)
		if err != nil {
			return nil, err
		}
		maxYearInt32 := int32(maxYear)
		filters.MaxYear = &maxYearInt32
	}

	limitParam := r.FormValue("limit")
	if len(limitParam) > 0 {
		limit, err := strconv.ParseInt(limitParam, 10, 64)
		if err != nil {
			return nil, err
		}
		limitInt32 := int32(limit)
		filters.Limit = &limitInt32
	}

	return filters, nil
}

func splitIdList(idList string) ([]int64, error) {
	ids := make([]int64, 0)
	idsText := strings.Split(idList, ",")
	for _, idText := range idsText {
		idInt, err := strconv.ParseInt(idText, 10, 64)
		if err != nil {
			return nil, err
		}
		ids = append(ids, idInt)
	}
	return ids, nil
}
