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

1. List all books (list)
2. List the books in which the given entry is in the titles of the books. (search)
3. Get book by ID (get)
4. Delete the book whose ID is given. (delete)
5. Buy the book with the ID given as many as you want and print the last information of the book on the screen.

Yanlış komut girildiğinde ekrana usage'ı yazdıracak. 


Concurrency ile ilgili medium yazısı yazılacak. 

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
# Requirements:
- README
- No third party package(s)
- Everything should be in English (Comments, Function names, File names, etc.)
- Use structs not maps
