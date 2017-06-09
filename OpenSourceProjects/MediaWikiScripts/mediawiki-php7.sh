/usr/share/hhvm/uninstall_fastcgi.sh
rm -rf /usr/share/nginx/html/*
mv mediawiki/* /usr/share/nginx/html/
cp -f mediawiki-php7 default
cp -f ./default /etc/nginx/sites-available/
service nginx restart
service php7.0-fpm restart
