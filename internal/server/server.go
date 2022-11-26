package server

// func (s *Server) Start() error {
// 	idleConnClosed := make(chan struct{})

// 	go func() {
// 		sigint := make(chan os.Signal, 1)
// 		signal.Notify(sigint, syscall.SIGINT, syscall.SIGTERM)
// 		<-sigint

// 		s.logger.Info("shutting down http server")

// 		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
// 		defer cancel()
// 		if err := s.httpServer.Shutdown(ctx); err != nil {
// 			s.logger.Error("failed to shutdown http server", zap.Error(err))
// 			return
// 		}

// 		s.logger.Info("http server is stopped")

// 		close(idleConnClosed)
// 	}()

// 	s.logger.Info("starting http server")
// 	if err := s.httpServer.ListenAndServe(); err != http.ErrServerClosed {
// 		return fmt.Errorf("failed to listen and serve: %v", err)
// 	}

// 	<-idleConnClosed

// 	s.afterShutdown()

// 	return nil
// }
