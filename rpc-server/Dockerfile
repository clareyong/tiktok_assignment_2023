FROM golang:1.18
WORKDIR /app
COPY . .
RUN sh ./build.sh
EXPOSE 8888
RUN chmod +x ./script/wait-for-it.sh
CMD ["./script/wait-for-it.sh",  "db:3306",  "--strict" , "--timeout=300" , "--" ,  "./output/bootstrap.sh"]
