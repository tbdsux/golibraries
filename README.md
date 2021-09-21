# golibraries

An API wrapper for Libraries.io (https://libraries.io/api)

## Usage

```go
package main

import (
	"fmt"
	"log"
	"os"

	"github.com/TheBoringDude/golibraries"
)

func main() {
	l := golibraries.New(os.Getenv("API_KEY"))

	q, err := l.ProjectSearch("fastapi", golibraries.ProjectSearchOptions{
		Sort: golibraries.SortRank,
	})

	if err != nil {
		log.Fatal(err)
	}

	for _, v := range q {
		fmt.Println(v.Name)
	}
}

```

##

#### &copy; 2021 | [LICENSE](./LICENSE)
