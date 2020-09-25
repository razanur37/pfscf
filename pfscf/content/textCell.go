package content

import (
	"fmt"

	"github.com/Blesmol/pfscf/pfscf/args"
	"github.com/Blesmol/pfscf/pfscf/canvas"
	"github.com/Blesmol/pfscf/pfscf/param"
	"github.com/Blesmol/pfscf/pfscf/preset"
	"github.com/Blesmol/pfscf/pfscf/stamp"
	"github.com/Blesmol/pfscf/pfscf/utils"
)

const (
	typeTextCell = "textCell"
)

// textCell is the final type to implement textCells.
// TODO switch to pointers to distinguish between unset values and zero values?
type textCell struct {
	Value    string
	X, Y     float64
	X2, Y2   float64
	Font     string
	Fontsize float64
	Align    string
	Canvas   string
	Presets  []string
}

func newTextCell() *textCell {
	var e textCell
	e.Presets = make([]string, 0)
	return &e
}

// isValid checks whether the current content object is valid and returns an
// error with details if the object is not valid.
func (e *textCell) isValid(paramStore *param.Store, canvasStore *canvas.Store) (err error) {
	err = utils.CheckFieldsAreSet(e, "Value", "Font", "Fontsize", "Canvas")
	if err != nil {
		return contentValErr(e, err)
	}

	err = utils.CheckFieldsAreInRange(e, 0.0, 100.0, "X", "Y", "X2", "Y2")
	if err != nil {
		return contentValErr(e, err)
	}

	if e.X == e.X2 {
		err = fmt.Errorf("Coordinates for X axis are equal: %v", e.X)
		return contentValErr(e, err)
	}

	if e.Y == e.Y2 {
		err = fmt.Errorf("Coordinates for Y axis are equal: %v", e.Y)
		return contentValErr(e, err)
	}

	if _, exists := canvasStore.Get(e.Canvas); !exists {
		err = fmt.Errorf("Canvas '%v' does not exist", e.Canvas)
		return contentValErr(e, err)
	}

	return nil
}

// resolve the presets for this content object.
func (e *textCell) resolve(ps preset.Store) (err error) {
	// check that required presets are not contradicting each other
	if err = ps.PresetsAreNotContradicting(e.Presets...); err != nil {
		err = fmt.Errorf("Error resolving content: %v", err)
		return
	}

	// apply presets
	for _, presetID := range e.Presets {
		preset, _ := ps.Get(presetID)
		if err = preset.FillPublicFieldsFromPreset(e, "Presets"); err != nil {
			err = fmt.Errorf("Error resolving content: %v", err)
			return
		}
	}

	return nil
}

// generateOutput generates the output for this textCell object.
func (e *textCell) generateOutput(s *stamp.Stamp, as *args.Store) (err error) {
	value := getValue(e.Value, as)
	if value == nil {
		return nil // nothing to do here...
	}

	y2 := s.DeriveY2(e.Canvas, e.Y, e.Y2, e.Fontsize)
	s.AddTextCell(e.Canvas, e.X, e.Y, e.X2, y2, e.Font, e.Fontsize, e.Align, *value, true)

	return nil
}

// deepCopy creates a deep copy of this entry.
// TODO create generic deep-copy function for public fields
func (e *textCell) deepCopy() Entry {
	var copy textCell
	utils.AddMissingValues(&copy, *e, "Presets")
	for _, preset := range e.Presets {
		copy.Presets = append(copy.Presets, preset)
	}

	return &copy
}
