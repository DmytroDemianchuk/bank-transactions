package repositories

import "github.com/jmoiron/sqlx"

type Repositories struct {
	RepoBank   *RepoBank
	RepoRemote *RepoRemote
}

func New(db *sqlx.DB) *Repositories {
	return &Repositories{
		RepoBank:   NewRepoBank(db),
		RepoRemote: NewRemoteRepo(db),
	}
}
