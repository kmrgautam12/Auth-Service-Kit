package authmech

import (
	"Pay-AI/financial-transaction-server/constantservice"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"
)

func GeneratePublicPrivateKeys() (*rsa.PrivateKey, *rsa.PublicKey, error) {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		fmt.Println("Error 1 - ", err)
		return nil, nil, err
	}
	publicKey := &privateKey.PublicKey
	SavePrivatePemKeys(privateKey)
	SavePublicKeyPem(publicKey)
	return privateKey, publicKey, err
}

func SavePrivatePemKeys(privateKey *rsa.PrivateKey) (bool, error) {
	fileOut, err := os.Create(constantservice.PEM_PRIVATE_PATH)
	if err != nil {
		return false, err
	}
	defer fileOut.Close()
	b := x509.MarshalPKCS1PrivateKey(privateKey)
	var pemFile = pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: b,
	}

	err = pem.Encode(fileOut, &pemFile)
	if err != nil {
		return false, err
	}
	return true, nil

}

func SavePublicKeyPem(publicKey *rsa.PublicKey) (bool, error) {

	fileOut, err := os.Create(constantservice.PEM_PUBLIC_PATH)
	if err != nil {
		return false, err
	}
	defer fileOut.Close()

	b := x509.MarshalPKCS1PublicKey(publicKey)

	p := pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: b,
	}
	err = pem.Encode(fileOut, &p)
	if err != nil {
		return false, err
	}
	return true, nil

}
