package testdata

import (
	"encoding/base64"
	"fmt"

	"github.com/gokalkan/gokalkan"
	"github.com/gokalkan/gokalkan/ckalkan"
	"github.com/gokalkan/gokalkan/types"
)

func ExampleClient_Sign() {
	opts := gokalkan.OptsTest

	cli, _ := gokalkan.NewClient(opts...)

	// Обязательно закрывайте клиент, иначе приведет к утечкам ресурсов
	defer cli.Close()

	//Тестовый RSA ключ от НУЦ РК
	keyPath := "./test/certs/gost1.p12"
	keyPassword := "Qwerty12"

	cli.LoadKeyStore(keyPath, keyPassword)

	signData, _ := cli.Sign(&types.SignInput{
		DataBytes:  []byte("Hello"),
		IsDetached: false,
		WithTSP:    false,
	})

	signInBase64 := base64.StdEncoding.EncodeToString(signData)

	fmt.Printf("Подписанные данные:\n%s", signInBase64)
	// Output:
	// Подписанные данные:
	// ........
}

func ExampleClient_Verify() {
	opts := gokalkan.OptsTest

	cli, _ := gokalkan.NewClient(opts...)

	// Обязательно закрывайте клиент, иначе приведет к утечкам ресурсов
	defer cli.Close()

	//Тестовый RSA ключ от НУЦ РК
	keyPath := "./test/certs/gost1.p12"
	keyPassword := "Qwerty12"

	cli.LoadKeyStore(keyPath, keyPassword)

	signData, _ := cli.Sign(&types.SignInput{
		DataBytes:  []byte("Hello"),
		IsDetached: false,
		WithTSP:    false,
	})

	ver, _ := cli.Verify(&types.VerifyInput{
		SignatureBytes:    signData,
		DataBytes:         []byte("Hello World!"),
		IsDetached:        false,
		MustCheckCertTime: false,
	})
	fmt.Println(ver)
	// Output:
	// Signature N 1
	// - CAdES-BES: verify signer certificate hash - OK.
	// Verify - OK
	// CMS Verify - OK
}

func ExampleClient_GetTimeFromSig() {
	opts := gokalkan.OptsTest

	cli, _ := gokalkan.NewClient(opts...)

	// Обязательно закрывайте клиент, иначе приведет к утечкам ресурсов
	defer cli.Close()

	//Тестовый RSA ключ от НУЦ РК
	keyPath := "./test/certs/gost1.p12"
	keyPassword := "Qwerty12"

	cli.LoadKeyStore(keyPath, keyPassword)

	signData, _ := cli.cli.Sign(&types.SignInput{
		DataBytes:  []byte("Hello"),
		IsDetached: false,
		WithTSP:    true,
	})

	signInBase64 := base64.StdEncoding.EncodeToString(signData)

	time, _ := cli.GetTimeFromSig(signInBase64, true)

	fmt.Printf("Время подписания: %s\n", time)
	// Output:
	// Время подписания: time....
}

func ExampleClient_SignWSSE() {
	opts := gokalkan.OptsTest

	cli, _ := gokalkan.NewClient(opts...)

	// Обязательно закрывайте клиент, иначе приведет к утечкам ресурсов
	defer cli.Close()

	//Тестовый RSA ключ от НУЦ РК
	keyPath := "./test/certs/gost1.p12"
	keyPassword := "Qwerty12"

	// Подгружаем хранилище с паролем
	cli.LoadKeyStore(keyPath, keyPassword)

	signedWSSE, _ := cli.SignWSSE("<root>this is sample</root>", "1")

	fmt.Printf("Подписанный документ WSSE: %s\n", signedWSSE)
	// Output:
	// Подписанный документ WSSE:
}

func ExampleClient_X509ExportCertificateFromStore() {
	opts := gokalkan.OptsTest

	cli, _ := gokalkan.NewClient(opts...)

	// Обязательно закрывайте клиент, иначе приведет к утечкам ресурсов
	defer cli.Close()

	//Тестовый RSA ключ от НУЦ РК
	keyPath := "./test/certs/gost1.p12"
	keyPassword := "Qwerty12"

	// Подгружаем хранилище с паролем
	cli.LoadKeyStore(keyPath, keyPassword)

	cert, _ := cli.X509ExportCertificateFromStore()

	fmt.Printf("Сертификат: %s\n", cert.Issuer)
	// Output:
	// Сертификат: CN=ҰКО 3.0 (GOST TEST),C=KZ
}

func ExampleClient_X509CertificateGetInfo() {
	opts := gokalkan.OptsTest

	cli, _ := gokalkan.NewClient(opts...)

	// Обязательно закрывайте клиент, иначе приведет к утечкам ресурсов
	defer cli.Close()

	//Тестовый RSA ключ от НУЦ РК
	keyPath := "./test/certs/gost1.p12"
	keyPassword := "Qwerty12"

	// Подгружаем хранилище с паролем
	cli.LoadKeyStore(keyPath, keyPassword)

	// Экспортируем сертификат из хранилища
	cert, _ := cli.X509ExportCertificateFromStore()

	info, _ := cli.X509CertificateGetInfo(cert)

	fmt.Printf("Информация по сертификату: %s\n", info)
	// Output:
	// Информация по сертификату:
}

func ExampleClient_GetCertFromXML() {
	opts := gokalkan.OptsTest

	cli, _ := gokalkan.NewClient(opts...)

	// Обязательно закрывайте клиент, иначе приведет к утечкам ресурсов
	defer cli.Close()

	//Тестовый RSA ключ от НУЦ РК
	keyPath := "./test/certs/gost1.p12"
	keyPassword := "Qwerty12"

	// Подгружаем хранилище с паролем
	cli.LoadKeyStore(keyPath, keyPassword)

	signedXML, _ := cli.SignXML(&types.SignXMLInput{
		Data:    "<root>Hello</root>",
		WithTSP: true,
	})

	cert, _ := cli.GetCertFromXML(signedXML, 0)

	fmt.Printf("Сертификат: %s\n", cert)
	// Output:
	// Информация по сертификату:
}

func ExampleClient_SignXML() {
	opts := gokalkan.OptsTest

	cli, _ := gokalkan.NewClient(opts...)

	// Обязательно закрывайте клиент, иначе приведет к утечкам ресурсов
	defer cli.Close()

	//Тестовый RSA ключ от НУЦ РК
	keyPath := "./test/certs/gost1.p12"
	keyPassword := "Qwerty12"

	cli.LoadKeyStore(keyPath, keyPassword)

	signData, _ := cli.SignXML(&types.SignXMLInput{
		Data:    "<root>Hello</root>",
		WithTSP: false,
	})

	fmt.Printf("Подписанные данные:\n%s", signData)
	// Output:
	// Подписанные данные:
	//<?xml version="1.0" encoding="UTF-8"?>
	//<root>Hello World!<ds:Signature xmlns:ds="http://www.w3.org/2000/09/xmldsig#" Id="1">
	//..................
}

func ExampleClient_VerifyXML() {
	opts := gokalkan.OptsTest

	cli, _ := gokalkan.NewClient(opts...)

	// Обязательно закрывайте клиент, иначе приведет к утечкам ресурсов
	defer cli.Close()

	//Тестовый RSA ключ от НУЦ РК
	keyPath := "./test/certs/gost1.p12"
	keyPassword := "Qwerty12"
	data := "<root>Hello World!</root>"

	cli.LoadKeyStore(keyPath, keyPassword)

	signData, _ := cli.SignXML(&types.SignXMLInput{
		Data:    "<root>Hello</root>",
		WithTSP: false,
	})

	ver, _ := cli.VerifyXML(&types.VerifyXMLInput{
		Signature:         sigXML,
		MustCheckCertTime: false,
	})

	fmt.Println(ver)
	// Output:
	//Signature N 1
	//Id = 1
	//certificateSerialNumber=.....................................
	//signatureAlgorithm=sha256WithRSAEncryption(1.2.840.113549.1.1.11)
	//serialNumber=IIN1234567891011
	//Signature is OK
}

func ExampleClient_SignHash() {
	opts := gokalkan.OptsTest

	cli, _ := gokalkan.NewClient(opts...)

	// Обязательно закрывайте клиент, иначе приведет к утечкам ресурсов
	defer cli.Close()

	//Тестовый RSA ключ от НУЦ РК
	keyPath := "./test/certs/gost1.p12"
	keyPassword := "Qwerty12"
	data := []byte("Hello world!")

	cli.LoadKeyStore(keyPath, keyPassword)

	inHash, _ := cli.HashSHA256([]byte(data))

	signData, _ := cli.SignHash(&types.SignHashInput{
		Algo:       ckalkan.HashAlgoSHA256,
		InHash:     inHash,
		IsDetached: true,
		WithTSP:    false,
	})

	fmt.Println(signData)
	// Output:
	//[48 130 18 93 ...]
}
