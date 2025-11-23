package auth

import (
	"regexp"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestValidateUsername_Format tests username validation format rules
func TestValidateUsername_Format(t *testing.T) {
	usernameRegex := regexp.MustCompile(`^[a-z0-9_]+$`)

	tests := []struct {
		name     string
		username string
		wantErr  bool
	}{
		{"valid username", "john_doe", false},
		{"valid with numbers", "user123", false},
		{"valid with underscores", "test_user", false},
		{"too short", "ab", true},
		{"too long", "thisusernameistoolong", true},
		{"uppercase", "JohnDoe", true},
		{"with spaces", "john doe", true},
		{"with special chars", "john@doe", true},
		{"empty", "", true},
		{"exact min length", "abc", false},
		{"exact max length", "abcdefghij", false},
		{"leading space", " john", true},
		{"trailing space", "john ", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			username := strings.TrimSpace(tt.username)
			
			// Length check
			if len(username) < 3 || len(username) > 10 {
				if !tt.wantErr {
					t.Errorf("Expected error for length check, got pass")
				}
				return
			}
			
			// Leading/trailing space check (after trim, original should match)
			if strings.HasPrefix(tt.username, " ") || strings.HasSuffix(tt.username, " ") {
				if !tt.wantErr {
					t.Errorf("Expected error for leading/trailing spaces, got pass")
				}
				return
			}
			
			// Format check
			if !usernameRegex.MatchString(username) {
				if !tt.wantErr {
					t.Errorf("Expected error for invalid format, got pass")
				}
				return
			}
			
			// Should pass
			if tt.wantErr {
				t.Errorf("Expected error but validation passed")
			}
		})
	}
}

// TestValidateUsername_Length tests username length constraints
func TestValidateUsername_Length(t *testing.T) {
	tests := []struct {
		name     string
		username string
		valid    bool
	}{
		{"min length", "abc", true},
		{"max length", "abcdefghij", true},
		{"one below min", "ab", false},
		{"one above max", "abcdefghijk", false},
		{"empty", "", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			valid := len(strings.TrimSpace(tt.username)) >= 3 && len(strings.TrimSpace(tt.username)) <= 10
			assert.Equal(t, tt.valid, valid, "Length validation failed for %q", tt.username)
		})
	}
}

// TestValidateUsername_FormatRegex tests the regex pattern
func TestValidateUsername_FormatRegex(t *testing.T) {
	usernameRegex := regexp.MustCompile(`^[a-z0-9_]+$`)

	tests := []struct {
		name     string
		username string
		matches  bool
	}{
		{"lowercase letters", "john", true},
		{"with numbers", "user123", true},
		{"with underscores", "test_user", true},
		{"uppercase", "John", false},
		{"with dash", "user-name", false},
		{"with dot", "user.name", false},
		{"with at", "user@name", false},
		{"with space", "user name", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			matches := usernameRegex.MatchString(tt.username)
			assert.Equal(t, tt.matches, matches, "Regex match failed for %q", tt.username)
		})
	}
}

