package util

import (
	"time"
)

type CI struct {
	PeriodMinutes int
	SourceDir     string
	OutFile       string
	ServiceName   string
}

func (ci *CI) Start() error {
	for {
		time.Sleep(time.Minute * time.Duration(ci.PeriodMinutes))
		err := ci.Run()
		if err != nil {
			return err
		}
	}
}

func (ci *CI) GitRepo() *GitRepo {
	return &GitRepo{
		Dir: ci.SourceDir,
	}
}

func (ci *CI) Codebase() *GoCodebase {
	return &GoCodebase{
		Dir: ci.SourceDir,
	}
}

func (ci *CI) Serivce() *SystemdService {
	return &SystemdService{
		Name: ci.ServiceName,
	}
}

func (ci *CI) Run() error {
	repo := ci.GitRepo()
	code := ci.Codebase()
	svc := ci.Serivce()

	// Pull from github.
	_, err := repo.Pull()
	if err != nil {
		return err
	}

	// Run `go get -u` update dependencies.
	changes, err := code.UpdateDeps()
	if err != nil {
		return err
	}

	// If there's been changes, rebuild and restart the server.
	if changes {
		err := code.Build(ci.OutFile)
		if err != nil {
			return err
		}

		err = svc.Restart()
		if err != nil {
			return err
		}
	}

	return nil
}
