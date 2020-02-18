#!/bin/bash

#-----------------------------
#    Name: sheep.sh
#  Author: YHY
# Version: 1.0
#   Usage: install sheep v1
#Platform: arm64

#--------------------Declare---------------------
gWorkPath="/opt/sheep"
gSoftPath="/opt/sheepsoftware"
gPkgName="sheep.tar.gz"
gNodeList="127.0.0.1"
gUserName="root"
gLocalUser=$(whoami)
gVersion=0

#-----------------function---------------
fHelp()
{
	echo " UserAges:"
	echo "            ./sheep.sh help                                 :帮助信息"
	echo "            ./sheep.sh package                              :打包安装包"
	echo "            ./sheep.sh installubuntupg                      :在线安装postgresql 默认 数据库sheep 用户root 密码password"
	echo "            ./sheep.sh remoteinstall 127.0.0.1 root         :远程安装服务"
	echo " Examples:"
	echo "            ./sheep.sh help"
}

fShowStep ()
{
	lColumns=59
	for ((i=0; i<${lColumns}; i++)); do
		printf '.'
	done
	printf "[ "
	if [[ $1 -eq 0 ]]; then
		printf "\033[1;34m%b\033[0m" " ok "
	else
		printf "\033[1;34m%b\033[0m" "fail\a"
	fi
	printf " ]"
	printf "\r${2}: \n"
}

fBuildGO()
{
	git pull	
	export GOROOT=/home/Y/install/go11/go
	PROJECT_NAME=sheep
	PROJECT_MAIN=*.go
	COMMIT_NUMBER=`git rev-list HEAD | wc -l | awk '{print $1}'`
	gVersion=$COMMIT_NUMBER
	
	export CGO_ENABLED=0 
	export GOOS=linux  
	export GO111MODULE=on
	export GOPROXY=https://goproxy.io
	go build -o $PROJECT_NAME $PROJECT_MAIN
}


fBuildUI()
{
	rm -rvf dist
	cd ../web/sheep
#	git reset --hard
	git pull
	npm run build

	mv -f dist ../../server
	cd ../../server
}

fPackage()
{
	timenow=$(date "+%F")
	
	#package
	tar -zcvf sheep.tar.gz conf dist sheep sheep.sh
	tar -zcvf sheep-$timenow.$gVersion.tar.gz sheep.tar.gz sheep.sh
	rm -rvf sheep.tar.gz
	
	#only for yhy self
	rm -rvf software/*
	mv *.tar.gz software/
	cd software/
	tar -zxvf *.tar.gz
	tar -zxvf sheep.tar.gz
}

fDistributeSshPubKey ()
{
	while true
	do
		printf "\033[1;34m%b\033[0m" "----- Distribute ssh public key -----\n"
		printf "\033[1;34m%b\033[0m" "do it <y>, skip <n> : "
		read lChoice
		case ${lChoice} in
			[yY] )
				break
				;;
			[nN] )
				return
				;;
			* )
				printf "\033[1;31m%b\033[0m" "please input [y|n]\n"
				continue
		esac
	done

	sed -i "/^.*${gNodeList}.*/d"  ~/.ssh/known_hosts
	ssh-keygen -t rsa
	chmod 755 ~/.ssh

	nodes=($gNodeList)
	for node in ${nodes[@]}
	do
		ssh-copy-id $gUserName@$node
		echo "scp from ${node} finished..."
	done
}

fDistributePkg()
{
	nodes=(${gNodeList})
	for node in ${nodes[@]}; do
		# copy soft
		isExistSoftPath=`ssh $gUserName@$node "if [ ! -d $gSoftPath ];then echo 0;else echo 1;fi"`
		if [ $isExistSoftPath -eq '0' ]; then
			ssh $gUserName@$node "mkdir -p ${gSoftPath};chmod -R 777 ${gSoftPath}"
		fi

		if [ "$1" = "remoteinstall" ];then
			ssh $gUserName@$node "rm -rvf ${gSoftPath}/${gPkgName}"
			scp $gPkgName $gUserName@${node}:${gSoftPath}
		fi

		fShowStep $? "copy package to ${node}"
	done
}

fDistributeInstall ()
{
	nodes=(${gNodeList})
	for node in ${nodes[@]}; do
		# 解压
		ssh $gUserName@$node "killall sheep;\
							  mkdir -p ${gWorkPath};\
							  chmod -R 777 $gWorkPath;\
							  tar -zxvf "${gSoftPath}/${gPkgName}" -C ${gWorkPath} --warning=no-timestamp;\
							  rm -rvf ${gSoftPath}"

		# install local
		ssh $gUserName@$node "chmod +x ${gWorkPath}/sheep.sh;${gWorkPath}/sheep.sh install"

		printf "\033[1;34m%b\033[0m" "install on Node:${node}:\n"
		
		fShowStep $? "install packages on ${node} finished"
	done
}


fInstall()
{
	#add x
	chmod +x ${gWorkPath}/sheep.sh
	chmod +x ${gWorkPath}/sheep

	#set config
#	sudo cp -rvf $gWorkPath/conf/env/sheep.service /etc/systemd/system

	#set localtime
#	sudo rm -rvf /etc/localtime
#	sudo ln -s /usr/share/zoneinfo/Asia/Shanghai  /etc/localtime

    # start
#	sudo systemctl daemon-reload
#	sudo systemctl enable sheep
#	sudo systemctl start sheep
	killall /opt/sheep/sheep
	/opt/sheep/sheep >/dev/null &
}

fInstallPGForUbuntuOnline()
{
	pgversion=`psql --version`
	if [ "${pgversion}" = "psql (PostgreSQL) 9.2.24" ]; then
		printf "psql (PostgreSQL) 9.2.24\n"
	else
		sudo apt-get clean
		sudo apt-get update
		sudo apt-get install postgresql -y --allow-unauthenticated

		sudo su - postgres -c "psql << EOF
CREATE USER root WITH SUPERUSER CREATEDB CREATEROLE PASSWORD 'password';
CREATE DATABASE sheep OWNER root;
EOF"
	fi	
}

fInstallPGForCentOSOnline()
{
	pgversion=`/usr/pgsql-11/bin/psql --version`
	if [ "${pgversion}" = "psql (PostgreSQL) 11.7" ]; then
		printf "psql (PostgreSQL) 11.7\n"
	else
		sudo yum install postgresql11-server -y
		sudo /usr/pgsql-11/bin/postgresql-11-setup initdb
		sudo systemctl start postgresql-11
		sudo systemctl enable postgresql-11
		sudo firewall-cmd --permanent --add-port=5432/tcp  
		sudo firewall-cmd --permanent --add-port=9001/tcp  
		sudo firewall-cmd --reload 
		sudo su - postgres -c "psql << EOF
CREATE USER root WITH SUPERUSER CREATEDB CREATEROLE PASSWORD 'password';
CREATE DATABASE sheep OWNER root;
EOF"
	fi	
}

fMain()
{
	if [ $# -eq 2 ];then
		gNodeList=$2
	fi

	if [ $# -eq 3 ];then
		gNodeList=$2
		gUserName=$3
	fi
	

	if [ "$1" = "help" ];then
		fHelp

	elif [ "$1" = "package" ];then
		fBuildGO

		fBuildUI

		fPackage

	elif [ "$1" = "install" ];then
		fInstall

	elif [ "$1" = "installcentospg" ];then
		fInstallPGForCentOSOnline

	elif [ "$1" = "installubuntupg" ];then
		fInstallPGForUbuntuOnline

	elif [ "$1" = "remoteinstall" ];then
		# config ssh remote with no password
		fDistributeSshPubKey
	
		fDistributePkg $1
	
		fDistributeInstall
	else
	
		fHelp
		
	fi

	printf "\033[1;34m%b\033[0m" "$1 finished\n"
}


fMain $*
