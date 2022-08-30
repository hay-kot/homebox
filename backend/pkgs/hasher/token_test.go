package hasher

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const ITERATIONS = 200

func Test_NewToken(t *testing.T) {
	t.Parallel()
	tokens := make([]Token, ITERATIONS)
	for i := 0; i < ITERATIONS; i++ {
		tokens[i] = GenerateToken()
	}

	// Check if they are unique
	for i := 0; i < 5; i++ {
		for j := i + 1; j < 5; j++ {
			if tokens[i].Raw == tokens[j].Raw {
				t.Errorf("NewToken() failed to generate unique tokens")
			}
		}
	}
}

func Test_HashToken_CheckTokenHash(t *testing.T) {
	t.Parallel()
	for i := 0; i < ITERATIONS; i++ {
		token := GenerateToken()

		// Check raw text is reltively random
		for j := 0; j < 5; j++ {
			assert.NotEqual(t, token.Raw, GenerateToken().Raw)
		}

		// Check token length is less than 32 characters
		assert.Less(t, len(token.Raw), 32)

		// Check hash is the same
		assert.Equal(t, token.Hash, HashToken(token.Raw))
	}
}
