/*
Copyright SecureKey Technologies Inc. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package defcore

import (
	"github.com/xuchengli/fabric-sdk-go/pkg/common/logging"
	"github.com/xuchengli/fabric-sdk-go/pkg/common/providers/core"
	"github.com/xuchengli/fabric-sdk-go/pkg/common/providers/fab"
	"github.com/xuchengli/fabric-sdk-go/pkg/core/logging/api"

	cryptosuiteimpl "github.com/xuchengli/fabric-sdk-go/pkg/core/cryptosuite/bccsp/sw"
	cryptosuiteimplGM "github.com/xuchengli/fabric-sdk-go/pkg/core/cryptosuite/bccsp/gm"
	signingMgr "github.com/xuchengli/fabric-sdk-go/pkg/fab/signingmgr"
	"github.com/xuchengli/fabric-sdk-go/pkg/fabsdk/provider/fabpvdr"

	"github.com/xuchengli/fabric-sdk-go/pkg/core/logging/modlog"
)

var logger = logging.NewLogger("fabsdk")

// ProviderFactory represents the default SDK provider factory.
type ProviderFactory struct {
}

// NewProviderFactory returns the default SDK provider factory.
func NewProviderFactory() *ProviderFactory {
	f := ProviderFactory{}
	return &f
}

// CreateCryptoSuiteProvider returns a new default implementation of BCCSP
func (f *ProviderFactory) CreateCryptoSuiteProvider(config core.CryptoSuiteConfig) (core.CryptoSuite, error) {
	if config.SecurityProvider() != "sw"{
		logger.Warnf("default provider factory doesn't support '%s' crypto provider", config.SecurityProvider())
	}

	if config.SecurityHashAlgorithm() == "SHA2" {
		cryptoSuiteProvider, err := cryptosuiteimpl.GetSuiteByConfig(config)
		return cryptoSuiteProvider, err
	}else if config.SecurityHashAlgorithm() == "SM3" {
		cryptoSuiteProvider, err := cryptosuiteimplGM.GetSuiteByConfig(config)
		return cryptoSuiteProvider, err
	}else{
		cryptoSuiteProvider, err := cryptosuiteimpl.GetSuiteWithDefaultEphemeral()
		return cryptoSuiteProvider, err
	}
}

// CreateSigningManager returns a new default implementation of signing manager
func (f *ProviderFactory) CreateSigningManager(cryptoProvider core.CryptoSuite) (core.SigningManager, error) {
	return signingMgr.New(cryptoProvider)
}

// CreateGMSigningManager return a new default implementation of signing manager of gm
func (f * ProviderFactory) CreateGMSigningManager(cryptoProvider core.CryptoSuite) (core.SigningManager, error) {
	return signingMgr.NewGM(cryptoProvider)
}

// CreateInfraProvider returns a new default implementation of fabric primitives
func (f *ProviderFactory) CreateInfraProvider(config fab.EndpointConfig) (fab.InfraProvider, error) {
	return fabpvdr.New(config), nil
}

// NewLoggerProvider returns a new default implementation of a logger backend
// This function is separated from the factory to allow logger creation first.
func NewLoggerProvider() api.LoggerProvider {
	return modlog.LoggerProvider()
}
