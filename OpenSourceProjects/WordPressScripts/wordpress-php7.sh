/usr/share/hhvm/uninstall_fastcgi.sh
rm -rf /usr/share/nginx/html/*
mv wordpress/* /usr/share/nginx/html/
cp -f wordpress-php7 default
cp -f ./default /etc/nginx/sites-available/
service nginx restart
service php7.0-fpm restart
