package user

import (
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/aribroo/go-ecommerce/entity"
	"github.com/aribroo/go-ecommerce/repository"
	"github.com/aribroo/go-ecommerce/utils"
	"github.com/gorilla/mux"
)

type Handler struct {
	userRepository repository.UserRepository
}

func NewHandler(userRepository repository.UserRepository) *Handler {
	return &Handler{userRepository: userRepository}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/register", h.handleRegister).Methods("POST")
	router.HandleFunc("/login", h.handleLogin).Methods("POST")
	router.HandleFunc("/users", h.findAll).Methods("GET")
	router.HandleFunc("/users/{id}", h.findOne).Methods("GET")
	router.HandleFunc("/users/{id}", h.update).Methods("PATCH")
	router.HandleFunc("/users/{id}", h.remove).Methods("DELETE")
}

func (h *Handler) handleRegister(w http.ResponseWriter, r *http.Request) {

	var payload *entity.RegisterUserPayload

	err := utils.ParseJSON(r, &payload)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	_, err = h.userRepository.FindByEmail(payload.Email)

	if err == nil {
		utils.WriteError(w, http.StatusBadRequest, errors.New("email already used"))
		return
	}

	hashedPassword, err := utils.Hash(payload.Password)

	if err != nil {
		log.Fatal(err)
		utils.WriteError(w, http.StatusInternalServerError, nil)
		return
	}

	user := entity.User{
		FirstName: payload.FirstName,
		LastName:  payload.LastName,
		Email:     payload.Email,
		Password:  hashedPassword,
	}

	err = h.userRepository.Insert(&user)

	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusCreated, "user created successfully")

}

func (h *Handler) handleLogin(w http.ResponseWriter, r *http.Request) {

	var payload *entity.RegisterUserPayload

	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	user, err := h.userRepository.FindByEmail(payload.Email)

	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, errors.New("email or password is wrong"))
		return
	}

	if ok := utils.Compare(payload.Password, user.Password); !ok {
		utils.WriteError(w, http.StatusBadRequest, errors.New("email or password is wrong"))
		return
	}

	utils.WriteJSON(w, http.StatusOK, "login user successfully")
}

func (h *Handler) findOne(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	userId, err := strconv.Atoi(vars["id"])

	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	user, err := h.userRepository.FindById(userId)

	if err != nil {
		utils.WriteError(w, http.StatusNotFound, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, user)
}

func (h *Handler) findAll(w http.ResponseWriter, r *http.Request) {
	users, err := h.userRepository.FindAll()

	if err != nil {
		utils.WriteError(w, http.StatusUnprocessableEntity, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, users)

}

func (h *Handler) update(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	userId, err := strconv.Atoi(params["id"])

	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	user, err := h.userRepository.FindById(userId)

	if err != nil {
		utils.WriteError(w, http.StatusNotFound, err)
		return
	}

	var payload *entity.UpdateUserPayload

	utils.ParseJSON(r, &payload)

	if payload.FirstName == "" {
		payload.FirstName = user.FirstName
	} else if payload.LastName == "" {
		payload.LastName = user.LastName
	}

	if err := h.userRepository.Update(userId, payload); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, "update user successfully")

}

func (h *Handler) remove(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)

	userId, err := strconv.Atoi(params["id"])
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	RowAffected, err := h.userRepository.Remove(userId)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	if RowAffected == 0 {
		utils.WriteError(w, http.StatusNotFound, errors.New("user not found"))
		return
	}

	utils.WriteJSON(w, http.StatusOK, "delete user successfully")

}
