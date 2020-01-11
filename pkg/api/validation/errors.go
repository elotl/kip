package validation

import (
	"fmt"

	"github.com/elotl/cloud-instance-provider/pkg/util/validation/field"
)

// +k8s:deepcopy-gen=false
type ValidationError struct {
	kind string
	name string
	errs field.ErrorList
}

func (e ValidationError) Error() string {
	return fmt.Sprintf("%s %s is invalid: %v",
		e.kind, e.name, e.errs.ToAggregate())
}

func NewError(kind, name string, errs field.ErrorList) *ValidationError {
	return &ValidationError{
		kind: kind,
		name: name,
		errs: errs,
	}
}
