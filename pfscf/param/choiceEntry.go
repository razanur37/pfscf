package param

import (
	"fmt"
	"strings"

	"github.com/razanur37/pfscf/pfscf/args"
	"github.com/razanur37/pfscf/pfscf/utils"
)

const (
	typeChoice = "choice"
)

type choiceEntry struct {
	commonFields

	TheExample     string   `yaml:"example"`
	TheDescription string   `yaml:"description"`
	TheChoices     []string `yaml:"choices"`
}

func (e *choiceEntry) Type() string {
	return typeChoice
}

func (e *choiceEntry) Example() string {
	if utils.IsSet(e.TheExample) {
		return e.TheExample
	}
	utils.Assert(len(e.TheChoices) > 0, "Validation should have ensured that there is at least one choice")
	return e.TheChoices[0]
}

func (e *choiceEntry) Description() string {
	return e.TheDescription
}

func (e *choiceEntry) AcceptedValues() []string {
	return e.TheChoices
}

func (e *choiceEntry) deepCopy() Entry {
	copy := *e

	copy.TheChoices = append(make([]string, 0), e.TheChoices...)

	return &copy
}

func (e *choiceEntry) isValid() (err error) {
	// missing example is ok, as we then simply take the first provided choice

	if !utils.IsSet(e.TheDescription) {
		return fmt.Errorf("Missing description")
	}
	if len(e.TheChoices) == 0 {
		return fmt.Errorf("Missing choices")
	}
	return nil
}

func (e *choiceEntry) validateAndProcessArgs(as *args.Store) error {
	argValue, exists := as.Get(e.ID())
	utils.Assert(exists, "Existence of entry should have been validated by caller")

	splitArgs := utils.SplitAndTrim(argValue, ",")

	for _, splitArg := range splitArgs {
		if !utils.Contains(e.TheChoices, splitArg) {
			return fmt.Errorf("Invalid choice '%v' was provided. Valid choices are: %v", splitArg, e.TheChoices)
		}
	}

	return nil
}

func (e *choiceEntry) describe(verbose bool) (result string) {
	var sb strings.Builder

	if !verbose {
		fmt.Fprintf(&sb, "- %v: %v\n", e.id, e.Description())
	} else {
		fmt.Fprintf(&sb, "- %v\n", e.id)
		fmt.Fprintf(&sb, "\tDesc: %v\n", e.Description())
		fmt.Fprintf(&sb, "\tType: %v\n", e.Type())
		fmt.Fprintf(&sb, "\tAllowed Choices: %v\n", e.TheChoices)
		fmt.Fprintf(&sb, "\tExample: %v\n", genericContentUsageExample(e.id, e.Example()))
	}

	return sb.String()
}
