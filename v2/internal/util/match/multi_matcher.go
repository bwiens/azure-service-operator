/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package match

import (
	"strings"

	"github.com/pkg/errors"
	kerrors "k8s.io/apimachinery/pkg/util/errors"
)

// multiMatcher is a matcher that contains multiple other matchers
type multiMatcher struct {
	matchers []StringMatcher
}

var _ StringMatcher = &multiMatcher{}

// newMultiMatcher returns a new matcher for multiple strings
func newMultiMatcher(matcher string) StringMatcher {
	splits := strings.Split(matcher, ";")
	matchers := make([]StringMatcher, 0, len(splits))
	for _, str := range splits {
		m := NewStringMatcher(str)
		matchers = append(matchers, m)
	}

	return &multiMatcher{
		matchers: matchers,
	}
}

// Matches returns true if the provided value is matched by any
func (mm *multiMatcher) Matches(value string) Result {
	result := matchNotFound()
	// MUST not short circuit - for the suggestions generated by any nested matchers to be useful,
	// each matcher must see all potential match candidates.
	for _, m := range mm.matchers {
		matchResult := m.Matches(value)
		if matchResult.Matched {
			// Only store the first match
			if result.Matched == false {
				result = matchResult
			}
		}
	}

	return result
}

// WasMatched returns nil if every nested matcher had a match, otherwise returns a diagnostic error
func (mm *multiMatcher) WasMatched() error {
	var errs []error
	count := 0
	for _, m := range mm.matchers {
		if err := m.WasMatched(); err != nil {
			errs = append(errs, err)
			count++
		}
	}

	return errors.Wrapf(
		kerrors.NewAggregate(errs),
		"%d of %d did not match",
		count,
		len(mm.matchers))
}

// IsRestrictive returns true if any of our contained matchers are restrictive
func (mm *multiMatcher) IsRestrictive() bool {
	for _, m := range mm.matchers {
		if m.IsRestrictive() {
			return true
		}
	}

	return false
}

// String returns all the matchers we contain
func (mm *multiMatcher) String() string {
	var builder strings.Builder
	for i, m := range mm.matchers {
		if i > 0 {
			builder.WriteString(";")
		}

		builder.WriteString(m.String())
	}

	return builder.String()
}

// HasMultipleMatchers returns true if the matcher contains multiple definitions
func HasMultipleMatchers(matcher string) bool {
	return strings.ContainsRune(matcher, ';')
}