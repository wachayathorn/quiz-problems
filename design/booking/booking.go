package booking

import (
	"context"
	"database/sql"
	"time"
)

type Booking struct {
	RoomID    string
	UserID    string
	StartTime time.Time
	EndTime   time.Time
}

type Repository interface {
	CreateBooking(ctx context.Context, booking Booking) error
}

type repositoryImpl struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repositoryImpl{db: db}
}

func (r *repositoryImpl) CreateBooking(ctx context.Context, booking Booking) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	_, err := r.db.ExecContext(ctx, "INSERT INTO bookings ...")
	return err
}

type Service interface {
	CreateBooking(ctx context.Context, booking Booking) error
}

type serviceImpl struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &serviceImpl{repo: repo}
}

func (s *serviceImpl) CreateBooking(ctx context.Context, booking Booking) error {
	return s.repo.CreateBooking(ctx, booking)
}

type Handler interface {
	CreateBooking(ctx context.Context, booking Booking) error
}

type handlerImpl struct {
	service Service
}

func NewHandler(service Service) Handler {
	return &handlerImpl{service: service}
}

func (h *handlerImpl) CreateBooking(ctx context.Context, booking Booking) error {
	return h.service.CreateBooking(ctx, booking)
}

func Setup(db *sql.DB) Handler {
	repo := NewRepository(db)
	svc := NewService(repo)
	return NewHandler(svc)
}
