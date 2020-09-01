package shared

import (
	"crypto"
	"crypto/tls"
	"crypto/x509"
	"errors"
	"fmt"
	"io/ioutil"
	"os"

	"go.mongodb.org/mongo-driver/mongo/options"
	"software.sslmate.com/src/go-pkcs12"
)

func setTLSOptions(opts *options.ClientOptions) (err error) {
	database := readenv("MONGODB_DATABASE")
	if database == "" {
		err = errors.New("MONGODB_DATABASE must be set")
		return
	}

	user := readenv("MONGODB_USER")
	password := readenv("MONGODB_PASSWORD")
	ca := readenvContents("MONGODB_CA_PATH")
	clientPfx := readenvContents("MONGODB_CLIENT_PFX_PATH")
	clientPfxPass := readenv("MONGODB_CLIENT_PFX_PASS")

	if len(user) != 0 || len(password) != 0 {
		opts.SetAuth(options.Credential{Username: user, Password: password, AuthSource: database})
	} else {
		fmt.Printf("No user / password authentication set.\n")
	}
	if len(ca) != 0 && len(clientPfx) != 0 {
		var key interface{}
		var cert *x509.Certificate
		if key, cert, _, err = pkcs12.DecodeChain(clientPfx, clientPfxPass); err != nil {
			return
		}
		tlsCert := tls.Certificate{
			Certificate: [][]byte{cert.Raw},
			PrivateKey:  key.(crypto.PrivateKey),
			Leaf:        cert,
		}
		roots := x509.NewCertPool()
		if ok := roots.AppendCertsFromPEM(ca); !ok {
			err = errors.New("failed to parse root certificate")
			return
		}

		cfg := &tls.Config{RootCAs: roots, Certificates: []tls.Certificate{tlsCert}}
		opts.SetTLSConfig(cfg)
	}

	return
}

func readenv(key string) (env string) {
	env = os.Getenv(key)
	if env != "" && env[0] == '<' {
		if contents, err := ioutil.ReadFile(env[1:]); err == nil {
			env = string(contents)
		} else {
			env = ""
		}
	}

	return
}

func readenvContents(key string) (contents []byte) {
	if env := readenv(key); env != "" {
		if contents, err := ioutil.ReadFile(env); err == nil {
			return contents
		}
	}

	return []byte{}
}
