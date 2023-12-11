// client create: GitBuilderClient
/*
  Created by /home/cnw/devel/go/yatools/src/golang.yacloud.eu/yatools/protoc-gen-cnw/protoc-gen-cnw.go
*/

/* geninfo:
   filename  : protos/golang.conradwood.net/apis/gitbuilder/gitbuilder.proto
   gopackage : golang.conradwood.net/apis/gitbuilder
   importname: ai_0
   clientfunc: GetGitBuilder
   serverfunc: NewGitBuilder
   lookupfunc: GitBuilderLookupID
   varname   : client_GitBuilderClient_0
   clientname: GitBuilderClient
   servername: GitBuilderServer
   gsvcname  : gitbuilder.GitBuilder
   lockname  : lock_GitBuilderClient_0
   activename: active_GitBuilderClient_0
*/

package gitbuilder

import (
   "sync"
   "golang.conradwood.net/go-easyops/client"
)
var (
  lock_GitBuilderClient_0 sync.Mutex
  client_GitBuilderClient_0 GitBuilderClient
)

func GetGitBuilderClient() GitBuilderClient { 
    if client_GitBuilderClient_0 != nil {
        return client_GitBuilderClient_0
    }

    lock_GitBuilderClient_0.Lock() 
    if client_GitBuilderClient_0 != nil {
       lock_GitBuilderClient_0.Unlock()
       return client_GitBuilderClient_0
    }

    client_GitBuilderClient_0 = NewGitBuilderClient(client.Connect(GitBuilderLookupID()))
    lock_GitBuilderClient_0.Unlock()
    return client_GitBuilderClient_0
}

func GitBuilderLookupID() string { return "gitbuilder.GitBuilder" } // returns the ID suitable for lookup in the registry. treat as opaque, subject to change.

func init() {
   client.RegisterDependency("gitbuilder.GitBuilder")
   AddService("gitbuilder.GitBuilder")
}


