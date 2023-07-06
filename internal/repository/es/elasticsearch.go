package es

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/estransport"
	log "github.com/sirupsen/logrus"
	"io"
	"item-search/internal/repository"
	"item-search/pkg/config"
	"net"
	"net/http"
	"os"
	"strings"
	"time"
)

var (
	result map[string]interface{}
	index  string
)

func Init() {
	esConfig := config.Conf.Elastic
	index = esConfig.Index
	address := fmt.Sprintf("%s://%s:%d", esConfig.Protocol, esConfig.Host, esConfig.Port)
	cfg := elasticsearch.Config{
		Addresses: []string{
			address,
		},
		Username: esConfig.Username,
		Password: esConfig.Password,
		Logger:   &estransport.ColorLogger{Output: os.Stdout},
		Transport: &http.Transport{
			MaxIdleConnsPerHost:   10,
			ResponseHeaderTimeout: time.Second,
			DialContext:           (&net.Dialer{Timeout: time.Second}).DialContext,
			TLSClientConfig: &tls.Config{
				MinVersion: tls.VersionTLS12,
			},
		},
	}

	es, err := elasticsearch.NewClient(cfg)
	if err != nil {
		log.Fatalf("Elasticsearch info error, %s", err)
	}
	repository.Es = es
}

func Search(data interface{}) {
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(data); err != nil {
		log.Fatalf("Error encoding query: %s", err)
	}
	res, err := repository.Es.Search(
		repository.Es.Search.WithContext(repository.Ctx),
		repository.Es.Search.WithIndex(index),
		repository.Es.Search.WithBody(&buf),
		repository.Es.Search.WithTrackTotalHits(true),
		repository.Es.Search.WithPretty(),
	)
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatalf("Error getting response: %s", err)
		}
	}(res.Body)

	if res.IsError() {
		var e map[string]interface{}
		if err := json.NewDecoder(res.Body).Decode(&e); err != nil {
			log.Fatalf("Error parsing the response body: %s", err)
		} else {
			// Print the response status and error information.
			log.Fatalf("[%s] %s: %s", res.Status(), e["error"].(map[string]interface{})["type"],
				e["error"].(map[string]interface{})["reason"])
		}
	}
	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		log.Fatalf("Error parsing the response body: %s", err)
	}
	// Print the response status, number of results, and request duration.
	log.Printf(
		"[%s] %d hits; took: %dms",
		res.Status(),
		int(result["hits"].(map[string]interface{})["total"].(map[string]interface{})["value"].(float64)),
		int(result["took"].(float64)),
	)
	// Print the ID and document source for each hit.
	for _, hit := range result["hits"].(map[string]interface{})["hits"].([]interface{}) {
		log.Printf(" * ID=%s, %s", hit.(map[string]interface{})["_id"], hit.(map[string]interface{})["_source"])
	}

	log.Println(strings.Repeat("=", 37))
}
