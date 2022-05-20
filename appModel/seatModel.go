package appModel

type SeatModel interface {
	GetAllSeat() ([]Seat, error)
	AddSeat(Seat) (Seat, error)
}
