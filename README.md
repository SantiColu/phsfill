# phsfill
Script to fill phishing sites databases with random data and prevent theft

## Installation & Build:
To install the library, run

```sh
go install -u  go get github.com/SantiColu/phsfill@lastest
```

## Usage:  
```sh
$ phsfill [OPTIONS]
```

```sh
-basic
    if it is true the data will be sended with basic-auth format
  
-delay int  
    delay between each request in milliseconds (default 1000)  
    
-form  
    if it is true the data will be sended as a form  
    
-requests int  
    number of requests (default 500)  
    
-type string  
    what action should the script do [valid options: login, register] (default "login")  
    
-url string  
    the scammer site url, where the fake data will be sended  
```

