# Create container
docker build -t gogo-quotes .

# Run container 

docker run -it --name gogo-quotes -e TOKEN="${TOKEN}" -e PORT="${PORT}" --net=host -p 3000:3000 gogo-quotes

docker exec -it gogo-quotes sh
