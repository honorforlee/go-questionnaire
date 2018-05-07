package sessx

// ReqParam searches for the effective value
// of the *request*, not in session.
// First among the POST fields.
// Then among the URL "path" parameters.
// Then among the URL GET parameters.
//
// It checks, whether whether any of the above had the param
// key set to *empty* string.
func (sess *SessT) ReqParam(key string, defaultVal ...string) (string, bool) {

	p := ""

	// Which to call: r.ParseForm() or r.ParseMultipartForm(1024*1024)
	// https://blog.saush.com/2015/03/18/html-forms-and-go/
	_ = sess.r.PostFormValue("impossibleKey") // hopefully causing the right parsing

	// POST Param overrides GET param
	posts := sess.r.PostForm
	if _, ok := posts[key]; ok {
		return posts.Get(key), true
	}

	// Path Param
	// [deleted]

	// URL Get Param
	gets := sess.r.URL.Query()
	if _, ok := gets[key]; ok {
		return gets.Get(key), true // if there are multiple GET params, this returns the *first* one
	}

	return p, false

}