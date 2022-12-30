### Server

#### 1. ServeMux
- a map between path & function
- `DefaultServeMux`: default ServeMux used by the server
- `HandleFunc`: register the handler function for a pattern in DefaultServeMux
- `servemux.Handle(<pattern>, <handler>)` bind a pattern with a handler
- `Handler` a function to handle a pattern. Should implement `ServeHTTP(http.ResponseWrite, *http.Request)` method.

#### 2. Server
- binded to and address:port
- binded to a ServeMux (default is `DefaultServeMux`)
- `server.ListenAndServe()` blocking function (can be called in another goroutine)

#### 3. Graceful Shutdown
- a pattern to gracefully shutdown the server without interrupting any active connections

### Data Handling

#### 1. JSON format
- convert any object to JSON with `json.Marshal()`
- use `json.encoder.Encode()` to directly put the encoded json to the stream/writer (more eficient, useful for large objects)
- similar logic `json.decoder.Decode()` directly put decoded json to reader
- use **struct tag** to indicate the format and filter attributes for json conversion

