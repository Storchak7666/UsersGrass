package event

import (
	"fmt"
	"github.com/upper/db/v4"
	"github.com/upper/db/v4/adapter/postgresql"
	"log"
)

var settings = postgresql.ConnectionURL{
	Database: `Userdb`,
	Host:     `localhost:5432`,
	User:     `postgres`,
	Password: `Olegoni678`,
}

type Repository interface {
	FindAll() ([]User, error)
	FindByName(name string) ([]User, error)
	CreateUser(name string, age int64, city string, country string) (*User, error)
	UpdateById(id int64, name string, age int64, city string, country string) (*User, error)
}

type repository struct {
	// Some internal data
}

func NewRepository() Repository {
	return &repository{}
}

func (r *repository) FindAll() ([]User, error) {
	sess, err := postgresql.Open(settings)
	if err != nil {
		log.Fatal("Open: ", err)
	}
	defer sess.Close()

	usersCol := sess.Collection("Users")

	// Find().All() maps all the records from the books collection.
	var users []User

	err = usersCol.Find().All(&users)
	if err != nil {
		log.Fatal("booksCol.Find: ", err)
	}

	// Print the queried information.
	fmt.Printf("Records in the %q collection:\n", usersCol.Name())
	for i := range users {
		fmt.Printf("record #%d: %#v\n", i, users[i])
	}

	return users, nil
}

func (r *repository) FindByName(name string) ([]User, error) {
	sess, err := postgresql.Open(settings)
	if err != nil {
		log.Fatal("Open: ", err)
	}
	defer sess.Close()

	usersCol := sess.Collection("Users")

	var users []User

	res := usersCol.Find()

	if err := res.All(&users); err != nil {
		log.Fatal("res.All: ", err)
	}
	WithName := res.And("Name LIKE", name) // WHERE ... AND title LIKE name'

	if err := WithName.All(&users); err != nil {
		log.Fatal("res.All: ", err)
	}

	return users, nil
}

func (r *repository) CreateUser(name string, age int64, city string, country string) (*User, error) {
	sess, err := postgresql.Open(settings)
	if err != nil {
		log.Fatal("Open: ", err)
	}
	db.LC().SetLevel(db.LogLevelDebug)
	defer sess.Close()
	_, err = sess.SQL().
		InsertInto("Users").
		Columns("Name", "Age", "City", "Country").Values(name, age, city, country). // Or Columns(c1, c2, c2, ...).Values(v1, v2, v2, ...).
		Exec()
	if err != nil {
		fmt.Printf("Query: %v. This is expected on the read-only sandbox.\n", err)
	}
	return &User{
		Id:      0,
		Name:    name,
		Age:     age,
		City:    city,
		Country: country,
	}, nil
}

func (r *repository) UpdateById(id int64, name string, age int64, city string, country string) (*User, error) {
	sess, err := postgresql.Open(settings)
	if err != nil {
		log.Fatal("Open: ", err)
	}
	db.LC().SetLevel(db.LogLevelDebug)

	defer sess.Close()

	if name != "" {
		_, err := sess.SQL().
			Update("Users").
			Set("Name = ?", name). // Or Set("first_name", "Edgar Allan").
			Where("id = ?", id).   // Or Where("id", eaPoe.ID)
			Exec()
		if err != nil {
			fmt.Printf("Query: %v. This is expected on the read-only sandbox.\n", err)
		}
	}

	if age >= 18 {
		_, err := sess.SQL().
			Update("Users").
			Set("Age = ?", age).
			Where("id = ?", id).
			Exec()
		if err != nil {
			fmt.Printf("Query: %v. This is expected on the read-only sandbox.\n", err)
		}
	}

	if city != "" {
		_, err := sess.SQL().
			Update("Users").
			Set("City = ?", city).
			Where("id = ?", id).
			Exec()
		if err != nil {
			fmt.Printf("Query: %v. This is expected on the read-only sandbox.\n", err)
		}
	}

	if country != "" {
		_, err := sess.SQL().
			Update("Users").
			Set("Country = ?", country).
			Where("id = ?", id).
			Exec()
		if err != nil {
			fmt.Printf("Query: %v. This is expected on the read-only sandbox.\n", err)
		}
	}

	var user User

	err = sess.SQL().
		SelectFrom("Users").
		Where("id", id). // Or Where("last_name = ?", "Poe")
		One(&user)
	if err != nil {
		log.Fatal("Query: ", err)
	}

	return &User{
		Id:      user.Id,
		Name:    user.Name,
		Age:     user.Age,
		City:    user.City,
		Country: user.Country,
	}, nil
}
