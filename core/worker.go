package core

import (
	"github.com/j3ssie/sfleet/libs"
	"github.com/j3ssie/sfleet/utils"
)

func ValidateWorker(host string, opt libs.Options) (wk Worker, err error) {
	utils.InforF("adding worker %v", host)
	runner, _ := InitProvider(host, opt)
	err = runner.Connect()
	if err != nil {
		utils.ErrorF("failed to connect to %v", host)
		return wk, err
	}

	wk = Worker{
		Name:       host,
		Host:       host,
		Port:       opt.Port,
		User:       opt.User,
		Key:        utils.NormalizePath(opt.SSHPrivateKey),
		PassPhrase: opt.Credentials,
	}
	return wk, nil
}
