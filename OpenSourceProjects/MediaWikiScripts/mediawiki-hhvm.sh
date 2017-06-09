/usr/share/hhvm/install_fastcgi.sh
rm -rf /usr/share/nginx/html/*
mv mediawiki/* /usr/share/nginx/html/
cp -f mediawiki-hhvm default
cp -f ./default /etc/nginx/sites-available/
service nginx restart
service hhvm restart
