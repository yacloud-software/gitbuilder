package builder

type StandardBuildInfo struct {
	Commit       string
	RepoID       uint64
	RepoName     string
	ArtefactName string
	Build        uint64
}

func (s *StandardBuildInfo) BuildNumber() uint64 {
	return s.Build
}
func (s *StandardBuildInfo) CommitID() string {
	return s.Commit
}
func (s *StandardBuildInfo) RepositoryName() string {
	return s.RepoName
}
func (s *StandardBuildInfo) RepositoryID() uint64 {
	return s.RepoID
}
func (s *StandardBuildInfo) RepositoryArtefactName() string {
	return s.ArtefactName
}
