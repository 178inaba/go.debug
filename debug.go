package debug

import (
	"net/http"
	"net/http/httputil"
	"net/url"
	"reflect"
	"runtime"
	"strings"

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

// DumpCookie is ...
func DumpCookie(client *http.Client, urlStr string) {
	log.Infof("---------- %s Start ----------", funcName(DumpCookie))

	url, err := url.Parse(urlStr)
	if err != nil {
		log.Panic(err)
	}

	log.Info(client.Jar.Cookies(url))
	log.Infof("---------- %s End ------------", funcName(DumpCookie))
}

// DumpResp is ...
func DumpResp(resp *http.Response) {
	log.Infof("---------- %s Start ----------", funcName(DumpResp))
	log.Infof("%+v", resp)
	log.Infof("---------- %s End ------------", funcName(DumpResp))
}

func funcName(i interface{}) string {
	longFuncName := runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
	splitFuncName := strings.Split(longFuncName, ".")

	return splitFuncName[len(splitFuncName)-1]
}
