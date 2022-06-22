package cmd

import (
	"sort"
	"sync"

	"github.com/j3ssie/sfleet/core"
	"github.com/j3ssie/sfleet/utils"
	"github.com/panjf2000/ants"
	"github.com/spf13/cobra"
)

func init() {
	var workerCmd = &cobra.Command{
		Use:     "worker",
		Aliases: []string{"w"},
		Short:   "Worker Utility",
		RunE:    runWorker,
	}
	RootCmd.AddCommand(workerCmd)
}

var workerPool = core.WorkerPool{}

func runWorker(cmd *cobra.Command, args []string) error {
	sort.Strings(args)
	if len(args) == 0 {
		args = append(args, "add")
	}

	existWorkerPool, err := core.GetConfig(options)
	if err == nil {
		workerPool = existWorkerPool
	}

	var wg sync.WaitGroup
	p, _ := ants.NewPoolWithFunc(options.Concurrency, func(i interface{}) {
		host := i.(string)
		workerHandle(host, args)
		wg.Done()
	}, ants.WithPreAlloc(true))
	defer p.Release()

	for _, host := range options.Inputs {
		wg.Add(1)
		_ = p.Invoke(host)
	}
	wg.Wait()

	// writing worker pool to config file
	core.WriteConfig(workerPool, options)

	return nil
}

func workerHandle(host string, args []string) {
	switch args[0] {
	case "add":
		// add worker
		worker, err := core.ValidateWorker(host, options)
		if err == nil {
			// check if it exist for not
			for _, w := range workerPool.Workers {
				if w.Host == host {
					utils.ErrorF("worker %v already exists", host)
					return
				}
			}
			workerPool.Workers = append(workerPool.Workers, worker)
		}
	case "clean":
		workerPool.Workers = []core.Worker{}
	}
}
