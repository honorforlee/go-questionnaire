// Package bootstrap provides identical initialization
// to main(), but also to system tests;
// initialization is gocloud-enabled so that config files
// can be loaded from an app engine bucket.
package bootstrap

import (
	"log"
	"net/http"
	"net/http/httptest"
	"time"

	"github.com/zew/go-questionnaire/cfg"
	"github.com/zew/go-questionnaire/cloudio"
	"github.com/zew/go-questionnaire/lgn"
	"github.com/zew/go-questionnaire/sessx"
	"github.com/zew/go-questionnaire/tpl"
	"github.com/zew/util"
)

// Config loads configuration and logins according to flags or env vars.
func Config() {

	fl := util.NewFlags()
	fl.Add(
		util.FlagT{
			Long:       "config_file",
			Short:      "cfg",
			DefaultVal: "config.json",
			Desc:       "JSON file containing config data",
		},
	)
	fl.Add(
		util.FlagT{
			Long:       "logins_file",
			Short:      "lgn",
			DefaultVal: "logins.json",
			Desc:       "JSON file containing logins data",
		},
	)
	fl.Gen()

	{
		cfg.CfgPath = fl.ByKey("cfg").Val

		fileName := cfg.CfgPath
		r, bucketClose, err := cloudio.Open(fileName)
		if err != nil {
			log.Fatalf("error opening writer to %v: %v", fileName, err)
		}
		defer func() {
			err := r.Close()
			if err != nil {
				log.Printf("error closing writer to bucket to %v: %v", fileName, err)
			}
		}()
		defer func() {
			err := bucketClose()
			if err != nil {
				log.Printf("error closing bucket of writer to %v: %v", fileName, err)
			}
		}()
		log.Printf("opened reader to cloud config %v", fileName)
		cfg.Load(r)

		err = cloudio.MarshalWriteFile(cfg.Example(), "config-example.json")
		if err != nil {
			log.Printf("config example save: %v", err)
		}
	}

	{
		lgn.LgnsPath = fl.ByKey("lgn").Val
		fileName := lgn.LgnsPath
		r, bucketClose, err := cloudio.Open(fileName)
		if err != nil {
			log.Fatalf("error opening writer to %v: %v", fileName, err)
		}
		defer func() {
			err := r.Close()
			if err != nil {
				log.Printf("error closing writer to bucket to %v: %v", fileName, err)
			}
		}()
		defer func() {
			err := bucketClose()
			if err != nil {
				log.Printf("error closing bucket of writer to %v: %v", fileName, err)
			}
		}()
		log.Printf("opened reader to cloud config %v", fileName)
		lgn.Load(r)

		err = cloudio.MarshalWriteFile(lgn.Example(), "logins-example.json")
		if err != nil {
			log.Printf("logins example save: %v", err)
		}

	}

	// template stuff
	{
		dummyReq, err := http.NewRequest("GET", "", nil)
		if err != nil {
			log.Fatalf("failed to create request for pre-loading assets %v", err)
		}
		respRec := httptest.NewRecorder()
		tpl.TemplatesPreparse(respRec, dummyReq)
		log.Printf("\n%v", respRec.Body.String())
	}

	if cfg.Get().SessionTimeout > 0 {
		sessx.Mgr().Lifetime = time.Duration(cfg.Get().SessionTimeout) * time.Hour // default is 24 hours
	}
	// true breaks session persistence in excel-db -
	// but not in go-count?down;
	// it also leads to session breakdown on iphone safari mobile,
	// maybe because appengine is pbberlin outside
	sessx.Mgr().Cookie.Secure = false // default is false

	// struc--2frm.S2F.Indent = 90

}
