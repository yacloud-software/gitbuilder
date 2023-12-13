package buildinfo

type BuildInfo interface {
	CommitID() string
	RepositoryID() uint64
	RepositoryName() string
	RepositoryArtefactName() string
	ArtefactID() uint64
	BuildNumber() uint64
	IsScriptIncluded(name string) bool
	GitURL() string
}




