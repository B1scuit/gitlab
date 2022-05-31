package project

import "github.com/B1scuit/gitlab/types"

type Project struct {
	List List
}

func NewProject(opts *types.ClientOptions) *Project {
	return &Project{
		List: newListFunc(opts),
	}
}
