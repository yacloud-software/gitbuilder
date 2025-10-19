package git

import "testing"

func TestLogMessages(t *testing.T) {
	testlogmessage(t, "upgraded", ` commit 1d52ebd0c733ed040f7b56059a0d18f771558015             
 Author: repomodifier yatools <automatic@yacloud.localdomain>
 Date:   Sun Oct 19 08:20:48 2025 +0000                      
                                                             
     upgraded                                                
`)

	testlogmessage(t, "testing logmessages", ` commit ec69c3718456e31f6c2f88b8fb96b4a8226fdac4             
 Author: Conrad Wood <cnw@conradwood.net>                    
 Date:   Sun Oct 19 09:20:46 2025 +0100                      
                                                             
     testing logmessages                                     
                                                             
     Change-Id: I8fdea54c81d9fb9329227ed7a72d4cb65f2da20d    
`)

}

func testlogmessage(t *testing.T, expected, teststr string) {
	n := tidyLogMessage(teststr)
	if n == expected {
		return
	}
	t.Logf("failed. on teststring:\n%s\n==== expected: =====\n\"%s\"\n===== but got: ======\n\"%s\"\n", teststr, expected, n)
	t.Fail()

}
