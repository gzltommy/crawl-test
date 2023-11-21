package persist

import (
	"github.com/gzltommy/crawl-test/05.DistributeCrawl/persist"
	"github.com/gzltommy/crawl-test/05.DistributeCrawl/types"
	"gopkg.in/olivere/elastic.v5"
)

type ItemService struct {
	Client *elastic.Client
}

func (s *ItemService) Save(item types.Item, result *string) error {
	err := persist.Save(s.Client, item)
	if err == nil {
		*result = "ok"
	}

	return err

}
