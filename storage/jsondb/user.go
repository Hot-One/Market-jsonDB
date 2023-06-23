package jsondb

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"os"

	"github.com/google/uuid"

	"market/models"
)

type UserRepo struct {
	fileName string
	file     *os.File
}

func NewUserRepo(fileName string, file *os.File) *UserRepo {
	return &UserRepo{
		fileName: fileName,
		file:     file,
	}
}

func (u *UserRepo) Create(req *models.UserCreate) (*models.User, error) {
	// Read File \\

	users, err := u.read()
	if err != nil {
		return nil, err
	}

	// Create Model of User \\
	var (
		id   = uuid.New().String()
		user = models.User{
			Id:      id,
			Name:    req.Name,
			Surname: req.Surname,
			Balance: req.Balance,
		}
	)

	users[id] = user

	// Write \\

	err = u.write(users)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (u *UserRepo) GetById(req *models.UserPrimaryKey) (*models.User, error) {
	// Read File \\

	users, err := u.read()

	if err != nil {
		return nil, err
	}

	// Check User Exist \\

	if _, have := users[req.Id]; !have {
		return nil, errors.New("User not found")
	}

	// Get By Id \\

	user := users[req.Id]

	return &user, nil
}

func (u *UserRepo) GetList(req *models.UserGetListRequest) (*models.UserGetListResponse, error) {

	var resp = &models.UserGetListResponse{}
	resp.Users = []*models.User{}

	// Read File \\

	userMap, err := u.read()
	if err != nil {
		return nil, err
	}

	// Fill the resp  \\

	resp.Count = len(userMap)
	for _, val := range userMap {
		users := val
		resp.Users = append(resp.Users, &users)
	}

	return resp, nil
}

func (u *UserRepo) Update(req *models.UserUpdate) (*models.User, error) {
	// Read File \\

	users, err := u.read()

	if err != nil {
		return nil, err
	}

	// Check User Exist \\

	if _, ok := users[req.Id]; !ok {
		return nil, errors.New("User not found")
	}

	// Update User \\

	users[req.Id] = models.User{
		Id:      req.Id,
		Name:    req.Name,
		Surname: req.Surname,
		Balance: req.Balance,
	}

	// Write Update User \\

	err = u.write(users)
	if err != nil {
		return nil, err
	}
	user := users[req.Id]

	return &user, nil
}

func (u *UserRepo) Delete(req *models.UserPrimaryKey) error {
	// Read File \\

	users, err := u.read()

	if err != nil {
		return err
	}

	// Delete User \\

	delete(users, req.Id)

	// Write \\

	err = u.write(users)
	if err != nil {
		return err
	}

	return nil
}

func (u *UserRepo) read() (map[string]models.User, error) {
	var (
		users   []*models.User
		userMap = make(map[string]models.User)
	)

	data, err := ioutil.ReadFile(u.fileName)
	if err != nil {
		log.Printf("Error while Read data: %+v", err)
		return nil, err
	}

	err = json.Unmarshal(data, &users)
	if err != nil {
		log.Printf("Error while Unmarshal data: %+v", err)
		return nil, err
	}

	for _, user := range users {
		userMap[user.Id] = *user
	}

	return userMap, nil
}

func (u *UserRepo) write(userMap map[string]models.User) error {
	var users []models.User

	for _, value := range userMap {
		users = append(users, value)
	}

	body, err := json.MarshalIndent(users, "", "	")
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(u.fileName, body, os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}
