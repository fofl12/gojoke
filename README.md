# gojoke
Lightweight unofficial Go wrapper for [JokeAPI](https://v2.jokeapi.dev/).

Basic usage:
```go
import (
  "github.com/snoo8/gojoke"
  "fmt"
)

func main() {
  joke, err := gojoke.Any()
  
  if err != nil {
    panic(err)
  }
  
  fmt.Println(joke)
}
```

## Reference
### `gojoke.Any () -> (joke string, error)`
Fetch any joke
```go
gojoke.Any() // > The generation of random numbers is too important to be left to chance.
```

### `gojoke.Safe () -> (joke string, error)`
Fetch a joke safe for everyone
```go
gojoke.Safe() // > What do you call a cow with no legs? ...
```

### `gojoke.Language (language string) -> (joke string, error)`
Fetch any joke in a certain language
```go
gojoke.Language("de") // > Warst du eigentlich schonmal in Las Vegas? ...
```

### `gojoke.Category (category string) -> (joke string, error)`
Fetch a joke in a certain category
```go
gojoke.Category("Pun") // > There are only 10 kinds of people in ...
```

### `gojoke.Joke (language string, category string, blacklist []string, single bool, twopart bool, safe bool) -> (joke string, error)`
Advanced joke search
| Argument    | Description                              | Example            |
|-------------|------------------------------------------|--------------------|
| `language`  | The language of the joke                 | "en"               |
| `category`  | The category of the joke                 | "Pun"              |
| `blacklist` | Blacklist flags                          | "religious","nsfw" |
| `single`    | If the type of the joke can be "single"  | true               |
| `twopart`   | If the type of the joke can be "twopart" | false              |
| `safe`      | If the joke must be safe for everyone    | true               |
