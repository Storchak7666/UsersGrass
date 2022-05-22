package controllers

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/test_server/internal/domain/event"
	"net/http"
	"strconv"
)

//"strconv"
type EventController struct {
	service *event.Service
}

func NewEventController(s *event.Service) *EventController {
	return &EventController{
		service: s,
	}
}

func (c *EventController) FindAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		events, err := (*c.service).FindAll()
		if err != nil {
			fmt.Printf("EventController.FindAll(): %s", err)
			err = internalServerError(w, err)
			if err != nil {
				fmt.Printf("EventController.FindAll(): %s", err)
			}
			return
		}
		err = success(w, events)
		if err != nil {
			fmt.Printf("EventController.FindAll(): %s", err)
		}
	}
}

func (c *EventController) FindByName() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		name := chi.URLParam(r, "name")
		event, err := (*c.service).FindByName(name)
		if err != nil {
			fmt.Printf("EventController.FindOne(): %s", err)
			err = internalServerError(w, err)
			if err != nil {
				fmt.Printf("EventController.FindOne(): %s", err)
			}
			return
		}
		err = success(w, event)
		if err != nil {
			fmt.Printf("EventController.FindOne(): %s", err)
		}
	}
}

func (c *EventController) CreateUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		name := chi.URLParam(r, "name")
		age, err := strconv.ParseInt(chi.URLParam(r, "age"), 10, 64)
		if err != nil {
			fmt.Printf("EventController.CreateUser(): %s", err)
			err = internalServerError(w, err)
			if err != nil {
				fmt.Printf("EventController.CreateUser(): %s", err)
			}
			return
		}
		city := chi.URLParam(r, "city")
		country := chi.URLParam(r, "country")

		event, err := (*c.service).CreateUser(name, age, city, country)
		if err != nil {
			fmt.Printf("EventController.CreateUser(): %s", err)
			err = internalServerError(w, err)
			if err != nil {
				fmt.Printf("EventController.CreateUser(): %s", err)
			}
			return
		}

		err = success(w, event)
		if err != nil {
			fmt.Printf("EventController.CreateUser(): %s", err)
		}
	}
}

func (c *EventController) UpdateById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		name := chi.URLParam(r, "name")
		age, err := strconv.ParseInt(chi.URLParam(r, "age"), 10, 64)
		id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
		if err != nil {
			fmt.Printf("EventController.CreateUser(): %s", err)
			err = internalServerError(w, err)
			if err != nil {
				fmt.Printf("EventController.CreateUser(): %s", err)
			}
			return
		}
		city := chi.URLParam(r, "city")
		country := chi.URLParam(r, "country")

		event, err := (*c.service).UpdateById(id, name, age, city, country)
		if err != nil {
			fmt.Printf("EventController.CreateUser(): %s", err)
			err = internalServerError(w, err)
			if err != nil {
				fmt.Printf("EventController.CreateUser(): %s", err)
			}
			return
		}

		err = success(w, event)
		if err != nil {
			fmt.Printf("EventController.CreateUser(): %s", err)
		}
	}
}
