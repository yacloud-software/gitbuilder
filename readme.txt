this is part of the CI/CD chain to build various repositories, e.g. java, golang, kicad etc.
This replaced the functionality built-in to the SVC backends (gitserver/gerrithooks) to build on-the-fly 

the scripts needs
apt-get install toilet make bzip2 git

needs "build-repo-client" in /usr/local/bin
needs "protorender-client" in /usr/local/bin
needs "gitserver-credentials" in /usr/local/bin


to compile singingcat firmware:

apt-get install gcc-arm-none-eabi gcc 
needs "binpatch" in /usr/local/bin

to compile espressif firmware:
/srv/singingcat/esp32
apt-get install cmake


==================== /etc/gitconfig ================
[credential]
        helper="/usr/local/bin/gitserver-credentials -registry=registry"
        useHttpPath = true
[pull]
        rebase = false




a new feature (true by default) ships with a gitconfig, so that the above is no longer required for gitbuilder.




========== rules ======
STANDARD_C
  expects a subdirectory 'c' and underneath subdirectories with a Makefile each.
  for example:
      c/src1/Makefile
      c/src2/Makefile
   the Makefile will be passed DIST=[distdir].
   all binaries must be compiled into ${DIST}/something

make will be invoked with the following variables:
 BUILD_NUMBER=123
 REPOSITORY_ID=123
 ARTEFACT_NAME=fooar
 GIT_URL=https://foo/bar

the above are also set as BUILDVARS:
BUILDVARS="-DBUILD_NUMBER=123 -DREPOSITORY_ID=123 -DGIT_URL='https://foo/bar' -DARTEFACT_NAME"

a typical Makefile is something like this:
BUILDVARS='-DBUILD_ID=1 -DREPOSITORY_ID=1 -DGIT_URL="notset" -DARTEFACT_NAME="none"'
build:
	gcc foo.c ${BUILDVARS}

and in foo.c:


#define STRINGIZE(x) #x
#define STRINGIZE_VALUE_OF(x) STRINGIZE(x)


int main(int argc,char **argv) {
    printf("Buildnumber: %s, RepositoryID: %s, Artefact: %s, Git: %s\n",STRINGIZE_VALUE_OF(BUILD_ID),STRINGIZE_VALUE_OF(REPOSIITORY_ID),STRINGIZE_VALUE_OF(ARTEFACT_NAME),STRINGIZE_VALUE_OF(GIT_URL));
    }
