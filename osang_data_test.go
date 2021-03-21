package osangdata

import (
	"fmt"
	"testing"
)

func TestCrawl(t *testing.T) {
	result, err := CrawlList(UrlRule, 1)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("list: %+v\n", result[0])
	fmt.Println(result[0].ID)
	detail, err := CrawlPage(UrlRule, result[0].ID)
	if err != nil {
		t.Fatal(err)
		return
	}
	t.Logf("detail: %+v\n", detail)
}
