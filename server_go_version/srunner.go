package main

import (
	"flag"
	"fmt"
	"github.com/gocql/gocql"
	"log"
	"net"
	"net/http"
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
	client CassandraClient
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
	client.Insert(id, "/home/a/b/trainingdata;/home/testingdata")
	storedPaths := client.Get(id)
	paths := strings.Split(storedPaths, ";") // split by ";"
	if len(paths) != 2 {                     // check paths
		log.Fatal("Paths in database is worng: " + storedPaths)
		return
	}
	w.Write([]byte("Job ID:" + jobId + "\n"))
	w.Write([]byte("Training Data: " + paths[0] + "\n"))
	w.Write([]byte("Testing Data:" + paths[1] + "\n"))
}

func NewCassandraClient(CassandraHostPort string) CassandraClient {
	cluster := gocql.NewCluster(CassandraHostPort)
	cluster.Keyspace = "honey"
	cluster.Consistency = gocql.Quorum
	session, _ := cluster.CreateSession()
	c := CassandraClient{
		session: session,
	}
	return c
}

func (c *CassandraClient) Close() {
	c.session.Close()
}

func (c *CassandraClient) Get(id int) string {
	var path string

	if err := c.session.Query(`SELECT id, path FROM data WHERE id = ?`,
		id).Consistency(gocql.One).Scan(&id, &path); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Result:", id, path)
	return path
}

func (c *CassandraClient) Insert(id int, path string) {
	if err := c.session.Query(`INSERT INTO data (id, path) VALUES (?, ?)`,
		id, path).Exec(); err != nil {
		log.Fatal(err)
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

	client = NewCassandraClient(CassandraHostPort)
	defer client.Close()

	log.Fatal(http.ListenAndServe(hostPort, nil))
}
