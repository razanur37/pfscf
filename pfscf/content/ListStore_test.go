package content

import (
	"testing"

	test "github.com/razanur37/pfscf/pfscf/testutils"
)

func TestStore_InheritFrom(t *testing.T) {
	t.Run("ensure deep copies", func(t *testing.T) {
		var tc text
		tc.X = 1.0
		tc.Presets = []string{"foo", "bar"}

		s1 := NewListStore()
		s1.add(&tc)

		s2 := NewListStore()
		s2.InheritFrom(s1)

		tcStore1 := s1[0].(*text)
		tcStore2 := s2[0].(*text)

		// modifications in an entry from one store should not be reflected in the other
		tcStore1.X = 2.0
		test.ExpectNotEqual(t, tcStore1.X, tcStore2.X)

		tcStore1.Presets[0] = "foobar"
		test.ExpectNotEqual(t, tcStore1.Presets[0], tcStore2.Presets[0])
	})
}
