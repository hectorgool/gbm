

curl -u elastic:changeme -X DELETE "http://172.17.0.2:9200/gbm"
curl -u elastic:changeme -X PUT "http://172.17.0.2:9200/gbm" -d '
{
    "settings": {
        "index": {
            "analysis": {
                "analyzer": {
                    "autocomplete": {
                        "tokenizer": "whitespace",
                        "filter": [
                            "lowercase",
                            "engram"
                        ]
                    }
                },
                "filter": {
                    "engram": {
                        "type": "edgeNGram",
                        "min_gram": 1,
                        "max_gram": 10
                    }
                }
            }
        }
    },
    "mappings": {
        "vehicle": {
            "properties": {  
                "location": {
                    "type": "geo_point"
                }
            }
        }
    }
}
'