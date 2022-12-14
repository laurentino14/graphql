package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/laurentino14/graphql/graph/generated"
	"github.com/laurentino14/graphql/graph/model"
	"github.com/laurentino14/graphql/prisma/client"
)

// CreateUser is the resolver for the createUser field.
func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (*model.User, error) {
	ct := client.NewClient()
	if err := ct.Prisma.Connect(); err != nil {
		return nil, err
	}

	defer func() {
		if err := ct.Prisma.Disconnect(); err != nil {
			panic(err)
		}
	}()
	req, err := ct.User.CreateOne(
		client.User.Firstname.Set(input.Firstname),
		client.User.Lastname.Set(input.Lastname),
		client.User.Email.Set(input.Email),
		client.User.Password.Set(input.Password),
		client.User.Cellphone.Set(input.Cellphone),
		client.User.BirthDate.Set(*input.BirthDate),
		client.User.Tokenuser.Set(*input.TokenUser),
		client.User.ID.Set(*input.ID),
	).Exec(ctx)
	if err != nil {
		return nil, err
	}
	user := &model.User{ID: req.ID,
		Firstname:   req.Firstname,
		Lastname:    req.Lastname,
		Password:    "***********",
		Email:       req.Email,
		Cellphone:   req.Cellphone,
		BirthDate:   &req.BirthDate,
		TokenUser:   &req.Tokenuser,
		AuthorOf:    []*model.Course{},
		Enrollments: []*model.Enrollment{}}
	return user, err
}

// Users is the resolver for the users field.
func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	ct := client.NewClient()
	if err := ct.Prisma.Connect(); err != nil {
		return nil, err
	}

	defer func() {
		if err := ct.Prisma.Disconnect(); err != nil {
			panic(err)
		}
	}()
	req, err := ct.User.FindMany().Exec(ctx)
	if err != nil {
		return nil, err
	}

	users := []*model.User{}
	for _, user := range req {
		users = append(users, &model.User{
			ID:          user.ID,
			Firstname:   user.Firstname,
			Lastname:    user.Lastname,
			Password:    "***********",
			Email:       user.Email,
			Cellphone:   user.Cellphone,
			BirthDate:   &user.BirthDate,
			TokenUser:   &user.Tokenuser,
			AuthorOf:    []*model.Course{},
			Enrollments: []*model.Enrollment{},
		})
	}

	return users, nil
}

// Courses is the resolver for the courses field.
func (r *queryResolver) Courses(ctx context.Context) ([]*model.Course, error) {
	panic(fmt.Errorf("not implemented: Courses - courses"))
}

// Modules is the resolver for the modules field.
func (r *queryResolver) Modules(ctx context.Context) ([]*model.Module, error) {
	panic(fmt.Errorf("not implemented: Modules - modules"))
}

// Lessons is the resolver for the lessons field.
func (r *queryResolver) Lessons(ctx context.Context) ([]*model.Lesson, error) {
	panic(fmt.Errorf("not implemented: Lessons - lessons"))
}

// Enrollments is the resolver for the enrollments field.
func (r *queryResolver) Enrollments(ctx context.Context) ([]*model.Enrollment, error) {
	panic(fmt.Errorf("not implemented: Enrollments - enrollments"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
