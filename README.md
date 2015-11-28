# Honeycomb-ServerTeam
## Team member:
  	Ningna Wang
  	Zhoucheng Li
  	Handi Xu
  	Long He
  	
  	
## Usage Instruction:
#### On BIC Cluster
```
CREATE KEYSPACE Honey
  WITH REPLICATION = { 'class' : 'SimpleStrategy', 'replication_factor' : 1 };

CREATE TABLE honey.data (
  id int PRIMARY KEY,
  path varchar
);

CREATE TABLE honey.result (
  id int PRIMARY KEY,
  path varchar
);
```

###### Go Version
```
export GOROOT=/home/honeycomb/Go/go
export GOPATH=/home/honeycomb/HoneyBuzzard/src/honeycomb
go build srunner
./srunner &
```

###### On terminal:
- javac HoneycombServer.java
- java HoneycombServer

###### open new terminal
- javac TestClient.java
- java TestClient

###### type in your input on terminal running TestClient
###### terminal that running HoneycombServer will response 
- "Honeycomb Server has received your input!"
   
