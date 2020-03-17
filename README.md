# Create container
docker build -t gogo .

# Run container 

docker run -it --name gogo -e TOKEN=${TOKEN} --net=host -p 3000:3000 gogo –rm

docker exec -it gogo sh
