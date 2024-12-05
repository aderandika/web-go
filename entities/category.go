package entities

import "time"

// Memanggil data dari tabel
type Category struct {
	Id        uint
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
