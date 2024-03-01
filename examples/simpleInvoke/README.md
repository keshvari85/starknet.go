Note: To run this example, you need a testnet endpoint.

Steps to run this example on testnet:


1. Install GO:
   
 ```bash
sudo apt-get install curl
 ```
 ```bash
VERSION="1.21.6"
 ```
 ```bash
ARCH="amd64"
 ```
 ```bash
curl -O -L "https://golang.org/dl/go${VERSION}.linux-${ARCH}.tar.gz"
 ```
 ```bash
wget -L "https://golang.org/dl/go${VERSION}.linux-${ARCH}.tar.gz"
 ```
 ```bash
wget -L "https://golang.org/dl/go${VERSION}.linux-${ARCH}.tar.gz"
 ```
 ```bash
curl -sL https://golang.org/dl/ | grep -A 5 -w "go${VERSION}.linux-${ARCH}.tar.gz"
 ```
 ```bash
tar -xf "go${VERSION}.linux-${ARCH}.tar.gz"
 ```
 ```bash
sudo chown -R root:root ./go
 ```
 ```bash
sudo mv -v go /usr/local
 ```
 ```bash
export GOPATH=$HOME/go
 ```
 ```bash
export PATH=$PATH:/usr/local/go/bin:$GOPATH/bin
 ```
 ```bash
source ~/.bash_profile
 ```
 ```bash
go version
 ```

2. Set and Run Script

 ```bash
git clone https://github.com/0xmoei/starknet.go
 ```
 ```bash
cd $HOME && cd starknet.go/examples/simpleInvoke
 ```
 ```bash
nano .env.testnet
 ```
Set INTEGRATION_BASE to your Lava Starknet Goerli testnet url 

 `CTRL + X`
 `Y + Enter`
 
 ```bash
nano main.go
 ```
Replace your Address + Public Key + Private Key

 `CTRL + X`
 `Y + Enter`
 
  ```bash
go mod tidy
 ```
 
  ```bash
go run main.go
 ```

You can Run with 'screen' if you want to close the terminal and run the script in the background
