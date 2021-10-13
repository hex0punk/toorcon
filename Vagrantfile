# -*- mode: ruby -*-
# vi: set ft=ruby :

Vagrant.configure("2") do |config|
  config.vm.network "forwarded_port", guest: 1984, host: 8080, protocol: "tcp"
  # Every Vagrant development environment requires a box. You can search for
  # boxes at https://vagrantcloud.com/search.
  config.vm.box = "hashicorp/bionic64"
  #config.vm.network "public_network", bridge: "en0: Wi-Fi (AirPort)"
  config.vm.provision :shell, :privileged => false, inline: <<-SHELL
	# update
	sudo apt-get -y update
	sudo apt-get -y upgrade
	# deps
	sudo apt-get -y install git build-essential
	sudo apt-get -y install python2.7 python-pip
	# Go
	mkdir -p go/src go/bin go/pkg
	curl -O https://dl.google.com/go/go1.17.2.linux-amd64.tar.gz
	sudo tar -C /usr/local -xzf go1.17.2.linux-amd64.tar.gz
	export GOPATH=$HOME/go
	export PATH=$PATH:/usr/local/go/bin:$HOME/go/bin
	echo "export GOPATH=$HOME/go"  >> ~/.bashrcc
	echo "export PATH=$PATH:/usr/local/go/bin:$HOME/go/bin" >> ~/.bashrc
	# Go-Fuzz
	go install github.com/dvyukov/go-fuzz/go-fuzz@latest
	go install github.com/dvyukov/go-fuzz/go-fuzz-build@latest
  	# GCatch
  	cd /home/vagrant/go/src
	mkdir github.com
	mkdir github.com/system-pclub
	cd github.com/system-pclub
	git clone https://github.com/system-pclub/GCatch.git
	cd GCatch/GCatch
	./installZ3.sh
	./install.sh
	
  SHELL
  
  config.vm.provider "virtualbox" do |vb|
      vb.memory = "10024"
      vb.cpus = 4
  end

end
