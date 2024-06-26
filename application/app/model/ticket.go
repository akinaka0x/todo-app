package model

import "gorm.io/gorm"

type Ticket struct {
	ID          uint
	Title       string
	Description string
	Completed   bool
}

func GetTickets(tickets *[]Ticket, db *gorm.DB) error {
	err := db.
		Model(&Ticket{}).
		Find(tickets).
		Error
	return err
}

func GetTicketsByStatus(tickets *[]Ticket, status bool, db *gorm.DB) error {
	err := db.
		Model(&Ticket{}).
		Where("Completed = ?", status).
		Find(tickets).
		Error
	return err
}

func GetTicket(ticket *Ticket, id uint, db *gorm.DB) error {
	err := db.
		Model(&Ticket{}).
		First(ticket, id).
		Error
	return err
}

func CreateTicket(ticket *Ticket, db *gorm.DB) error {
	err := db.
		Model(&Ticket{}).
		Create(ticket).
		Error
	return err
}

func CompleteTicket(id uint, db *gorm.DB) error {
	err := db.
		Model(&Ticket{}).
		Where("id = ?", id).
		Update("Completed", true).
		Error
	return err
}

func DeleteTicket(id uint, db *gorm.DB) error {
	err := db.Delete(&Ticket{}, id).Error
	return err
}
