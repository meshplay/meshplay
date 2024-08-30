package main

import "github.com/layer5io/meshkit/errors"

// Please reference the following before contributing an error code:
// https://docs.meshplay.khulnasofy.com/project/contributing/contributing-error
// https://github.com/meshplay/meshkit/blob/master/errors/errors.go
const (
	ErrCreatingUUIDInstanceCode                   = "meshplay-server-1001"
	ErrRegisteringMeshplayOAMTraitsCode            = "meshplay-server-1002"
	ErrRegisteringMeshplayOAMWorkloadsCode         = "meshplay-server-1003"
	ErrRetrievingUserHomeDirectoryCode            = "meshplay-server-1004"
	ErrCreatingUserDataDirectoryCode              = "meshplay-server-1005"
	ErrCreatingMapPreferencePersisterInstanceCode = "meshplay-server-1006"
	ErrDatabaseAutoMigrationCode                  = "meshplay-server-1007"
	ErrInvalidURLSkippingProviderCode             = "meshplay-server-1008"
	ErrListenAndServeCode                         = "meshplay-server-1009"
	ErrCleaningUpLocalProviderCode                = "meshplay-server-1010"
	ErrClosingDatabaseInstanceCode                = "meshplay-server-1011"
	ErrInitializingRegistryManagerCode            = "meshplay-server-1012"
	ErrInitializingKeysRegistrationCode           = "meshplay-server-1013"
	ErrCreatingOPAInstanceCode                    = "meshplay-server-1323"
)

func ErrInitializingRegistryManager(err error) error {
	return errors.New(ErrInitializingRegistryManagerCode, errors.Fatal, []string{"could not initialize registry manager"}, []string{err.Error()}, []string{"could not migrate tables into the database"}, []string{"make sure the database instance passed is not nil"})
}

func ErrCreatingUUIDInstance(err error) error {
	return errors.New(ErrCreatingUUIDInstanceCode, errors.Fatal, []string{"Unable to create UUID instance"}, []string{"Unable to create UUID instance: ", err.Error()}, []string{}, []string{})
}

func ErrRegisteringMeshplayOAMTraits(err error) error {
	return errors.New(ErrRegisteringMeshplayOAMTraitsCode, errors.Alert, []string{"Error registering local OAM traits"}, []string{"Error registering local OAM traits: ", err.Error()}, []string{}, []string{})
}

func ErrRegisteringMeshplayOAMWorkloads(err error) error {
	return errors.New(ErrRegisteringMeshplayOAMWorkloadsCode, errors.Alert, []string{"Error registering local OAM workloads"}, []string{"Error registering local OAM workloads: ", err.Error()}, []string{}, []string{})
}

func ErrRetrievingUserHomeDirectory(err error) error {
	return errors.New(ErrRetrievingUserHomeDirectoryCode, errors.Fatal, []string{"Unable to retrieve the user's home directory"}, []string{"Unable to retrieve the user's home directory: ", err.Error()}, []string{}, []string{})
}

func ErrCreatingUserDataDirectory(dir string) error {
	return errors.New(ErrCreatingUserDataDirectoryCode, errors.Fatal, []string{"Unable to create the directory for storing user data at: ", dir}, []string{"Unable to create the directory for storing user data at: ", dir}, []string{}, []string{})
}

func ErrCreatingMapPreferencePersisterInstance(err error) error {
	return errors.New(ErrCreatingMapPreferencePersisterInstanceCode, errors.Fatal, []string{"Unable to create a new MapPreferencePersister instance"}, []string{"Unable to create a new MapPreferencePersister instance: ", err.Error()}, []string{}, []string{})
}

func ErrDatabaseAutoMigration(err error) error {
	return errors.New(ErrDatabaseAutoMigrationCode, errors.Fatal, []string{"Unable to auto migrate to database"}, []string{"Unable to auto migrate to database: ", err.Error()}, []string{}, []string{})
}

func ErrInvalidURLSkippingProvider(url string) error {
	return errors.New(ErrInvalidURLSkippingProviderCode, errors.Alert, []string{url, " is invalid url skipping provider"}, []string{url, " is invalid url skipping provider"}, []string{}, []string{})
}

func ErrListenAndServe(err error) error {
	return errors.New(ErrListenAndServeCode, errors.Fatal, []string{"ListenAndServe Error"}, []string{"ListenAndServe Error: ", err.Error()}, []string{}, []string{})
}

func ErrCleaningUpLocalProvider(err error) error {
	return errors.New(ErrCleaningUpLocalProviderCode, errors.Alert, []string{"Error cleaning up local provider"}, []string{"Error cleaning up local provider: ", err.Error()}, []string{}, []string{})
}

func ErrClosingDatabaseInstance(err error) error {
	return errors.New(ErrClosingDatabaseInstanceCode, errors.Alert, []string{"Error closing database instance"}, []string{"Error closing database instance: ", err.Error()}, []string{}, []string{})
}

func ErrInitializingKeysRegistration(err error) error {
	return errors.New(ErrInitializingKeysRegistrationCode, errors.Fatal, []string{"could not initialize keys registry manager"}, []string{err.Error()}, []string{"could not migrate tables into the database"}, []string{"make sure the database instance passed is not nil"})
}

func ErrCreatingOPAInstance(err error) error {
	return errors.New(ErrCreatingOPAInstanceCode, errors.Alert, []string{"Error creating OPA Instance."}, []string{err.Error()}, []string{"Unable to create OPA instance, policies will not be evaluated."}, []string{"Ensure relationships are registered"})
}