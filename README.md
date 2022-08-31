# A toy project when applying to PI

### Quick start
```
make example
```
### Test
```
make test
```
### Run with docker-compose
First, build image
```
docker build -t pitest .
```
Then, run it
```
docker-compose up
```
In terms of parameters, change them from `env.list`