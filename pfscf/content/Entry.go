package content

import (
	"fmt"

	"github.com/razanur37/pfscf/pfscf/args"
	"github.com/razanur37/pfscf/pfscf/canvas"
	"github.com/razanur37/pfscf/pfscf/param"
	"github.com/razanur37/pfscf/pfscf/preset"
	"github.com/razanur37/pfscf/pfscf/stamp"
)

// Entry is an interface for the content. D'oh!
type Entry interface {
	isValid(*param.Store, *canvas.Store) (err error)
	resolve(ps preset.Store) (err error)
	generateOutput(s *stamp.Stamp, as *args.Store) (err error)
	deepCopy() Entry
}

// TODO now with no ID we should print all fields of the respective entry instead (don't forget the type)
func contentValErr(ce Entry, errIn error) (errOut error) {
	return fmt.Errorf("Error validating content: %v; complete content entry is: %v", errIn, ce)
}
