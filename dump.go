package dump

import (
	"net/http"
	"net/http/httputil"
	"net/url"
	"reflect"
	"runtime"
	"strings"

	log "github.com/Sirupsen/logrus"
)

// Dump is ...
func Dump(v interface{}) {
	log.Infof("---------- %s Start ----------", funcName(Dump))
	log.Infof("%+v", v)
	log.Infof("---------- %s End ------------", funcName(Dump))
}

// ReqAll is ...
func ReqAll(req *http.Request) {
	log.Infof("---------- %s Start ----------", funcName(ReqAll))

	dump, err := httputil.DumpRequestOut(req, true)
	if err != nil {
		log.Panic(err)
	}

	log.Infof("%s", dump)
	log.Infof("---------- %s End ------------", funcName(ReqAll))
}

// RespAll is ...
func RespAll(resp *http.Response) {
	log.Infof("---------- %s Start ----------", funcName(RespAll))

	dump, err := httputil.DumpResponse(resp, true)
	if err != nil {
		log.Panic(err)
	}

	log.Infof("%s", dump)
	log.Infof("---------- %s End ------------", funcName(RespAll))
}

// Cookie is ...
func Cookie(client *http.Client, rawurl string) {
	log.Infof("---------- %s Start ----------", funcName(Cookie))

	url, err := url.Parse(rawurl)
	if err != nil {
		log.Panic(err)
	}

	log.Info(client.Jar.Cookies(url))
	log.Infof("---------- %s End ------------", funcName(Cookie))
}

func funcName(i interface{}) string {
	longFuncName := runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
	splitFuncName := strings.Split(longFuncName, ".")

	return splitFuncName[len(splitFuncName)-1]
}
