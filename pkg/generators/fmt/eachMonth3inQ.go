package fmt

import (
	"fmt"
	"log"

	"github.com/zew/go-questionnaire/pkg/css"
	"github.com/zew/go-questionnaire/pkg/qst"
	"github.com/zew/go-questionnaire/pkg/trl"
)

func eachMonth3inQ(q *qst.QuestionnaireT) error {

	include := false
	include = include || q.Survey.Year == 2021 && q.Survey.Month == 9
	include = include || q.Survey.Year == 2021 && q.Survey.Month == 12
	include = include || q.Survey.Year == 2022 && q.Survey.Month == 3
	include = include || q.Survey.Year == 2022 && q.Survey.Month == 6

	// !!!
	// update the 3 month reference "compared to [now minus three months]"
	// !!!

	if !include {
		return nil
	}

	// q1
	var rowLabelsAssetClassesEuroZoneQ3 = []trl.S{
		{
			"de": "Aktien",
			"en": "Stocks",
		},
		{
			"de": "Staats&shy;anleihen",
			"en": "Govt. bonds",
		},
		{
			"de": "Unter&shy;nehmens&shy;anleihen",
			"en": "Corporate bonds",
		},
		{
			"de": "Immobilien",
			"en": "Real estate",
		},
	}

	var inputNamesAssetClassesEuroZoneQ3 = []string{
		"ass_euro_stocks",
		"ass_euro_bonds_govt",
		"ass_euro_bonds_corp",
		"ass_euro_re",
	}

	// q2
	var lblsQ2 = []trl.S{
		{
			"de": "Aktien",
			"en": "Stocks",
		},
		{
			"de": "Staats&shy;anleihen",
			"en": "Govt. bonds",
		},
		{
			"de": "Unter&shy;nehmens&shy;anleihen",
			"en": "Corporate bonds",
		},
		{
			"de": "Immobilien",
			"en": "Real estate",
		},
		{
			"de": "Gold",
			"en": "Gold",
		},
		{
			"de": "Rohstoffe",
			"en": "Commodities",
		},
		{
			"de": "Krypto&shy;währungen",
			"en": "Crypto currencies",
		},
	}

	var namesQ2 = []string{
		"ass_global_stocks",
		"ass_global_bonds_govt",
		"ass_global_bonds_corp",
		"ass_global_re",
		"ass_global_gold",
		"ass_global_raw_materials",
		"ass_global_crypto",
	}

	// q3
	var namesQ3Assets = []string{
		"chg_euro_stocks",
		"chg_euro_bonds_govt",
		"chg_euro_bonds_corp",
		"chg_euro_re",
	}
	var namesQ3Influence = []string{
		"economy",    // overall economic outlook
		"ecb",        // monetary policy ecb
		"fed",        // monetary policy fed
		"inflation",  // outlook inflation
		"politics",   // political framework
		"valuation",  // market valuation
		"warukraine", //
		"covid19",    //
	}

	var labelsQ3Assets = []trl.S{
		{
			"de": "Aktien",
			"en": "Stocks",
		},
		{
			"de": "Staats&shy;anleihen",
			"en": "Sovereign bonds",
		},
		{
			"de": "Unter&shy;nehmens&shy;anleihen",
			"en": "Corporate bonds",
		},
		{
			"de": "Immobilien",
			"en": "Real estate",
		},
	}

	var labelsQ3Influences = []trl.S{
		{
			"de": "Gesamtwirtschaftlicher Ausblick",
			"en": "Economic outlook",
		},
		{
			"de": "Geldpolitik der EZB",
			"en": "ECB monetary policy",
		},
		{
			"de": "Geldpolitik der US-Notenbank",
			"en": "US Federal Reserve monetary policy",
		},
		{
			"de": "Ausblick Inflation",
			"en": "Inflation outlook ",
		},
		{
			"de": "Politische Rahmen&shy;bedingungen",
			"en": "Political framework",
		},
		{
			"de": "Aktuelle Markt&shy;bewertung",
			// "en": "Market valuation",
			"en": "Current valuation multiples",
		},
		{
			"de": "Krieg Russ&shy;land - Ukraine",
			"en": "Russia's war with Ukraine",
		},
		{
			"de": "Covid-19 Pandemie",
			"en": "Covid-19 pandemic",
		},
	}

	//
	//
	//
	//
	//
	{
		page := q.AddPage()
		page.Label = trl.S{
			"de": "Sonderfrage: Anlageklassen im Eurogebiet und weltweit",
			// "en": "Special: Asset Classes in the Euro Area",
			"en": "Additional questions on the attractiveness of different asset classes",
		}
		page.Short = trl.S{
			"de": "Sonderfrage:<br>Anlageklassen",
			"en": "Special:<br>Asset classes",
		}
		page.WidthMax("46rem")

		//
		// gr1
		{
			var columnTemplateLocal = []float32{
				3.0, 1,
				0.0, 1,
				0.0, 1,
				0.0, 1,
				0.5, 1,
			}
			gb := qst.NewGridBuilderRadios(
				columnTemplateLocal,
				positiveNegative5(),
				inputNamesAssetClassesEuroZoneQ3,
				radioVals6,
				rowLabelsAssetClassesEuroZoneQ3,
			)

			gb.MainLabel = trl.S{
				"de": `
				<p style=''>
					<b>1.</b> &nbsp;

					Mit Blick auf die nächsten sechs Monate, 
					wie beurteilen Sie das Rendite-Risko-Profil 
					der folgenden Anlageklassen?
					
					Orientieren Sie sich an breit gestreuten Indizes 
					für das <b>Eurogebiet</b>.
				</p>

				<p style=''>
					Das Rendite-Risiko-Profil beurteile ich …
				</p>
				`,
				"en": `
				<p style=''>
					<b>1.</b> &nbsp;
					How do you assess the return-risk profile of the following asset classes 
					in the <b><i>euro area</i></b> for the next 6 months? 

					Please consider well-diversified indices.
					
					
				</p>
				<p style=''>
					My assessment of the return-risk profile is …
				</p>
				`,
			}

			gr := page.AddGrid(gb)
			gr.OddRowsColoring = true
		}

		//
		// gr2
		{
			var columnTemplateLocal = []float32{
				3.0, 1,
				0.0, 1,
				0.0, 1,
				0.0, 1,
				0.5, 1,
			}
			gb := qst.NewGridBuilderRadios(
				columnTemplateLocal,
				positiveNegative5(),
				namesQ2,
				radioVals6,
				lblsQ2,
			)

			gb.MainLabel = trl.S{
				"de": `
				<p style=''>
					<b>2.</b> &nbsp;
					Mit Blick auf die nächsten sechs Monate, 
					wie beurteilen Sie das Rendite-Risko-Profil 
					der folgenden Anlageklassen?

					Orientieren Sie sich an breit gestreuten <b>globalen</b> Indizes.
				</p>

				<p style=''>
					Das Rendite-Risiko-Profil beurteile ich …
				</p>
				`,
				"en": `
				<p style=''>
					<b>2.</b> &nbsp;
					How do you assess the return-risk profile  
					of the following <b><i>global</i></b> asset classes for the next 6 months? 

					Please consider well-diversified indices.
					
					
				</p>
				<p style=''>
					My assessment of the return-risk profile is …
				</p>
				`,
			}

			gr := page.AddGrid(gb)
			gr.OddRowsColoring = true
		}

		//
		// gr3
		{
			gr := page.AddGroup()
			gr.Cols = 1
			gr.BottomVSpacers = 1
			gr.Style = css.NewStylesResponsive(gr.Style)
			gr.Style.Desktop.StyleBox.WidthMax = "30rem"
			gr.Style.Mobile.StyleBox.WidthMax = "100%"
			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.Label = trl.S{
					"de": `
				<p style=''>
					<b>3.</b>  &nbsp;
					Haben Entwicklungen der folgenden Faktoren 
					Sie zu einer Revision Ihrer Einschätzungen 
					zum Rendite-Risiko-Profil der einzelnen Assetklassen
					im <b>Eurogebiet</b> 
					gegenüber März 2022 bewogen?
				</p>

				<p style=''>
					und wenn ja, nach oben (+) oder unten (-)
				</p>

					`,

					"en": `
				<p style=''>
					<b>3.</b>  &nbsp;

					Did developments in the following areas
					lead you to change your assessment 
					of the return-risk profiles
					of the following four asset classes 
					(compared to March 2022)
					in the <b>euro area</b>?

				</p>

				<p style=''>
					(+) = upward change, (-) = downward change
				</p>
					

					`,
				}
				inp.ColSpanLabel = 1
			}

		}

		//
		//
		//
		// gr4 ... gr11
		// questions 3.1 ... 3.4
		var columnTemplateLocal = []float32{
			3.6, 1,
			0.0, 1,
			0.0, 1,
			0.0, 1,
			0.0, 1,
			0.4, 1,
		}
		// additional row below each block
		colsBelow := append([]float32{1.0}, columnTemplateLocal...)

		// the first column - with a width of 4.6  (3.6+1)
		// 		is separated into two cols:
		// 			1.4, 2.2  and 0, 1
		// 		adding up to
		// 			3.6       and 1
		//
		// the default GapColumn = "0.4rem" skewed this; so we previously used
		//    		1.38, 2.1
		colsBelow = []float32{
			1.4, 2.2,
			0.0, 1,
			0.0, 1,
			0.0, 1,
			0.0, 1,
			0.0, 1,
			0.4, 1,
		}
		colsBelowPairs := []float32{} // adding up colspan label and colspan control
		for i := 0; i < len(colsBelow); i += 2 {
			colsBelowPairs = append(colsBelowPairs, colsBelow[i]+colsBelow[i+1])
		}

		for idx, assCl := range namesQ3Assets {

			names := []string{}
			for _, nm := range namesQ3Influence {
				names = append(names, assCl+"__"+nm)
			}

			lbl := labelsQ3Assets[idx]

			{
				gb := qst.NewGridBuilderRadios(
					columnTemplateLocal,
					improvedDeterioratedPlusMinus6(),
					names,
					radioVals6,
					labelsQ3Influences,
				)

				gb.MainLabel = trl.S{
					"de": fmt.Sprintf(`
					<p style='position: relative; top: 0.8rem'>
						<span>3.%v.</span> &nbsp;
						%v
						 &nbsp; - &nbsp;  Eurogebiet 
					</p>
					`,
						idx+1,
						lbl.Tr("de"),
					),
					"en": fmt.Sprintf(`
					<p style='position: relative; top: 0.8rem'>
						<span>3.%v.</span> &nbsp;

						%v
						 &nbsp; - &nbsp;  euro area 
					</p>
					`,
						idx+1,
						lbl.Tr("en"),
					),
				}

				gr := page.AddGrid(gb)
				gr.BottomVSpacers = 1
			}

			{

				gr := page.AddGroup()
				gr.Cols = 7
				gr.BottomVSpacers = 2
				if idx == 3 {
					gr.BottomVSpacers = 4
				}
				stl := ""
				for colIdx := 0; colIdx < len(colsBelowPairs); colIdx++ {
					stl = fmt.Sprintf(
						"%v   %vfr ",
						stl,
						colsBelowPairs[colIdx],
					)
				}
				gr.Style = css.NewStylesResponsive(gr.Style)
				if gr.Style.Desktop.StyleGridContainer.TemplateColumns == "" {
					gr.Style.Desktop.StyleBox.Display = "grid"
					gr.Style.Desktop.StyleGridContainer.TemplateColumns = stl
					// log.Printf("fmt special 2021-09: grid template - %v", stl)
				} else {
					log.Printf("GridBuilder.AddGrid() - another TemplateColumns already present.\nwnt%v\ngot%v", stl, gr.Style.Desktop.StyleGridContainer.TemplateColumns)
				}

				{
					inp := gr.AddInput()
					inp.Type = "text"
					inp.Name = assCl + "__other_label"
					inp.MaxChars = 17
					inp.ColSpan = 1
					inp.ColSpanLabel = 2.4
					inp.ColSpanControl = 4
					// inp.Placeholder = trl.S{"de": "Andere: Welche?", "en": "Other: Which?"}
					inp.Label = trl.S{
						"de": "Andere",
						"en": "Other",
					}
				}

				//
				for idx := 0; idx < len(improvedDeterioratedPlusMinus6()); idx++ {
					rad := gr.AddInput()
					rad.Type = "radio"

					rad.Name = assCl + "__other"
					rad.ValueRadio = fmt.Sprint(idx + 1)

					rad.ColSpan = 1
					rad.ColSpanLabel = colsBelow[2*(idx+1)]
					rad.ColSpanControl = colsBelow[2*(idx+1)] + 1

				}

			}

		}

	}

	return nil

}
