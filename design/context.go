package design

import (
	"context"
	"database/sql"
	"time"
)

// Model section
type Booking struct {
	RoomID    string
	UserID    string
	StartTime time.Time
	EndTime   time.Time
}

// Repository section
type BookingRepository interface {
	CreateBooking(ctx context.Context, booking Booking) error
}

type bookingImpl struct {
	db *sql.DB
}

func NewBookingRepository(db *sql.DB) BookingRepository {
	return &bookingImpl{db: db}
}

func (r *bookingImpl) CreateBooking(ctx context.Context, booking Booking) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	_, err := r.db.ExecContext(ctx, "INSERT INTO bookings ...")
	return err
}

// Service section
type BookingService interface {
	CreateBooking(ctx context.Context, booking Booking) error
}

type bookingService struct {
	bookingRepo BookingRepository
}

func NewBookingService(bookingRepo BookingRepository) BookingService {
	return &bookingService{bookingRepo: bookingRepo}
}

func (s *bookingService) CreateBooking(ctx context.Context, booking Booking) error {
	return s.bookingRepo.CreateBooking(ctx, booking)
}

// Handler section
type BookingHandler interface {
	CreateBooking(ctx context.Context, booking Booking) error
}

type bookingHandler struct {
	bookingService BookingService
}

func NewBookingHandler(bookingService BookingService) BookingHandler {
	return &bookingHandler{bookingService: bookingService}
}

func (h *bookingHandler) CreateBooking(ctx context.Context, booking Booking) error {
	return h.bookingService.CreateBooking(ctx, booking)
}

// Main section
func setup() {
	db, err := sql.Open("mysql", "user:password@tcp(localhost:3306)/dbname")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	bookingRepo := NewBookingRepository(db)
	bookingService := NewBookingService(bookingRepo)
	bookingHandler := NewBookingHandler(bookingService)

	// This ctx assume is from HTTP request
	ctx := context.Background()
	booking := Booking{
		RoomID:    "room1",
		UserID:    "user1",
		StartTime: time.Now(),
		EndTime:   time.Now().Add(time.Hour),
	}
	err = bookingHandler.CreateBooking(ctx, booking)
	if err != nil {
		panic(err)
	}
}
