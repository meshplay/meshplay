package core

import (
	"io"
	"os"
	"path/filepath"

	log "github.com/sirupsen/logrus"

	"github.com/pkg/errors"
)

var (
	// ManifestsFolder is where the Kubernetes manifests are stored
	ManifestsFolder = "manifests"
	// MeshplayFolder is the default relative location of the meshplay config
	// related configuration files.
	MeshplayFolder = ".meshplay"

	ReleaseTag string
)

// SafeClose is a helper function help to close the io
func SafeClose(co io.Closer) {
	if cerr := co.Close(); cerr != nil {
		log.Error(cerr)
	}
}

// CreateManifestsFolder creates a new folder (.meshplay/manifests)
func CreateManifestsFolder() error {
	meshplayManifestFolder := filepath.Join(MeshplayFolder, ManifestsFolder)

	log.Debug("deleting " + ManifestsFolder + " folder...")
	// delete manifests folder if it already exists
	if err := os.RemoveAll(meshplayManifestFolder); err != nil {
		return err
	}
	log.Debug("creating " + ManifestsFolder + "folder...")
	// create a manifests folder under ~/.meshplay to store the manifest files
	if err := os.MkdirAll(meshplayManifestFolder, os.ModePerm); err != nil {
		return errors.Wrapf(err, "failed to make %s directory", ManifestsFolder)
	}
	log.Debug("created manifests folder...")

	return nil
}
