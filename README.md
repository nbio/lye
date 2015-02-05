# lye

`go get github.com/nbio/lye`

Minimal precursor for [SOAP](http://en.wikipedia.org/wiki/SOAP). Extracted from and used in production at [domainr.com](https://domainr.com).

## Why?

1. [Because](http://www.theatlantic.com/technology/archive/2013/11/english-has-a-new-preposition-because-internet/281601/) [XML](http://harmful.cat-v.org/software/xml/).
2. “It’s only after we’ve lost everything that we’re free to do anything.” — [Tyler Durden](http://genius.com/David-fincher-fight-club-chemical-burn-scene-annotated)

## Usage

```go
req := struct {
	XMLName  xml.Name `xml:"https://example.com/soapnamespace Command"`
	Question string   `xml:"Nested>Question"`
}{
	Question: "Who am I?",
}
x, err := lye.NewRequest(req).Marshal()
if err != nil {
	return err
}
res, err := http.Post("https://api.example.com/", lye.ContentType, bytes.NewBuffer(x))
if err != nil {
	return err
}
```

## TODO

- [ ] Tests
- [ ] Examples
- [ ] Responses?

## Author

© 2015 nb.io, LLC
