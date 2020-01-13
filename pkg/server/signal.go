package server

// var onlyOneSignalHandler = make(chan struct{})

// var shutdownSignals = []os.Signal{os.Interrupt, syscall.SIGTERM}

// // SetupSignalHandler registered for SIGTERM and SIGINT. A stop
// // channel is returned which is closed on one of these signals. If a
// // second signal is caught, the program is terminated with exit code
// // 1.
// func SetupSignalHandler() (<-chan struct{}, *sync.WaitGroup) {
// 	close(onlyOneSignalHandler) // panics when called twice

// 	quitTimeout := time.Duration(10)
// 	quit := make(chan struct{})
// 	wg := &sync.WaitGroup{}
// 	waitGroupDone := make(chan struct{})
// 	c := make(chan os.Signal, 2)
// 	signal.Notify(c, shutdownSignals...)

// 	go func() {
// 		<-c
// 		glog.Warningln("Caught shutdown signal in signal handler")
// 		close(quit)
// 		go waitForWaitGroup(wg, waitGroupDone)
// 		select {
// 		case <-waitGroupDone:
// 			os.Exit(0)
// 		case <-c:
// 			glog.Errorln("Shutdown called twice, forcing exit")
// 			os.Exit(1)
// 		case <-time.After(time.Second * quitTimeout):
// 			glog.Errorf(
// 				"Loops were still running after %d seconds, forcing exit",
// 				quitTimeout)
// 			os.Exit(2)
// 		}
// 		// if we get a second signal, exit directly
// 	}()

// 	return quit, wg
// }

// func waitForWaitGroup(wg *sync.WaitGroup, waitGroupDone chan struct{}) {
// 	wg.Wait()
// 	glog.Info("All controllers have exited")
// 	waitGroupDone <- struct{}{}
// }
