package v1

import (
	"net/url"
	"strconv"

	"github.com/google/uuid"
)

func queryUUIDList(params url.Values, key string) []uuid.UUID {
	var ids []uuid.UUID
	for _, id := range params[key] {
		uid, err := uuid.Parse(id)
		if err != nil {
			continue
		}
		ids = append(ids, uid)
	}
	return ids
}

func queryIntOrNegativeOne(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		return -1
	}
	return i
}

func queryBool(s string) bool {
	b, err := strconv.ParseBool(s)
	if err != nil {
		return false
	}
	return b
}
