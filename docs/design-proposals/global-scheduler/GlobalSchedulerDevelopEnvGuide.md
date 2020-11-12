## Set up developer environment

Nov-11-2020, Wang Jun

Note: tested on AWS EC2 Ubuntu 18.04 x86 image, because the build process need large memory, at least 8G memorys are needed.

### Install Golang
```
GOLANG_VERSION=${GOLANG_VERSION:-"1.12.17"}
wget https://dl.google.com/go/go${GOLANG_VERSION}.linux-amd64.tar.gz -P /tmp
sudo tar -C /usr/local -xzf /tmp/go${GOLANG_VERSION}.linux-amd64.tar.gz

# export go path in your shell
export GOROOT=/usr/local/go
export PATH=$PATH:$GOROOT/bin
```

### Install make & gcc
```
sudo apt -y install make
sudo apt -y install gcc
sudo apt -y install jq
```

### Install docker
```
sudo apt -y install docker.io
```

### Fork and clone repo
```
# fork global resource scheduler repo



```

