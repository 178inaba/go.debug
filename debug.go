package debug

import (
	"net/http"
	"net/http/httputil"
	"net/url"
	"reflect"
	"runtime"

	log "github.com/Sirupsen/logrus"
)

// DumpReqAll is ...
func DumpReqAll(req *http.Request) {
	log.Infof("---------- %s Start ----------", funcName(DumpReqAll))

	dump, err := httputil.DumpRequestOut(req, true)
	if err != nil {
		log.Panic(err)
	}

	log.Infof("%s", dump)
	log.Infof("---------- %s End ------------", funcName(DumpReqAll))
}

// DumpCookies is ...
func DumpCookies(client *http.Client, urlStr string) {
	log.Infof("---------- %s Start ----------", funcName(DumpCookies))

	url, err := url.Parse(urlStr)
	if err != nil {
		log.Panic(err)
	}

	log.Info(client.Jar.Cookies(url))
	log.Infof("---------- %s End ------------", funcName(DumpCookies))
}

// DumpResp is ...
func DumpResp(resp *http.Response) {
	log.Infof("---------- %s Start ----------", funcName(DumpResp))
	log.Infof("%+v", resp)
	log.Infof("---------- %s End ------------", funcName(DumpResp))
}

func funcName(i interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
}
