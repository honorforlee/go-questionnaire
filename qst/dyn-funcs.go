package qst

import (
	"fmt"
	"time"
)

// Statistics returns the percentage of
// answers responded to
func (q *QuestionaireT) Statistics() (int, int, float64) {
	responses := 0
	inputs := 0
	for _, p := range q.Pages {
		for _, gr := range p.Groups {
			for _, i := range gr.Inputs {
				if i.IsLayout() {
					continue
				}
				if i.Type == "textarea" {
					continue
				}
				inputs++
				if i.Response != "" {
					responses++
				}
			}
		}

	}
	return responses, inputs, float64(responses / inputs)
}

// RepsonseStatistics returns the percentage of
// answers responded to
func RepsonseStatistics(q *QuestionaireT) (string, error) {

	responses, inputs, pct := q.Statistics()
	ct := q.ClosingTime
	// ct = ct.Truncate(time.Hour)
	cts := ct.Format("02.01.2006 15:04")
	nextDay := q.ClosingTime.Add(24 * time.Hour)
	nextDayS := nextDay.Format("02.01.2006")

	ret := ""
	if q.LangCode == "de" {
		s1 := fmt.Sprintf("Sie haben %v von %v beantwortet. %2.1f%%  <br>\n", responses, inputs, pct)
		s2 := fmt.Sprintf("Umfrage endet am %v. Veröffentlichung am %v  <br>\n", cts, nextDayS)
		ret = s1 + s2
	} else if q.LangCode == "en" {
		s1 := fmt.Sprintf("You answered %v out of %v questions. %2.1f%%  <br>\n", responses, inputs, pct)
		s2 := fmt.Sprintf("Survey will finish at %v. Publication will be at %v<br>\n", cts, nextDayS)
		ret = s1 + s2
	}
	return ret, nil
}

type dynFuncT func(*QuestionaireT) (string, error)

var DynFuncs = map[string]dynFuncT{
	"RepsonseStatistics": RepsonseStatistics,
}
