# Github Query API

NOTE: Because we have implemented the exercise in Go, 
all possible issue solutions will be presented in the Go way

## Technical issues

1. ### Main
   1. Get the file path param from config.
   2. Should use the common logger that we are using  
      in our company
   3. Add logging to the modules
   4. Add monitoring to the modules

2. ### Process Flow
   1. Current implementation doesn't use the multithreading approach.
      We can easily implement it and run any HTTP request in the independent
      thread. In this case we should take care about: 
      1. MAX num of subprocesses are running
      2. Close the subprocesses that are running too long
         (please feet the solution to the [...up to 10 requests per minute...](#integration-issues))
      
3. ### Data Feeder
   1. Data feeder - currently implemented as 'Read At All at Once'  
      there could be an issue in case of huge data source file, so  
      IMHO it good idea to send the URL via go-channel.  
      NOTE: the file will be locked during the reading process
   2. It's not bad idea to add the error handler in case of `GetData(filename)` func
      will be finished with an error

4. ### HTTP Client
   1. In the current implementation, we use the standard HTTP client.  
      The main issue is in establishing the connection to the  
      GitHub API (it costs a lot of time). So we should improve/find 
      the HTTP client who allows run the request with `"Keep Alive"` option

## Integration issues

1. ### Integration HTTP Client
   1. GitHub REST search API has a couple of limitation
      For example:  
      _"...up to 1,000 results for each search"_  
      _"... the rate limit allows you to make up to 10 requests per minute"_
      So...
   2. _"...1,000 results..."_ is pretty huge number, so we can keep   
      the GitHub approach, **BUT** we don't know the customer needs, than  
      it is good idea to discus about that with **our product and customer**.
      In case we should fetch ALL data the **pagination engine** will be 
      implemented (maybe we should add a config option to  
      allow run our fetcher in 2 modes)
   3. _"...up to 10 requests per minute..."_ additional topic to discuss with  
      product and customer. In our case (because the app is running as dedicated process)  
      I think we may implement the **retry engine** (as part of HTTPClient) 
      