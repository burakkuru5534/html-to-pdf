# Html To Pdf Service

## Introduction

In this project, we will be building a html to pdf service.

### Languages and frameworks

Technologies used in this project:

Golang,
postgresql

Test Environments:

postman,
golang testing library,

### Database

Postgresql was used as the database language.

Tables created:

table html_content:document
columns:
id: serial primary key
html_content: text
pdf_file_html_content: text (unique)

## Problem solution

We should be able to create, read, and delete,list documents and also download it.
We should prevent document from creating duplicate pdf file html_contents.

### Create Document

Create Document request url example:

Method: POST

http://localhost:8080/api/document

request Body Example:

 ```json
{
  "html_content": "<h1>hello world</h1>",
  "pdf_file_name": "hello_world.pdf"
}
 ```

response example:

for 200:

 ```json
{
  "id": 1,
  "html_content": "<h1>hello world</h1>",
  "pdf_file_name": "hello_world.pdf"
}
 ```

for 400:
    
```json
{"error": "Bad request"}
```

for 403:
```json
{"error": "Document with that pdf_file_name already exists"}
```

for 500:
```json
{"error": "server error"}
```

### Get Document 

Get Document request url example:

Method: GET

http://localhost:8080/Document?id=1

id: this id should be one of the Document's ids.

request Body:

response example:

for 200:
 ```json
{
  "id": 1,
  "html_content": "<h1>hello world</h1>",
  "pdf_file_name": "hello_world.pdf"
}
 ```

for 400:

```json
{"error": "Bad request"}
```

for 404:
```json
{"error": "Document with that id does not exist"}
```

for 500:
```json
{"error": "server error"}
```


### Delete Document

Delete Document request url example:
Method: DELETE

http://localhost:8080/api/document?id=1

id: this id should be one of the Document's ids.

response example:

for 200: "ok"


for 400:

```json
{"error": "Bad request"}
```

for 404:
```json
{"error": "Document with that id does not exist"}
```

for 500:
```json
{"error": "server error"}
```

### Document List

Document List request url example:
Method: GET

http://localhost:8080/document

response example:

for 200:
 ```json
[{
  "html_content":"<h1>hello world</h1>",
  "pdf_file_name":"hello_world.pdf"
}]
 ```
for 400:

```json
{"error": "Bad request"}
```

for 500:
```json
{"error": "server error"}
```

### Download Document

Download Document request url example:

Method: GET

http://localhost:8080/document/download?id=1

id: this id should be one of the Document's ids.

response example:

for 200: "ok"

for 400:

```json
{"error": "Bad request"}
```

for 404:
```json
{"error": "Document with that id does not exist"}
```

for 500:
```json
{"error": "server error"}
```

### Test

I used postman and also golang testing libary to test these rest APIs

you can run test by typing:

go test -v

## Conclusion

We have successfully implemented the HTML to PDF service.

note: I was going to use chi router to get id from url but I couldn't get it to work for testing library. So I used query params instead.



