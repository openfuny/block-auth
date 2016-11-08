package main

import (
	. "block-auth/adminhandler"
	_ "block-auth/routers"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/astaxie/beego"
)

const (
	KeyFile    = "conf/keys/server.key"
	CertFile   = "conf/keys/server.crt"
	CaCertFile = "conf/keys/ca.crt"
)

func init() {
}

func startAdminListen() {
	adminHanlder := &ManagerHandler{}

	certPool := x509.NewCertPool()
	cartPath := CaCertFile
	caCrt, err := ioutil.ReadFile(cartPath)
	if err != nil {
		fmt.Println("ReadFile err:", err)
		return
	}
	certPool.AppendCertsFromPEM(caCrt)
	managerServer := &http.Server{
		Addr:    ":1443",
		Handler: adminHanlder,
		TLSConfig: &tls.Config{
			ClientCAs:  certPool,
			ClientAuth: tls.RequireAndVerifyClientCert,
		},
	}
	adminErr := managerServer.ListenAndServeTLS(CertFile, KeyFile)
	if adminErr != nil {
		fmt.Println("ListenAndServeTLS err:", err)
	}

}

func main() {
	beego.BConfig.Listen.EnableHTTP = false
	beego.BConfig.Listen.EnableHTTPS = true
	beego.BConfig.Listen.HTTPSPort = 9443
	beego.BConfig.Listen.HTTPSKeyFile = KeyFile
	beego.BConfig.Listen.HTTPSCertFile = CertFile

	go startAdminListen()

	beego.Run()
}
