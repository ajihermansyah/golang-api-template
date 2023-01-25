package user

import (
	"golang-api-template/model/entity"
	"golang-api-template/repository"

	"gopkg.in/mgo.v2"
)

const (
	collectionUser string = "users"
)

type userRepository struct {
	dbSession *mgo.Session
	database  string
}

func NewUserRepository(sess *mgo.Session, database string) repository.UserRepositoryInterface {
	return &userRepository{sess, database}
}

func (repo *userRepository) CreateUser(user entity.User) error {
	var err error

	ds := repo.dbSession.Copy()
	defer ds.Close()

	table := ds.DB(repo.database).C(collectionUser)

	index := mgo.Index{
		Key: []string{"email"},
	}
	err = table.EnsureIndex(index)
	if err != nil {
		return err
	}

	err = table.Insert(user)
	if err != nil {
		return err
	}
	return err
}
