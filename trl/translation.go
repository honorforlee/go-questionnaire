// Package trl is a central store for application environment translations;
// While specific objects such as QuestionaireT contain translations for its contents,
// we need some global store too.
package trl

import "strings"

// LangCodes for returning multiple translations.
// When no langCode is available, then the first entry rules.
// A call to All() returns explicitly all key-values.
// LangCodes will be initialized in cfg.Load().LangCodes; we prevent circular dependency
var LangCodes = []string{"de", "en"}

// S stores a multi lingual string.
// Contains one value for each language code.
type S map[string]string

const noTrans = "multi lingual string not initialized."

// Tr translates by key.
// Defaults           to english.
// Defaults otherwise to first lang code.
// Returns a 'signifiers string' if no translation exists,
// to help to uncover missing translations.
func (s S) Tr(langCode string) string {
	if val, ok := s[langCode]; ok {
		return val
	}
	if val, ok := s[LangCodes[0]]; ok {
		return val
	}
	if s == nil {
		return noTrans
	}
	return noTrans
}

// TrSilent gives no warning - if the translation is not set.
// Good if we do not require a translation string.
// Good for i.e. HTML title attribute - where errors are easy to overlook.
func (s S) TrSilent(langCode string) string {
	ret := s.Tr(langCode)
	if ret == noTrans {
		return ""
	}
	return ret
}

// String is the default "stringer" implementation
// Similar to Tr() - except third try: take whatever exists
func (s S) String() string {
	if val, ok := s["en"]; ok {
		return val
	}
	if val, ok := s[LangCodes[0]]; ok {
		return val
	}
	for _, val := range s {
		return val
	}
	return ""
}

// All returns all translations
// ordered by lang codes
func (s S) All() string {
	ret := ""
	for _, key := range LangCodes {
		if val, ok := s[key]; ok {
			ret += val
			ret += "\n\n"
		}
	}
	return ret
}

// Set checks whether s is not empty
func (s S) Set() bool {
	if len(s) < 1 { // also covers s == nil
		return false
	}
	return true
}

// Map - Translations Type
// Usage in templates
// 		{{.Trls.imprint.en                     }}  // directly accessin a specific translation; chaining the map keys
// 		{{.Trls.imprint.Tr       .Sess.LangCode}}  // using .Tr(langCode)
// 		{{.Trls.imprint.TrSilent .Sess.LangCode}}  //
type Map map[string]S

//
//
// Hyphenization
// =================
//
// hyphm is a map with words and their hyphenized form as value.
// hyphm is filled during app initialization from hyph, below.
var hyphm = map[string]string{}

//
// hyph is a slice _with_ hyphenized words.
// During application init - we use it to fill hyphm
var hyph = []string{

	"An&shy;gabe",

	"Ausfall&shy;risiken",
	"Risiko&shy;trag&shy;fähig&shy;keit",
	"Re&shy;finanz&shy;ierung",

	"Wett&shy;be&shy;werbs&shy;sit&shy;uation",
	"Regu&shy;lierung",
	"Groß&shy;unter&shy;nehmen",
	"Immob&shy;ilien&shy;kredite",
	"Kon&shy;sum&shy;enten&shy;kredite",

	"Regierungs&shy;bildung Deutschland",
	"Kredit&shy;nachfrage",
	"Kredit&shy;angebot",

	"mittel&shy;fristig",
}

func init() {
	for _, v := range hyph {
		hyphm[strings.Replace(v, "&shy;", "", -1)] = v
	}
}

func hyphenizeUnused(s string) string {
	if _, ok := hyphm[s]; ok {
		return hyphm[s]
	}
	return s
}

// HyphenizeText replaces "mittelfristig" with "mittel&shy;fristig"
// We do it _once_ during creation of the questionare JSON template.
func HyphenizeText(s string) string {
	for k, v := range hyphm {
		s = strings.Replace(s, k, v, -1)
	}
	return s
}
