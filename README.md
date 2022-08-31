# PI.EXCHANGE application toy project
### Brief introduction
This is a toy project done when applying to PI.EXCHANGE. To be concise, this project is used to generate emails from the predefined templates with a given list of customers.

Please notice these information:
- The input and output files are saved in `asset`
- The parameters are defined in `env.list` which is then exported to be ENV variables for the program to use
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