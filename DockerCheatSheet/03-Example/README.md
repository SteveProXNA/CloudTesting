##### README.md
###### 31/08/2021
```
01. Setup
touch main.cpp
touch Dockerfile
```
```
02. Run normally
docker build . -t cpp_test:1.0
docker run --rm cpp_test:1.0
```
```
03. Run interactively
docker run -it cpp_test:1.0 bash
./Test
```
```
04. Cleanup
docker rmi $(docker images -qa) --force
```
