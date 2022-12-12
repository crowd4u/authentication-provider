cd /usr/local/apache2/test
composer install

while true
do
  sleep 3
  echo "try startup apache."
  if [[ -f  /var/run/httpd/httpd.pid ]]; then
    kill `cat /var/run/httpd/httpd.pid` 2> /dev/null
    sleep 1
    rm /var/run/httpd/httpd.pid
    sleep 1
  fi
  httpd -DFOREGROUND
done
