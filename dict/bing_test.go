package dict

import "testing"

func TestBingQuery(t *testing.T) {
	result := BingQuery("hello")
	t.Log(result)
}

func TestSougouQuery(t *testing.T) {
	result := SougouQuery("hi")
	t.Log(result)
}
