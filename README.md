# libattio

[![Go Reference](https://pkg.go.dev/badge/github.com/sid6mathur/libattio/attio.svg)](https://pkg.go.dev/github.com/sid6mathur/libattio/attio)
[![Regenerate](https://github.com/sid6mathur/libattio/actions/workflows/regenerate.yml/badge.svg)](https://github.com/sid6mathur/libattio/actions/workflows/regenerate.yml)

Unofficial Go client for the [Attio CRM REST API](https://docs.attio.com/rest-api), generated with
[openapi-generator](https://openapi-generator.tech) from Attio's [public OpenAPI spec](https://api.attio.com/openapi/api).
A daily job regenerates it and tags a new minor version when anything changed.

```sh
go get github.com/sid6mathur/libattio/attio
```

```go
client := attio.NewClient(os.Getenv("ATTIO_API_KEY"))

lists, _, err := client.ListsAPI.V2ListsGet(ctx).Execute()

person, _, err := client.RecordsAPI.V2ObjectsObjectRecordsPut(ctx, "people").
	MatchingAttribute("email_addresses").
	V2ObjectsObjectRecordsPutRequest(attio.V2ObjectsObjectRecordsPutRequest{
		Data: attio.V2ObjectsObjectRecordsPutRequestData{
			Values: map[string][]any{"email_addresses": {"ada@example.com"}},
		},
	}).Execute()
```

Method and type names follow the endpoint paths, because the spec has no operation IDs.
Types named `...AnyOfN` are numbered by their position in the spec and can be renumbered
when Attio reorders it; for write values, prefer plain maps or your own structs.

## Layout

| Path | What |
| --- | --- |
| [`attio/`](attio) | The Go package. All files except `attio.go` are generated. |
| [`spec/`](spec) | `openapi.json`, Attio's spec after the fixes in `patch.jq`. |
| [`Makefile`](Makefile) | `make` downloads the spec, regenerates `attio/`, and runs the tests. Needs curl, jq, Java and Go. |

## License

[BSD 2-Clause](LICENSE).
