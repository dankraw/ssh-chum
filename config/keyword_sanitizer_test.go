package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShouldSanitizeKeywords(t *testing.T) {
	t.Parallel()

	// given
	compiler := newKeywordSanitizer()
	entries := []struct {
		input    string
		expected string
	}{
		{"identity_file", "IdentityFile"},
		{"port", "Port"},
		{"hash_known_hosts", "HashKnownHosts"},
		{"MACs", "MACs"},
		{"RhostsRSAAuthentication", "RhostsRSAAuthentication"},
	}

	for _, e := range entries {
		// when
		actual := compiler.sanitize(e.input)

		// then
		assert.Equal(t, actual, e.expected)
	}
}
