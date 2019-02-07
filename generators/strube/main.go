package strube

import (
	"github.com/zew/go-questionnaire/qst"
	"github.com/zew/go-questionnaire/trl"
)

var xx = trl.S{
	"de": "...",
	"en": "...",
	"fr": "...",
	"it": "...",
}

func labelsGoodBad19() []trl.S {

	tm := []trl.S{
		{
			"de": "lehne ab<span class='ordinals'><br>-4</span>",
			"en": "Disagree<span class='ordinals'><br>-4</span>",
			"fr": "Pas d’accord<span class='ordinals'><br>-4</span>",
			"it": "Non favorevole<span class='ordinals'><br>-4</span>",
		},
		{
			"de": "<span class='ordinals'><br>-3</span>",
			"en": "<span class='ordinals'><br>-3</span>",
			"fr": "<span class='ordinals'><br>-3</span>",
			"it": "<span class='ordinals'><br>-3</span>",
		},
		{
			"de": "<span class='ordinals'><br>-2</span>",
			"en": "<span class='ordinals'><br>-2</span>",
			"fr": "<span class='ordinals'><br>-2</span>",
			"it": "<span class='ordinals'><br>-2</span>",
		},
		{
			"de": "<span class='ordinals'><br>-1</span>",
			"en": "<span class='ordinals'><br>-1</span>",
			"fr": "<span class='ordinals'><br>-1</span>",
			"it": "<span class='ordinals'><br>-1</span>",
		},
		{
			"de": "unentschieden<span class='ordinals'><br>0</span>",
			"en": "Undecided<span class='ordinals'><br>0</span>",
			"fr": "Indifférent<span class='ordinals'><br>0</span>",
			"it": "Indeciso<span class='ordinals'><br>0</span>",
		},
		{
			"de": "<span class='ordinals'><br>1</span>",
			"en": "<span class='ordinals'><br>1</span>",
			"fr": "<span class='ordinals'><br>1</span>",
			"it": "<span class='ordinals'><br>1</span>",
		},
		{
			"de": "<span class='ordinals'><br>2</span>",
			"en": "<span class='ordinals'><br>2</span>",
			"fr": "<span class='ordinals'><br>2</span>",
			"it": "<span class='ordinals'><br>2</span>",
		},
		{
			"de": "<span class='ordinals'><br>3</span>",
			"en": "<span class='ordinals'><br>3</span>",
			"fr": "<span class='ordinals'><br>3</span>",
			"it": "<span class='ordinals'><br>3</span>",
		},
		{
			"de": "stimme zu<span class='ordinals'><br>4</span>",
			"en": "Agree<span class='ordinals'><br>4</span>",
			"fr": "D’accord<span class='ordinals'><br>4</span>",
			"it": "Favorevole<span class='ordinals'><br>4</span>",
		},
	}

	return tm

}

// Create creates a questionnaire with a few pages and inputs.
func Create(params []qst.ParamT) (*qst.QuestionnaireT, error) {
	q := qst.QuestionnaireT{}
	q.Survey = qst.NewSurvey("strube")
	q.Survey.Params = params
	q.Variations = 0 // attention => shuffles submit buttons if > 0

	q.LangCodes = map[string]string{
		"de": "Deutsch",
		"en": "English",
		"fr": "Français",
		"it": "Italiano",
	}
	q.LangCodesOrder = []string{
		"en",
		"fr",
		"de",
		"it",
	}
	q.LangCode = "en" // No default; forces usage of UserLangCode()

	q.Survey.Org = trl.S{
		"de": "ZEW",
		"en": "ZEW",
		"fr": "ZEW",
		"it": "ZEW",
	}
	q.Survey.Name = trl.S{
		"de": "Umfrage: Gestaltung der Europäischen Union",
		"en": "Survey: Design of the European Union",
		"fr": "Questionnaire: Design de l’Union Européenne",
		"it": "Questionario: Design dell'Unione Europea",
	}

	i1 := "[attr-country]"
	i2 := "[attr-has-euro]"
	_, _ = i1, i2

	//
	// Page 1
	{

		p := q.AddPage()
		// p.NoNavigation = true
		p.Width = 70
		p.Section = trl.S{
			"de": "Allgemeine Einstellungen zum Euro und zur Wirtschaftspolitik",
			"en": "General attitudes on euro and economic policy",
			"fr": "Attitudes générales vis-à-vis de l'euro et de la politique économique",
			"it": "Posizioni generali sull'Euro e sulla politica economica europea.",
		}
		p.Label = trl.S{
			"de": "",
			"en": "",
			"fr": "",
			"it": "",
		}
		p.Desc = trl.S{
			"de": "Inwieweit stimmen Sie den folgenden Aufgaben zu?",
			"en": "Do you agree with the following statements?",
			"fr": "Approuvez-vous les propositions suivantes ?",
			"it": "In che misura si trova d’accordo con le seguenti affermazioni?",
		}
		p.Short = trl.S{
			"de": "Allgemeine Einstellungen",
			"en": "General attitudes",
			"fr": "Attitudes générales",
			"it": "Posizioni generali",
		}

		// 11
		{
			// Dynamic header for following headless radio matrix
			gr := p.AddGroup()
			gr.BottomVSpacers = 0
			inp := gr.AddInput()
			inp.Type = "dynamic"
			inp.DynamicFunc = "HasEuroQuestion"
		}
		{
			// Headless radio matrix
			names1stMatrix := []string{"benefit"}
			emptyRowLabels := []trl.S{}
			gr2 := p.AddRadioMatrixGroup(labelsGoodBad19(), names1stMatrix, emptyRowLabels, 1)
			gr2.Cols = 9 // necessary, otherwise no vspacers
			gr2.OddRowsColoring = true

		}

		// 12
		{
			names1stMatrix := []string{"supply"}
			emptyRowLabels := []trl.S{}
			gr := p.AddRadioMatrixGroup(labelsGoodBad19(), names1stMatrix, emptyRowLabels, 1)
			gr.Cols = 9 // necessary, otherwise no vspacers
			gr.OddRowsColoring = true
			gr.Label = trl.S{
				"de": "Nachfrage- vs. angebotsbezogene Politikmaßnahmen<br>",
				"en": "Demand-side versus supply-side policies<br>",
				"fr": "Politiques axées sur la demande et politiques axées sur l’offre<br>",
				"it": "Politiche della domanda e dell'offerta<br>",
			}
			gr.Desc = trl.S{
				"de": "Regierungen können versuchen, das Wirtschaftswachstum durch verschiedene Instrumente zu stimulieren. Einige argumentieren, dass nachfragebezogene Politikmaßnahmen (z. B. eine Erhöhung der durch Schulden finanzierten öffentlichen Ausgaben) wirksamer sind als angebotsbezogene Maßnahmen (z. B. eine Verringerung der Regulierung der Arbeits- und Gütermärkte).",
				"en": "Governments can try to stimulate economic growth through different instruments. Some argue that demand-side policies (e.g. an increase of debt-financed public spending) are more effective than supply-side policies (e.g. a reduction of regulation in labour and good markets).",
				"fr": "Les gouvernements peuvent essayer de stimuler la croissance économique grâce à différents instruments. Certains font valoir que les politiques axées sur la demande (par exemple, une augmentation des dépenses publiques financées par l’endettement) sont plus efficaces que les politiques axées sur l'offre (par exemple, une réduction de la réglementation sur le marché du travail et des biens).",
				"it": "I governi possono cercare di stimolare la crescita economica attraverso diversi strumenti. Alcuni sostengono che le politiche di sostegno della domanda (ad esempio un aumento della spesa pubblica finanziato in deficit) siano più efficaci delle politiche di sostegno dell’offerta (ad esempio un alleggerimento della regolamentazione del mercato del lavoro e dei beni).",
			}
		}

	}

	//
	// Page 2
	{

		p := q.AddPage()
		// p.NoNavigation = true
		p.Width = 70
		p.Section = trl.S{
			"de": "Zur Aufgaben- und Kompetenzverteilung in Europa",
			"en": "EU competencies",
			"fr": "Répartition des missions et des compétences en Europe",
			"it": "Competenze dell'Unione europea (UE). È d'accordo con le seguenti affermazioni?",
		}
		p.Label = trl.S{
			"de": "",
			"en": "",
			"fr": "",
			"it": "",
		}
		p.Desc = trl.S{
			"de": "",
			"en": "",
			"fr": "",
			"it": "",
		}
		p.Short = trl.S{
			"de": "Aufgaben- und Kompetenzverteilung",
			"en": "EU competencies",
			"fr": "Répartition des compétences",
			"it": "Competenze UE",
		}

		// 21
		{
			names1stMatrix := []string{"tax"}
			emptyRowLabels := []trl.S{}
			gr := p.AddRadioMatrixGroup(labelsGoodBad19(), names1stMatrix, emptyRowLabels, 1)
			gr.Cols = 9 // necessary, otherwise no vspacers
			gr.OddRowsColoring = true
			gr.Label = trl.S{
				"de": "Steuerpolitik<br>",
				"en": "Tax policy<br>",
				"fr": "Politique fiscale<br>",
				"it": "Tassazione<br>",
			}
			gr.Desc = trl.S{
				"de": "Der Rat der EU sollte mit qualifizierter Mehrheit anstelle von Einstimmigkeit über Steuern beschließen können (z.B. über verbindliche Ober- oder Untergrenzen für Unternehmenssteuern).",
				"en": "The European Council should be able to vote on tax issues with a qualified majority instead of una-nimity (e.g. common caps or floors for corporate taxes binding for member states).",
				"fr": "Le Conseil européen devrait pouvoir statuer avec une majorité qualifiée sur les impôts perçus dans les États membres (par exemple sur des taux planchers et plafonds de l’impôt sur les Sociétés).",
				"it": "Il Consiglio europeo dovrebbe poter votare su questioni tributarie a maggioranza qualificata invece che all’unanimità (ad esempio su limiti massimi e minimi per le imposte sulle aziende comuni a tutti gli Stati membri).",
			}
		}

		// 22
		{
			names1stMatrix := []string{"redist"}
			emptyRowLabels := []trl.S{}
			gr := p.AddRadioMatrixGroup(labelsGoodBad19(), names1stMatrix, emptyRowLabels, 1)
			gr.Cols = 9 // necessary, otherwise no vspacers
			gr.OddRowsColoring = true
			gr.Label = trl.S{
				"de": "Umverteilung<br>",
				"en": "Redistribution<br>",
				"fr": "Redistribution des revenus<br>",
				"it": "Ridistribuzione<br>",
			}
			gr.Desc = trl.S{
				"de": "Es sollte mehr Umverteilung von reichen zu armen EU-Mitgliedstaaten geben.",
				"en": "There should be more redistribution from richer to poorer EU member states.",
				"fr": "Il devrait y avoir davantage de redistribution des États membres de l'UE les plus riches vers les plus pauvres.",
				"it": "Ci dovrebbe essere maggiore ridistribuzione di risorse dagli Stati membri più ricchi a quelli più po-veri dell'UE.",
			}
		}

		// 23
		{
			names1stMatrix := []string{"immig"}
			emptyRowLabels := []trl.S{}
			gr := p.AddRadioMatrixGroup(labelsGoodBad19(), names1stMatrix, emptyRowLabels, 1)
			gr.Cols = 9 // necessary, otherwise no vspacers
			gr.OddRowsColoring = true
			gr.Label = trl.S{
				"de": "Einwanderungspolitik<br>",
				"en": "Immigration policy<br>",
				"fr": "Politique d’immigration<br>",
				"it": "Immigrazione<br>",
			}
			gr.Desc = trl.S{
				"de": "Die EU sollte eine stärkere Rolle in der Einwanderungspolitik erhalten (z.B. Aufnahmestandards festlegen oder über die Verteilung von Flüchtlingen entscheiden).",
				"en": "The EU should get a stronger role in immigration policy (e.g. decisions over admission standards or allocation of refugees across member states).",
				"fr": "L’UE devrait jouer un rôle renforcé dans la politique d’immigration des États membres (par exemple en fixant les normes d’accueil ou en décidant de la répartition des réfugiés).",
				"it": "L’UE dovrebbe avere un ruolo più incisivo sull’immigrazione (ad esempio sulle decisioni relative ai criteri di ammissione o sulla distribuzione dei rifugiati tra i Paesi membri).",
			}
		}

		// 24
		{
			names1stMatrix := []string{"defense"}
			emptyRowLabels := []trl.S{}
			gr := p.AddRadioMatrixGroup(labelsGoodBad19(), names1stMatrix, emptyRowLabels, 1)
			gr.Cols = 9 // necessary, otherwise no vspacers
			gr.OddRowsColoring = true
			gr.Label = trl.S{
				"de": "Verteidigungspolitik<br>",
				"en": "Defence policy<br>",
				"fr": "Politique de défense<br>",
				"it": "Difesa<br>",
			}
			gr.Desc = trl.S{
				"de": "Eine unter dem Befehl der EU stehende und aus ihrem Haushalt finanzierte europäische Armee sollte Aufgaben der nationalen Streitkräfte für internationale Kriseneinsätze übernehmen.",
				"en": "A European army under the command of the EU and financed from its budget should take over duties from national armies regarding international conflict deployments.",
				"fr": "L’UE pourrait constituer une armée européenne placée sous son commandement et financée par son budget avec pour mission d’assurer les opérations extérieures à la place des armées nationales.",
				"it": "Un esercito Europeo sotto il comando dell’UE, finanziato dal budget europeo, dovrebbe subentrare alle forze armate nazionali nei conflitti internazionali.",
			}
		}

		// {
		// 	gr := p.AddGroup()
		// 	gr.Cols = 1
		// 	gr.Width = 99
		// 	{
		// 		inp := gr.AddInput()
		// 		inp.Type = "button"
		// 		inp.Name = "submitBtn"
		// 		inp.Response = "2"
		// 		inp.Label = cfg.Get().Mp["next"]
		// 		inp.AccessKey = "n"
		// 		inp.ColSpanControl = 1
		// 		inp.HAlignControl = qst.HRight
		// 	}
		// }

	}

	//
	// Page 3
	{

		p := q.AddPage()
		// p.NoNavigation = true
		p.Width = 70
		p.Section = trl.S{
			"de": "Europäische Währungsunion (EWU)",
			"en": "European Monetary Union (EMU)",
			"fr": "L'Union économique et monétaire de l'Union européenne (UEM)",
			"it": "Unione monetaria europea (UME)",
		}
		p.Label = trl.S{
			"de": "",
			"en": "",
			"fr": "",
			"it": "",
		}
		p.Desc = trl.S{
			"de": "",
			"en": "",
			"fr": "",
			"it": "",
		}
		p.Short = trl.S{
			"de": "Währungsunion",
			"en": "Monetary Union",
			"fr": "L'Union monétaire",
			"it": "Unione monetaria",
		}

		// 31
		{
			names1stMatrix := []string{"insure"}
			emptyRowLabels := []trl.S{}
			gr := p.AddRadioMatrixGroup(labelsGoodBad19(), names1stMatrix, emptyRowLabels, 1)
			gr.Cols = 9 // necessary, otherwise no vspacers
			gr.OddRowsColoring = true
			gr.Label = trl.S{
				"de": "Europäische Arbeitslosenversicherung<br>",
				"en": "European unemployment insurance<br>",
				"fr": "Assurance chômage européenne<br>",
				"it": "Indennità europea di disoccupazione<br>",
			}
			gr.Desc = trl.S{
				"de": "Die EWU braucht fiskalische Stabilisierungsmechanismen, um Mitgliedstaaten gegen asymmetri-sche Schocks abzusichern (z.B. eine gemeinsame europäische Arbeitslosenversicherung).",
				"en": "The EMU needs fiscal stabilization systems to insure member states against asymmetric shocks (e.g. a common European unemployment insurance).",
				"fr": "L’UEM a besoin de systèmes de stabilisation budgétaire pour assurer les États membres contre les chocs asymétriques (par exemple, une assurance chômage européenne commune)",
				"it": "La UME necessita di sistemi di stabilizzazione fiscale atti a proteggere gli Stati membri da shock asimmetrici (ad esempio un’indennità di disoccupazione europea comune).",
			}
		}

		// 32
		{
			names1stMatrix := []string{"eurobonds"}
			emptyRowLabels := []trl.S{}
			gr := p.AddRadioMatrixGroup(labelsGoodBad19(), names1stMatrix, emptyRowLabels, 1)
			gr.Cols = 9 // necessary, otherwise no vspacers
			gr.OddRowsColoring = true
			gr.Label = trl.S{
				"de": "Eurobonds<br>",
				"en": "Eurobonds<br>",
				"fr": "Euro-obligations<br>",
				"it": "Eurobond<br>",
			}
			gr.Desc = trl.S{
				"de": "Für Eurobonds haften alle Euro-Staaten gemeinsam und alle Euro-Staaten zahlen den gleichen Zins. Die EWU sollte Eurobonds ausgeben.",
				"en": "All euro countries are jointly liable for Eurobonds and all euro countries pay the same interest. The EMU should issue Eurobonds.",
				"fr": "La zone euro devrait émettre des euro-obligations et les États membres s’en porter tous garants solidairement et bénéficier du même taux d’intérêt.",
				"it": "Gli Eurobond sono titoli pubblici di debito di cui tutti i Paesi Euro sono collettivamente responsabili e su cui tutti i Paesi Euro pagano gli stessi interessi. La UME dovrebbe iniziare ad emettere Euro-bond.",
			}
		}

		// 33
		{
			names1stMatrix := []string{"stability"}
			emptyRowLabels := []trl.S{}
			gr := p.AddRadioMatrixGroup(labelsGoodBad19(), names1stMatrix, emptyRowLabels, 1)
			gr.Cols = 9 // necessary, otherwise no vspacers
			gr.OddRowsColoring = true
			gr.Label = trl.S{
				"de": "Stabilitäts- und Wachstumspakt (SWP)<br>",
				"en": "Stability and Growth Pact (SGP)<br>",
				"fr": "Pacte de Stabilité et de Croissance (PSC)<br>",
				"it": "Patto di Stabilità e Crescita (PSC)<br>",
			}
			gr.Desc = trl.S{
				"de": "Der SWP definiert Defizit- und Schuldengrenzen für EU-Mitgliedsstaaten. Der SWP schränkt die Fiskalpolitik der Mitgliedsstaaten unangemessen stark ein und sollte gelockert werden.",
				"en": "The SGP defines deficit and debt limits for EU member states. The SGP inappropriately constrains fiscal policy in member states and should be relaxed.",
				"fr": "Le PSC définit des limites aux déficits et aux dettes des États membres. Le PSC représente une con-trainte excessive sur les politiques fiscales des États membres et devrait être assoupli.",
				"it": "Il PSC definisce i limiti per il deficit e il debito pubblico dei Paesi membri dell’UE. Il PSC limita ecces-sivamente la politica fiscale degli Stati membri e dovrebbe essere allentato.",
			}
		}

		// 34
		{
			names1stMatrix := []string{"bankruptcy"}
			emptyRowLabels := []trl.S{}
			gr := p.AddRadioMatrixGroup(labelsGoodBad19(), names1stMatrix, emptyRowLabels, 1)
			gr.Cols = 9 // necessary, otherwise no vspacers
			gr.OddRowsColoring = true
			gr.Label = trl.S{
				"de": "Insolvenzverfahren für Eurostaaten<br>",
				"en": "Insolvency procedure for euro member states<br>",
				"fr": "Procédure d'insolvabilité pour les États membres de l'euro<br>",
				"it": "Procedura di insolvenza per gli Stati dell'Eurozona<br>",
			}
			gr.Desc = trl.S{
				"de": "Es sollte ein klares Insolvenzverfahren für Eurostaaten mit unhaltbaren Schulden geben.",
				"en": "There should be an explicit sovereign insolvency procedure for euro member states with unsus-tainable debt. ",
				"fr": "Il devrait exister une procédure d'insolvabilité souveraine explicite pour les États membres de la zone euro ayant une dette insoutenable. ",
				"it": "Dovrebbe esistere una esplicita procedura di insolvenza per i Paesi Euro con debito insostenibile.",
			}
		}

		// 35
		{
			names1stMatrix := []string{"purchase"}
			emptyRowLabels := []trl.S{}
			gr := p.AddRadioMatrixGroup(labelsGoodBad19(), names1stMatrix, emptyRowLabels, 1)
			gr.Cols = 9 // necessary, otherwise no vspacers
			gr.OddRowsColoring = true
			gr.Label = trl.S{
				"de": "Wertpapierkaufprogramm der EZB<br>",
				"en": "Asset purchase programme of ECB<br>",
				"fr": "Achats d’emprunts par la BCE<br>",
				"it": "Programma di acquisti di attività finanziarie da parte della Banca centrale europea (BCE) <br>",
			}
			gr.Desc = trl.S{
				"de": "Die Europäische Zentralbank (EZB) hat in den zurückliegenden Jahren durch den Kauf von Staatsan-leihen von Euro-Staaten eine sehr aktive Rolle gespielt. Diese starke Rolle der EZB sollte fortgesetzt werden.",
				"en": "The European Central Bank (ECB) has taken a strongly active position in recent years by purchasing sovereign bonds of euro countries. This strongly active position of the ECB should continue.",
				"fr": "Pour combattre la crise, la Banque centrale européenne s’est engagée fortement dans les années passées en achetant des emprunts d’États de la zone euro. Cet engagement devrait se poursuivre.",
				"it": "Negli ultimi anni la BCE ha attuato una politica monetaria molto espansiva comprando titoli di Stato dei Paesi Euro. Questa politica della BCE dovrebbe continuare in futuro.",
			}
		}

		// 36
		{
			names1stMatrix := []string{"bankunion"}
			emptyRowLabels := []trl.S{}
			gr := p.AddRadioMatrixGroup(labelsGoodBad19(), names1stMatrix, emptyRowLabels, 1)
			gr.Cols = 9 // necessary, otherwise no vspacers
			gr.OddRowsColoring = true
			gr.Label = trl.S{
				"de": "Vollendung der Bankenunion<br>",
				"en": "Completion of Banking Union<br>",
				"fr": "Union bancaire<br>",
				"it": "Completamento dell’Unione bancaria europea<br>",
			}
			gr.Desc = trl.S{
				"de": "Für ein angemessenes Funktionieren sollte die europäische Bankenunion durch die Europäische Einlagensicherung (European Deposit Insurance System: EDIS) vollendet werden.",
				"en": "For its proper functioning, the European Banking Union should be completed through the Europe-an Deposit Insurance Scheme (EDIS). ",
				"fr": "Pour son bon fonctionnement, l’Union bancaire européenne devrait être complétée par le Système Européen de Garanties des Dépôts (SEGD).",
				"it": "Per funzionare correttamente, l’Unione bancaria europea dovrebbe essere completata tramite l’introduzione di un sistema europeo di garanzia dei depositi.",
			}
		}

	}

	//
	// Page 4
	{

		p := q.AddPage()
		// p.NoNavigation = true
		p.Width = 70
		p.Section = trl.S{
			"de": "Persönliche Fragen",
			"en": "Personal questions",
			"fr": "Questions personnelles",
			"it": "Domande personali",
		}
		p.Label = trl.S{
			"de": "",
			"en": "",
			"fr": "",
			"it": "",
		}
		p.Desc = trl.S{
			"de": "",
			"en": "",
			"fr": "",
			"it": "",
		}
		p.Short = trl.S{
			"de": "Über Sie",
			"en": "About you",
			"fr": "Q. personnelles",
			"it": "D. personali",
		}

		// 41
		{
			gr := p.AddGroup()
			gr.Cols = 3
			gr.OddRowsColoring = true

			{
				inp := gr.AddInput()
				inp.Name = "birth"
				inp.Type = "text"
				inp.ColSpanControl = 2
				inp.ColSpanLabel = 1
				inp.MaxChars = 4
				inp.Label = trl.S{
					"de": "In welchem Jahr wurden Sie geboren?",
					"en": "In which year were you born?",
					"fr": "Quelle est votre année de naissance ?",
					"it": "In quale anno è nato?",
				}

			}

			{
				inp := gr.AddInput()
				inp.Name = "nationality"
				inp.Type = "dropdown"

				inp.DD = &qst.DropdownT{}
				inp.DD.Add("", "Please choose")
				inp.DD.Add("ger", "Germany")
				inp.DD.Add("uk", "UK")

				inp.ColSpanControl = 2
				inp.ColSpanLabel = 1
				inp.Label = trl.S{
					"de": "Welche Nationalität haben Sie?",
					"en": "What is your nationality?",
					"fr": "Quelle est votre nationalité ?",
					"it": "Qual è la sua nazionalità?",
				}

			}

		}

	}

	//
	// Page Finish
	{
		p := q.AddPage()
		p.NoNavigation = true
		p.Label = trl.S{
			"de": "Vielen Dank",
			"en": "Thank you",
			"fr": "Nous vous remercions d'avoir répondu à nos questions.",
			"it": "Grazie per aver risposto al nostro questionario.",
		}

		{
			// Only one group => shuffling is no problem
			gr := p.AddGroup()
			gr.Cols = 1

			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.CSSLabel = "special-input-vert-wider"
				inp.Desc = trl.S{
					"de": "Danke für Ihre Teilnahme an unserer Umfrage.",
					"en": "Thank you for your participation in our survey.",
					"fr": "Nous vous remercions d'avoir répondu à nos questions.",
					"it": "Grazie per aver risposto al nostro questionario.",
				}
			}

			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.CSSLabel = "special-input-vert-wider"
				inp.Desc = trl.S{
					"de": "<span style='font-size: 100%;'>Ihre Eingaben wurden gespeichert.</span>",
					"en": "<span style='font-size: 100%;'>Your entries have been saved.</span>",
					"fr": "<span style='font-size: 100%;'>Vos réponses ont été sauvegardées.</span>",
					"it": "<span style='font-size: 100%;'>Le Sue risposte sono state salvate.</span>",
				}
			}

		}

	}

	(&q).Hyphenize()
	(&q).ComputeMaxGroups()
	if err := (&q).TranslationCompleteness(); err != nil {
		return &q, err
	}
	return &q, nil
}