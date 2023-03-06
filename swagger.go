/*
Swagger for Go lang is a tool to generate Swagger 2.0 documentation for Go lang applications.
*/
package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/firdavsDev/go-quest/models"
	"github.com/firdavsDev/go-quest/utils"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

// GetAllQuests Get All Quests
func GetAllQuests(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var quests []models.Quest
	models.DB.Find(&quests)

	json.NewEncoder(w).Encode(quests)
}

// GetQuest GET {ID}
func GetQuest(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id := mux.Vars(r)["id"]
	var quest models.Quest

	if err := models.DB.Where("id = ?", id).First(&quest).Error; err != nil {
		utils.RespondWithError(w, http.StatusNotFound, "Quest not found")
		return
	}

	json.NewEncoder(w).Encode(quest)
}

// Post
var validate *validator.Validate

type QuestInput struct {
	Title       string `json:"title" validate:"required"`
	Description string `json:"description" validate:"required"`
	Reward      int    `json:"reward" validate:"required"`
}

func CreateQuest(w http.ResponseWriter, r *http.Request) {
	var input QuestInput

	body, _ := ioutil.ReadAll(r.Body)
	_ = json.Unmarshal(body, &input)

	validate = validator.New()
	err := validate.Struct(input)

	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Validation Error")
		return
	}

	quest := models.Quest{Title: input.Title, Description: input.Description, Reward: input.Reward}
	models.DB.Create(&quest)

	utils.RespondWithJSON(w, http.StatusCreated, quest)
}

// Put
func UpdateQuest(w http.ResponseWriter, r *http.Request) {
	var input QuestInput

	body, _ := ioutil.ReadAll(r.Body)
	_ = json.Unmarshal(body, &input)

	validate = validator.New()
	err := validate.Struct(input)

	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Validation Error")
		return
	}

	id := mux.Vars(r)["id"]
	var quest models.Quest

	if err := models.DB.Where("id = ?", id).First(&quest).Error; err != nil {
		utils.RespondWithError(w, http.StatusNotFound, "Quest not found")
		return
	}

	models.DB.Model(&quest).Updates(input)

	utils.RespondWithJSON(w, http.StatusOK, quest)
}

// Delete
func DeleteQuest(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var quest models.Quest

	if err := models.DB.Where("id = ?", id).First(&quest).Error; err != nil {
		utils.RespondWithError(w, http.StatusNotFound, "Quest not found")
		return
	}

	models.DB.Delete(&quest)

	utils.RespondWithJSON(w, http.StatusOK, map[string]string{"message": "Quest deleted"})
}

// Path: main.go
