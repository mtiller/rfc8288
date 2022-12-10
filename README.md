### Disclaimer

This repository is effectively a fork of
https://github.com/tniswong/go.rfcx/tree/master/rfc8288 by
[`tniswong`](https://github.com/tniswong). That respository hadn't been touched
in many years. So I forked it so I could upgrade the code to leverage Go modules
and in the process I've switched the testing framework involved to
[Testify](https://github.com/stretchr/testify). I have, of course, retained
Tim's copyright in the [`LICENSE.md`](./LICENSE.md) file to reflect that the
majority of the work here was his.

## Usage

To use the library, do:

```
go get github.com/mtiller/rfc8288
```

### Parsing Links

```go
result, err := ParseLink(`<https://www.google.com>; rel="next"; hreflang="en"; title="title"; title*="title*"; type="type"; extension="value"`)
```

In this case, you'll get a `Link` structure back that will look like this:

```go
Link{
    HREF:          url.URL{
        Scheme: "https",
        Host: "www.google.com",
    },
    Rel:           "next",
    HREFLang:      "en",
    Title:         "title",
    TitleStar:     "title*",
    Type:          "type",
    extensionKeys: []string{"extension"},
    extensions: map[string]interface{}{
        "extension": "value",
    },
},
```

### Constructing Link

According to RFC8288, certain keys are reserved. These are all part of the
`Link` structure so you can create a link and directly set their values, _e.g._:

```go
Link{
    HREF:   url.URL{
        Scheme: "https",
        Host: "www.google.com",
    },
    Rel:       "rel",
    HREFLang:  "hreflang",
    Media:     "media",
    Title:     "title",
    TitleStar: "title*",
    Type:      "type",
}
```

If you want to add a non-standard key, you need to use the `Extends` method on `Link`, _e.g.,_:

```go
l.Extend("extension", "value")
```

It is an **error** to use `Extend` to add a reserved key.

### Link Header Values

To serialize a `Link` instance into the value used by
[RFC8288](https://www.rfc-editor.org/rfc/rfc8288.html) Link Headers,
just use the `String()` method on the `Link` type, _e.g._,

```go
link := Link{
    HREF: parseURL("https://www.google.com", t),
    Type: "type",
}
fmt.Printfln(link.String())
```

This will output:

```
<https://www.google.com>; type="type"
```

Note, this it the link header **value**, not the link header itself.

### JSON

The `Link` structure defines special `MarshalJSON` and `Links can be marshaled
into JSON and`UnmarshalJSON`JSON methods so you can use the standard`json`package to marshal and unmarshal`Link` instances.
