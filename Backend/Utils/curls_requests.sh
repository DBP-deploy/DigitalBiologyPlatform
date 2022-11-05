#Login
curl 'http://localhost:8080/user/login' -X POST -H 'User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:106.0) Gecko/20100101 Firefox/106.0' -H 'Accept: text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,*/*;q=0.8' -H 'Accept-Language: en-US,en;q=0.5' -H 'Accept-Encoding: gzip, deflate, br' -H 'Connection: keep-alive' -H 'Upgrade-Insecure-Requests: 1' -H 'Sec-Fetch-Dest: document' -H 'Sec-Fetch-Mode: navigate' -H 'Sec-Fetch-Site: cross-site' -H 'Content-Type: application/json' -H 'Origin: null' -H 'Pragma: no-cache' -H 'Cache-Control: no-cache' --data-raw '{"username" : "usr", "password" : "pw"}'

#Create user
curl 'http://localhost:8080/user/' -X PUT -H 'User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:106.0) Gecko/20100101 Firefox/106.0' -H 'Accept: text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,*/*;q=0.8' -H 'Accept-Language: en-US,en;q=0.5' -H 'Accept-Encoding: gzip, deflate, br' -H 'Connection: keep-alive' -H 'Upgrade-Insecure-Requests: 1' -H 'Sec-Fetch-Dest: document' -H 'Sec-Fetch-Mode: navigate' -H 'Sec-Fetch-Site: cross-site' -H 'Sec-Fetch-User: ?1' -H 'Content-Type: application/json' -H 'Origin: null' -H 'Pragma: no-cache' -H 'Cache-Control: no-cache' --data-raw '{"username" : "usr", "password" : "pw", "email": "ofi@a.c"}'

#Get bearer infos authenticated
curl 'http://localhost:8080/user/me' -H 'User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:106.0) Gecko/20100101 Firefox/106.0' -H 'Accept: text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,*/*;q=0.8' -H 'Accept-Language: en-US,en;q=0.5' -H 'Accept-Encoding: gzip, deflate, br' -H 'Connection: keep-alive' -H 'Upgrade-Insecure-Requests: 1' -H 'Sec-Fetch-Dest: document' -H 'Sec-Fetch-Mode: navigate' -H 'Sec-Fetch-Site: cross-site' -H 'Pragma: no-cache' -H 'Cache-Control: no-cache' -H 'Content-Type: application/json' -H 'Authorization: Bearer eyJ1IjogInVzciIsICJ0IiA6ICI0MWE4ZmUyYThjODU1NWJjYzJlYmQ3ZWYzYWRkZDkzZTQ5OWEzMWYwMDNlODI3MDIyYzc2ODg3ZDhkZmNjNjY1IiB9'

#Get protocols list
curl 'http://localhost:8080/protocol/me' -H 'User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:106.0) Gecko/20100101 Firefox/106.0' -H 'Accept: text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,*/*;q=0.8' -H 'Accept-Language: en-US,en;q=0.5' -H 'Accept-Encoding: gzip, deflate, br' -H 'Connection: keep-alive' -H 'Upgrade-Insecure-Requests: 1' -H 'Sec-Fetch-Dest: document' -H 'Sec-Fetch-Mode: navigate' -H 'Sec-Fetch-Site: cross-site' -H 'Content-Type: application/json' -H 'Authorization: Bearer eyJ1IjogInRlc3R1c2VyIiwgInQiIDogImU1YjdkYjQ5ZDkyOGYzYjQyMTQyY2U4MTZlMzA1Y2Y0MmFjODA0ZmI3ZjhjNTVlNmUxYzZjOTNkYWI2Nzg4NDUiIH0=' -H 'Pragma: no-cache' -H 'Cache-Control: no-cache'