package elog

import (
	"fmt"
	"os"
	"time"

	"github.com/treetopllc/elastilog"
	"github.com/goadesign/goa"
)

// Logger
type ElastiLogger struct {
	count int
	client elastilog.Client
}

// New returns a goa log adapter backed by a nul logger.
func New(logger ElastiLogger) goa.LogAdapter {
	return &adapter{Logger: logger}
}

func NewElastiLogger() ElastiLogger {
	uri := "http://localhost:9200"
	return ElastiLogger{
		client: elastilog.NewClient(uri, "goaelog"),
		count: 0,
	}
}

type ElasticEntry struct {
	level string
	entry elastilog.Entry
}

func NewElasticEntry() *ElasticEntry {
	hostname, _ := os.Hostname()
	return &ElasticEntry{
		level: "info",
		entry: elastilog.Entry{
			Host:       hostname,
			Timestamp:  time.Now(),
			Attributes: make(elastilog.Attributes),
		},
	}
}
func (n *ElastiLogger) send(msg string, keyvals ...interface{}) {
	n.count++
	e := NewElasticEntry()
	e.entry.Log = msg
	fmt.Println(len(keyvals))
	data2el(e.entry.Attributes, keyvals)
	n.client.Send(e.entry)
}

func data2el(attrs elastilog.Attributes, keyvals []interface{}) {
	fmt.Println(len(keyvals))
	for i := 0; i < len(keyvals); i += 2 {
		k := keyvals[i]
		var v interface{} = goa.ErrMissingLogValue
		if i+1 < len(keyvals) {
			v = keyvals[i+1]
		}
		fmt.Printf("%#v=%#v\n", k, v)
		attrs[fmt.Sprintf("%v", k)] = fmt.Sprintf("%v", v)
	}
	return
}
// LogAdapter
type adapter struct {
	Logger ElastiLogger
}

func (a *adapter) Info(msg string, keyvals ...interface{}) {
	a.Logger.send(msg, keyvals)
}
func (a *adapter) Error(msg string, keyvals ...interface{}) {
	a.Logger.send(msg, keyvals)
}
func (a *adapter) New(keyvals ...interface{}) goa.LogAdapter {
	return &adapter{Logger: NewElastiLogger()}
}
