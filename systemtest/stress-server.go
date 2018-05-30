// Package systemtest contains system tests;
// these are however run from the app dir one above.
// Working dir will be initially /go-questionaire/systemtest,
// but we will step up one dir in the code below.
package systemtest

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/zew/go-questionaire/bootstrap"
	"github.com/zew/go-questionaire/cfg"
	"github.com/zew/go-questionaire/ctr"
	"github.com/zew/go-questionaire/generators"
	"github.com/zew/go-questionaire/handlers"
	"github.com/zew/go-questionaire/lgn"
	"github.com/zew/go-questionaire/muxwrap"
	"github.com/zew/go-questionaire/qst"
	"github.com/zew/go-questionaire/sessx"
	"github.com/zew/util"
)

// We need this file and this empty func to avoid
// "no buildable Go source files" on travis
func main() {

}

// StartServer starts a server almost like main().
// It is currently unused.
// For coverage, tests must be run from app root.
// and they must call main().
// See app root main_test for more details.
func StartServer(t *testing.T, doChDirUp bool) {

	// For database files, static files and templates relative paths to work,
	// as if running from main app dir:
	if doChDirUp {
		os.Chdir("..")
	}

	wd, _ := os.Getwd()
	t.Logf("test directory one up: %v ; should be app main dir", wd)

	rand.Seed(time.Now().UTC().UnixNano())
	log.SetFlags(log.Lshortfile | log.Ldate | log.Ltime)

	os.Setenv("GO_TEST_MODE", "true")
	defer os.Setenv("GO_TEST_MODE", "false")
	bootstrap.Config()

	//
	generators.FMT()

	//
	// Start the server
	{
		mux1 := http.NewServeMux()
		mux1.HandleFunc(cfg.Pref("/login-primitive"), lgn.LoginPrimitiveH)
		mux1.HandleFunc(cfg.Pref("/"), handlers.MainH)
		mux1.HandleFunc(cfg.PrefWTS("/"), handlers.MainH)
		mux2 := muxwrap.NewHandlerMiddleware(mux1)
		sessx.Mgr().Lifetime(2 * time.Hour) // default is 24 hours
		sessx.Mgr().Persist(false)
		mux3 := sessx.Mgr().Use(mux2)

		IPPort := fmt.Sprintf("%v:%v", cfg.Get().BindHost, cfg.Get().BindSocket)
		t.Logf("starting http server at %v ...", IPPort)

		chSuccess := make(chan error)
		bootFunc := func(ch chan error) {
			err := http.ListenAndServe(IPPort, mux3)
			ch <- err
		}

		go bootFunc(chSuccess)

		select {
		case errBoot := <-chSuccess:
			if errBoot != nil {
				t.Fatalf("\nCould not start test server. \nLive system running? \nError %v", errBoot)
				return
			}
		case <-time.After(1100 * time.Millisecond):
			t.Logf("Test server came up without error")
		}
		// time.Sleep(1100 * time.Millisecond) // wait for application to come up
	}

}

// SimulateLoad logs in as 'systemtest'
// and performs some requests.
func SimulateLoad(t *testing.T) {

	port := cfg.Get().BindSocket

	host := fmt.Sprintf("http://localhost:%v", port)

	urlLogin := host + cfg.Pref("/login-primitive")
	t.Logf("url import %v", urlLogin)

	urlMain := host + cfg.Pref()
	t.Logf("url main   %v", urlMain)

	//
	// Login and save session cookie
	var sessCook *http.Cookie
	{
		t.Logf(" ")
		t.Logf("Getting cookie")
		t.Logf("==================")
		urlReq := urlLogin

		vals := url.Values{}
		vals.Set("username", "systemtest")
		vals.Set("password", "systemtest")
		req, err := http.NewRequest("POST", urlReq, bytes.NewBufferString(vals.Encode())) // <-- URL-encoded payload
		if err != nil {
			t.Errorf("error creating request for %v: %v", urlReq, err)
		}
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		client := util.HttpClient()
		resp, err := client.Do(req)
		if err != nil {
			t.Errorf("error requesting cookie from %v: %v; %v", urlReq, err, resp)
		}
		defer resp.Body.Close()
		for _, v := range resp.Cookies() {
			if v.Name == "session" {
				sessCook = v
			}
		}

		t.Logf("Cookie is %+v \ngleaned from %v", sessCook, req.URL)
		if sessCook == nil {
			t.Fatal("we need a session cookie to continue")
		}

		respBytes, _ := ioutil.ReadAll(resp.Body)
		if strings.Contains(string(respBytes), "logged in as systemtest") {
			t.Fatalf("Response must contain 'logged in as systemtest' \n\n%v", string(respBytes))
		}

	}

	ctr.Reset()
	//
	//
	// Post values and check the response
	{
		t.Logf(" ")
		t.Logf("Main view")
		t.Logf("==================")
		urlReq := urlMain

		waveID := qst.NewWaveID().String()

		vals := url.Values{}
		vals.Set("wave_id", waveID)
		vals.Set("y0_ez", ctr.IncrementStr()) // Don't forget to reset; otherwise depending on generate.FMT() the result is not deterministic
		vals.Set("y0_deu", ctr.IncrementStr())
		vals.Set("token", lgn.FormToken())
		t.Logf("POST requesting %v?%v", urlReq, vals.Encode())
		resp, err := util.Request("POST", urlReq, vals, []*http.Cookie{sessCook})
		if err != nil {
			t.Errorf("error requesting %v: %v", urlReq, err)
		}

		respStr := string(resp)

		// character distance -
		// must be large - since `name='y0_deu'` is found for val='1' ... val='2'
		// but only the second/third radio has value='2' checked="checked"
		scope := 340
		{
			needle1 := `name='y0_ez'`
			needle2 := `value='1' checked="checked"`
			pos1 := strings.Index(respStr, needle1)
			pos2 := strings.Index(respStr, needle2)
			t.Logf("Response should contain: %v ... %v \n%v %v => %v",
				needle1, needle2, pos1, pos2, pos2-pos1,
			)
			if pos1 < 1 || pos2 < 1 || (pos2-pos1) > scope {
				t.Fatal("Failed")
			}
		}
		{
			// needle := `name='y0_deu' id='y0_deu' title=' Deutschland' value='2' checked="checked"`
			needle1 := `name='y0_deu'`
			needle2 := `value='2' checked="checked"`
			pos1 := strings.Index(respStr, needle1)
			pos2 := strings.Index(respStr, needle2)
			t.Logf("Response should contain: %v ... %v \n%v %v => %v",
				needle1, needle2, pos1, pos2, pos2-pos1,
			)
			if pos1 < 1 || pos2 < 1 || (pos2-pos1) > scope {
				t.Fatal("Failed")
			}
		}

	}

}