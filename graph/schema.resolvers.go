package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"demo/graph/common"
	"demo/graph/generated"
	"demo/graph/model"
	"strconv"
	"time"

	"gorm.io/gorm/clause"
)

// CreateTodo is the resolver for the createTodo field.
func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	context := common.GetContext(ctx)
	var myuser *model.User
	context.Database.FirstOrCreate(&myuser, model.User{
		ID:   input.UserID,
		Name: "test",
	})

	todo := &model.Todo{
		Text: input.Text,
		Done: false,
		User: myuser,
	}
	err := context.Database.Create(&todo).Error
	if err != nil {
		return nil, err
	}

	return todo, nil
}

// CreateSchedule is the resolver for the createSchedule field.
func (r *mutationResolver) CreateSchedule(ctx context.Context, input *model.NewSchedule) (*model.Schedule, error) {
	context := common.GetContext(ctx)

	last := &model.Schedule{}
	context.Database.Clauses(clause.OrderBy{
		Expression: clause.Expr{SQL: "CAST(id as INTEGER) DESC"},
	}).Find(last)
	id, err1 := strconv.Atoi(last.ID)
	if err1 != nil {
		id = 0
	}
	print(id)

	schedule := &model.Schedule{
		ID:      strconv.Itoa(id + 1),
		Summary: input.Summary,
		Start:   input.Start,
		End:     input.End,
	}

	err := context.Database.Create(&schedule).Error
	if err != nil {
		return nil, err
	}

	return schedule, nil
}

// DeleteSchedule is the resolver for the deleteSchedule field.
func (r *mutationResolver) DeleteSchedule(ctx context.Context, input *model.ScheduleID) (bool, error) {
	db := common.GetContext(ctx).Database
	err := db.Delete(&model.Schedule{}, input.ID).Error
	return err == nil, err
}

// ModifySchedule is the resolver for the modifySchedule field.
func (r *mutationResolver) ModifySchedule(ctx context.Context, input *model.ScheduleTime) (bool, error) {
	db := common.GetContext(ctx).Database
	err := db.Model(&model.Schedule{}).Where("id = ?", input.ID).Updates(model.Schedule{
		Start: input.Start,
		End:   input.End,
	}).Error

	return err == nil, err
}

// ModifyScheduleTitle is the resolver for the modifyScheduleTitle field.
func (r *mutationResolver) ModifyScheduleTitle(ctx context.Context, input *model.ScheduleTitle) (bool, error) {
	db := common.GetContext(ctx).Database
	err := db.Model(&model.Schedule{}).Where("id = ?", input.ID).Updates(model.Schedule{
		Summary: input.Summary,
	}).Error

	return err == nil, err
}

// Todos is the resolver for the todos field.
func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	var items []*model.Todo
	rs, err := common.GetAll(ctx, items)
	return rs.([]*model.Todo), err
}

// Schedules is the resolver for the schedules field.
func (r *queryResolver) Schedules(ctx context.Context, start *time.Time, end *time.Time) ([]*model.Schedule, error) {
	var items []*model.Schedule
	context := common.GetContext(ctx)
	var err error
	if start != nil && end != nil {
		err = context.Database.Where("datetime(start) >= datetime(?) and datetime(end)  <= datetime(?) ", start, end).Find(&items).Error

		if err == nil {
			return items, err
		}

		return nil, err
	}
	rs, err := common.GetAll(ctx, items)
	return rs.([]*model.Schedule), err

}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
