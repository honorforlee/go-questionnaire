package lgn

import (
	"fmt"
	"html"
	"log"
	"net/http"
	"path"
	"sort"
	"strings"

	"github.com/zew/go-questionnaire/cfg"
	"github.com/zew/go-questionnaire/cloudio"
	"github.com/zew/go-questionnaire/qst"
	"github.com/zew/go-questionnaire/sessx"
)

// LoginByHash first checks for direct login;
// extra short and preconfigured in config.json;
// last part of the path - moved into h;
//
// LoginByHash then takes request values "u" and hash "h" - and not any password;
// it checks the hash against the values except "h";
// any other request parameters are sorted and included into the hashing check;
// extended by loginsT.Salt.
// reduced  by those in 'exempted';
//
// request values "h" without an "u" are considered direct login attempts,
// where the hash actually represents a user ID.
//
// On success, function creates a logged in user
// out of nothing by the name of u.
// This user gets assigned the params/values as
// role-names/values.
//
// On wrong hashes, it returns the difference as error.
// Be careful not to show the error to the end user.
func LoginByHash(w http.ResponseWriter, r *http.Request) (bool, error) {

	sess := sessx.New(w, r)

	err := r.ParseForm()
	if err != nil {
		return false, err
	}

	u := r.Form.Get("u")
	u = html.EscapeString(u) // XSS prevention
	h := r.Form.Get("h")     // hash

	if u == "" && h == "" {
		return false, nil
	}

	// First - try direct login
	if _, isSet := r.Form["u"]; !isSet { // Note: r.Form[key] contains GET *and* POST values
		if _, isSet := r.Form["h"]; isSet {
			// => userId is not set - but hash is set

			//
			surveyID := ""
			parts := strings.Split(h, "--") // h coming from anonymous id
			if len(parts) > 1 {
				surveyID = strings.ToLower(parts[0])
				h = parts[1]
			}

			userID := fmt.Sprint(HashIDDecodeFirst(h))

			if userID > "0" {
				log.Printf("Trying anonymous login - surveyID | hashID | userID - %v | %v | %v", surveyID, h, userID)
				for _, dlr := range cfg.Get().DirectLoginRanges {
					cmp := userID
					if (surveyID != "" && surveyID == dlr.SurveyID) ||
						cmp >= dlr.Start && cmp <= dlr.Stop {
						log.Printf("Matching survey %v - or direct login range %v <=  %v <=  %v",
							dlr.SurveyID, dlr.Start, cmp, dlr.Stop)
						l := LoginT{}
						l.User = userID
						l.IsInitPassword = false // roles become effective only for non init passwords
						l.Roles = map[string]string{}
						l.Attrs = map[string]string{}
						l.Attrs["survey_id"] = dlr.SurveyID
						l.Attrs["wave_id"] = dlr.WaveID
						for pk, pv := range dlr.Profile {
							l.Attrs[pk] = pv
						}
						sess.PutObject("login", l)
						log.Printf("login saved to session as %T from loginByHash", l)
						return true, nil
					}
				}
			}
		}
		return false, nil
	}

	// Second - try login from user database

	l := LoginT{}
	l.User = u
	l.IsInitPassword = false // roles become effective only for non init passwords
	l.Roles = map[string]string{}
	l.Attrs = map[string]string{}

	chkKeys := []string{}
	for key := range r.Form {
		if _, ok := exempted[key]; ok {
			continue
		}
		chkKeys = append(chkKeys, key)
	}

	sort.Strings(chkKeys)
	checkStr := ""
	for _, key := range chkKeys {
		val := r.Form.Get(key)
		checkStr += val + "-"
	}
	checkStr += lgns.Salt
	log.Printf("trying hash login over chkKeys %v-salt: '%v' ", strings.Join(chkKeys, "-"), checkStr)
	hExpected := Md5Str([]byte(checkStr))
	if hExpected != h {
		return false, fmt.Errorf("hash over check string unequal hash argument\n%v\n%v", hExpected, h)
	}

	for key := range r.Form {
		// same key - multiple values
		// attrs=country:Sweden&attrs=height:176
		if val, ok := userAttrs[key]; ok {
			if key == "attrs" {
				// for i := 0; i < len(r.Form[key]); i++ { // instead of val := r.Form.Get(key)
				// 	val := r.Form[key][i]
				// 	kv := strings.Split(val, ":")
				// 	if len(kv) == 2 {
				// 		l.Attrs[kv[0]] = kv[1]
				// 	}
			} else if key == "p" {
				profileKey := r.Form.Get("sid") + r.Form.Get(key)
				prof, ok := cfg.Get().Profiles[profileKey]
				if !ok {
					log.Printf("Did not find profile %v", profileKey)
					continue
				}
				for pk, pv := range prof {
					log.Printf("\tprofile to attr key-val  %-20v - %v", pk, pv)
					l.Attrs[pk] = pv
				}
			} else {
				l.Attrs[val] = r.Form.Get(key)
			}
		}
	}

	log.Printf("logging in as %v with attrs %v type %T", u, l.Attrs, l)
	sess.PutObject("login", l)
	log.Printf("login saved to session as %T from loginByHash", l)

	return true, nil
}

// ReloadH removes the existing questioniare from the session,
// reading it anew from the questionnaire template JSON file,
// allowing to start anew
func ReloadH(w http.ResponseWriter, r *http.Request) {

	sess := sessx.New(w, r)

	log.Printf("reset start")

	_, err := LoginByHash(w, r)
	if err != nil {
		log.Printf("Login by hash error 1: %v", err)
		// Don't show the revealing original error
		s := cfg.Get().Mp["login_by_hash_failed"].All()
		s += "LoginByHash error."
		log.Print(s)
		w.Write([]byte(s))
		return
	}
	l, isLoggedIn, err := LoggedInCheck(w, r)
	if err != nil {
		log.Printf("Login by hash error 2: %v", err)
		s := cfg.Get().Mp["login_by_hash_failed"].All()
		s += "LoggedInCheck error."
		log.Print(s)
		w.Write([]byte(s))
		return
	}
	if !isLoggedIn {
		log.Printf("Login by hash error 3: %v", "not logged in")
		s := cfg.Get().Mp["login_by_hash_failed"].All()
		s += "You are not logged in."
		log.Print(s)
		w.Write([]byte(s))
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	msg := ""
	if r.Method == "POST" {
		l.DeleteFiles()
		sess.Remove(r.Context(), "questionnaire")
		log.Printf("removed quest session")
		msg = "User files deleted. Questionnaire deleted from session."
	} else {
		msg = "Not a POST request. No delete action taken."
	}

	if sess.EffectiveStr("skip_validation") != "" {
		sess.PutString("skip_validation", "true")
	} else {
		sess.Remove(r.Context(), "skip_validation")
	}

	relForm := r.Form // relevant Form
	if len(r.PostForm) > 5 {
		relForm = r.PostForm
	}

	attrsStr := ""
	for _, val := range relForm["attrs"] {
		if val != "" {
			attrsStr += fmt.Sprintf("<input type=\"text\" name=\"attrs\" value=\"%v\" /> <br>\n", val)
		}
	}

	fmt.Fprintf(w, `
	<html>
	  <head>
		<meta charset="utf-8" />
		<title>Reset entries</title>
		<style>
		* {font-family: monospace;}
		</style>
	  </head>
      <body style='white-space:pre'>
        <b>%v</b>
        <form method="POST" class="survey-edit-form" >
            User ID          <input type="text"   name="u"                   value="%v"   /> <br>
            Survey ID        <input type="text"   name="sid"                 value="%v"   /> <br>
            Wave ID          <input type="text"   name="wid"                 value="%v"   /> <br>
            User profile #   <input type="text"   name="p"                   value="%v"   /> country name, currency etc.<br>
            Hash             <input type="text"   name="h"    size=40        value="%v"   /> <br>
            Lang code        <input type="text"   name="lang_code"  size=6   value="%v"   /> 'en', 'de' ...<br>
            Page             <input type="text"   name="page"                value="%v"   /> zero-indexed <br>
            Mobile           <input type="text"   name="mobile"              value="%v"   /> 0-auto, 1-mobile, 2-desktop <br>
            Skip validation  <input type="text"   name="skip_validation"     value="%v"   /> <br>
            %v
                             <input type="submit" name="submit" id="submit"  value="Submit" accesskey="s"  /> <br>
		</form>        
		<script> document.getElementById('submit').focus(); </script>  `,
		msg,
		l.User,
		l.Attrs["survey_id"],
		l.Attrs["wave_id"],
		relForm.Get("p"),
		relForm.Get("h"),
		relForm.Get("lang_code"),
		relForm.Get("page"),
		relForm.Get("mobile"),
		relForm.Get("skip_validation"),
		attrsStr,
	)

	queryString := Query(
		relForm.Get("u"), relForm.Get("sid"), relForm.Get("wid"), relForm.Get("p"), relForm.Get("h"),
	)
	if relForm.Get("lang_code") != "" {
		queryString += "&lang_code=" + relForm.Get("lang_code")
	}
	if relForm.Get("page") != "" {
		queryString += "&page=" + relForm.Get("page")
	}
	if relForm.Get("mobile") != "" {
		queryString += "&mobile=" + relForm.Get("mobile")
	}
	for _, attr := range relForm["attrs"] {
		if attr != "" {
			queryString += "&attrs=" + attr
		}
	}

	url := fmt.Sprintf("%v?%v", cfg.PrefTS(), queryString)

	fmt.Fprintf(w, "<a href='%v'  target='_blank'>Start questionnaire (again)<a> <br>\n", url)
	if r.Method == "POST" {
		fmt.Fprintf(w,
			`
		<SCRIPT language="JavaScript1.2">
			//var win = window.open('%s','qst','menubar=1,resizable=1,width=350,height=250,target=q');
			var win = window.open('%s', 'qst');
			win.focus();
			console.log('window opened')
		</SCRIPT>`,
			url,
			url,
		)
	}

	fmt.Fprint(w, "\t</body>\n</html>")

}

// QuestPath returns the path to the JSON questionnaire,
// Similar to qst.QuestionnaireT.FilePath1()
// See also userAttrs{}
func (l *LoginT) QuestPath() string {

	userSurveyType := ""
	userWaveID := ""
	for attr, val := range l.Attrs {
		if attr == "survey_id" {
			userSurveyType = val
		}
		if attr == "wave_id" {
			userWaveID = val
		}
	}

	if userSurveyType == "" || userWaveID == "" {
		log.Printf("Error constructing path for user questionnaire file; userSurveyType or userWaveID is empty: %v - %v", userSurveyType, userWaveID)
	}

	pth := path.Join(".", qst.BasePath(), userSurveyType, userWaveID, l.User) + ".json"
	return pth
}

// DeleteFiles deletes all JSON files
func (l *LoginT) DeleteFiles() {

	userSurveyType := ""
	userWaveID := ""
	for attr, val := range l.Attrs {
		if attr == "survey_id" {
			userSurveyType = val
		}
		if attr == "wave_id" {
			userWaveID = val
		}
	}

	if userSurveyType == "" || userWaveID == "" {
		log.Printf("Error deleting questionnaire file;  userSurveyType or userWaveID is empty: %v - %v", userSurveyType, userWaveID)
		return
	}

	// pth1 := path.Join(".", qst.BasePath(), userSurveyType, userWaveID, l.User) + "_joined.json"
	// pth2 := path.Join(".", qst.BasePath(), userSurveyType, userWaveID, l.User) + "_split.json"
	pth3 := path.Join(".", qst.BasePath(), userSurveyType, userWaveID, l.User) + ".json"
	// pth4 := path.Join(".", qst.BasePath(), userSurveyType, userWaveID, l.User) + ".json.attrs"

	pths := []string{pth3}

	for _, pth := range pths {
		err := cloudio.Delete(pth)
		if err != nil {
			if !cloudio.IsNotExist(err) {
				log.Printf("Error deleting questionnaire file: %v", err)
			}
		} else {
			log.Printf("removed quest file %v", pth)
		}
	}

}