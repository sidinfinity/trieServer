package main

import (
    "fmt"
    "log"
    "net/http"
    "os"
    "strings"
)

type TrieNode struct {
    child [26]*TrieNode
    isEnd bool
}

type trie struct {
    root *TrieNode
}

var t *trie

func initTrie() *trie {
    return &trie{
        root: &TrieNode{},
    }
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

func isEmpty(node *TrieNode) bool {
  for i:=0; i < 26; i++ {
    if (node.child[i] != nil) {return false}
  }
  return true
}

func (t *trie) delete(node *TrieNode, word string, count int) *TrieNode {

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
  if (node.child[index] != nil) {
    node.child[index] = t.delete(node.child[index], word, count+1)
  }

  if (isEmpty(node) && node.isEnd == false) {
    node = nil
  }

  return node
}

func (t *trie) displayTrie(node *TrieNode, s string, w http.ResponseWriter) bool {
  cur := node;

  if cur.isEnd {
    fmt.Fprintf(w, s + " ")
  }
  for i:=0; i < 26; i++ {
    if cur.child[i] != nil {
        temp := s;
        t.displayTrie(cur.child[i], temp + string(i + 97), w)
    }
  }
  return true
}

func (t *trie) autoComplete(word string, w http.ResponseWriter) bool {
  cur := t.root
  wordLength := len(word)

  for i:=0; i < wordLength; i++ {
    index := word[i] - 'a'
    if cur.child[index] == nil {
      fmt.Fprintf(w, "No Words Found")
      return false;
    }
    cur = cur.child[index]
  }
  t.displayTrie(cur, word, w)
  return true
}

func (t *trie) insert(word string) bool {
    if (len(word) == 0) {return false}

    wordLength := len(word)
    cur := t.root

    for i:=0; i < wordLength; i++ {
        index := word[i] - 'a'
        if cur.child[index] == nil {
            cur.child[index] = &TrieNode{}
        }
        cur = cur.child[index]
      }
    cur.isEnd = true

    return true
}

func add(w http.ResponseWriter, r *http.Request) {
    fmt.Println(w, "Welcome to the Add!")
    fmt.Println(r.URL.Path)
    word := strings.Split(r.URL.Path, "/")[2]
    fmt.Println(word)
    t.insert(word)

    /*
    Adds word to file

    if t.insert(word) {
        f, err := os.OpenFile("words.txt", os.O_APPEND|os.O_WRONLY, 0755)
        defer f.Close()
        if err != nil {
            log.Fatal(err)
        }

        if _,err = f.WriteString(word + "\n"); err != nil {
            log.Fatal(err)
        }
    }
    */

}

/*
Deletes word from a file

func delFromFile(word string) {
  f, err := os.Open("words.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	var bs []byte
	buf := bytes.NewBuffer(bs)

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		if scanner.Text() != word {
			_, err := buf.Write(scanner.Bytes())
			if err != nil {
				log.Fatal(err)
			}
			_, err = buf.WriteString("\n")
			if err != nil {
				log.Fatal(err)
			}
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	err = ioutil.WriteFile("words.txt", buf.Bytes(), 0666)
  if err != nil {
    log.Fatal(err)
  }

}
*/

func delete(w http.ResponseWriter, r *http.Request) {
    fmt.Println(w, "Delete")
    fmt.Println(r.URL.Path)
    word := strings.Split(r.URL.Path, "/")[2]
    fmt.Println(word)
    t.delete(t.root, word, 0)
}

func display(w http.ResponseWriter, r *http.Request) {
    fmt.Println(w, "Display")
    t.displayTrie(t.root, "", w)
}

func homePage(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "Welcome to the HomePage!")
    fmt.Println("Endpoint Hit: homePage")
}

func search(w http.ResponseWriter, r *http.Request) {
    fmt.Println(w, "Search")
    fmt.Println(r.URL.Path)
    word := strings.Split(r.URL.Path, "/")[2]
    fmt.Println(word)
    if (t.search(word)) {
      fmt.Fprintf(w, "Word Found")
    } else {
      fmt.Fprintf(w, "Word Not Found")
    }
}

func autoComplete(w http.ResponseWriter, r *http.Request) {
    fmt.Println(w, "AutoComplete")
    fmt.Println(r.URL.Path)
    word := strings.Split(r.URL.Path, "/")[2]
    fmt.Println(word)
    if (!t.autoComplete(word, w)) {
      fmt.Fprintf(w, "No Word Exists")
    }
}

func handleRequests() {
    http.HandleFunc("/", homePage)
    http.HandleFunc("/add/", add)
    http.HandleFunc("/delete/", delete)
    http.HandleFunc("/display", display)
    http.HandleFunc("/search/", search)
    http.HandleFunc("/autocomplete/", autoComplete)

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

func main() {
    t = initTrie()
    /*
    f, err := os.OpenFile("words.txt", os.O_RDWR, 0644)
    if err != nil {
        log.Fatal(err)
    }
    defer f.Close()

    scanner := bufio.NewScanner(f)
    // optionally, resize scanner's capacity for lines over 64K, see next example
    for scanner.Scan() {
        t.insert(scanner.Text())
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
    */
    handleRequests()
}
