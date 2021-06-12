package zookeeper

import (
	"fmt"
	"strconv"
	"time"

	"github.com/logrusorgru/aurora"
	"github.com/samuel/go-zookeeper/zk"
)

var zkClient *zk.Conn

// Connect ...
func Connect(uri string) error {
	c, _, err := zk.Connect([]string{uri}, time.Second*30)
	if err != nil {
		fmt.Println("Error when connect to Zookeeper", uri, err)
		return err
	}

	fmt.Println(aurora.Green("*** CONNECTED TO ZOOKEEPER: " + uri))

	// Set client
	zkClient = c

	return nil
}

// GetStringValue ...
func GetStringValue(path string) string {
	value, _, err := zkClient.Get(path)
	if err != nil {
		fmt.Printf("Get value from path %s err: %v\n", path, err)
	}
	return string(value)
}

// GetFloat64Value ...
func GetFloat64Value(path string) float64 {
	s := GetStringValue(path)
	v, _ := strconv.ParseFloat(s, 64)
	return v
}

// GetIntValue ...
func GetIntValue(path string) int {
	s := GetStringValue(path)
	v, _ := strconv.Atoi(s)
	return v
}
