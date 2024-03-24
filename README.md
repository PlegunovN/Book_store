My first go project 


I found a small server in the open spaces with one file.
I corrected minor inaccuracies and the server started working.
I divided individual functions into separate files and packages.
Created a Postgres database in a Docker container.
Converted functions to work with the database.


Everything was done solely for educational purposes.

DB tables schema:
	book:
	id bigint PRIMARY KEY
	title VARCHAR(50)
	author bigint
	
	author:
	id bigint PRIMARY KEY 		= 	book.author
	firstname VARCHAR(50)
	lastname VARCHAR(50)



Create new book:
URL: 127.0.0.1:8080/book
Method: Post
JSON body: 
{
    "title": "Mu-mu",
    "author": {
        "firstname": "Ivan",
        "lastname": "Turgeniev"
    }
}
Response: nil



Get one book:
URL: 127.0.0.1:8080/book/{book id}
Method:Get
Response:
 {
        "id": 11,
        "title": "Mu-mu",
        "author": 13,
        "firstname": "Ivan",
        "lastname": "Turgeniev"
    }
    
    
    
Get one author:
URL: 127.0.0.1:8080/author/{author id}
Method: Get
Response:
{
    "id": 13,
    "firstname": "Ivan",
    "lastname": "Turgeniev"
}



Update book:
URL:127.0.0.1:8080/update/book
Method: Put
Json body:
{
    "title": "Mu-mu2",
    "id": 11
}
Response: nil



Update author:
URL:127.0.0.1:8080/update/author
Method: Put
Json body:
{
    "id": 9,
    "firstname": "newFirs",
    "lastname": "newLast"
}
Response: nil



Get books:
URL: 127.0.0.1:8080/books?limit=10&offset=0
Method: Get
Query params: limit, offset
Order by: book.id
Response: some books

































