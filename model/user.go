package model

import (
	supa "github.com/nedpals/supabase-go"
)

type User struct {
	*supa.User
}
