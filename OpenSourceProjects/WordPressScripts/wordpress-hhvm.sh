/usr/share/hhvm/install_fastcgi.sh
rm -rf /usr/share/nginx/html/*
mv wordpress/* /usr/share/nginx/html/
cp -f wordpress-hhvm default
cp -f ./default /etc/nginx/sites-available/
service nginx restart
service hhvm restart
