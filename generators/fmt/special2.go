package fmt

import (
	"github.com/zew/go-questionaire/qst"
	"github.com/zew/go-questionaire/trl"
)

func addSeasonal2(q *qst.QuestionaireT) {

	if monthOfQuarter() != 2 && false {
		return
	}

	p := q.AddPage()
	p.Section = trl.S{"de": "Sonderfrage", "en": "Special"}
	p.Label = trl.S{"de": "Prognosetreiber Inflation und Geldpolitik", "en": "Inflation and monetary policy drivers"}
	p.Short = trl.S{"de": "Inflation, Geldpolitik", "en": "Inflation, monetary<br>policy"}
	p.Width = 90

	{
		gr := p.AddGroup()
		gr.Cols = 9
		gr.Label = trl.S{
			"de": "1.",
			"en": "1.",
		}
		gr.Desc = trl.S{
			"de": "Punktprognose der jährlichen Inflationsrate im Euroraum",
			"en": "Forecast yearly inflation rate in the Euro area",
		}

		{
			inp := gr.AddInput()
			inp.Type = "textblock"
			inp.Desc = trl.S{
				"de": "Anstieg des HICP von Jan bis Dez; Erwartungswert",
				"en": "HICP  increase from Jan to Dec; expected value",
			}
			inp.ColSpanLabel = 3
		}

		{
			inp := gr.AddInput()
			inp.Type = "text"
			inp.Name = "p1_y1"
			inp.MaxChars = 4
			inp.Validator = "inRange20"
			inp.Desc = trl.S{
				"de": nextY(0),
				"en": nextY(0),
			}
			inp.Suffix = trl.S{
				"de": "%",
				"en": "pct",
			}
			inp.HAlignLabel = qst.HRight

		}
		{
			inp := gr.AddInput()
			inp.Type = "text"
			inp.Name = "p1_y2"
			inp.MaxChars = 4
			inp.Validator = "inRange20"
			inp.Desc = trl.S{
				"de": nextY(1),
				"en": nextY(1),
			}
			inp.Suffix = trl.S{
				"de": "%",
				"en": "pct",
			}
			inp.HAlignLabel = qst.HRight
		}
		{
			inp := gr.AddInput()
			inp.Type = "text"
			inp.Name = "p1_y3"
			inp.MaxChars = 4
			inp.Validator = "inRange20"
			inp.Desc = trl.S{
				"de": nextY(2),
				"en": nextY(2),
			}
			inp.Suffix = trl.S{
				"de": "%",
				"en": "pct",
			}
			inp.HAlignLabel = qst.HRight

		}

	}

	// gr1
	{
		labels123Matrix := []trl.S{
			{
				"de": "Konjunkturdaten Euroraum",
				"en": "Business cycle data Euro area",
			},
			{
				"de": "Konjunkturdaten global",
				"en": "Business cycle data globally",
			},
			{
				"de": "Löhne Euroraum",
				"en": "Wages Euro area",
			},
			{
				"de": "Rohstoffpreise",
				"en": "Raw material prices",
			},
			{
				"de": "Wechselkurse",
				"en": "Exchange rates",
			},
			{
				"de": "EZB-Geldpolitik",
				"en": "ECB monetary policy",
			},
		}
		// ioi => impact on inflation
		names1stMatrix := []string{
			"ioi_cycle_data_ea",
			"ioi_cycle_data_gl",
			"ioi_wages_ea",
			"ioi_rmp",
			"ioi_exch_rates",
			"ioi_mp_ecb",
		}
		gr := p.AddRadioMatrixGroup(labelsStronglyPositiveStronglyNegativeInfluence(),
			names1stMatrix, labels123Matrix, 2)
		gr.Cols = 8 // necessary, otherwise no vspacers
		gr.OddRowsColoring = true
		gr.Label = trl.S{
			"de": "2.",
			"en": "2.",
		}
		gr.Desc = trl.S{
			"de": "Haben Entwicklungen in den folgenden Bereichen Sie zu einer Revision Ihrer Inflationsprognosen (ggü. Vormonat) für den Euroraum bewogen und wenn ja in welche Richtung?",
			"en": "Which developments have lead you to change your assessment of the inflation outlook for the Euro are compared to the previous month",
		}
	}

	{
		gr := p.AddGroup()
		gr.Label = trl.S{"de": "3.", "en": "3."}
		gr.Desc = trl.S{
			"de": "Den Hauptrefinanzierungssatz der EZB (am 01.02.2018: 0,0%) erwarte ich auf Sicht von <br>\n &nbsp;",
			"en": "I expect the main refinance rate of the ECB (01.02.2018: 0,0%) in <br>\n &nbsp;",
		}
		gr.Cols = 100

		{
			inp := gr.AddInput()
			inp.Type = "textblock"
			inp.ColSpanLabel = 8
			inp.Desc = trl.S{
				"de": "6 Monaten",
				"en": "6 months",
			}
			inp.HAlignLabel = qst.HRight
		}

		{

			inp := gr.AddInput()
			inp.Type = "text"
			inp.Name = "i_ez_06_low"
			inp.MaxChars = 4
			inp.Validator = "inRange20"

			inp.ColSpanLabel = 7
			inp.CSSLabel = "vert-wider"
			inp.ColSpanControl = 9
			inp.Desc = trl.S{
				"de": "zwischen&nbsp;",
				"en": "between&nbsp;",
			}
			inp.Suffix = trl.S{"de": "%", "en": "pct"}
			inp.HAlignLabel = qst.HRight
			inp.HAlignControl = qst.HLeft
		}

		{
			inp := gr.AddInput()
			inp.Type = "text"
			inp.Name = "i_ez_06_high"
			inp.MaxChars = 4
			inp.Validator = "inRange20"

			inp.ColSpanLabel = 3
			inp.ColSpanControl = 73
			inp.Desc = trl.S{
				"de": "und",
				"en": "and",
			}
			inp.Suffix = trl.S{"de": "%", "en": "pct"}
			inp.HAlignLabel = qst.HLeft
			inp.HAlignControl = qst.HLeft
		}

		{
			inp := gr.AddInput()
			inp.Type = "textblock"
			inp.ColSpanLabel = 8
			inp.Desc = trl.S{
				"de": " 24 Monaten",
				"en": " 24 months",
			}
			inp.HAlignLabel = qst.HRight
		}

		// Second row
		{
			inp := gr.AddInput()
			inp.Type = "text"
			inp.Name = "i_ez_24_low"
			inp.MaxChars = 4
			inp.Validator = "inRange20"

			inp.ColSpanLabel = 7
			inp.CSSLabel = "vert-wider"
			inp.ColSpanControl = 9
			inp.Desc = trl.S{
				"de": "zwischen&nbsp;",
				"en": "between&nbsp;",
			}
			inp.Suffix = trl.S{"de": "%", "en": "pct"}
			inp.HAlignLabel = qst.HRight
			inp.HAlignControl = qst.HLeft
		}

		{
			inp := gr.AddInput()
			inp.Type = "text"
			inp.Name = "i_ez_24_high"
			inp.MaxChars = 4
			inp.Validator = "inRange20"

			inp.ColSpanLabel = 3
			inp.ColSpanControl = 73
			inp.Desc = trl.S{
				"de": "und",
				"en": "and",
			}
			inp.Suffix = trl.S{"de": "%", "en": "pct"}
			inp.HAlignLabel = qst.HLeft
			inp.HAlignControl = qst.HLeft
		}

		{
			inp := gr.AddInput()
			inp.Type = "textblock"
			inp.ColSpanLabel = 15
			inp.Desc = trl.S{
				"de": " &nbsp;",
				"en": " &nbsp;",
			}
		}
		{
			inp := gr.AddInput()
			inp.Type = "textblock"
			inp.ColSpanLabel = 85
			inp.Desc = trl.S{
				"de": " [zentrales 90% Konfidenzintervall]",
				"en": " [central 90pct confidence interval]",
			}
			inp.CSSLabel = "textblock-smaller"
		}

	}

}
