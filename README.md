# trieServer
Hosts a public http server on google cloud, allowing users to modify and view a trie data structure.

The server is hosted on google cloud with the their App Engine. I created the server using the language go, and their "net/http" package allows me to create a server. To execute commands on the trie, I have seperate request handlers, each with a specific url. To get the users input for insertion, deletion, and searching, they can add the word onto the end of the url such as https://....../add/word Unfortunatly, I wasn't able to complete the client CLI, but my plan was to create a python script and use their requests module to send an input, and get an output. 
