package rest

import "tech-challenge-auth/internal/canonical"

func (c *LoginRequest) toCanonical() canonical.Login {
	return canonical.Login{
		Login:    c.Login,
		Password: c.Password,
	}
}
