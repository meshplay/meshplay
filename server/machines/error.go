package machines

import (
	"fmt"

	"github.com/layer5io/meshkit/errors"
)

const (
	ErrInvalidTransitionCode      = "meshplay-server-1214"
	ErrInvalidTransitionEventCode = "meshplay-server-1215"
	ErrInititalizeK8sMachineCode  = "meshplay-server-1216"
	ErrAssetMachineCtxCode        = "meshplay-server-1217"
	ErrInvalidTypeCode            = "meshplay-server-1218"
)

func ErrInvalidTransition(from, to StateType) error {
	return errors.New(ErrInvalidTransitionCode, errors.Alert, []string{fmt.Sprintf("transition restricted from \"%s\" to \"%s\"", from, to)}, []string{}, []string{}, []string{})
}

func ErrInvalidTransitionEvent(from StateType, event EventType) error {
	return errors.New(ErrInvalidTransitionEventCode, errors.Alert, []string{fmt.Sprintf("unsupported transition event received \"%s\" for the state \"%s\"", event, from)}, []string{}, []string{}, []string{})
}

func ErrInititalizeK8sMachine(err error) error {
	return errors.New(ErrInititalizeK8sMachineCode, errors.Alert, []string{"Provided connection id is invalid"}, []string{err.Error()}, []string{"Provided ID is not a valid uuid."}, []string{"Hard delete and reinitialise the connection process."})
}

func ErrAssertMachineCtx(err error) error {
	return errors.New(ErrAssetMachineCtxCode, errors.Alert, []string{"type assertion of context to *kubernetes.MachineCtx failed"}, []string{err.Error()}, []string{"The machine context has become invalid."}, []string{"Retry the opereation or try archiving and reinitializing the connection"})
}

func ErrInvalidType(err error) error {
	return errors.New(ErrInvalidTypeCode, errors.Alert, []string{"Provided connection id is invalid"}, []string{err.Error()}, []string{"Provided ID is not a valid uuid."}, []string{"Hard delete and reinitialise the connection process."})
}
