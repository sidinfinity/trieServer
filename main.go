package main

import (
  "fmt"
  "log"
  "net/http"
  "strings"
  "os"
)

type trieNode struct {
  child [26]*trieNode
  isEnd bool
}

type trie struct {
    root *trieNode
}

func initTrie() *trie {
    return &trie{
        root: &trieNode{},
    }
}

func (t *trie) insert(word string) bool {
  if (len(word) == 0) {return false}

  wordLength := len(word)
  cur := t.root

  for i:=0; i < wordLength; i++ {
    index := word[i] - 'a'
    if cur.child[index] == nil {
      cur.child[index] = &trieNode{}
    }
    cur = cur.child[index]
  }
  cur.isEnd = true

  return true
}

func (t *trie) search(word string) bool {
  wordLength := len(word)
  cur := t.root

  for i:=0; i < wordLength; i++ {
      index := word[i] - 'a'
      if cur.child[index] == nil {
        return false
      }
      cur = cur.child[index]
  }

  return cur.isEnd

}

func isEmpty(node *trieNode) bool {
  for i:=0; i < 26; i++ {
    if (node.child[i] != nil) {return false}
  }
  return true
}

func (t *trie) delete(node *trieNode, word string, count int) *trieNode {

  if count == len(word)  {
    if (node.isEnd) {
      node.isEnd = false
    }

    if(isEmpty(node)) {
      node = nil
    }
    return node
  }

  index := word[count] - 'a'
  node.child[index] = t.delete(node.child[index], word, count+1)


  if (isEmpty(node) && node.isEnd == false) {
    node = nil
  }

  return node
}

func (t *trie) displayTrie(node *trieNode, s string) bool {
  cur := node;

  if cur.isEnd {
    fmt.Println(s)
  }
  for i:=0; i < 26; i++ {
    if cur.child[i] != nil {
        temp := s;
        t.displayTrie(cur.child[i], temp + fmt.Sprint(i + 97))
    }
  }
  return true
}

func (t *trie) autoComplete(word string) bool {
  cur := t.root
  wordLength := len(word)

  for i:=0; i < wordLength; i++ {
    index := word[i] - 'a'
    if cur.child[index] == nil {
      fmt.Println("No Words Found")
      return false;
    }
    cur = cur.child[index]
  }
  t.displayTrie(cur, word)
  return true
}


func homePage(w http.ResponseWriter, r *http.Request){
  fmt.Println(w, "Welcome to the HomePage! %s", *r)
  fmt.Println("Endpoint Hit: homePage")
}

func add(w http.ResponseWriter, r *http.Request) {
  fmt.Println(w, "Welcome to the Add!")
  fmt.Println(r.URL.Path)
  word := strings.Split(r.URL.Path, "/")[2]
  fmt.Println(word)


}

func delete(w http.ResponseWriter, r *http.Request) {
  fmt.Println(w, "Delete")
  fmt.Println(r.URL.Path)
  word := strings.Split(r.URL.Path, "/")[2]
  fmt.Println(word)



}

func display(w http.ResponseWriter, r *http.Request) {
  fmt.Println(w, "Display")

}

func search(w http.ResponseWriter, r *http.Request) {
  fmt.Println(w, "Search")
  fmt.Println(r.URL.Path)
  word := strings.Split(r.URL.Path, "/")[2]
  fmt.Println(word)
}

func autoComplete(w http.ResponseWriter, r *http.Request) {
  fmt.Println(w, "AutoComplete")
  fmt.Println(r.URL.Path)
  word := strings.Split(r.URL.Path, "/")[2]
  fmt.Println(word)
}

func handleRequests() {
  http.HandleFunc("/", homePage)
  http.HandleFunc("/add/", add)
  http.HandleFunc("/delete", delete)
  http.HandleFunc("/display", display)
  http.HandleFunc("/search", search)
  http.HandleFunc("/autocomplete", autoComplete)
}

func main() {
  handleRequests()

  port := os.Getenv("PORT")
    if port == "" {
      port = "8080"
      log.Printf("Defaulting to port %s", port)
    }

    log.Printf("Listening on port %s", port)
    if err := http.ListenAndServe(":"+port, nil); err != nil {
            log.Fatal(err)
    }
}
