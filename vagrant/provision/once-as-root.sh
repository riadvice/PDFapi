#!/usr/bin/env bash

source /app/vagrant/provision/common.sh

#== Import script args ==

timezone=$(echo "$1")

#== Provision script ==

info "Provision-script user: $(whoami)"

export DEBIAN_FRONTEND=noninteractive

info "Configure timezone"
timedatectl set-timezone ${timezone} --no-ask-password

info "Update OS software"
sudo apt-get update
sudo apt-get upgrade -y

info "Install ubuntu tools"
sudo apt-get install -y wget gnupg2 lsb-release curl zip unzip nginx-full bc ntp xmlstarlet bash-completion

info "Install and configure golang"
sudo wget https://dl.google.com/go/go1.17.1.linux-amd64.tar.gz
sudo tar -xvf go1.17.1.linux-amd64.tar.gz
sudo mv go /usr/local

info "Install Pythong3 and package"
sudo apt install python3-cairosvg

info "Configure NGINX"
sudo sed -i 's/user www-data/user vagrant/g' /etc/nginx/nginx.conf
echo "Done!"

info "Enabling site configuration"
sudo ln -s /app/vagrant/dev/nginx/app.conf /etc/nginx/sites-enabled/app.conf
echo "Done!"
