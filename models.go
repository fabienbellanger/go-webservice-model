package gwm

import "fmt"

// SuperUser represents a SuperUser in database.
type SuperUser struct {
	ID       string  `db:"suse_id"`
	Login    string  `db:"suse_login"`
	Password string  `db:"suse_password"`
	Salt     *string `db:"suse_salt"` // Nullable
}

func (su SuperUser) String() string {
	return fmt.Sprintf("SuperUser={id=%v, login: %v, password: %v, salt: %v}\n",
		su.ID, su.Login, su.Password, su.Salt)
}
