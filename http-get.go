package main

import (
  "fmt"
  "net/http"
  "time"
  "net"
  "os"
  "io/ioutil"
  "strings"
   flag "github.com/spf13/pflag"
   
)

var (
    input = flag.StringP("input", "i", "", "Input path. '-' means stdin\n")
    output = flag.StringP("output", "o", "", "Output path. '-' means stdout\n")
    userAgent = flag.StringP("userAgent", "u", "", "User agent for HTTP requests\n")
    maxDownloadTime = flag.IntP("maxDownloadTime", "t", 10, "Max download time (s)\n")
    maxConnectTime = flag.IntP("maxConnectTime", "T", 5, "Max connect time (s)\n")
)

 	

func keepLines(s string, n int) string {
	result := strings.Join(strings.Split(s, "\n")[:n], "\n")
	return strings.Replace(result, "\r", "", -1)
}


func check(e error) {
    if e != nil {
        panic(e)
        os.Exit(1)
    }
}

func saveFile(contents []byte, path string) {
    d1 := []byte(contents)
    err := ioutil.WriteFile(path, d1, 0644)
    check(err)
}

func main() {

    //see https://medium.com/@nate510/don-t-use-go-s-default-http-client-4804cb19f779
    //also https://gist.github.com/ijt/950790/fca88967337b9371bb6f7155f3304b3ccbf3946f

    flag.Parse()
    args := flag.Args() //get trailing arguments after any flags
    url := args[0]      //url is the last argument

    
    connectTimeout := time.Second * time.Duration(*maxConnectTime)
    clientTimeout := time.Second * time.Duration(*maxDownloadTime)
    
    
    //create custom transport with timeout
    var netTransport = &http.Transport {
      Dial: (&net.Dialer {
        Timeout: connectTimeout,
      }).Dial,
      TLSHandshakeTimeout: connectTimeout,
    }

    
    //create custom client with timeout
    var netClient = &http.Client{
      Timeout: clientTimeout,
      Transport: netTransport,
    }
        
    
    //fmt.Println("making request")
    req, err := http.NewRequest("GET", url, nil)
    check(err)
    
    //set user agent if specified
    if (*userAgent != "") {
        req.Header.Add("User-Agent", *userAgent)
    }
    
    response, err := netClient.Do(req)
    check(err)

    //fmt.Println("done request")
    
    //final response (may have redirected)
    if (url != response.Request.URL.String()) {
        //notify of target location on stderr
        //see https://stackoverflow.com/questions/16784419/in-golang-how-to-determine-the-final-url-after-a-series-of-redirects
         fmt.Fprintln(os.Stderr, "redirected: " + response.Request.URL.String())
    }
    
    defer response.Body.Close()
    contents, err := ioutil.ReadAll(response.Body)
    check(err)
    
    
    //process the output
    if (*output == "-") || (*output == "") {
        fmt.Printf("%s", string(contents))
        
    } else {
        //save to the specified output
        httpBytes := []byte(contents)            
        saveFile(httpBytes, *output)
    }
        

}