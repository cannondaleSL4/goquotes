# Create container
docker build -t gogo .

# Run container 

docker run -it --name gogo -e TOKEN="${TOKEN}" PORT="${PORT}"  --net=host -p 3000:3000 gogo

docker exec -it gogo sh
