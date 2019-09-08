# GolangYoti

# Why I chose Golang Instead of C++

Golang is a language im interested in and because thise role will require me to use Golang I thoguht this project was a great opportuniy to use GoLang.

# Thought Process 

First I thought about how I wanted to design the server, once I had an idea I started off by writing all the functionlity in server.go and then started to refactor.
In the process of refactoring I tried to adhere to the SOLID principles because if there are any modifications or extentions in the future it makes life easier if the code is clean 
and each file has one purpose. This is why I have files such as the router.go and routes.go which do exactly one thing and they are easy to extend.

# How to test the server

I used Postman to test the requests and responses examples are below.


Encryption Request 

{
  "ID": "1",
  "Data": "Yoti"
}

Encryption Response

{
  "Result": "Data encrypted succesfully!",
  "Key": "A6ylprDz3VIiQyt2imVMWQ=="
}

Decryption Request

{
  "ID": "1",
  "Key": "A6ylprDz3VIiQyt2imVMWQ=="
}

Decryption Response 

{
    "Result": "Data decrypted succesfully!",
    "Data": "Yoti"
}