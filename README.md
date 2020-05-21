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
scanner := bufio.NewScanner(os.Stdin)

for scanner.Scan() {
    fmt.Println(scanner.Text())
}

if err := scanner.Err(); err != nil {
    log.Println(err)
}
```