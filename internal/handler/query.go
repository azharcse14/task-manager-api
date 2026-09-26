package handler

import (
	"net/url"
	"strconv"
)

// intQuery: URL থেকে একটা সংখ্যা পড়ে। না দিলে def ফেরত দেয়।
func intQuery(q url.Values, key string, def int) (int, error) {
	s := q.Get(key)
	if s == "" {
		return def, nil
	}
	return strconv.Atoi(s)
}

// boolQuery: URL থেকে true/false পড়ে। না দিলে nil ফেরত দেয়।
func boolQuery(q url.Values, key string) (*bool, error) {
	s := q.Get(key)
	if s == "" {
		return nil, nil
	}
	v, err := strconv.ParseBool(s)
	if err != nil {
		return nil, err
	}
	return &v, nil
}
