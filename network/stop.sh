#!/bin/bash
docker-compose down
docker volume prune
rm -rf channel-artifacts
rm -rf crypto-config
rm -rf mychannel.block
rm log.txt
sudo rm -rf chaincode
rm *.tar.gz
