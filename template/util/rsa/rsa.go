package rsa

import "github.com/devfeel/dotcode/template"

func init(){
	template.RegisterTemplate(new(File))
}


type File struct{}

func(t *File) Path() string{
	return "util/rsa"
}

func (t *File) FileName() string{
	return "rsa.go"
}

func (t *File) String() string {
	return `package _rsa

import (
	"crypto/x509"
	"crypto/rsa"
	"crypto/rand"
	"encoding/pem"
	"errors"
)

// RsaEncrypt
func RsaEncrypt(origData []byte, publicKey []byte) ([]byte, error) {
	block, _ := pem.Decode(publicKey)
	if block == nil {
		return nil, errors.New("public key error")
	}
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	pub := pubInterface.(*rsa.PublicKey)
	return rsa.EncryptPKCS1v15(rand.Reader, pub, origData)
}

// RsaDecrypt
// key-len: 1024
// PKCS#1
func RsaDecrypt(ciphertext []byte, privateKey []byte) ([]byte, error) {
	block, _ := pem.Decode(privateKey)
	if block == nil {
		return nil, errors.New("private key error!")
	}
	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	return rsa.DecryptPKCS1v15(rand.Reader, priv, ciphertext)
}`
}


