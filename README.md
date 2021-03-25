# gqlgen-add-bson-mappings

If you are using the marvellous gqlgen from https://github.com/99designs/gqlgen in combination with mongodb you may have had the trouble of mapping the mongodb `_id` field to an ID in your graphQL.

There are other things like `omitempty` that would be nice to add to some fields. Through the plugin API of gqlgen it is possible to add this functionality. 

This repository contains a tiny demo how how you can add `bson: ".."` style qualifiers to the generated models automatically and specifically for certain models and fields. 

## How to use this repo

Clone it, run `go get` and then execute the `regenerate.sh` example. This *should* work with no problems. Check the `gql_models.go` file which contains:

```
type ExampleType1 struct {
	ID          *string `json:"id" bson:"_id"`
	Name        *string `json:"name"`
	Description *string `json:"description" bson:"omitempty"`
}

type ExampleType2 struct {
	ID          *string `json:"id" bson:"_id"`
	Name        *string `json:"name"`
	Description *string `json:"description"`
	Something   *string `json:"something" bson:"omitempty"`
}
```

The interesting bits are that because of the `bson:"_id"` part, now the ID is automatically marshalled and unmarshalled properly into the struct from mongodb bson. To demonstrate some random fields have `omitempty` added. 

## How it works

In the generate folder you find the `gql_generate.go` file which uses gqlgen to generate, but before that, it parses a mappings JSON file. This contains definitions like this:

```
[
  { "model": "*", "field": "id", "tagPostfix": " bson:\"_id\"" },
  {
    "model": "ExampleType1",
    "field": "description",
    "tagPostfix": " bson:\"omitempty\""
  },
  {
    "model": "ExampleType2",
    "field": "something",
    "tagPostfix": " bson:\"omitempty\""
  }
]
```
Where:
- `model` specifies a specific model or `*` can be used to target all
- `field` specifies the field
- `tagPrefix` specifies what to append to the tag

This is the file you would be editing to easily specify some of the things you need for your bson mappings. 

## Usage

Well, use the code as you need it in your project. This may serve as a basis to solve other problems too, it could be extended to allow for regex matching etc, but it serves the main purpose for now and this repo should get you up and running with that approach. 

We at qubidu.com want to contribute back to the awesome dev world. If something takes us a while to find a solution for and it turns out there are only partial answers, we try to share it with you, once we make it work. 

Thanks to the awesome work of gqlgen, our favorite go schema-first graphql lib.




