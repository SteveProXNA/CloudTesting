##### README.md
###### 31/08/2021
```
01. Setup
go mod init dockercheatsheet
touch app.go
go mod tidy
```
```
02. Docker
Ctrl + Shift + P | Add Docker Files to Workspace
Go | 5000 | No
Right click Dockerfile | Build Image... | 02example:latest
Right click 02example:latest | Run Interactive
curl http://localhost:8081/test
```
```
03. Testing
Launch browswer | http://localhost:5000
```
```
04. Cleanup
docker stop $(docker ps -q)
docker rmi $(docker images -qa) --force
```
