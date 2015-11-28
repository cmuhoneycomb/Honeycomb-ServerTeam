package main

import (
	"flag"
	"fmt"
	"github.com/gocql/gocql"
	"log"
	"net"
	"net/http"
	"os/exec"
	"strconv"
	"strings"
)

type CassandraClient struct {
	session *gocql.Session
}

const (
	defaultHostPort = "32768"
	//CassandraHostPort = "128.2.7.38"
	CassandraHostPort = "127.0.0.1"
)

var (
	port   = flag.String("port", defaultHostPort, "port number to listen on")
	client *CassandraClient
)

func init() {
	log.SetFlags(log.Lshortfile | log.Lmicroseconds)
}

func runSparkJob(w http.ResponseWriter, r *http.Request) {
	jobId := r.URL.Query().Get("id")
	id, err := strconv.Atoi(jobId)
	if err != nil {
		w.Write([]byte("Wrong id!\n"))
		return
	}
	// test database insert
	client.Insert(id, "/home/a/b/trainingdata;/home/testingdata")
	// test database lookup
	storedPaths, err := client.Get(id)
	if err != nil {
		w.Write([]byte("Job id does not exist!\n"))
		return
	}
	paths := strings.Split(storedPaths, ";") // split by ";"
	if len(paths) != 2 {                     // check paths
		log.Println("Paths in database is worng: " + storedPaths)
		w.Write([]byte("Paths in database is worng: " + storedPaths))
		return
	}
	w.Write([]byte("Job ID:" + jobId + "\n"))
	w.Write([]byte("Training Data: " + paths[0] + "\n"))
	w.Write([]byte("Testing Data:" + paths[1] + "\n"))

	// call external python script
	go func(id int) {
		output, err := exec.Command("python", "test.py").Output()
		if err != nil {
			log.Fatal(err) // caution: log.Fatal may terminate the program
		} else {
			log.Println(string(output))
			client.Insert(id, string(output))
		}
	}(id)
}

func NewCassandraClient(CassandraHostPort string) (*CassandraClient, error) {
	cluster := gocql.NewCluster(CassandraHostPort)
	cluster.Keyspace = "honey"
	cluster.Consistency = gocql.Quorum
	session, err := cluster.CreateSession()
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	c := &CassandraClient{
		session: session,
	}
	return c, nil
}

func (c *CassandraClient) Close() {
	c.session.Close()
}

func (c *CassandraClient) Get(id int) (string, error) {
	var path string

	if err := c.session.Query(`SELECT id, path FROM data WHERE id = ?`,
		id).Consistency(gocql.One).Scan(&id, &path); err != nil {
		log.Println(err)
		return "", err
	}
	fmt.Println("Result:", id, path)
	return path, nil
}

func (c *CassandraClient) Insert(id int, path string) {
	if err := c.session.Query(`INSERT INTO result (id, path) VALUES (?, ?)`,
		id, path).Exec(); err != nil {
		log.Println(err)
	}
}

func main() {
	flag.Parse()
	if *port == "" {
		// If port string is empty, then use the default host port
		*port = defaultHostPort
	}
	hostPort := net.JoinHostPort("", *port)
	// sample request: curl localhost:32768/?id=2
	http.HandleFunc("/", runSparkJob)

	var err error
	client, err = NewCassandraClient(CassandraHostPort)
	if err != nil {
		log.Fatal("failed to create Cassandra Client")
		return
	}
	defer client.Close()

	log.Fatal(http.ListenAndServe(hostPort, nil))
}
