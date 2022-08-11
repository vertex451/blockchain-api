# GraphQL

[Main Readme](/README.md) > [Backend Readme](../../../../../README.md)


To send the first query find the address in service logs and open it in browser.
Then send next request(Place any name instead of "Bob"):
```
query hello {
  hello(name: "Bob") {
    response
  }
}
```

You should get in response:
``` 
{
  "data": {
    "hello": {
      "response": "Hello, Bob"
    }
  }
}
```
