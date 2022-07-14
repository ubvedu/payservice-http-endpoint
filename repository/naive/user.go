package naive

import "github.com/google/uuid"

type (
    User struct {
        Email    string
        Password string
    }

    UserRepo struct {
        users map[uuid.UUID]User
    }
)

func NewUserRepo() *UserRepo {
    return &UserRepo{
        users: map[uuid.UUID]User{},
    }
}

func (r *UserRepo) AddUser(email, password string) uuid.UUID {
    id := uuid.New()
    r.users[id] = User{Email: email, Password: password}
    return id
}

func (r *UserRepo) HasId(id uuid.UUID) bool {
    _, ok := r.users[id]
    return ok
}
