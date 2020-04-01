# Tago SDK para GO

## Inserindo dados

```
go run main.go --token="your_token" --payload="template.json"
```
ou (linux)

```
./bin/tago --token="your_token" --payload="template.json"
```

## Estrutura

* SDK
    * Admin
        * Device
        * Bucket
    * Device
        * Data
        * PubSub
            * Sub
            * Pub
            * Debug

## Coisas uteis

```go
q, _ := url.ParseQuery("")
q.Add("filter[tags][0][key]", "gateway")
q.Add("filter[tags][0][value]", "gw-01")

uri := &url.URL{
    Path:     "/device",
    RawQuery: q.Encode(),
}

fmt.Println(uri.String())
```