package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/Peyton232/freelance-auction/graph/generated"
	"github.com/Peyton232/freelance-auction/graph/model"
)

func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (*model.User, error) {
	return db.CreateUser(&input), nil
}

func (r *mutationResolver) RemoveUser(ctx context.Context, userid string) (*model.User, error) {
	return db.RemoveUser(userid), nil
}

func (r *mutationResolver) CreateAuction(ctx context.Context, input model.NewAuction) (*model.Auction, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) RemoveAuction(ctx context.Context, auctionID string) (*model.Auction, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) BidOnAuction(ctx context.Context, userID string, auctionID string) (*model.Auction, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) AllUsers(ctx context.Context) ([]*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) UserByID(ctx context.Context, userID string) (*model.User, error) {
	return db.FindUserByID(userID), nil
}

func (r *queryResolver) AllAuctions(ctx context.Context) ([]*model.Auction, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) AuctionByID(ctx context.Context, auctionID string) (*model.Auction, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
