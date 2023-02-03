package builder

import (
	pb "golang.conradwood.net/apis/gitbuilder"
)

type StandardBuildInfo struct {
	Req *pb.BuildRequest
	/*
		Commit       string
		RepoID       uint64
		RepoName     string
		ArtefactName string
		Build        uint64
	*/
}

func (s *StandardBuildInfo) BuildNumber() uint64 {
	return s.Req.BuildNumber
}
func (s *StandardBuildInfo) CommitID() string {
	return s.Req.CommitID
}
func (s *StandardBuildInfo) RepositoryName() string {
	return s.Req.RepoName
}
func (s *StandardBuildInfo) RepositoryID() uint64 {
	return s.Req.RepositoryID
}
func (s *StandardBuildInfo) RepositoryArtefactName() string {
	return s.Req.ArtefactName
}
func (s *StandardBuildInfo) IsScriptIncluded(scriptname string) bool {
	if len(s.Req.IncludeBuildScripts) == 0 && len(s.Req.ExcludeBuildScripts) == 0 {
		return true
	}
	for _, cn := range s.Req.ExcludeBuildScripts {
		if cn == scriptname {
			return false
		}
	}
	return true
}
