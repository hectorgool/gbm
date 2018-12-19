package elasticsearch

import (
	"context"
	"fmt"
	elastic "gopkg.in/olivere/elastic.v5"
	"log"
	"os"
	"github.com/hectorgool/mvp_gbm/common"
	"github.com/satori/go.uuid"
	"encoding/json"
	//debug "github.com/davecgh/go-spew/spew"
)

var (
	client *elastic.Client
)

type(
	SearchResult struct{
		VehicleID string `json:"vehicleid"`
		Location Location `json:"location"`
		Date string `json:"date"`
	}
	Location struct{
    	Latitude string `json:"lat"`
		Longitude string `json:"lon"`
	}
)

func init() {

	var err error

	client, err = elastic.NewClient(
		elastic.SetSniff(false), 
		elastic.SetURL(os.Getenv("ELASTICSEARCH_ENTRYPOINT")),
		elastic.SetBasicAuth(os.Getenv("ELASTICSEARCH_USERNAME"), os.Getenv("ELASTICSEARCH_PASSWORD")),
	)
	common.CheckError(err)
	indexExists()

}

func indexExists(){
	exists, err := client.IndexExists(os.Getenv("ELASTICSEARCH_INDEX")).Do(context.Background())
	common.CheckError(err)
	if !exists {
		createIndex()
	}
}

func createIndex(){
	mapping := `{
		"settings":{
			"number_of_shards":1,
			"number_of_replicas":0
		},
		"mappings": {
			"vehicle": { 
				"properties": {
					"vehicleid": {
						"type": "text",
						"store": true
					},
					"location": {
						"type": "geo_point"
					},
					"date": {
						"type": "date" 
					}
				}
			}
		}
	}`
	
	ctx := context.Background()
	createIndex, err := client.CreateIndex(os.Getenv("ELASTICSEARCH_INDEX")).BodyString(mapping).Do(ctx)
	common.CheckError(err)
	if !createIndex.Acknowledged {
		log.Println("create index not allow!")
	}
}

// Ping fuction
func Ping() (string, error) {

	ctx := context.Background()
	info, code, err := client.Ping(os.Getenv("ELASTICSEARCH_ENTRYPOINT")).Do(ctx)
	common.CheckError(err)

	msg := fmt.Sprintf("Elasticsearch returned with code %d and version %s", code, info.Version.Number)
	return msg, nil

}

// Read Document
func ReadDocument(id string) {

	ctx := context.Background()
	get, err := client.Get().
		Index(os.Getenv("ELASTICSEARCH_INDEX")).
		Type(os.Getenv("ELASTICSEARCH_TYPE")).
	    Id(id).
	    Do(ctx)
	common.CheckError(err)

	if get.Found {
	    log.Printf("Got document %s in version %d from index %s, type %s\n", get.Id, get.Version, get.Index, get.Type)
	}

}

// CreateDocument fuction save json in the server
func CreateDocument( vehicleid string, latitude, longitude float64 ) error {

	id := uuid.Must(uuid.NewV4()).String()
	jsonData := fmt.Sprintf(
	`{
		"vehicleid": "%v",
		"location": { 
			"lat": "%v",
			"lon": "%v"
		},
		"date": "%v"

	}`, vehicleid, latitude, longitude, common.MakeTimestamp() )

	ctx := context.Background()
	doc, err := client.Index().
		Index(os.Getenv("ELASTICSEARCH_INDEX")).
		Type(os.Getenv("ELASTICSEARCH_TYPE")).
	    Id(id).
	    BodyString(jsonData).
	    Do(ctx)
	common.CheckError(err)
	log.Printf("Indexed geolocation %s to index %s, type %s\n", doc.Id, doc.Index, doc.Type)
	return nil
}

func queryTerm(q string) string{
	
	var searchJSON string
	
	if( q == ""){
		searchJSON = `{
			"query": { 	
				"match_all": {}
			}	
		}`	
	} else{	
		searchJSON = fmt.Sprintf(`{
			"query" : {
				"match": {
					"vehicleid" : "%v"
				}
			}
		}`, q)
	}
	return searchJSON
}


// GetDocuments fuction save json in the server
func GetDocuments(q string) (*elastic.SearchResult, error) {

	searchJSON := queryTerm(q)
	ctx := context.Background()
	searchResult, err := client.Search().
		Index(os.Getenv("ELASTICSEARCH_INDEX")).
		Type(os.Getenv("ELASTICSEARCH_TYPE")).
		Source(searchJSON).
		Do(ctx)
	common.CheckError(err)

	return searchResult, nil
}

func DisplayResults( searchResult *elastic.SearchResult ) ([]*SearchResult, error) {

    var Documents []*SearchResult
    for _, hit := range searchResult.Hits.Hits {
        var d SearchResult	
		//parses *hit.Source into the instance of the Document struct
        err := json.Unmarshal(*hit.Source, &d)
		common.CheckError(err)

        Documents = append(
			Documents, 
			&SearchResult{
				VehicleID: d.VehicleID,
				Location: Location{
					Latitude: d.Location.Latitude, 
					Longitude: d.Location.Longitude,
				},
				Date: d.Date,
			},
		)
	}
    return Documents, nil

}

func Search(q string) ([]*SearchResult, error) {

	searchResult, err := GetDocuments(q)
	common.CheckError(err)
	result, err := DisplayResults(searchResult)
	common.CheckError(err)

	return result, nil

}