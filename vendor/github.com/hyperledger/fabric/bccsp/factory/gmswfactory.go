package factory

import (
	"errors"
	"fmt"

	"github.com/hyperledger/fabric/bccsp"
	"github.com/hyperledger/fabric/bccsp/sw"
)

const (
	// GuomiBasedFactoryName is the name of the factory of the software-based BCCSP implementation
	GMSoftwareBasedFactoryName = "GMSW"
)

// GMFactory is the factory of the guomi-based BCCSP.
type GMSWFactory struct{}

// Name returns the name of this factory
func (f *GMSWFactory) Name() string {
	return GMSoftwareBasedFactoryName
}

// Get returns an instance of BCCSP using Opts.
func (f *GMSWFactory) Get(config *FactoryOpts) (bccsp.BCCSP, error) {
	// Validate arguments
	if config == nil || config.SwOpts == nil {
		return nil, errors.New("Invalid config. It must not be nil.")
	}

	gmswOpts := config.SwOpts

	var ks bccsp.KeyStore
	if gmswOpts.Ephemeral == true {
		ks = sw.NewDummyKeyStore()
	} else if gmswOpts.FileKeystore != nil {
		fks, err := sw.NewFileBasedKeyStore(nil, gmswOpts.FileKeystore.KeyStorePath, false)
		if err != nil {
			return nil, fmt.Errorf("Failed to initialize gm software key store: %s", err)
		}
		ks = fks
	} else {
		// Default to DummyKeystore
		ks = sw.NewDummyKeyStore()
	}

	return sw.NewWithParams(gmswOpts.SecLevel, gmswOpts.HashFamily, ks)
}