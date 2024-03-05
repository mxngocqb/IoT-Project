package elastic

import "github.com/elastic/go-elasticsearch/v8"

type elasticService struct {
	ec *elasticsearch.Client
}

