# mylibrary
<img src="./library.jpg" alt="library" width="468" height="318">

MyLibrary was created as an exercise to build a library website with CRUD features.

## The REST APIs
+ GET /library/v1/books

  Returns a json document containing information for all available books in the following format:
  ```
  [
       {
         "bookId": string,
         "title":  string,
         "Author": string,
         "Pages":  int
       },
       ...
  ]
  ```
+ GET /library/v1/book/\<bookId\>

  Returns the book's contents in addition to the book's information:
  ```
    {
      "info": {
        "bookId": string,
        "title":  string,
        "Author": string,
        "Pages":  int
      },
      "Content": string
    }
    ```
  The reply also contains a 'bid' attribute identifying the book's id as a positive confirmation.
  
 + POST /library/v1/book
 
   This uploads a new book using the same json form returned by 'GET /library/v1/book/<bookId>'.
   The reply contains a 'bid' attribute identifying the new book's id.
   
 + PATCH /library/v1/book/\<bookId\>
 
   This updates an existing book using the elements found in the json form returned by 'GET /library/v1/book/<bookId>'.
   Only those elements found in the body's json document will be used to update the book.
   
 + DELETE /library/v1/book/\<bookId\>
 
   Deletes an exiting book

## Examples
```curl http://localhost:12250/library/v1/books```

```curl http://localhost:12250/library/v1/book/2```

```curl --header "Content-Type: application/json" --request POST --data @peacock.json http://localhost:12250/library/v1/book```

```curl --header "Content-Type: application/json" --request PATCH  --data @fixfox.json http://localhost:12250/library/v1/book/1```

```curl  --request DELETE  http://localhost:12250/book/3```

``````