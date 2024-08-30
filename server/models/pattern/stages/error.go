package stages

import "github.com/khulnasoft/meshkit/errors"

const (
	ErrResolveReferenceCode = "meshplay-server-1361"
)

func ErrResolveReference(err error) error {
	return errors.New(ErrResolveReferenceCode, errors.Alert, []string{}, []string{err.Error()}, []string{}, []string{})
}
