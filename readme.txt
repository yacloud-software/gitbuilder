this is part of the CI/CD chain to build various repositories, e.g. java, golang, kicad etc.
This replaced the functionality built-in to the SVC backends (gitserver/gerrithooks) to build on-the-fly 

needs:
apt-get install toilet make


needs "build-repo-client" in /usr/local/bin
