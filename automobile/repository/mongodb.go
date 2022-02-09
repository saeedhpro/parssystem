package repository

import "context"

type MongoDB interface {
	GetAutomobileList(ctx context.Context) error
	GetAutomobileByID(ctx context.Context, id string) error
}
