# Workload Agent 

`Workload Agent` is used to launch encrypted workloads on a trusted host.

## Key features
- Start a guest VM
- Stop a quest VM
- Create VM/container trust report
- Fetch a flavor from Workload Service


## System Requirements
- RHEL 8.1
- Epel 8 Repo
- Proxy settings if applicable

## Software requirements
- git
- makeself
- `go` version >= `go1.12.1` & <= `go1.14.4`
- docker 18.06 or higher
- libvirt
- cryptsetup

# Step By Step Build Instructions

## Install required shell commands

### Install tools from `yum`
```shell
sudo yum install -y git wget makeself
```

### Install `go` version >= `go1.12.1` & <= `go1.14.4`
The `Workload Agent` requires Go version 1.12.1 that has support for `go modules`. The build was validated with the latest version go1.14.4 of `go`. It is recommended that you use go1.14.4 version of `go`. You can use the following to install `go`.
```shell
wget https://dl.google.com/go/go1.14.4.linux-amd64.tar.gz
tar -xzf go1.14.4.linux-amd64.tar.gz
sudo mv go /usr/local
export GOROOT=/usr/local/go
export PATH=$GOPATH/bin:$GOROOT/bin:$PATH
```

### Install Docker-CE
Docker-CE is only required for the container confidentiality with secure-docker-daemon use case. Red Hat no longer supports Docker-CE natively, so the packages must be manually installed. The build was validated with version 19.03 of Docker-CE. 
```shell
yum install https://download.docker.com/linux/centos/7/x86_64/stable/Packages/containerd.io-1.2.6-3.3.el7.x86_64.rpm
yum config-manager --add-repo=https://download.docker.com/linux/centos/docker-ce.repo
yum install docker-ce
systemctl disable firewalld
systemctl enable --now docker
```

## Build Workload Agent(WLA)

- Git clone the WLA
- Run scripts to build the WLA

### Build full installer

Supports the following use cases:
- Virtual machine confidentiality
- Container confidentiality using secure-docker-daemon
- Container confidentiality using skopeo and cri-o

```shell
git clone https://github.com/intel-secl/workload-agent.git
cd workload-agent
make installer
```

### Build installer without the secure docker daemon
### This is used for Skopeo encryption and Cri-o decryption for container confidentiality.
### Note: This installer should also be used if exclusively using ISecL for virtual machine confidentiality.

```shell
git clone https://github.com/intel-secl/workload-agent.git
cd workload-policy-manager
make installer-no-docker
```

# Third Party Dependencies

## WLA

### Direct dependencies

| Name                  | Repo URL                        | Minimum Version Required           |
| ----------------------| --------------------------------| :--------------------------------: |
| logrus                | github.com/sirupsen/logrus      | v1.4.2                             |
| testify               | github.com/stretchr/testify     | v1.3.0                             |
| yaml.v2               | gopkg.in/yaml.v2                | v2.2.2                             |
| fs notify             | github.com/fsnotify/fsnotify    | v1.4.7                             |


### Indirect Dependencies

| Repo URL                          | Minimum version required           |
| ----------------------------------| :--------------------------------: |
| github.com/Gurpartap/logrus-stack | v0.0.0-20170710170904-89c00d8a28f4 |
| github.com/facebookgo/stack       | v0.0.0-20160209184415-751773369052 |
| golang.org/x/net                  | v0.0.0-20190206173232-65e2d4e15006 |

*Note: All dependencies are listed in go.mod*

# Links
# Links
https://01.org/intel-secl/
