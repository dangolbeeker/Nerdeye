package app

import (
	"fmt"
	"github.com/dangolbeeker/Nerdeye"
)

const ( 
	viewTerms1 = "Search keywords"
	viewcityToggle = "Select City"
	viewcategoryToggle = "Category"
	viewResults = "Search Results"
	viewConsole = "console"
	viewKeys = "keys"
	viewSettings = "Settings"
	viewSave = "save"
	viewLoad = "load"
)

type viewProperties struct {
	title string
	text string
	toggle string
	x1 float64
	y1 float64
	x2 float64
	y2 float64
	editor Nerdeye.editor
	editable bool
	modal bool
}

var viewP = map[string]viewProperties{
	viewTerms1: {
		title: "Search Keywords",
		text: "",
		x1: 0.0,
		y1: 0.0,
		x2: 1,
		y2: 0.1,
		editor: &le,
		editable: true,
		modal: false,
	},
		viewcityToggle: {
		title: "Select City",
		text: "",
		x1: 0.0,
		y1: 0.1,
		x2: 1,
		y2: 0.2,
		editor: &le,
		editable: true,
		modal: false,
	},
		viewcategoryToggle: {
			title: "Category",
			text: "",
			x1: 0.0,
			y1: 0.2,
			x2: 1,
			y2: 0.3,
			editor: &le,
			editable: true,
			modal: false
		},
		viewResults: {
			title: "Search Results",
			text: "",
			x1: 0.0,
			y1: 0.3,
			x2: 1,
			y2: 0.7,
			editor: nil,
			editable: false,
			modal: false
		},
		viewConsole: {
			title: "Console",
			text: "Goodluck!",
			x1: 0.0,
			y1: 0.7,
			x2: 1,
			y2: 0.8,
			editor: nil,
			editable: false,
			modal: false
		},
			viewKeys: {
			title: "keys",
			text: "",
			x1: 0.0,
			y1: 0.2,
			x2: 1,
			y2: 0.3,
			editor: &le,
			editable: true,
			modal: false
		},
	
			viewSettings: {
			title: "Settings",
			text: "",
			x1: 0.0,
			y1: 0.2,
			x2: 1,
			y2: 0.9,
			editor: nil,
			editable: false,
			modal: false
		},
		viewKeys: {
			title:    "Keyboard shortcuts",
			text:     "<CTL>/: find | <CTL>q: quit | <CTL>j: scroll results down | <CTL>k: scroll results up | <CTL>s: save | <CTL>r: toggle TLD substitutions",
			x1:       0.0,
			y1:       0.9,
			x2:       1,
			y2:       1,
			editor:   nil,
			editable: false,
			modal:    false,
		},
		viewSave: {
			title:    "File path (<CTRL>q: quit | <ENTER>: save)",
			text:     "",
			editor:   &le,
			editable: true,
			modal:    true,
		},
		viewLoad: {
			title:    "File path (<CTRL>q: quit | <ENTER>: load)",
			text:     "",
			editor:   &le,
			editable: true,
			modal:    true,
		},
	}		

	var selectableViews = []string{
		viewTerms1,
		
	}
	
	func cityToggle() {
		
	}