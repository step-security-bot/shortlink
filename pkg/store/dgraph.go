package store

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/batazor/shortlink/pkg/link"
	"github.com/dgraph-io/dgo/v2"
	"github.com/dgraph-io/dgo/v2/protos/api"
	"google.golang.org/grpc"
	"log"
)

type DGraphLink struct {
	Uid string `json:"uid,omitempty"`
	link.Link
	DType []string `json:"dgraph.type,omitempty"`
}

type DGraphLinkResponse struct {
	Link []struct {
		link.Link
		Uid string
	}
}

type DGraphLinkList struct {
	client *dgo.Dgraph
}

func (dg *DGraphLinkList) Init() error {
	conn, err := grpc.Dial("localhost:9080", grpc.WithInsecure())
	if err != nil {
		return err
	}
	dg.client = dgo.NewDgraphClient(api.NewDgraphClient(conn))

	if err = dg.migration(); err != nil {
		return err
	}

	return nil
}

func (dg *DGraphLinkList) get(id string) (*DGraphLinkResponse, error) {
	ctx := context.Background()
	txn := dg.client.NewTxn()
	defer txn.Discard(ctx)

	q := `
query all($a: string) {
	link(func: eq(hash, $a)) {
		uid
		url
		hash
		describe
	}
}`

	val, err := txn.QueryWithVars(ctx, q, map[string]string{"$a": id})
	if err != nil {
		fmt.Println("Err: ", err)
		return nil, &link.NotFoundError{Link: link.Link{Url: id}, Err: errors.New(fmt.Sprintf("Not found id: %s", id))}
	}

	var response DGraphLinkResponse

	if err = json.Unmarshal(val.Json, &response); err != nil {
		fmt.Println("Err: ", err)
		return nil, &link.NotFoundError{Link: link.Link{Url: id}, Err: errors.New(fmt.Sprintf("Failed parse link: %s", id))}
	}

	return &response, nil
}

func (dg *DGraphLinkList) Get(id string) (*link.Link, error) {
	ctx := context.Background()
	txn := dg.client.NewTxn()
	defer txn.Discard(ctx)

	response, err := dg.get(id)
	if err != nil {
		return nil, &link.NotFoundError{Link: link.Link{Url: id}, Err: errors.New(fmt.Sprintf("Not found id: %s", id))}
	}

	if len(response.Link) == 0 {
		return nil, &link.NotFoundError{Link: link.Link{Url: id}, Err: errors.New(fmt.Sprintf("Not found id: %s", id))}
	}

	if response.Link[0].Url == "" {
		return nil, &link.NotFoundError{Link: link.Link{Url: id}, Err: errors.New(fmt.Sprintf("Not found id: %s", id))}
	}

	return &response.Link[0].Link, nil
}

func (dg *DGraphLinkList) Add(data link.Link) (*link.Link, error) {
	hash := data.GetHash([]byte(data.Url), []byte("secret"))
	data.Hash = hash[:7]

	ctx := context.Background()
	txn := dg.client.NewTxn()
	defer txn.Discard(ctx)

	item := DGraphLink{
		Uid:   fmt.Sprintf(`_:%s`, data.Hash),
		Link:  data,
		DType: []string{"Link"},
	}

	pb, err := json.Marshal(item)
	if err != nil {
		log.Fatal(err)
	}

	mu := &api.Mutation{
		SetJson:   pb,
		CommitNow: true,
		// TODO: Add condition
		//Cond: `@if(eq(len(hash), 1))`,
		//SetNquads: []byte(fmt.Sprintf(`uid(hash) <hash> "%s" .`, data.Hash)),
	}
	_, err = txn.Mutate(ctx, mu)
	if err != nil {
		return nil, &link.NotFoundError{Link: data, Err: errors.New(fmt.Sprintf("Failed save link: %s", data.Url))}
	}

	return &data, nil
}

func (dg *DGraphLinkList) Update(data link.Link) (*link.Link, error) {
	return nil, nil
}

func (dg *DGraphLinkList) Delete(id string) error {
	ctx := context.Background()
	txn := dg.client.NewTxn()
	defer txn.Discard(ctx)

	links, err := dg.get(id)
	if err != nil {
		return &link.NotFoundError{Link: link.Link{Url: id}, Err: errors.New(fmt.Sprintf("Not found id: %s", id))}
	}

	if len(links.Link) == 0 {
		return nil
	}

	mu := &api.Mutation{
		CommitNow: true,
	}
	for _, link := range links.Link {
		dgo.DeleteEdges(mu, link.Uid, "hash")
	}

	_, err = txn.Mutate(ctx, mu)
	if err != nil {
		fmt.Println(err)
		return &link.NotFoundError{Link: link.Link{Url: id}, Err: errors.New(fmt.Sprintf("Not found id: %s", id))}
	}

	return nil
}

func (dg *DGraphLinkList) migration() error {
	ctx := context.Background()
	txn := dg.client.NewTxn()
	defer txn.Discard(ctx)

	op := &api.Operation{
		Schema: `
type Link {
    url: string
    hash: string
    describe: string
}

hash: string @index(term) @lang .
`,
	}

	return dg.client.Alter(ctx, op)
}
