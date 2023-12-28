package models
import (
	"time"
	"example.com/restAPI/db"
)
type Event struct {
	ID int64
	Name string `binding:"required"`
	Description string `binding:"required"`
	Location string `binding:"required"`
	DateTime time.Time
	UserID int
}
func (e Event) Save() error{
	// later add it to db
	query := `INSERT INTO events(name, description, location, dateTime, user_id)
			  VALES(?, ?, ?, ?, ?)`
	stmt, err := db.DB.Prepare(query)
	if err!= nil {
		return err
	}
	defer stmt.Close()
	result, err := stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserID)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	e.ID = id
	return err
}

func GetAllEvents() ([]Event, error){
	query := "SELECT * from events"
	rows, err := db.DB.Query(query)
	if err!= nil {
		return nil, err
	}
	defer rows.Close()
	var events []Event
	for rows.Next() {
		var event Event
		err := rows.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)
		if err != nil {
			return nil, err
		}
		events = append(events, event)
	}
	return events, nil
}

func GetEventById(ind int64) (*Event, error){
	query := "SELECT * from events WHERE id = ?"
	row := db.DB.QueryRow(query)
	var event Event
	err := row.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)
	if err!=nil {
		return nil, err
	}
	return &event, nil
}