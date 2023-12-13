package appinfo

import (
	ai "golang.conradwood.net/go-easyops/appinfo"
)

// this file is automatically updated by the build server
// whenever a non-go-framework is build and has file package vendored

var (
	APP_BUILD_NUMBER        = 0                                                                                              // replaceme
	APP_BUILD_DESCRIPTION   = "Build #0 of test_reponame at 2021-11-08 08:58:58.325912374 +0000 GMT m=+0.142678563 on host " // replaceme
	APP_BUILD_TIMESTAMP     = 1636361938                                                                                     // replaceme
	APP_BUILD_REPOSITORY    = "test_reponame"                                                                                // replaceme
	APP_BUILD_REPOSITORY_ID = 1                                                                                              // replaceme
	APP_BUILD_COMMIT        = ""                                                                                             // replaceme
)

func init() {
	avi := &ai.AppVersionInfo{
		Number:         uint64(APP_BUILD_NUMBER),
		Description:    APP_BUILD_DESCRIPTION,
		Timestamp:      int64(APP_BUILD_TIMESTAMP),
		RepositoryID:   uint64(APP_BUILD_REPOSITORY_ID),
		RepositoryName: APP_BUILD_REPOSITORY,
		CommitID:       APP_BUILD_COMMIT,
	}
	ai.RegisterAppInfo(avi)
}




