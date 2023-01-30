package user

import (
	"fmt"
	"golang-api-template/model/entity"
	"golang-api-template/repository"

	"gopkg.in/mgo.v2/bson"

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

func (repo *userRepository) FindAllUser(limit int, page int, filterText string, keyword string) ([]entity.User, int, error) {
	var (
		err           error
		users         = make([]entity.User, 0)
		offset        = (page - 1) * limit
		totalRecord   int
		sortBy        = bson.M{"$sort": bson.M{"created_at": -1}}
		filterKeyword = bson.M{}
		pipeline      = []bson.M{}
	)

	// filtering
	if filterText != "" {
		if keyword == "email" {
			filterKeyword = bson.M{"email": bson.M{"$regex": filterText, "$options": "i"}}
		} else if keyword == "name" {
			filterKeyword = bson.M{"name": bson.M{"$regex": filterText, "$options": "i"}}
		} else if keyword == "username" {
			filterKeyword = bson.M{"username": bson.M{"$regex": filterText, "$options": "i"}}
		}
	}

	pipeline = []bson.M{
		{"$match": bson.M{
			"$and": []bson.M{
				filterKeyword,
			},
		}},
		sortBy,
		{"$skip": offset},
		{"$limit": limit},
	}

	fmt.Println(pipeline)

	pipelineCount := bson.M{
		"$and": []bson.M{
			filterKeyword,
		},
	}

	ds := repo.dbSession.Copy()
	defer ds.Close()
	table := ds.DB(repo.database).C(collectionUser)

	//indexing
	index := mgo.Index{Key: []string{"created_at"}}
	_ = table.EnsureIndex(index)

	index = mgo.Index{Key: []string{"email"}}
	_ = table.EnsureIndex(index)

	index = mgo.Index{Key: []string{"name"}}
	_ = table.EnsureIndex(index)

	index = mgo.Index{Key: []string{"username"}}
	_ = table.EnsureIndex(index)

	pipe := table.Pipe(pipeline)
	if err = pipe.All(&users); err != nil {
		return nil, 0, err
	}

	totalRecord, err = table.Find(pipelineCount).Count()

	fmt.Println(&users, totalRecord)

	return users, totalRecord, nil
}

func (repo *userRepository) UpdateUser(input entity.User) error {
	var err error

	ds := repo.dbSession.Copy()
	defer ds.Close()
	table := ds.DB(repo.database).C(collectionUser)

	err = table.Update(
		bson.M{"_id": bson.ObjectIdHex(input.ID.Hex())},
		bson.M{"$set": bson.M{
			"name":       input.Name,
			"email":      input.Email,
			"username":   input.Username,
			"age":        input.Age,
			"password":   input.Password,
			"gender":     input.Gender,
			"updated_at": input.UpdatedAt,
		}},
	)
	return err
}

func (repo *userRepository) FindUserById(userId string) (entity.User, error) {
	ds := repo.dbSession.Copy()
	defer ds.Close()
	table := ds.DB(repo.database).C(collectionUser)

	objResult := entity.User{}
	err := table.Find(bson.M{"_id": bson.ObjectIdHex(userId)}).One(&objResult)

	return objResult, err
}
