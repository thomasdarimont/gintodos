package model

type RepositoryInitializer interface {
	Initialize(r TodoRepository)
}
