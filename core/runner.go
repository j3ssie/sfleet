package core

import (
	"time"

	"github.com/cenkalti/backoff/v4"
	"github.com/fatih/color"
	"github.com/j3ssie/sfleet/libs"
	"github.com/j3ssie/sfleet/utils"
	"github.com/melbahja/goph"
)

type Runner struct {
	Host          string
	SSHPrivateKey string
	SSHUser       string
	PassPhrase    string
	Commands      []string

	BackOff *backoff.ExponentialBackOff
	Client  *goph.Client

	Opt libs.Options
}

// InitProvider init provider object to easier interact with cloud provider
func InitProvider(host string, opt libs.Options) (Runner, error) {
	var runner Runner
	runner.Host = host
	runner.SSHUser = "root"
	runner.PassPhrase = opt.Credentials
	runner.SSHPrivateKey = utils.NormalizePath(runner.Opt.SSHPrivateKey)

	if runner.SSHPrivateKey == "" {
		runner.SSHPrivateKey = utils.NormalizePath("~/.ssh/id_rsa")
	}

	// command
	runner.Commands = opt.Commands

	// for retry
	b := backoff.NewExponentialBackOff()
	// It never stops if MaxElapsedTime == 0.
	b.MaxElapsedTime = 1200 * time.Second
	b.Multiplier = 2.0
	b.InitialInterval = 30 * time.Second
	runner.BackOff = b

	return runner, nil
}

func (r *Runner) Start() error {
	err := r.Action(connect)
	if err == nil {
		err = r.Action(run)
	}
	return err
}

const (
	connect = "connect"
	run     = "run"
)

func (r *Runner) Action(action string) (err error) {
	utils.InforF("action: %v -- %v", color.HiBlueString(action), r.Host)
	operation := func() error {
		switch action {
		case connect:
			err = r.Connect()
		case run:
			for _, cmd := range r.Commands {
				r.RunCommnad(cmd)
			}
		}
		return err
	}
	err = backoff.Retry(operation, r.BackOff)
	if err != nil {
		utils.WarnF("error create instance action %v", r.Host)
		return err
	}
	return nil
}
