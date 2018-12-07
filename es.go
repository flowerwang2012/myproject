package main

import (
	"github.com/olivere/elastic"
	"fmt"
	"context"
	"time"
	"encoding/json"
)
// Tweet is a structure used for serializing/deserializing data in Elasticsearch.
type Tweet struct {
	User     string                `json:"user"`
	Message  string                `json:"message"`
	Retweets int                   `json:"retweets"`
	Image    string                `json:"image,omitempty"`
	Created  time.Time             `json:"created,omitempty"`
	Tags     []string              `json:"tags,omitempty"`
	Location string                `json:"location,omitempty"`
	Suggest  *elastic.SuggestField `json:"suggest_field,omitempty"`
}

const mapping = `
{
	"settings":{
		"number_of_shards": 1,
		"number_of_replicas": 0
	},
	"mappings":{
		"tweet":{
			"properties":{
				"user":{
					"type":"keyword"
				},
				"message":{
					"type":"text",
					"store": true,
					"fielddata": true
				},
				"image":{
					"type":"keyword"
				},
				"created":{
					"type":"date"
				},
				"tags":{
					"type":"keyword"
				},
				"location":{
					"type":"geo_point"
				},
				"suggest_field":{
					"type":"completion"
				}
			}
		}
	}
}`

func main() {
	ctx := context.Background()
	client, err := elastic.NewClient(elastic.SetURL("http://127.0.0.1:9200"))
	if err != nil {
		// Handle error
		panic(err)
	}
	// Ping the Elasticsearch server to get e.g. the version number
	info, code, err := client.Ping("http://127.0.0.1:9200").Do(ctx)
	if err != nil {
		// Handle error
		panic(err)
	}
	fmt.Printf("Elasticsearch returned with code %d and version %s\n", code, info.Version.Number)

	// Getting the ES version number is quite common, so there's a shortcut
	esversion, err := client.ElasticsearchVersion("http://127.0.0.1:9200")
	if err != nil {
		// Handle error
		panic(err)
	}
	fmt.Printf("Elasticsearch version %s\n", esversion)

	// Use the IndexExists service to check if a specified index exists.
	exists, err := client.IndexExists("twitter").Do(ctx)
	if err != nil {
		// Handle error
		panic(err)
	}
	if !exists {
		// Create a new index.
		createIndex, err := client.CreateIndex("twitter").BodyString(mapping).Do(ctx)
		if err != nil {
			// Handle error
			panic(err)
		}
		if !createIndex.Acknowledged {
			// Not acknowledged
			// 执行结果 { "acknowledged":true }
		}
	} else {
		fmt.Println("twitter is exists")
	}

	// Index a tweet (using JSON serialization)
	//tweet1 := Tweet{User: "olivere", Message: "Take Five", Retweets: 0}
	//put1, err := client.Index().
	//	Index("twitter").
	//	Type("tweet").
	//	Id("1").
	//	BodyJson(tweet1).
	//	Do(ctx)
	//if err != nil {
	//	// Handle error
	//	panic(err)
	//}
	//fmt.Printf("Indexed tweet %s to index %s, type %s\n", put1.Id, put1.Index, put1.Type)
	//for i := 2; i < 24; i++ {
	//	// Index a second tweet (by string)
	//	tweet2 := `{"user" : "olivere", "message" : "It's a Raggy Waltz"}`
	//	put2, err := client.Index().
	//		Index("twitter").
	//		Type("tweet").
	//		Id(strconv.Itoa(i)).
	//		BodyString(tweet2).
	//		Do(ctx)
	//	if err != nil {
	//		// Handle error
	//		panic(err)
	//	}
	//	fmt.Printf("Indexed tweet %s to index %s, type %s\n", put2.Id, put2.Index, put2.Type)
	//}
	// Get tweet with specified ID
	get1, err := client.Get().
		Index("twitter").
		Type("tweet").
		Id("1").
		Do(ctx)
	if err != nil {
		// Handle error
		panic(err)
	}
	if get1.Found {
		fmt.Printf("Got document %s in version %d from index %s, type %s\n", get1.Id, get1.Version, get1.Index, get1.Type)
	}
	// Flush to make sure the documents got written.
	_, err = client.Flush().Index("twitter").Do(ctx)
	if err != nil {
		panic(err)
	}

	// Search with a term query
	//termQuery := elastic.NewTermQuery("user", "olivere")
	boolQuery := elastic.NewBoolQuery().Must()
	//boolQuery = boolQuery.Must(elastic.NewMatchQuery("user", "olivere"))
	//boolQuery = boolQuery.Must(elastic.NewMatchQuery("message","Take Five"))
	searchResult, err := client.Search().
		Index("twitter").   // search in index "twitter"
		Query(boolQuery).   // specify the query
		//Sort("user", true). // sort by "user" field, ascending
		From(0).Size(10).   // take documents 0-9
		Pretty(true).       // pretty print request and response JSON
		Do(ctx)             // execute
	if err != nil {
		// Handle error
		panic(err)
	}
	// searchResult is of type SearchResult and returns hits, suggestions,
	// and all kinds of other information from Elasticsearch.
	fmt.Printf("Query took %d milliseconds\n", searchResult.TookInMillis)

	// Each is a convenience function that iterates over hits in a search result.
	// It makes sure you don't need to check for nil values in the response.
	// However, it ignores errors in serialization. If you want full control
	// over iterating the hits, see below.
	//var ttyp Tweet
	//for _, item := range searchResult.Each(reflect.TypeOf(ttyp)) {
	//	if t, ok := item.(Tweet); ok {
	//		fmt.Printf("Tweet by %s: %s\n", t.User, t.Message)
	//	}
	//}
	// TotalHits is another convenience function that works even when something goes wrong.
	fmt.Printf("Found a total of %d tweets\n", searchResult.TotalHits())

	// Here's how you iterate through results with full control over each step.
	if searchResult.Hits.TotalHits > 0 {

		// Iterate through results
		for _, hit := range searchResult.Hits.Hits {
			// hit.Index contains the name of the index

			// Deserialize hit.Source into a Tweet (could also be just a map[string]interface{}).
			var t Tweet
			err := json.Unmarshal(*hit.Source, &t)
			if err != nil {
				// Deserialization failed
			}

			// Work with tweet
			fmt.Printf("Tweet id%s by %s: %s\n", hit.Id, t.User, t.Message)
		}
	} else {
		// No hits
		fmt.Print("Found no tweets\n")
	}

	// Update a tweet by the update API of Elasticsearch.
	// We just increment the number of retweets.
	//update, err := client.Update().Index("twitter").Type("tweet").Id("1").
	//	Script(elastic.NewScriptInline("ctx._source.retweets += params.num").Lang("painless").Param("num", 1)).
	//	//Upsert(map[string]interface{}{"retweets": 0}). upsert还不清楚如何使用
	//	Do(ctx)
	//if err != nil {
	//	// Handle error
	//	panic(err)
	//}
	//fmt.Printf("New version of tweet %q is now %d\n", update.Id, update.Version)

	// Delete an index.
	//deleteIndex, err := client.DeleteIndex("logstash-2018.05.30").Do(ctx)
	//if err != nil {
	//	// Handle error
	//	panic(err)
	//}
	//if !deleteIndex.Acknowledged {
	//	// Not acknowledged
	//	fmt.Println("Not acknowledged")
	//}

	//searchResult1, err := client.Search().
	//	Index("logstash-*").   // search in index "twitter"
	//	From(0).Size(10).   // take documents 0-9
	//	Pretty(true).       // pretty print request and response JSON
	//	Do(ctx)             // execute
	//if err != nil {
	//	// Handle error
	//	panic(err)
	//}
	//// searchResult is of type SearchResult and returns hits, suggestions,
	//// and all kinds of other information from Elasticsearch.
	//fmt.Printf("Query took %d milliseconds\n", searchResult1.TookInMillis)
	//fmt.Printf("Found a total of %d tweets\n", searchResult1.Hits.TotalHits)

	aggregation := elastic.NewTermsAggregation().Field("json.queue_name").Size(10000)
	searchResult2, err := client.Search().
		Index("filebeat-*").
		Query(elastic.NewMatchAllQuery()).
		Aggregation("all_queue_name", aggregation).
		Size(10000).
		Pretty(true).
		Do(ctx)
	if err != nil {
		// Handle error
		panic(err)
	}
	agg, found := searchResult2.Aggregations.Terms("all_queue_name")
	if !found {
		fmt.Printf("we should have a terms aggregation called %q", "timeline")
	}
	for _, bucket := range agg.Buckets {
		// Every bucket should have the user field as key.
		queuename := bucket.Key
		fmt.Println(queuename)
	}
	fmt.Println(len(searchResult2.Hits.Hits))
}