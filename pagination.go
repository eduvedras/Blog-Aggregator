package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func getLimitAndOffset(r *http.Request) (int, int, error) {
	limit := r.URL.Query().Get("limit")
	offset := r.URL.Query().Get("offset")

	var limitInt int
	var err error
	if limit == "" {
		limitInt = 20
	} else {
		limitInt, err = strconv.Atoi(limit)
		if err != nil {
			return 0, 0, fmt.Errorf("Limit must be an integer, received %v", limit)
		}
	}

	var offsetInt int
	if offset == "" {
		offsetInt = 0
	} else {
		offsetInt, err = strconv.Atoi(offset)
		if err != nil {
			return 0, 0, fmt.Errorf("Offset must be an integer, received %v", offset)
		}
	}

	return limitInt, offsetInt, nil
}
