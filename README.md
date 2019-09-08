# GolangYoti

## Overal Imrprovements that can be made

Refactoring im pretty big on clean code and generally refactor my code once the functiionlity is there, due to time I did not have chance to refactor everything, handler.go
is a file id like to clean up. Its not messy but I feel like theres a cleaner way to handle requests in Go.

Add the Database if there was more time.

On the client code I would of changed the code to be a console menu system that way the user can continully send requests.

## Why I chose Golang Instead of C++

Golang is a language im interested in and because thise role will require me to use Golang I thought this project was a great opportunity to use GoLang.

## Thought Process 

First I thought about how I wanted to design the server, once I had an idea I started off by writing all the functionlity in server.go and then started to refactor.
In the process of refactoring I tried to adhere to the SOLID principles because if there are any modifications or extentions in the future it makes life easier if the code is clean 
and each file has one purpose. This is why I have files such as the router.go and routes.go which do exactly one thing and they are easy to extend.

## How to test the server

I used Postman to test the requests and responses examples are below.


### Encryption Request 

```json
{
  "ID": "1",
  "Data": "Yoti"
}
```

### Encryption Response

```json
{
  "Result": "Data encrypted succesfully!",
  "Key": "A6ylprDz3VIiQyt2imVMWQ=="
}
```

### Decryption Request

```json
{
  "ID": "1",
  "Key": "A6ylprDz3VIiQyt2imVMWQ=="
}
```

### Decryption Response 

```json
{
    "Result": "Data decrypted succesfully!",
    "Data": "Yoti"
}
```

### Testing

I created tests for the server such as health checks to confirm the correct status is being returned and I created a integration test which tests the whole service.

### Extra Credits

I didn't get time to implement this part due to family events over the weekend but I wrote down an implementation for having the repo as a serperate instace/service

I would have a NoSQL database such as MongoDB running as its own instance where the server would make a connection to on startup, once this connection is made we can then peform CRUD 
operations on that database.

Why MongoDB, it doesn't have to be MongoDB it could be another NoSQL database my reason for this is because I dont think the information we are saving at the moment warrents a SQL daabase, In the future we could add/remove infomation we want to be stored and its easier to accomodate this in a NoSQL database.

### Performance Improved

If there was a database which we were sending requests to then that would take some time and adds some latency, a way to improve this would be to add an in memory cache like Redis,
we could store the same information in the cache and that way when a user sends a decryption request we can read the cipher text from the cache which would be much faster than making a call to the database.