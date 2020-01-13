package lib

import "testing"

func TestWRequest(t *testing.T) {
	targetURL := "http://www.baidu.com"
	resp, err := WRequest(targetURL)

	if err != nil {
		t.Fatal(err)
	}
	t.Fatal(resp)
}
