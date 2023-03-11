# service-factory

## Overview

Simple service factory boilerplate to actuate your long-running structs' lifecycle. 

## Installation

```
go get github.com/nicholastcs/service-factory
```

## Recommended Usage

Below are the common use of this service factory implementation.

```go
package main

import (
    servicefactory "github.com/nicholastcs/service-factory"
)

func main() {
    // channel to listen for SIGTERM
    sig := make(chan os.Signal, 1)
    signal.Notify(sig, syscall.SIGTERM)

    // logger implements serviceFactory.Logger
    var logger = ...

    // services below implements servicefactory.Service interface
    // func Start() error {...} must be non-blocking.
    api := &ApiService{ ... }
    auth := &AuthService{ ... }
    payment := &PaymentService{ ... }

    // aggregrate into servicefactory.ServiceFactory, they implements
    // servicefactory.Service so, it is possible to have nested
    // service factories.
    factory := servicefactory.NewServiceFactory(logger)
    factory.Add("api", api)
    factory.Add("auth", auth)
    factory.Add("payment", payment)
    
    err := factory.Start()
    if err != nil {
        // log error then terminate app...
    }

    s := <-sig
    err := factory.Stop()
    if err != nil {
        // log error then terminate app...
    }
}
```

