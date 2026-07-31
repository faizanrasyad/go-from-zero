package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

func Home (w http.ResponseWriter, r*http.Request) {
	fmt.Fprintln(w, "Welcome to BrewHub Coffee API!")
}

func GetMenuByID (w http.ResponseWriter, r*http.Request) {
	w.Header().Set("Content-Type", "application/json")
	
	id, _ := strconv.Atoi(strings.TrimPrefix(r.URL.Path, "/menu/"))

	var menu MenuItem
	for i := 0; i < len(Menus); i++ {
		if (Menus[i].ID == id) {
			menu = Menus[i]
			i = len(Menus)
		}
	}
	
	if (menu.ID != 0) {
		json.NewEncoder(w).Encode(menu)
	} else {
		http.Error(w, "Menu not found", http.StatusNotFound)
	}
	
}

func GetMenu (w http.ResponseWriter, r*http.Request) {
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(Menus)
	
}


