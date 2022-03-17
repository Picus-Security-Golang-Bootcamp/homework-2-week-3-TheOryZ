## Homework | Week 3


### Book list application with Go

We have books struct object and the fields are;
```
- Book ID
- Book Title
- Number of Page
- Number of Stock
- Price
- Stock Code
- ISBN
- Author (ID ve Name)
```

This application has five functions.
1. List all books (list)
2. List the books in which the given entry is in the titles of the books. (search)
3. Get book by ID (get)
4. Delete the book whose ID is given. (delete)
5. Buy the book with the ID given as many as you want and print the last information of the book on the screen.

You can access it by typing --help to see the usage patterns of the correct commands.
```
go run main.go --help
```


### list command
```
go run main.go list
```

### search command 
```
go run main.go search <bookName>
go run main.go search Lord of the Ring: The Return of the King
```

### get command
```
go run main.go get <bookID>
go run main.go get 5
```

### delete command
```
go run main.go delete <bookID>
go run main.go delete 5
```

### buy command
```
go run main.go buy <bookID> <quantity>
go run main.go buy 5 2
```

### 
There is a .json file in the application where the book information is kept.

```json
{
   "Books":[
      {
         "Id":1,
         "Title":"Dune",
         "Number_Of_Pages":412,
         "Number_Of_Stocks":22,
         "Price":20.5,
         "ISBN":"ISBN987654321",
         "Release_Date":"1965",
         "Author":{
            "Id":1,
            "Name":"Frank Herbet"
         },
         "IsDeleted":false
      },
      {
         "Id":2,
         "Title":"The Lord of The Rings",
         "Number_Of_Pages":1137,
         "Number_Of_Stocks":5,
         "Price":24.9,
         "ISBN":"ISBN987654321",
         "Release_Date":"1955",
         "Author":{
            "Id":2,
            "Name":"J. R. R. Tolkien"
         },
         "IsDeleted":true
      },
      {
         "Id":3,
         "Title":"1984",
         "Number_Of_Pages":328,
         "Number_Of_Stocks":10,
         "Price":18.5,
         "ISBN":"ISBN987654321",
         "Release_Date":"1949",
         "Author":{
            "Id":3,
            "Name":"George Orwell"
         },
         "IsDeleted":false
      },
      {
         "Id":4,
         "Title":"Brave New World",
         "Number_Of_Pages":311,
         "Number_Of_Stocks":1,
         "Price":19,
         "ISBN":"ISBN987654321",
         "Release_Date":"1932",
         "Author":{
            "Id":4,
            "Name":"Aldous Huxley"
         },
         "IsDeleted":false
      },
      {
         "Id":5,
         "Title":"Foundation",
         "Number_Of_Pages":255,
         "Number_Of_Stocks":8,
         "Price":32.5,
         "ISBN":"ISBN987654321",
         "Release_Date":"1942",
         "Author":{
            "Id":5,
            "Name":"Isaac Asimov"
         },
         "IsDeleted":false
      }
   ]
}
```