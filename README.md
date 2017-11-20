# Go Explore

Text based adventure game framework written in Go

## Getting Started

1. [Download and install Golang](https://golang.org/doc/install)
2. Set up and navigate to your `%GOPATH%` root
    1. [Follow the instructions here if you haven't done this before](https://golang.org/doc/install)
3. Run `go get github.com/lukemerrett/go-explore` to clone into `src/github.com/lukemerrett/go-explore`
4. CD into `src/github.com/lukemerrett/go-explore`
5. Run `go build`
6. Run `go-explore.exe`

## Configuring The Scenes

All the scenes and transitions are configured through a YAML file found in `resources/game-data.yaml`.

The format is:

```yaml
Scenes:
  Scene Key:
    Title: Title of Scene
    Body: Body Text of Scene
    Transitions:
      Scene Key: Text For Transition
```

For example:

```yaml
Scenes:
  Home:
    Title: Your home
    Body: A house on a street in a neighbourhood
    Transitions:
      Garden: Go to the garden
      Hallway: Go inside the house
  Hallway:
    Title: Hallway
    Body: The inside hallway of the house
    Transitions:
      Home: Go back outside
  Garden:
    Title: Your garden
    Body: The garden at the back of the house
    Transitions:
      Home: Go to the front door
```