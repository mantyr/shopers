# Golang Shopers - replacer for Category name for shops

[![Build Status](https://travis-ci.org/mantyr/shopers.svg?branch=master)](https://travis-ci.org/mantyr/shopers) [![GoDoc](https://godoc.org/github.com/mantyr/shopers?status.png)](http://godoc.org/github.com/mantyr/shopers) [![Software License](https://img.shields.io/badge/license-The%20Not%20Free%20License,%20Commercial%20License-brightgreen.svg)](LICENSE.md)

This stable version

## Testing
```GO
type shoperTest struct{
    a []string
    b []string
}

var shoperTests = []shoperTest{
    {[]string{"shoes", "Basketball"},   []string{"Обувь", "Кроссовки баскетбольные"}},
    {[]string{"Обувь", "Basketball"},   []string{"Обувь", "Кроссовки баскетбольные"}},
    {[]string{"обувь", "Basketball"},   []string{"Обувь", "Кроссовки баскетбольные"}},
    {[]string{"clothes", "Basketball"}, []string{"Одежда", "Для баскетбола"}},
    {[]string{"Basketball"},            []string{"Баскетбол"}},
    {[]string{"Slides"},                []string{"Сланцы, балетки"}},
    {[]string{"обувь", "сланцы"},       []string{"Обувь", "Сланцы, балетки"}},
    {[]string{"обувь", "сланцы", "летние", "Красивые", "Великолепные"},  []string{"Обувь", "Сланцы, балетки", "Летние", "Красивые", "Великолепные балетки"}},
}
```

## Installation

    $ go get github.com/mantyr/shopers
    $ go get github.com/mantyr/conf
    $ go get github.com/mantyr/runner

## Example
```GO
package main

import (
    "github.com/mantyr/shopers"
    "github.com/mantyr/conf"
    "fmt"
)

func main() {
    conf.SetDefaultCatalog("./testdata")

    fields := NewShopers()
    fields.LoadConf("category_replace")
    fields.IsUcFirst(true)

    fmt.Println(fields.Get([]string{"обувь", "basketball"}))     // print []string{"Обувь", "Кроссовки баскетбольные"}
    fmt.Println(fields.Get([]string{"clothes", "Basketball"}))   // print []string{"Одежда", "Для баскетбола"}
}
```

## Author

[Oleg Shevelev][mantyr]

[mantyr]: https://github.com/mantyr
