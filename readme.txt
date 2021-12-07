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



==================== /etc/gitconfig ================
[credential]
        helper="/usr/local/bin/gitserver-credentials -registry=registry"
        useHttpPath = true
[pull]
        rebase = false









