<h1 align="center">
  <br>
  <a href="http://www.amitmerchant.com/electron-markdownify"><img src="./resources/images/logo.svg" alt="Markdownify" height="100"></a>
</h1>

<h3 align="center">A collection of go modules based on popular packages written in different languages (mostly JS)</h3><br/>

<!-- <p align="center">
  <a href="https://badge.fury.io/js/electron-markdownify">
    <img src="https://badge.fury.io/js/electron-markdownify.svg"
         alt="Gitter">
  </a>
  <a href="https://gitter.im/amitmerchant1990/electron-markdownify"><img src="https://badges.gitter.im/amitmerchant1990/electron-markdownify.svg"></a>
  <a href="https://saythanks.io/to/amitmerchant1990">
      <img src="https://img.shields.io/badge/SayThanks.io-%E2%98%BC-1EAEDB.svg">
  </a>
  <a href="https://www.paypal.me/AmitMerchant">
    <img src="https://img.shields.io/badge/$-donate-ff69b4.svg?maxAge=2592000&amp;style=flat">
  </a>
</p> -->

<p align="center">
  <a href="#key-features">Packages</a> •
  <a href="#how-to-use">How To Use</a> •
  <a href="#credits">Why</a> •
  <a href="#related">Related</a> •
  <a href="#license">License</a>
</p>

## Packages

- **find** - Find a file or directory by walking up parent directories - _Based on [sidresorhus's find-up](https://github.com/sindresorhus/find-up)_

## How To Use

You should be able to use this package just like any go module

Importing the **find** package

```go
package main

import (
    "github.com/ajhenry/go-wild/find"
    "fmt"
)

func main() {
    fmt.Println(find.Up(".git", find.Options{}))
}
```

## License

MIT