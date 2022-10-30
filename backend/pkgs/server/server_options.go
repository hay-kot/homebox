package server

import "time"

type Option = func(s *Server) error

func WithMiddleware(mw ...Middleware) Option {
	return func(s *Server) error {
		s.mw = append(s.mw, mw...)
		return nil
	}
}

func WithWorker(w Worker) Option {
	return func(s *Server) error {
		s.Worker = w
		return nil
	}
}

func WithHost(host string) Option {
	return func(s *Server) error {
		s.Host = host
		return nil
	}
}

func WithPort(port string) Option {
	return func(s *Server) error {
		s.Port = port
		return nil
	}
}

func WithReadTimeout(seconds int) Option {
	return func(s *Server) error {
		s.readTimeout = time.Duration(seconds) * time.Second
		return nil
	}
}

func WithWriteTimeout(seconds int) Option {
	return func(s *Server) error {
		s.writeTimeout = time.Duration(seconds) * time.Second
		return nil
	}
}

func WithIdleTimeout(seconds int) Option {
	return func(s *Server) error {
		s.idleTimeout = time.Duration(seconds) * time.Second
		return nil
	}
}
