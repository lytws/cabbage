package driver

import (
	"math/rand"
	"time"

	"cbg/adapters/gwimpl"
)

type UserModel struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

var (
	userList = []UserModel{
		{"John", "fgy24f"},
		{"Zywoo", "f34gqrg"},
		{"Nancy", "f3rqt34"},
		{"Beyond", "f3ct4e"},
		{"Walker", "vearca4"},
		{"Luke", "f2c4qfawc"},
		{"Wahson", "cq34fe"},
	}
)

type MockUserDB struct {
	list []UserModel

	random *rand.Rand
}

func NewMockUserDB() *MockUserDB {
	return &MockUserDB{
		list:   userList,
		random: rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

func (d MockUserDB) randUserIndex() int {
	return d.random.Intn(len(d.list))
}

func (d MockUserDB) GetRandomUser() UserModel { return d.list[d.randUserIndex()] }

func (t MockUserDB) GetUser(username string) (*gwimpl.UserModel, error) {
	for _, u := range t.list {
		if u.Username == username {
			return &gwimpl.UserModel{Username: u.Username, Password: u.Password}, nil
		}
	}
	return nil, gwimpl.ErrUserNotFound
}
