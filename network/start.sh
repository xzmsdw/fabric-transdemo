#!/bin/bash
export PATH=${PWD}/bin:$PATH
#生成证书文件
cryptogen generate --config=crypto-config.yaml
./ccp-generate.sh
#区块生成
configtxgen  -profile TwoOrgsOrdererGenesis -outputBlock ./channel-artifacts/genesis.block -channelID fabric-channel
#生成通道文件
configtxgen -profile TwoOrgsChannel -outputCreateChannelTx ./channel-artifacts/channel.tx  -channelID mychannel
#生成组织锚节点
configtxgen -outputAnchorPeersUpdate ./channel-artifacts/Org1MSPanchors.tx -channelID mychannel -profile TwoOrgsChannel -asOrg Org1MSP
configtxgen -outputAnchorPeersUpdate ./channel-artifacts/Org2MSPanchors.tx -channelID mychannel -profile TwoOrgsChannel -asOrg Org2MSP
#节点创建
docker-compose up -d

sleep 10
#创建通道
docker exec -it cli1 bash -c "peer channel create -o orderer.example.com:7050 -c mychannel -f ./channel-artifacts/channel.tx --tls true --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/example.com/msp/tlscacerts/tlsca.example.com-cert.pem"
#复制通道文件
docker cp cli1:/opt/gopath/src/github.com/hyperledger/fabric/peer/mychannel.block ./
docker cp ./mychannel.block cli2:/opt/gopath/src/github.com/hyperledger/fabric/peer
#加入通道
docker exec -it cli2 bash -c "peer channel join -b mychannel.block"
docker exec -it cli1 bash -c "peer channel join -b mychannel.block"
#更新锚节点
docker exec -it cli1 bash -c "peer channel update -o orderer.example.com:7050 -c mychannel -f ./channel-artifacts/Org1MSPanchors.tx --tls --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem"
docker exec -it cli2 bash -c "peer channel update -o orderer.example.com:7050 -c mychannel -f ./channel-artifacts/Org2MSPanchors.tx --tls --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem"

sudo cp -rf /home/iris/newnode/tmpcodes/farmer/ ~/newnode/chaincode/go/farmer
sudo cp -rf /home/iris/newnode/tmpcodes/driver/ ~/newnode/chaincode/go/driver
sudo cp -rf /home/iris/newnode/tmpcodes/material/ ~/newnode/chaincode/go/material
sudo cp -rf /home/iris/newnode/tmpcodes/sell/ ~/newnode/chaincode/go/sell
sudo cp -rf /home/iris/newnode/tmpcodes/seller/ ~/newnode/chaincode/go/seller
sudo cp -rf /home/iris/newnode/tmpcodes/capital/ ~/newnode/chaincode/go/capital
sudo cp -rf /home/iris/newnode/tmpcodes/users/ ~/newnode/chaincode/go/users
#打包链码
docker exec -it cli1 bash -c "cd /opt/gopath/src/github.com/hyperledger/fabric-cluster/chaincode/go/farmer;go env -w GOPROXY=https://goproxy.cn,direct;go env -w GO111MODULE=auto;go mod init farmer;go mod tidy;go mod vendor;cd /opt/gopath/src/github.com/hyperledger/fabric/peer;peer lifecycle chaincode package farmer.tar.gz --path github.com/hyperledger/fabric-cluster/chaincode/go/farmer --label farmer_1.0"
docker exec -it cli1 bash -c "cd /opt/gopath/src/github.com/hyperledger/fabric-cluster/chaincode/go/driver;go env -w GOPROXY=https://goproxy.cn,direct;go env -w GO111MODULE=auto;go mod init driver;go mod tidy;go mod vendor;cd /opt/gopath/src/github.com/hyperledger/fabric/peer;peer lifecycle chaincode package driver.tar.gz --path github.com/hyperledger/fabric-cluster/chaincode/go/farmer --label driver_1.0"
docker exec -it cli1 bash -c "cd /opt/gopath/src/github.com/hyperledger/fabric-cluster/chaincode/go/material;go env -w GOPROXY=https://goproxy.cn,direct;go env -w GO111MODULE=auto;go mod init material;go mod tidy;go mod vendor;cd /opt/gopath/src/github.com/hyperledger/fabric/peer;peer lifecycle chaincode package material.tar.gz --path github.com/hyperledger/fabric-cluster/chaincode/go/farmer --label material_1.0"
docker exec -it cli1 bash -c "cd /opt/gopath/src/github.com/hyperledger/fabric-cluster/chaincode/go/sell;go env -w GOPROXY=https://goproxy.cn,direct;go env -w GO111MODULE=auto;go mod init sell;go mod tidy;go mod vendor;cd /opt/gopath/src/github.com/hyperledger/fabric/peer;peer lifecycle chaincode package sell.tar.gz --path github.com/hyperledger/fabric-cluster/chaincode/go/farmer --label sell_1.0"
docker exec -it cli1 bash -c "cd /opt/gopath/src/github.com/hyperledger/fabric-cluster/chaincode/go/seller;go env -w GOPROXY=https://goproxy.cn,direct;go env -w GO111MODULE=auto;go mod init seller;go mod tidy;go mod vendor;cd /opt/gopath/src/github.com/hyperledger/fabric/peer;peer lifecycle chaincode package seller.tar.gz --path github.com/hyperledger/fabric-cluster/chaincode/go/farmer --label seller_1.0"
docker exec -it cli1 bash -c "cd /opt/gopath/src/github.com/hyperledger/fabric-cluster/chaincode/go/capital;go env -w GOPROXY=https://goproxy.cn,direct;go env -w GO111MODULE=auto;go mod init capital;go mod tidy;go mod vendor;cd /opt/gopath/src/github.com/hyperledger/fabric/peer;peer lifecycle chaincode package capital.tar.gz --path github.com/hyperledger/fabric-cluster/chaincode/go/farmer --label capital_1.0"
docker exec -it cli1 bash -c "cd /opt/gopath/src/github.com/hyperledger/fabric-cluster/chaincode/go/users;go env -w GOPROXY=https://goproxy.cn,direct;go env -w GO111MODULE=auto;go mod init users;go mod tidy;go mod vendor;cd /opt/gopath/src/github.com/hyperledger/fabric/peer;peer lifecycle chaincode package users.tar.gz --path github.com/hyperledger/fabric-cluster/chaincode/go/farmer --label users_1.0"
#复制到本地,组织二
docker cp cli1:/opt/gopath/src/github.com/hyperledger/fabric/peer/farmer.tar.gz ./
docker cp farmer.tar.gz  cli2:/opt/gopath/src/github.com/hyperledger/fabric/peer
docker cp cli1:/opt/gopath/src/github.com/hyperledger/fabric/peer/driver.tar.gz ./
docker cp driver.tar.gz  cli2:/opt/gopath/src/github.com/hyperledger/fabric/peer
docker cp cli1:/opt/gopath/src/github.com/hyperledger/fabric/peer/material.tar.gz ./
docker cp material.tar.gz  cli2:/opt/gopath/src/github.com/hyperledger/fabric/peer
docker cp cli1:/opt/gopath/src/github.com/hyperledger/fabric/peer/sell.tar.gz ./
docker cp sell.tar.gz  cli2:/opt/gopath/src/github.com/hyperledger/fabric/peer
docker cp cli1:/opt/gopath/src/github.com/hyperledger/fabric/peer/seller.tar.gz ./
docker cp seller.tar.gz  cli2:/opt/gopath/src/github.com/hyperledger/fabric/peer
docker cp cli1:/opt/gopath/src/github.com/hyperledger/fabric/peer/capital.tar.gz ./
docker cp capital.tar.gz  cli2:/opt/gopath/src/github.com/hyperledger/fabric/peer
docker cp cli1:/opt/gopath/src/github.com/hyperledger/fabric/peer/users.tar.gz ./
docker cp users.tar.gz  cli2:/opt/gopath/src/github.com/hyperledger/fabric/peer
#安装链码
docker exec -it cli1 bash -c "peer lifecycle chaincode install farmer.tar.gz"
docker exec -it cli2 bash -c "peer lifecycle chaincode install farmer.tar.gz"
#查询获取packageID
docker exec -it cli1 bash -c "peer lifecycle chaincode queryinstalled " >&log.txt
cat log.txt
PACKAGE_ID=$(sed -n "/farmer_1.0/{s/^Package ID: //; s/, Label:.*$//; p;}" log.txt)
#组织批准
docker exec -it cli1 bash -c "peer lifecycle chaincode approveformyorg --channelID mychannel --name farmer --version 1.0 --package-id $PACKAGE_ID --sequence 1 --tls true --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem"
docker exec -it cli2 bash -c "peer lifecycle chaincode approveformyorg --channelID mychannel --name farmer --version 1.0 --package-id $PACKAGE_ID --sequence 1 --tls true --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem"
#链码提交
docker exec -it cli1 bash -c "peer lifecycle chaincode commit -o orderer.example.com:7050 --channelID mychannel --name farmer --version  1.0 --sequence 1 --tls true --cafile  /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem  --peerAddresses peer0.org1.example.com:7051  --tlsRootCertFiles   /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt  --peerAddresses peer0.org2.example.com:9051  --tlsRootCertFiles   /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org2.example.com/peers/peer0.org2.example.com/tls/ca.crt"
#安装链码
docker exec -it cli1 bash -c "peer lifecycle chaincode install driver.tar.gz"
docker exec -it cli2 bash -c "peer lifecycle chaincode install driver.tar.gz"
#查询获取packageID
docker exec -it cli1 bash -c "peer lifecycle chaincode queryinstalled " >&log.txt
cat log.txt
PACKAGE_ID=$(sed -n "/driver_1.0/{s/^Package ID: //; s/, Label:.*$//; p;}" log.txt)
#组织批准
docker exec -it cli1 bash -c "peer lifecycle chaincode approveformyorg --channelID mychannel --name driver --version 1.0 --package-id $PACKAGE_ID --sequence 1 --tls true --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem"
docker exec -it cli2 bash -c "peer lifecycle chaincode approveformyorg --channelID mychannel --name driver --version 1.0 --package-id $PACKAGE_ID --sequence 1 --tls true --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem"
#链码提交
docker exec -it cli1 bash -c "peer lifecycle chaincode commit -o orderer.example.com:7050 --channelID mychannel --name driver --version  1.0 --sequence 1 --tls true --cafile  /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem  --peerAddresses peer0.org1.example.com:7051  --tlsRootCertFiles   /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt  --peerAddresses peer0.org2.example.com:9051  --tlsRootCertFiles   /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org2.example.com/peers/peer0.org2.example.com/tls/ca.crt"
#安装链码
docker exec -it cli1 bash -c "peer lifecycle chaincode install material.tar.gz"
docker exec -it cli2 bash -c "peer lifecycle chaincode install material.tar.gz"
#查询获取packageID
docker exec -it cli1 bash -c "peer lifecycle chaincode queryinstalled " >&log.txt
cat log.txt
PACKAGE_ID=$(sed -n "/material_1.0/{s/^Package ID: //; s/, Label:.*$//; p;}" log.txt)
#组织批准
docker exec -it cli1 bash -c "peer lifecycle chaincode approveformyorg --channelID mychannel --name material --version 1.0 --package-id $PACKAGE_ID --sequence 1 --tls true --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem"
docker exec -it cli2 bash -c "peer lifecycle chaincode approveformyorg --channelID mychannel --name material --version 1.0 --package-id $PACKAGE_ID --sequence 1 --tls true --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem"
#链码提交
docker exec -it cli1 bash -c "peer lifecycle chaincode commit -o orderer.example.com:7050 --channelID mychannel --name material --version  1.0 --sequence 1 --tls true --cafile  /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem  --peerAddresses peer0.org1.example.com:7051  --tlsRootCertFiles   /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt  --peerAddresses peer0.org2.example.com:9051  --tlsRootCertFiles   /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org2.example.com/peers/peer0.org2.example.com/tls/ca.crt"
#安装链码
docker exec -it cli1 bash -c "peer lifecycle chaincode install sell.tar.gz"
docker exec -it cli2 bash -c "peer lifecycle chaincode install sell.tar.gz"
#查询获取packageID
docker exec -it cli1 bash -c "peer lifecycle chaincode queryinstalled " >&log.txt
cat log.txt
PACKAGE_ID=$(sed -n "/sell_1.0/{s/^Package ID: //; s/, Label:.*$//; p;}" log.txt)
#组织批准
docker exec -it cli1 bash -c "peer lifecycle chaincode approveformyorg --channelID mychannel --name sell --version 1.0 --package-id $PACKAGE_ID --sequence 1 --tls true --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem"
docker exec -it cli2 bash -c "peer lifecycle chaincode approveformyorg --channelID mychannel --name sell --version 1.0 --package-id $PACKAGE_ID --sequence 1 --tls true --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem"
#链码提交
docker exec -it cli1 bash -c "peer lifecycle chaincode commit -o orderer.example.com:7050 --channelID mychannel --name sell --version  1.0 --sequence 1 --tls true --cafile  /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem  --peerAddresses peer0.org1.example.com:7051  --tlsRootCertFiles   /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt  --peerAddresses peer0.org2.example.com:9051  --tlsRootCertFiles   /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org2.example.com/peers/peer0.org2.example.com/tls/ca.crt"
#安装链码
docker exec -it cli1 bash -c "peer lifecycle chaincode install seller.tar.gz"
docker exec -it cli2 bash -c "peer lifecycle chaincode install seller.tar.gz"
#查询获取packageID
docker exec -it cli1 bash -c "peer lifecycle chaincode queryinstalled " >&log.txt
cat log.txt
PACKAGE_ID=$(sed -n "/seller_1.0/{s/^Package ID: //; s/, Label:.*$//; p;}" log.txt)
#组织批准
docker exec -it cli1 bash -c "peer lifecycle chaincode approveformyorg --channelID mychannel --name seller --version 1.0 --package-id $PACKAGE_ID --sequence 1 --tls true --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem"
docker exec -it cli2 bash -c "peer lifecycle chaincode approveformyorg --channelID mychannel --name seller --version 1.0 --package-id $PACKAGE_ID --sequence 1 --tls true --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem"
#链码提交
docker exec -it cli1 bash -c "peer lifecycle chaincode commit -o orderer.example.com:7050 --channelID mychannel --name seller --version  1.0 --sequence 1 --tls true --cafile  /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem  --peerAddresses peer0.org1.example.com:7051  --tlsRootCertFiles   /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt  --peerAddresses peer0.org2.example.com:9051  --tlsRootCertFiles   /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org2.example.com/peers/peer0.org2.example.com/tls/ca.crt"
#安装链码
docker exec -it cli1 bash -c "peer lifecycle chaincode install capital.tar.gz"
docker exec -it cli2 bash -c "peer lifecycle chaincode install capital.tar.gz"
#查询获取packageID
docker exec -it cli1 bash -c "peer lifecycle chaincode queryinstalled " >&log.txt
cat log.txt
PACKAGE_ID=$(sed -n "/capital_1.0/{s/^Package ID: //; s/, Label:.*$//; p;}" log.txt)
#组织批准
docker exec -it cli1 bash -c "peer lifecycle chaincode approveformyorg --channelID mychannel --name capital --version 1.0 --package-id $PACKAGE_ID --sequence 1 --tls true --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem"
docker exec -it cli2 bash -c "peer lifecycle chaincode approveformyorg --channelID mychannel --name capital --version 1.0 --package-id $PACKAGE_ID --sequence 1 --tls true --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem"
#链码提交
docker exec -it cli1 bash -c "peer lifecycle chaincode commit -o orderer.example.com:7050 --channelID mychannel --name capital --version  1.0   --sequence 1 --tls true --cafile  /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem  --peerAddresses peer0.org1.example.com:7051  --tlsRootCertFiles   /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt  --peerAddresses peer0.org2.example.com:9051  --tlsRootCertFiles   /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org2.example.com/peers/peer0.org2.example.com/tls/ca.crt"
#安装链码
docker exec -it cli1 bash -c "peer lifecycle chaincode install users.tar.gz"
docker exec -it cli2 bash -c "peer lifecycle chaincode install users.tar.gz"
#查询获取packageID
docker exec -it cli1 bash -c "peer lifecycle chaincode queryinstalled " >&log.txt
cat log.txt
PACKAGE_ID=$(sed -n "/users_1.0/{s/^Package ID: //; s/, Label:.*$//; p;}" log.txt)
#组织批准
docker exec -it cli1 bash -c "peer lifecycle chaincode approveformyorg --channelID mychannel --name users --version 1.0 --package-id $PACKAGE_ID --sequence 1 --tls true --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem"
docker exec -it cli2 bash -c "peer lifecycle chaincode approveformyorg --channelID mychannel --name users --version 1.0 --package-id $PACKAGE_ID --sequence 1 --tls true --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem"
#链码提交
docker exec -it cli1 bash -c "peer lifecycle chaincode commit -o orderer.example.com:7050 --channelID mychannel --name users --version  1.0 --sequence 1 --tls true --cafile  /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem  --peerAddresses peer0.org1.example.com:7051  --tlsRootCertFiles   /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt  --peerAddresses peer0.org2.example.com:9051  --tlsRootCertFiles   /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org2.example.com/peers/peer0.org2.example.com/tls/ca.crt"

docker exec -it cli1 bash -c "peer chaincode invoke -o orderer.example.com:7050 --tls --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem -C mychannel -n farmer --peerAddresses peer0.org1.example.com:7051 --tlsRootCertFiles /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt --peerAddresses peer0.org2.example.com:9051 --tlsRootCertFiles /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org2.example.com/peers/peer0.org2.example.com/tls/ca.crt -c '{\"function\":\"RecordCropsGrow\",\"Args\":[\"1\",\"1\",\"1\",\"1\",\"1\",\"1\",\"1\",\"1\",\"1\"]}'"

docker exec -it cli1 bash -c "peer chaincode invoke -o orderer.example.com:7050 --tls --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem -C mychannel -n farmer --peerAddresses peer0.org1.example.com:7051 --tlsRootCertFiles /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt --peerAddresses peer0.org2.example.com:9051 --tlsRootCertFiles /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org2.example.com/peers/peer0.org2.example.com/tls/ca.crt -c '{\"function\":\"QueryCropsProcessByCropsId\",\"Args\":[\"1\"]}'"

