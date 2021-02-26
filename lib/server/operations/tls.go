/*
Copyright IBM Corp All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package operations

import (
	//"crypto/tls"
	"github.com/Hyperledger-TWGC/ccs-gm/tls"
	"github.com/Hyperledger-TWGC/ccs-gm/x509"
	"github.com/Hyperledger-TWGC/ccs-gm/sm2"
	//"crypto/x509"
	"io/ioutil"
)

var (
	// DefaultTLSCipherSuites is the list of default cipher suites
	DefaultTLSCipherSuites = []uint16{
		tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,
		tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
		tls.TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256,
		tls.TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384,
		tls.TLS_RSA_WITH_AES_128_GCM_SHA256,
		tls.TLS_RSA_WITH_AES_256_GCM_SHA384,
	}
	// strong GM TLS cipher suites
	DefaultGMTLSCipherSuites = []uint16{
		tls.GMTLS_SM2_WITH_SM4_SM3,
		tls.GMTLS_ECDHE_SM2_WITH_SM4_SM3,
	}
)

// TLS contains the TLS configuration for the operations system serve
type TLS struct {
	Enabled            bool
	CertFile           string
	KeyFile            string
	ClientCertRequired bool
	ClientCACertFiles  []string
}

// Config returns TLS configuration
func (t *TLS) Config() (*tls.Config, error) {
	var tlsConfig *tls.Config

	if t.Enabled {
		cert, err := tls.LoadX509KeyPair(t.CertFile, t.KeyFile)
		if err != nil {
			return nil, err
		}
		caCertPool := x509.NewCertPool()
		for _, caPath := range t.ClientCACertFiles {
			caPem, err := ioutil.ReadFile(caPath)
			if err != nil {
				return nil, err
			}
			caCertPool.AppendCertsFromPEM(caPem)
		}
		// tlsConfig = &tls.Config{
		// 	Certificates: []tls.Certificate{cert},
		// 	CipherSuites: DefaultTLSCipherSuites,
		// 	ClientCAs:    caCertPool,
		// }

		tlsConfig := &tls.Config{}
		_, ok := cert.PrivateKey.(*sm2.PrivateKey)
		if ok {
			tlsConfig = &tls.Config{
				Certificates: []tls.Certificate{cert},
				CipherSuites: DefaultGMTLSCipherSuites,
				ClientCAs:    caCertPool,
				GMSupport: &tls.GMSupport{},
			}
		} else {
			tlsConfig = &tls.Config{
				Certificates: []tls.Certificate{cert},
				CipherSuites: DefaultTLSCipherSuites,
				ClientCAs:    caCertPool,
			}
		}
		if t.ClientCertRequired {
			tlsConfig.ClientAuth = tls.RequireAndVerifyClientCert
		} else {
			tlsConfig.ClientAuth = tls.VerifyClientCertIfGiven
		}
	}

	return tlsConfig, nil
}
