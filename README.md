# go-workshop

1. Start http server on port 8080, use `http.ListenAndServe()` function
2. Add a handler returning `Hello Stranger` text - you need to add struct implementing `Handler` interface
3. Extract this struct to standalone package, implement constructor and Interface for this struct as well, with method `Start`
4. Create model struct `Saying` with one `string` field called `Name` return this struct instead of `Hello Stranger` use `json.Marshall` to return json
5. Pass the name as an argument from cmdline use `flag` package

## Deployment
params: 
1. ssh username
2. FI name
```
sh deploy-script.sh fdvoracek go-heroes
```
