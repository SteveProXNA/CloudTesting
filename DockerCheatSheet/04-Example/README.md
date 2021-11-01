##### README.md
###### 31/08/2021
```
01. Setup
touch envoy-demo.yaml
```
```
02. Docker
```
docker run --rm -it \
      -v $(pwd)/envoy-demo.yaml:/envoy-demo.yaml \
      -p 9901:9901 \
      -p 10000:10000 \
      envoyproxy/envoy-dev:1acf02f70c75a7723d0269b7f375b3a94cb0fbf0 \
          -c envoy-demo.yaml
```
03. Testing
Launch browswer | http://localhost:10000
```
```
04. Cleanup
docker stop $(docker ps -q)
docker rmi $(docker images -qa) --force
```
