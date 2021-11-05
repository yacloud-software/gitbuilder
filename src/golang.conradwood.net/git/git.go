package git

type LocalRepo struct {
	url       string
	fetchurls []string
	inuse     bool
}

// clone a repo, check it out to current head in master and fetch optional urls too
func GetLocalRepo(url string, fetchurls []string) (*LocalRepo, error) {
	lr := &LocalRepo{
		url:       url,
		fetchurls: fetchurls,
		inuse:     true,
	}
	return lr, nil
}

func (lr *LocalRepo) Release() {
	lr.inuse = false
}
