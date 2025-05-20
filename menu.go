package main

import (
	"log"

	"github.com/charmbracelet/huh"
	"gorm.io/gorm"
)

func start_menu(db *gorm.DB) {

	var servers []Server
	if err := db.Find(&servers).Error; err != nil {
		log.Fatal("Query failed:", err)
	}

	var options []huh.Option[Server]
	for _, server := range servers {
		options = append(options, huh.NewOption(server.AdminDomain, server))
	}

	var selected []Server
	menu := huh.NewMultiSelect[Server]().
		Title("Select Servers:").
		Options(options...).
		Value(&selected)

	menu.Run()
}
