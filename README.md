# Fabric CA Developer's Guide

This is the Developer's Guide for Fabric CA, which is a Certificate Authority for Hyperledger Fabric.

See [User's Guide for Fabric CA](https://hyperledger-fabric-ca.readthedocs.io) for information on how to use Fabric CA.

## Prerequisites

* Go 1.12.12 installation or later
* **GOPATH** environment variable is set correctly
* docker version 17.06 or later
* docker-compose version 1.14 or later
* A Linux Foundation ID  (see [create a Linux Foundation ID](https://identity.linuxfoundation.org/))


## 说明
  功能说明：
   *  支持国密算法
   *  支持生成国密算法x509证书
   *  支持国密算法https
    
## 编译说明
    1、编译二进制
        make native or make release
        
    2、编译镜像
        make docker
        
    3、除了改造支持国密算法，其它使用方法和正常社区版一致，参考社区文档即可。   
        
## 使用说明
    服务端
    $export FABRIC_CA_SERVER_HOME=xxx //根据实际路径
    $export  FABRIC_CA_SERVER_TLS_ENABLED=true //启用https(可不使用) 
    $fabric-ca-server start -b admin:adminpw
    
    客户端
    $export FABRIC_CA_CLIENT_HOME=xxx  //根据实际路径
    $export  FABRIC_CA_SERVER_TLS_ENABLED=true //启用tls(可不使用)    
    $export FABRIC_CA_CLIENT_TLS_CERTFILES=xxx //设置tls根证书路径
    $fabric-ca-client enroll -u https://admin:adminpw@localhost:7054
    
    使用镜像方式，参考社区文档即可。

## License
版权所有 前海联合网络科技有限公司

Copyright 2021- 
FORESEA UNITED NETWORK TECH. All Rights Reserved. Licensed under the Apache License, Version 2.0 (the "License");

you may not use this file except in compliance with the License. You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0 Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.

See the License for the specific language governing permissions and limitations under the License.