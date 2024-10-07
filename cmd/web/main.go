package main

import (
	"english-resources-vaughan/internal/sections"
	"fmt"
	"log"
	"net/http"
	"os"
)

const (
	// Paths
	toDoExpressionsPath      = "/to-do"
	mostImportantVerbsPath   = "/most-important-verbs"
	importantVerbsPath       = "/important-verbs"
	shouldKnowVerbsPath      = "/should-know-verbs"
	shouldKnowVerbs2Path     = "/should-know-verbs-2"
	pronounsPath             = "/pronouns"
	conjunctionsPath         = "/conjunctions"
	adjectivesAndAdverbsPath = "/adjetives-and-adverbs"
	prepositionsPath         = "/prepositions"
	verbTensesPath           = "/verb-tenses"
)

func main() {

	// Home
	dataHome := buildLinks()
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		sections.RenderWithData(w, "index.page.gohtml", dataHome)
	})

	//Expressions section
	dataToDoExpressions := sections.ReadJSON[ExpressionsPage]("common_expressions.json")
	http.HandleFunc(toDoExpressionsPath, func(w http.ResponseWriter, r *http.Request) {
		sections.RenderWithData(w, "to-do-expressions.page.gohtml", dataToDoExpressions)
	})

	// Verbs sections
	dataMostImportantVerbs := sections.ReadJSON[ListPage]("most_important_verbs.json")
	http.HandleFunc(mostImportantVerbsPath, func(w http.ResponseWriter, r *http.Request) {
		sections.RenderWithData(w, "list.page.gohtml", dataMostImportantVerbs)
	})

	dataImportantVerbs := sections.ReadJSON[ListPage]("important_verbs.json")
	http.HandleFunc(importantVerbsPath, func(w http.ResponseWriter, r *http.Request) {
		sections.RenderWithData(w, "list.page.gohtml", dataImportantVerbs)
	})

	dataVerbsShouldKnow := sections.ReadJSON[ListPage]("should_know_verbs.json")
	http.HandleFunc(shouldKnowVerbsPath, func(w http.ResponseWriter, r *http.Request) {
		sections.RenderWithData(w, "list.page.gohtml", dataVerbsShouldKnow)
	})

	dataVerbsShouldKnow2 := sections.ReadJSON[ListPage]("should_know_verbs_2.json")
	http.HandleFunc(shouldKnowVerbs2Path, func(w http.ResponseWriter, r *http.Request) {
		sections.RenderWithData(w, "list.page.gohtml", dataVerbsShouldKnow2)
	})

	// Pronouns
	dataPronouns := sections.ReadJSON[ListsWithHeader]("the_35_pronouns.json")
	http.HandleFunc(pronounsPath, func(w http.ResponseWriter, r *http.Request) {
		sections.RenderWithData(w, "list-columns.page.gohtml", dataPronouns)
	})

	// Conjunctions
	dataConjunctions := sections.ReadJSON[ListsWithHeader]("the_31_conjunctions.json")
	http.HandleFunc(conjunctionsPath, func(w http.ResponseWriter, r *http.Request) {
		sections.RenderWithData(w, "list-columns.page.gohtml", dataConjunctions)
	})

	// Adjectives and Adverbs
	dataAdjAndAdv := sections.ReadJSON[ListsWithHeader]("adjectives_and_adverbs.json")
	http.HandleFunc(adjectivesAndAdverbsPath, func(w http.ResponseWriter, r *http.Request) {
		sections.RenderWithData(w, "list-columns.page.gohtml", dataAdjAndAdv)
	})

	// Prepositions
	dataPrepositions := sections.ReadJSON[ListPage]("prepositions.json")
	http.HandleFunc(prepositionsPath, func(w http.ResponseWriter, r *http.Request) {
		sections.RenderWithData(w, "list.page.gohtml", dataPrepositions)
	})

	// Verb tenses
	type VerbTenses struct {
		Title       string `json:"title"`
		Description string `json:"description"`
		Data        []struct {
			Type   string `json:"type"`
			Values [][]struct {
				Tense   string `json:"tense"`
				Example string `json:"example"`
			} `json:"values"`
		} `json:"data"`
	}
	dataVerbTenses := sections.ReadJSON[VerbTenses]("verb_tenses.json")
	http.HandleFunc(verbTensesPath, func(w http.ResponseWriter, r *http.Request) {
		sections.RenderWithData(w, "verb-tenses.page.gohtml", dataVerbTenses)
	})

	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}
	log.Printf("Starting app at port %d", port)
	fs := http.FileServer(http.Dir("./internal/assets"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	err := http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
	if err != nil {
		log.Panic(err)
	}
}

type Link struct {
	Path        string
	Description string
}

type ExpressionsPage struct {
	Title       string   `json:"title"`
	Description []string `json:"description"`
	Data        []struct {
		PhraseSpanish string `json:"phrase_esp"`
		PhraseEnglish string `json:"phrase_eng"`
	} `json:"data"`
}

type ListPage struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Data        []struct {
		Value string `json:"value"`
	} `json:"data"`
}

type ListsWithHeader struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Data        []struct {
		Type   string   `json:"type"`
		Values []string `json:"values"`
	} `json:"data"`
}

func buildLinks() []Link {
	return []Link{
		{
			Path:        toDoExpressionsPath,
			Description: "Expresiones con 'to do'",
		},
		{
			Path:        mostImportantVerbsPath,
			Description: "Verbos más importantes",
		},
		{
			Path:        importantVerbsPath,
			Description: "Verbos importantes",
		},
		{
			Path:        shouldKnowVerbsPath,
			Description: "Verbos a conocer",
		},
		{
			Path:        shouldKnowVerbs2Path,
			Description: "Verbos a conocer II",
		},
		{
			Path:        pronounsPath,
			Description: "Los 35 Pronombres",
		},
		{
			Path:        conjunctionsPath,
			Description: "Las 31 Conjunciones",
		},
		{
			Path:        adjectivesAndAdverbsPath,
			Description: "Adjetivos y Adverbios",
		},
		{
			Path:        prepositionsPath,
			Description: "Preposiciones",
		},
		{
			Path:        verbTensesPath,
			Description: "Tiempos Verbales",
		},
	}
}
