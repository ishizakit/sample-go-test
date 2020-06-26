package model

import "time"

type User struct {
	ID         int       `db:"id"`
	Name       string    `db:"name"`
	Email      string    `db:"email"`
	LastActive time.Time `db:"registered_at"`
}

const (
	// activeMonthより前の
	activeMonth = 12
)

// BeforeYearsAgo LastActiveがactiveMonthヶ月より後ならtrue
func (u *User) IsActive() bool {
	return u.LastActive.After(
		time.Now().AddDate(0, -activeMonth, 0),
	)
}
