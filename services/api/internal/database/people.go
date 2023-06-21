package database

import (
	"context"
	"log"
	"sync"
)

var (
	prefixMap     map[string]struct{}
	prefixMapInit sync.Once
	suffixMap     map[string]struct{}
	suffixMapInit sync.Once
)

// IsValidPrefix returns true if the prefix is valid.
func (q *Queries) IsValidPrefix(ctx context.Context, prefix string) bool {
	prefixMapInit.Do(func() {
		prefixes, err := q.ListPrefixes(ctx)
		if err != nil {
			log.Fatal(err)
		}

		prefixMap = map[string]struct{}{}
		for _, p := range prefixes {
			prefixMap[p] = struct{}{}
		}
	})

	_, ok := prefixMap[prefix]
	return ok
}

// IsValidSuffix returns true if the suffix is valid.
func (q *Queries) IsValidSuffix(ctx context.Context, suffix string) bool {
	suffixMapInit.Do(func() {
		suffixes, err := q.ListSuffixes(ctx)
		if err != nil {
			log.Fatal(err)
		}

		suffixMap = map[string]struct{}{}
		for _, s := range suffixes {
			suffixMap[s] = struct{}{}
		}
	})

	_, ok := suffixMap[suffix]
	return ok
}
