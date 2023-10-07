package role

import (
	"errors"
	"github.com/prongbang/user-service/internal/localizations"
)

type UseCase interface {
	Count() int64
	GetList(filter Filter) []Role
	GetListByUnderLevel(level int) []Role
	GetListByUnderRoles(roles []string) []Role
	GetById(id string) Role
	GetByName(name string) Role
	Add(data *CreateRole) error
	Update(data *UpdateRole) error
	Delete(id string) error
}

type useCase struct {
	Repo Repository
}

func (u *useCase) Count() int64 {
	return u.Repo.Count()
}

func (u *useCase) GetList(filter Filter) []Role {
	return u.Repo.GetList(filter)
}

func (u *useCase) GetListByUnderLevel(level int) []Role {
	return u.Repo.GetListByUnderLevel(level)
}

func (u *useCase) GetListByUnderRoles(roles []string) []Role {
	if len(roles) == 0 {
		return []Role{}
	}
	return u.Repo.GetListByUnderRoles(roles)
}

func (u *useCase) GetById(id string) Role {
	return u.Repo.GetById(id)
}

func (u *useCase) GetByName(name string) Role {
	return u.Repo.GetByName(name)
}

func (u *useCase) Add(data *CreateRole) error {
	if rs := u.Repo.GetByName(data.Name); rs.ID != "" {
		return errors.New(localizations.CommonDataIsDuplicated)
	}
	return u.Repo.Add(data)
}

func (u *useCase) Update(data *UpdateRole) error {
	if rs := u.Repo.GetByName(data.Name); rs.ID != "" && rs.ID != data.ID {
		return errors.New(localizations.CommonDataIsDuplicated)
	}
	return u.Repo.Update(data)
}

func (u *useCase) Delete(id string) error {
	return u.Repo.Delete(id)
}

func NewUseCase(repo Repository) UseCase {
	return &useCase{
		Repo: repo,
	}
}
