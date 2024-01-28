package canonical

import "time"

type Login struct {
	Login    string
	Password string
}

type User struct {
	Id            string
	Login         string
	Password      string
	AccessLevelID int
}

type Access struct {
	ID            string
	AccessLevelID int
	USERID        string
	AccessedAt    time.Time
}
