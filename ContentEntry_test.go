package main

import "testing"

func init() {
	SetIsTestEnvironment(true)
}

func getContentDataWithDummyData(t *testing.T, cdType string) (cd ContentData) {
	cd.Type = cdType
	cd.Desc = "Some Description"
	cd.X1 = 12.0
	cd.Y1 = 12.0
	cd.X2 = 24.0
	cd.Y2 = 24.0
	cd.XPivot = 36.0
	cd.Font = "Helvetica"
	cd.Fontsize = 14.0
	cd.Align = "LB"
	cd.Example = "Some Example"

	expectAllExportedSet(t, cd) // to be sure that we also get all new fields

	return cd
}

func TestNewContentEntry(t *testing.T) {
	cd := getContentDataWithDummyData(t, "myType")
	ce := NewContentEntry("myId", cd)

	expectEqual(t, ce.ID(), "myId")
	expectEqual(t, ce.Type(), "myType")
	expectEqual(t, ce.Description(), "Some Description")
	expectEqual(t, ce.X1(), 12.0)
	expectEqual(t, ce.Y1(), 12.0)
	expectEqual(t, ce.X2(), 24.0)
	expectEqual(t, ce.Y2(), 24.0)
	expectEqual(t, ce.XPivot(), 36.0)
	expectEqual(t, ce.Font(), "Helvetica")
	expectEqual(t, ce.Fontsize(), 14.0)
	expectEqual(t, ce.Align(), "LB")
	expectEqual(t, ce.Example(), "Some Example")
}

func TestContentEntry_IsValid(t *testing.T) {

	t.Run("valid type", func(t *testing.T) {
		cd := getContentDataWithDummyData(t, "textCell")
		ce := NewContentEntry("id", cd)
		isValid, err := ce.IsValid()

		expectEqual(t, isValid, true)
		expectNoError(t, err)

	})

	t.Run("invalid type", func(t *testing.T) {
		cd := getContentDataWithDummyData(t, "textCellX")
		ce := NewContentEntry("id", cd)
		isValid, err := ce.IsValid()

		expectEqual(t, isValid, false)
		expectError(t, err)
	})

	t.Run("textCell with missing values", func(t *testing.T) {
		cd := getContentDataWithDummyData(t, "textCell")
		cd.Font = ""
		ce := NewContentEntry("id", cd)

		isValid, err := ce.IsValid()

		expectEqual(t, isValid, false)
		expectError(t, err)
	})
}

func TestEntriesAreNotContradicting(t *testing.T) {
	var err error

	cdEmpty := ContentData{}
	ceEmpty := NewContentEntry("idEmpty", cdEmpty)

	cdAllSet := getContentDataWithDummyData(t, "type")
	ceAllSet := NewContentEntry("idAllSet", cdAllSet)

	t.Run("no self-contradiction", func(t *testing.T) {
		// a given CE with values should not contradict itself
		err = EntriesAreNotContradicting(&ceAllSet, &ceAllSet)
		expectNoError(t, err)
	})

	t.Run("empty contradicts nothing", func(t *testing.T) {
		// a given CE with no values should contradict nothing
		err = EntriesAreNotContradicting(&ceEmpty, &ceEmpty)
		expectNoError(t, err)
		err = EntriesAreNotContradicting(&ceAllSet, &ceEmpty)
		expectNoError(t, err)
		err = EntriesAreNotContradicting(&ceEmpty, &ceAllSet)
		expectNoError(t, err)
	})

	t.Run("non-overlapping", func(t *testing.T) {
		// Have two partly-set objects with non-overlapping content
		cdLeft := ContentData{X1: 1.0, Desc: "desc"}
		ceLeft := NewContentEntry("idLeft", cdLeft)
		cdRight := ContentData{X2: 2.0, Font: "font"}
		ceRight := NewContentEntry("idRight", cdRight)
		err = EntriesAreNotContradicting(&ceLeft, &ceRight)
		expectNoError(t, err)
	})

	t.Run("conflicting string attribute", func(t *testing.T) {
		cdLeft := getContentDataWithDummyData(t, "type")
		cdLeft.Font = cdLeft.Font + "foo" // <= conflicting data
		ceLeft := NewContentEntry("idLeft", cdLeft)
		cdRight := getContentDataWithDummyData(t, "type")
		ceRight := NewContentEntry("idRight", cdRight)

		err = EntriesAreNotContradicting(&ceLeft, &ceRight)
		expectError(t, err)
	})

	t.Run("conflicting float64 attribute", func(t *testing.T) {
		cdLeft := getContentDataWithDummyData(t, "type")
		cdLeft.Fontsize = cdLeft.Fontsize + 1.0 // <= conflicting data
		ceLeft := NewContentEntry("idLeft", cdLeft)
		cdRight := getContentDataWithDummyData(t, "type")
		ceRight := NewContentEntry("idRight", cdRight)

		err = EntriesAreNotContradicting(&ceLeft, &ceRight)
		expectError(t, err)
	})
}

func TestContentEntry_AddMissingValuesFromOther(t *testing.T) {

	cdEmpty := ContentData{}
	cdAllSet := getContentDataWithDummyData(t, "type")

	t.Run("fill empty set from full set", func(t *testing.T) {
		ceSrc := NewContentEntry("idAllSet", cdAllSet)
		ceDst := NewContentEntry("idEmpty", cdEmpty)

		ceDst.AddMissingValuesFrom(&ceSrc)
		expectAllExportedSet(t, ceDst)

	})

	t.Run("do not overwrite existing data", func(t *testing.T) {
		ceSrc := NewContentEntry("src", ContentData{Desc: "srcDesc", Font: "srcFont", X1: 1.0, Y1: 2.0})
		ceDst := NewContentEntry("dst", ContentData{Desc: "dstDesc", X1: 3.0, X2: 4.0})
		ceDst.AddMissingValuesFrom(&ceSrc)

		expectEqual(t, ceDst.Description(), "dstDesc")
		expectEqual(t, ceDst.Font(), "srcFont")
		expectEqual(t, ceDst.X1(), 3.0)
		expectEqual(t, ceDst.Y1(), 2.0)
		expectEqual(t, ceDst.X2(), 4.0)
	})
}
