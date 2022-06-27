# trieServer
Hosts a public http server on google cloud, allowing users to modify and view a trie data structure.

The server is hosted on google cloud with the their App Engine. I created the server using the language go, and their "net/http" package allows me to create a server. To execute commands on the trie, I have seperate request handlers, each with a specific url. To get the users input for insertion, deletion, and searching, they can add the word onto the end of the url such as https://....../add/word 

The client side of server uses the python requests module to send a request to the server when the user executes the command. 

Room for Improvement
- More error handling
- Concurrency
- Backing up the Server
- Letting the client know what the issue is if there is one
