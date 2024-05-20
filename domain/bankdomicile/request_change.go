package bankdomicile

import (
	"context"

	"github.com/aronrcihter/monorepo/domain/vo"
)

type RequestChangeUseCase struct {
	repo RequestChangeReposotory
}

func NewRequestChangeUseCase(repo RequestChangeReposotory) *RequestChangeUseCase {
	return &RequestChangeUseCase{
		repo: repo,
	}
}

type RequestChangeInput struct {
	Name  string
	Email string
}

type RequestChangeOutput struct {
	ID string
}

type RequestChangeReposotory interface {
	Create(ctx context.Context, input vo.RequestChangeInput) (vo.RequestChangeOutput, error)
}

func (uc RequestChangeUseCase) RequestChange(ctx context.Context, input RequestChangeInput) (RequestChangeOutput, error) {
	// start input

	output, err := uc.repo.Create(ctx, vo.RequestChangeInput{
		//some input
	})
	if err != nil {
		//set span error
		return RequestChangeOutput{}, err
	}

	return RequestChangeOutput{
		// some return
	}, nil
}
