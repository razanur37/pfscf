package content

import (
	"testing"

	"github.com/razanur37/pfscf/pfscf/canvas"
	"github.com/razanur37/pfscf/pfscf/param"
	"github.com/razanur37/pfscf/pfscf/stamp"
	test "github.com/razanur37/pfscf/pfscf/testutils"
)

func getTextCellWithDummyData(presets ...string) (tc *text) {
	tc = newText()

	tc.Value = "Some value"
	tc.X = 12.0
	tc.Y = 12.0
	tc.X2 = 24.0
	tc.Y2 = 24.0
	tc.Font = "Helvetica"
	tc.Fontsize = 14.0
	tc.Align = "CB"
	tc.Canvas = "test"
	for _, preset := range presets {
		tc.Presets = append(tc.Presets, preset)
	}

	return tc
}

func TestTextCell_IsValid(t *testing.T) {
	paramStore := param.NewStore()
	canvasStore := canvas.NewStore()
	canvas := canvas.NewEntry()
	testCoord := 10.0
	canvas.X2 = &testCoord
	canvas.Y2 = &testCoord
	canvasStore.Add("test", &canvas)

	t.Run("errors", func(t *testing.T) {
		t.Run("missing value", func(t *testing.T) {
			tc := getTextCellWithDummyData()
			tc.Font = "" // "Unset" one required value

			err := tc.isValid(&paramStore, &canvasStore)
			test.ExpectError(t, err, "Missing value", "Font")
		})

		t.Run("value out of range", func(t *testing.T) {
			tc := getTextCellWithDummyData()
			tc.Y2 = 101.0

			err := tc.isValid(&paramStore, &canvasStore)
			test.ExpectError(t, err, "out of range", "Y2")
		})

		t.Run("invalid canvas", func(t *testing.T) {
			tc := getTextCellWithDummyData()
			tc.Canvas = "foobar"

			err := tc.isValid(&paramStore, &canvasStore)
			test.ExpectError(t, err, "Canvas 'foobar' does not exist")
		})
	})

	t.Run("valid", func(t *testing.T) {
		tc := getTextCellWithDummyData()
		tc.X = 0.0 // set something to "zero", which is also acceptable

		err := tc.isValid(&paramStore, &canvasStore)
		test.ExpectNoError(t, err)
	})
}

func TestTextCell_Resolve(t *testing.T) {
	ps := getTestPresetStore(t)

	t.Run("errors", func(t *testing.T) {
		t.Run("non-existant preset", func(t *testing.T) {
			tc := getTextCellWithDummyData("non-existing preset")

			err := tc.resolve(*ps)
			test.ExpectError(t, err, "does not exist")
		})

		t.Run("conflicting presets", func(t *testing.T) {
			tc := getTextCellWithDummyData("conflict1", "conflict2")

			err := tc.resolve(*ps)
			test.ExpectError(t, err, "Contradicting values", "font", "conflict1", "conflict2")
		})
	})

	t.Run("valid", func(t *testing.T) {
		tc := getTextCellWithDummyData("sameData1", "sameData2")
		tc.Font = ""

		err := tc.resolve(*ps)
		test.ExpectNoError(t, err)

		test.ExpectIsSet(t, tc.Font)
	})
}

func TestTextCell_generateOutput(t *testing.T) {
	stamp := stamp.NewStamp(100.0, 100.0, 0.0, 0.0)
	stamp.AddCanvas("test", 0.0, 0.0, 100.0, 100.0)
	testArgName := "someId"
	testArgValue := "foobar"
	as := getTestArgStore(testArgName, testArgValue)

	t.Run("valid", func(t *testing.T) {
		tc := getTextCellWithDummyData()
		tc.Value = "param:someId"

		err := tc.generateOutput(stamp, as)
		test.ExpectNoError(t, err)
	})
}

func TestTextCell_deepCopy(t *testing.T) {
	e1 := newText()
	e1.Value = "t1"
	e1.Presets = append(e1.Presets, "t1")

	e2 := e1.deepCopy().(*text)
	e2.Value = "t2"
	e2.Presets[0] = "t2"

	test.ExpectNotEqual(t, e1.Value, e2.Value)
	test.ExpectNotEqual(t, e1.Presets[0], e2.Presets[0])
}
