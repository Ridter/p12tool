package work

import (
	"bytes"
	"crypto/tls"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"github.com/urfave/cli/v2"
	"io/ioutil"
	"log"
	"p12tool/util"
	"p12tool/vars"

	"golang.org/x/crypto/pkcs12"
)

func ParseP12file(ctx *cli.Context) (err error)  {
	Parse(ctx)
	vars.Logger = util.NewLogger(vars.DebugMode, "")
	p12Bytes, err := ioutil.ReadFile(vars.Cert)
	if err != nil {
		vars.Logger.Log.Error("[-] Please input cert file")
		return nil
	}
	pass := vars.Pass
	if len(pass) == 0{
		vars.Logger.Log.Error("[-] Please input a password")
		return nil
	}
	// P12 to PEM
	blocks,  err := pkcs12.ToPEM(p12Bytes, pass)

	if err != nil{
		if err == pkcs12.ErrIncorrectPassword{
			vars.Logger.Log.Error("[-] Password incorrect")
		}else{
			vars.Logger.Log.Error("[-] Check your file is P12 cert file !!")
		}
		return nil
	}
	// Append all PEM Blocks together
	var pemData []byte
	for _, b := range blocks {
		pemData = append(pemData, pem.EncodeToMemory(b)...)
	}
	vars.Logger.Log.Notice("[+] Parse cert file ok ;)")
	//println(PemPrivateKeyFromPem(string(pemData)))
	//println(PemCertFromPem(string(pemData)))
	block, rest := pem.Decode([]byte(PemCertFromPem(string(pemData))))
	if block == nil || len(rest) > 0 {
		log.Fatal("Certificate decoding error")
		return nil
	}
	cert, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("=============================INFO===========================\nCommonName:\t\t%s\nCountry:\t\t%s\nOrganization:\t\t%s\nLocality:\t\t%s\nNotAfter:\t\t%s\n============================================================\n",cert.Subject.CommonName,cert.Subject.Country,cert.Subject.Organization,cert.Subject.Locality,cert.NotAfter)
	for i := range cert.ExtKeyUsage{
		if cert.ExtKeyUsage[i] == 0 || cert.ExtKeyUsage[i] == 3{
			vars.Logger.Log.Notice("[+] Can be used for code signing！！;)")
			return nil
		}
	}
	vars.Logger.Log.Error("[-] Not suitable for code signing ;(")
	return err
}


func PemPrivateKeyFromPem(data string) string {
	pemBytes := []byte(data)

	// Use tls lib to construct tls certificate and key object from PEM data
	// The tls.X509KeyPair function is smart enough to parse combined cert and key pem data
	certAndKey, err := tls.X509KeyPair(pemBytes, pemBytes)
	if err != nil {
		panic(err)
	}

	// Get parsed private key as PKCS8 data
	privBytes, err := x509.MarshalPKCS8PrivateKey(certAndKey.PrivateKey)
	if err != nil {
		panic(fmt.Sprintf("Unable to marshal private key: %v", err))
	}

	// Encode just the private key back to PEM and return it
	var privPem bytes.Buffer
	if err := pem.Encode(&privPem, &pem.Block{Type: "PRIVATE KEY", Bytes: privBytes}); err != nil {
		panic(fmt.Sprintf("Failed to write data: %s", err))
	}

	return privPem.String()
}

func PemCertFromPem(data string) string {
	pemBytes := []byte(data)

	// Use tls lib to construct tls certificate and key object from PEM data
	// The tls.X509KeyPair function is smart enough to parse combined cert and key pem data
	certAndKey, err := tls.X509KeyPair(pemBytes, pemBytes)
	if err != nil {
		panic(fmt.Sprintf("Error generating X509KeyPair: %v", err))
	}

	leaf, err := x509.ParseCertificate(certAndKey.Certificate[0])
	if err != nil {
		panic(err)
	}

	// Encode just the leaf cert as pem
	var certPem bytes.Buffer
	if err := pem.Encode(&certPem, &pem.Block{Type: "CERTIFICATE", Bytes: leaf.Raw}); err != nil {
		panic(fmt.Sprintf("Failed to write data: %s", err))
	}

	return certPem.String()
}
