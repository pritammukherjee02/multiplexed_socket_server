#!/bin/bash
curl -OL https://go.dev/dl/go1.21.0.linux-amd64.tar.gz

rm -rf /usr/local/go && tar -C /usr/local -xzf go1.21.0.linux-amd64.tar.gz

echo export PATH=$PATH:/usr/local/go/bin > ~/.profile

source ~/.profile

go version
