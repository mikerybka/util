package util

import (
	"fmt"
	"time"
)

type CI struct {
	PeriodMinutes int
	SourceDir     string
	OutFile       string
	ServiceName   string
	TwilioClient  *TwilioClient
	AdminPhone    string
}

func (ci *CI) Start() {
	for {
		ok, err := ci.Run()
		if err != nil {
			fmt.Println(time.Now(), "ERROR", err)
		}
		if ok {
			fmt.Println(time.Now(), "DEPLOY")
			ci.TwilioClient.SendSMS(ci.AdminPhone, "New deploy!")
		}
		time.Sleep(time.Minute * time.Duration(ci.PeriodMinutes))
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

func (ci *CI) Run() (bool, error) {
	repo := ci.GitRepo()
	code := ci.Codebase()
	svc := ci.Serivce()

	// Pull from github.
	_, err := repo.Pull()
	if err != nil {
		return false, err
	}

	// Run `go get -u` update dependencies.
	changes, err := code.UpdateDeps()
	if err != nil {
		return false, err
	}

	// If no changes, we're done.
	if !changes {
		return false, nil
	}

	// Build the new code.
	err = code.Build(ci.OutFile)
	if err != nil {
		return false, err
	}

	// Restart the server.
	err = svc.Restart()
	if err != nil {
		return false, err
	}

	// Commit and push the changes upstream.
	err = repo.AddAll()
	if err != nil {
		return false, err
	}
	err = repo.Commit("update deps")
	if err != nil {
		return false, err
	}
	err = repo.Push()
	if err != nil {
		return false, err
	}

	return true, nil
}
