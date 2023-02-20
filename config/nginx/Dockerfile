FROM nginx
RUN rm /etc/nginx/conf.d/default.conf
COPY nginx.conf /etc/nginx/conf.d/default.conf
# should have index.html as bad page routing
COPY index.html /usr/share/nginx/html/index.html