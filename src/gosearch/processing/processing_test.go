package processing

import (
  "fmt"
  "testing"
)

func TestSimpleTokenizer(t *testing.T) {
  ref := []string{"See", "John", "Markoff", "Apple", "Adopts", "Open", "Source", "for", "its", "Server", "Computers", "New", "York", "Times"}
  tokenized := SimpleTokenizer(`
      #See John Markoff, “Apple Adopts ‘Open Source’ for its Server Computers, ‘New York Times’”
     `, 128)

  if fmt.Sprintf("%v",tokenized) != fmt.Sprintf("%v",ref) {
    t.Errorf("expected %v == %v", tokenized, ref)
  }
}

func TestSimpleTokenizerApostrophe(t *testing.T) {
  ref := []string{"There's","not","a","problem","that","I","can't","fix"}
  tokenized := SimpleTokenizer(`
     'There's not a problem that I can't fix'
  `, 128)
  if fmt.Sprintf("%v",tokenized) != fmt.Sprintf("%v",ref) {
    t.Errorf("expected %v == %v", tokenized, ref)
  }
}

func TestSimpleTokenizerCJK(t *testing.T) {
  ref := []string{"宮","崎","駿"}
  tokenized := SimpleTokenizer(`
     宮崎 駿
  `, 128)
  if fmt.Sprintf("%v",tokenized) != fmt.Sprintf("%v",ref) {
    t.Errorf("expected %v == %v", tokenized, ref)
  }
}

func TestSimpleTokenizerEmptyString(t *testing.T) {
  tokenized := SimpleTokenizer("", 128)
  if len(tokenized) != 0 {
    t.Errorf("expected empty array, but got %v", tokenized)
  }
}

func TestSimpleTokenizerCap(t *testing.T) {
  ref := []string{"not","a","that","I","can't","fix"}
  tokenized := SimpleTokenizer(`
     'There's not a problem that I can't fix'
  `, 4)
  if fmt.Sprintf("%v",tokenized) != fmt.Sprintf("%v",ref) {
    t.Errorf("expected %v == %v", tokenized, ref)
  }
}

func TestLowercaseFilter(t *testing.T) {
  src := []string{"Compare", "FOO"}
  ref := []string{"compare", "foo"}
  result := LowercaseFilter(src)
  if fmt.Sprintf("%v",result) != fmt.Sprintf("%v",ref) {
    t.Errorf("expected %v == %v", result, ref)
  }
}

func TestStopWordsFilter(t *testing.T) {
  filter := CreateStopWordsFilter([]string{"foo","bar","baz"})
  ref := []string{"quick","brown","fox","jumps","over","lazy","dog"}
  result := filter([]string{"foo","quick","bar","baz","brown","foo","fox","jumps","over","bar","lazy","dog","baz"})
  if fmt.Sprintf("%v",result) != fmt.Sprintf("%v",ref) {
    t.Errorf("expected %v == %v", result, ref)
  }
}
