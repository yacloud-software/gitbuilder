package buildinfo

type BuildInfo interface {
	CommitID() string
	RepositoryID() uint64
	RepositoryName() string
	RepositoryArtefactName() string
	BuildNumber() uint64
}
