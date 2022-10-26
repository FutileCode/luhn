luhn package checks the validity of a number against the luhn algorithm

```
go get github.com/FutileCode/luhn
```
```
import (
  "github.com/FutileCode/luhn"
)
```
```
valid, reason := luhn.Valid("1234 5678 9876 5432")
```
