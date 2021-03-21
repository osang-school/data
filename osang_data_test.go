package osangdata

import (
	"testing"
)

func TestCrawl(t *testing.T) {
	result, err := CrawlList(UrlRule, 1)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("list: %+v\n", result[0])
	detail, err := CrawlPage(UrlNotice, 196655955)
	if err != nil {
		t.Fatal(err)
		return
	}
	t.Logf("detail: %+v\n", detail)
}
