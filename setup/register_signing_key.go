package setup

import (
	"bytes"
	"crypto/tls"
	b "encoding/base64"
	"encoding/json"
	"fmt"
	csetup "intel/isecl/lib/common/setup"
	conf "intel/isecl/wlagent/config"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

type RegisterSigningKey struct {
}

type SigningKey struct {
	Version        int    `json:"Version"`
	KeyAttestation string `json:"KeyAttestation"`
	PublicKey      string `json:"PublicKey"`
	KeySignature   string `json:"KeySignature"`
	KeyName        string `json:"KeyName"`
}
type SigningKeyCert struct {
	SigningKeyCertificate string `json:"signing_key_der_certificate"`
}

func (rs RegisterSigningKey) Run(c csetup.Context) error {

	var url string
	var requestBody []byte
	var signingkey SigningKey
	var tpmVersion string
	var originalNameDigest []byte
	var signingKeyCert SigningKeyCert

	url = "https://10.105.168.177:8443/mtwilson/v2/rpc/certify-host-signing-key"

	fileName := conf.GetSigningKeyFileName()
	_, err := os.Stat(fileName)
	if os.IsNotExist(err) {
		log.Fatal("signingkey file does not exist")
	}
	jsonFile, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
	}

	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)

	err = json.Unmarshal(byteValue, &signingkey)
	if err != nil {
		fmt.Println("error:", err)
	}

	tpmCertifyKeyBytes, _ := b.StdEncoding.DecodeString(strings.TrimSpace(signingkey.KeyAttestation))
	tpmCertifyKey := b.StdEncoding.EncodeToString(tpmCertifyKeyBytes[2:])

	originalNameDigest, _ = b.StdEncoding.DecodeString(strings.TrimSpace(signingkey.KeyName))
	originalNameDigest = originalNameDigest[1:]
	for i := 0; i < 34; i++ {
		originalNameDigest = append(originalNameDigest, 0)
	}

	nameDigest := b.StdEncoding.EncodeToString(originalNameDigest)

	operatingSystem := detectOS()
	fmt.Println(operatingSystem)
	if signingkey.Version == 2 {
		tpmVersion = "2.0"
	} else {
		tpmVersion = "1.2"
	}

	request := `{
		 "public_key_modulus":"` + signingkey.PublicKey + `",
	 	 "tpm_certify_key":"` + tpmCertifyKey + `",
	     "tpm_certify_key_signature":"` + signingkey.KeySignature + `",
	 	 "aik_der_certificate":"MIICzjCCAbagAwIBAgIGAWfSQrjuMA0GCSqGSIb3DQEBCwUAMBsxGTAXBgNVBAMTEG10d2lsc29uLXBjYS1haWswHhcNMTgxMjIxMTkzNDA3WhcNMjgxMjIwMTkzNDA3WjAAMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAlEFutiERjj4TfP92T2YAmvCPnnb04ht+n0mrKB2/PvAjufgogS/1Vds8mGuT0gl8uvSaBI02HVMHAQTLlCCYcgo689ArlvrmPA9nwKhv7gb22GC64tU+4CgDyp5V8Km3w/ho0xl0m3QUqKO6l8Zwzl8kUUQWoz22pQsO7Yz61p0a+GOziRLdYCvR8W/QNbNlPSfWwVocVSo0V4itnxC3aX3J1wdw8XyyHW/2rS9wjcDOpZ45Fc5Itkxc0gKrUxHkvMiFW/Uy+fsuKDNxju3rPA+49xSeoVxp3IlyQLVxpR2Jr2/a53OZjwBOl5AigCesqKY/Ityq56Zi/STjyEEnEwIDAQABozMwMTAvBgNVHREBAf8EJTAjgSEAC/1n/f1Q/Tv9/WwIeVH9/f11fDz9/XD9YFUm+P0eRi8wDQYJKoZIhvcNAQELBQADggEBAFrsvWI/1fI6J2swpgUiIhfds3vMjc0J31BJp46a900Vd+awko726Lbsx43xwV0jlrTRiWX4StpEEQXcVF+TTDIgd4GSc5qXN8N4vcOQDl5j4Yg2tsLm3FAFppVCLO8rC1D9UdhM0K63sY8Xz92IIGINnqQTslHPmlGPJ9lTgBkWOu/rzicY394g/czdVa1l36KSLkCpwnB5b1RQAfPUVWSGlzdKIvmb/+F9Ur6VOPZ1CpuIJLgtVhiZVMscZYHSX0kyT3ayQj1tJTT9x5VgZW15Pdnj3lh6TL36TUmd/KpEFN37jFdnLDXynE/QDyj+neyPrM5g3rHFvsagJr5rkTY=",
	 	 "name_digest":"` + nameDigest + `",
	     "tpm_version":"` + tpmVersion + `",
		 "operating_system":"Linux"}`
	fmt.Println(request)
	requestBody = []byte(`{
		 "public_key_modulus":"` + signingkey.PublicKey + `",
	 	 "tpm_certify_key":"` + tpmCertifyKey + `",
	     "tpm_certify_key_signature":"` + signingkey.KeySignature + `",
	 	 "aik_der_certificate":"MIICzjCCAbagAwIBAgIGAWfSQrjuMA0GCSqGSIb3DQEBCwUAMBsxGTAXBgNVBAMTEG10d2lsc29uLXBjYS1haWswHhcNMTgxMjIxMTkzNDA3WhcNMjgxMjIwMTkzNDA3WjAAMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAlEFutiERjj4TfP92T2YAmvCPnnb04ht+n0mrKB2/PvAjufgogS/1Vds8mGuT0gl8uvSaBI02HVMHAQTLlCCYcgo689ArlvrmPA9nwKhv7gb22GC64tU+4CgDyp5V8Km3w/ho0xl0m3QUqKO6l8Zwzl8kUUQWoz22pQsO7Yz61p0a+GOziRLdYCvR8W/QNbNlPSfWwVocVSo0V4itnxC3aX3J1wdw8XyyHW/2rS9wjcDOpZ45Fc5Itkxc0gKrUxHkvMiFW/Uy+fsuKDNxju3rPA+49xSeoVxp3IlyQLVxpR2Jr2/a53OZjwBOl5AigCesqKY/Ityq56Zi/STjyEEnEwIDAQABozMwMTAvBgNVHREBAf8EJTAjgSEAC/1n/f1Q/Tv9/WwIeVH9/f11fDz9/XD9YFUm+P0eRi8wDQYJKoZIhvcNAQELBQADggEBAFrsvWI/1fI6J2swpgUiIhfds3vMjc0J31BJp46a900Vd+awko726Lbsx43xwV0jlrTRiWX4StpEEQXcVF+TTDIgd4GSc5qXN8N4vcOQDl5j4Yg2tsLm3FAFppVCLO8rC1D9UdhM0K63sY8Xz92IIGINnqQTslHPmlGPJ9lTgBkWOu/rzicY394g/czdVa1l36KSLkCpwnB5b1RQAfPUVWSGlzdKIvmb/+F9Ur6VOPZ1CpuIJLgtVhiZVMscZYHSX0kyT3ayQj1tJTT9x5VgZW15Pdnj3lh6TL36TUmd/KpEFN37jFdnLDXynE/QDyj+neyPrM5g3rHFvsagJr5rkTY=",
	 	 "name_digest":"` + nameDigest + `",
	     "tpm_version":"` + tpmVersion + `",
		 "operating_system":"Linux"}`)

	// set POST request Accept, Content-Type and Authorization headers
	httpRequest, err := http.NewRequest("POST", url, bytes.NewBuffer(requestBody))
	httpRequest.Header.Set("Accept", "application/json")
	httpRequest.Header.Set("Content-Type", "application/json")

	//*******Remove hardcoded values
	httpRequest.SetBasicAuth("admin", "password")
	httpResponse, err := sendRequest(httpRequest)
	if err != nil {
		log.Fatal("Error in signing key registration.", err)
	}
	_ = json.Unmarshal([]byte(httpResponse), &signingKeyCert)
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Println(bindingKeyCert.BindingKeyCertificate)
	aikPem := "-----BEGIN CERTIFICATE-----" + "\n" + signingKeyCert.SigningKeyCertificate + "\n" + "-----END CERTIFICATE-----" + "\n"
	file, err := os.Create("signingkeycert.pem")
	if err != nil {
		log.Fatal("Error in creating file.", err)
	}
	_, err = file.Write([]byte(aikPem))
	if err != nil {
		log.Fatal("Error in writing to file.", err)
	}

}
func detectOS() string {
	if os.PathSeparator == '\\' && os.PathListSeparator == ';' {
		return "Windows"
	}
	return "Linux"

}
func sendRequest(req *http.Request) ([]byte, error) {
	tlsConfig := tls.Config{
		InsecureSkipVerify: true,
	}
	transport := http.Transport{
		TLSClientConfig: &tlsConfig,
	}
	client := &http.Client{
		Transport: &transport,
	}
	response, err := client.Do(req)
	if err != nil {
		log.Println("Error in sending request.", err)
		return nil, err
	}
	defer response.Body.Close()

	//create byte array of HTTP response body
	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		return nil, err
	}

	log.Println("status code returned : ", strconv.Itoa(response.StatusCode))
	return body, nil
}

// Validate checks whether or not the Register Signing Key task was completed successfully
func (rs RegisterSigningKey) Validate(c csetup.Context) error {
	return nil
}
