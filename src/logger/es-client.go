package logger

import (
	"fmt"
	"log"
	"sample_go_app/src/config"

	"github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/esutil"
)

var client *elasticsearch.Client

func init() {
	var err error

	envVar := config.DefaultEnvironmentalVariable

	cfg := elasticsearch.Config{
		Addresses: []string{
			envVar.ELASTICSEARCH_END_URL,
		},
		Username: envVar.ELASTICSEARCH_USER,
		Password: envVar.ELASTICSEARCH_PASSWORD,
	}

	client, err = elasticsearch.NewClient(cfg)
	if err != nil {
		log.Fatalf("Error creating the client: %s", err)
	}

}

func InsertDocs(docs *LogMessage) error {
	envVar := config.DefaultEnvironmentalVariable

	insertResponse, _ := client.Index(envVar.ELASTICSEARCH_LOG_INDEX, esutil.NewJSONReader(docs))

	fmt.Println("res", insertResponse)
	defer insertResponse.Body.Close()
	return nil
}
