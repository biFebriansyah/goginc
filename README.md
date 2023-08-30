## Deployment

1. cd to nginx folder
   cd /etc/nginx/sites-available/

2. create file
   sudo vi fwg.fazztrack.com

3. link file
   sudo ln -s /etc/nginx/sites-available/fwg.fazztrack.com /etc/nginx/sites-enabled/

4. test file config
   sudo nginx -t

5. reload nginx service
   sudo nginx -s reload

6. create certificate ssl
   sudo certbot --nginx
