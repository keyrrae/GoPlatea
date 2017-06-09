apt-get update
apt-get install -y unzip vim git-core curl wget build-essential
add-apt-repository -y ppa:nginx/mainline
apt-get install -y nginx

apt-get install -y software-properties-common

apt-key adv --recv-keys --keyserver hkp://keyserver.ubuntu.com:80 0x5a16e7281be7a449
add-apt-repository "deb http://dl.hhvm.com/ubuntu $(lsb_release -sc) main"
apt-get update
apt-get install -y hhvm

apt-get install -y mysql-server

apt-get install -y php7.0 php7.0-fpm php7.0-mysql php7.0-cli
apt-get install -y php5.6 php5.6-fpm php5.6-mysql php5.6-cli
