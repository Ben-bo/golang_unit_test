package helper

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type dataUser struct {
	name     string
	request  string
	expected string
}

// before test (only run per package)
func TestMain(t *testing.M) {

	fmt.Println("before all test") // func in here will run before all test run

	t.Run() //run all test

	fmt.Println(" after all test ") // func in here will run after all test
}

func TestHello(t *testing.T) {
	// require = memamngil failnow (program berhenti pada saat gagal)
	// assert = memamngil fail (program tetap jalan sampe habis pada saat gagal)
	// t.skip() untuk skip test

	result := Hello("beni")

	assert.Equal(t, "hello beni", result)
	fmt.Println("after test hello")
}

func TestSubTest(t *testing.T) {
	testedUser2 := dataUser{
		name:     "user2",
		request:  "user2",
		expected: "hello user2",
	}
	users := []dataUser{
		{
			name:     "user1",
			request:  "user1",
			expected: "hello user1",
		},
		testedUser2,
	}

	for _, user := range users {
		t.Run("testing "+user.name, func(t *testing.T) {
			result := Hello(user.request)

			require.Equal(t, user.expected, result)
		})
	}

	t.Run("testing Beni", func(t *testing.T) {
		result := Hello("beni")

		require.Equal(t, "hello beni", result)
	})
	t.Run("testing Rendi", func(t *testing.T) {
		result := Hello("rendi")

		require.Equal(t, "hello rendi", result)
	})
}
