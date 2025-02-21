filtered stream example
===

This is sample code that uses the `filteredstream` package.

# Steps to use Filtered Stream API

## 1. Create search stream rules.

Suppose you want to retrieve streams filtered by the keyword "Twitter API v2", create a search stream rule with the following code.

```go
func CreateSearchStreanmRules(c *gotwitter.GoTwitter) {
	p := &types.CreateRulesInput{
		Add: []types.AddingRule{
			{Value: String("Twitter API v2"), Tag: String("example rule")},
		},
	}

	res, err := filteredstream.CreateRules(context.TODO(), c, p)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	for _, r := range res.Data {
		fmt.Printf("ID: %s, Value: %s, Tag: %s\n", StringValue(r.ID), StringValue(r.Value), StringValue(r.Tag))
	}
}
```

## 2. Search stream

To retrieve the streams to which you have applied the rules (filters) you have created, implement as follows.

```go
func SearchStream(c *gotwi.GoTwitter) {
	p := &types.SearchStreamInput{}
	s, err := filteredstream.SearchStream(context.Background(), c, p)
	if err != nil {
		fmt.Println(err)
		return
	}

	cnt := 0
	for s.Receive() {
		cnt++
		t, err := s.Read()
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(StringValue(t.Data.ID), StringValue(t.Data.Text))
		}

		if cnt > 10 {
			s.Stop()
			break
		}
	}
}
```

# Run example code

1. Create a search stream rule.

    ```bash
    go run . create 'Twitter'
    ```
    
2. Call filtered stream API.

    ```bash
    go run . stream
    ```
    
3. List search stream rules.

    ```bash
    go run . list
    ```
    
4. Delete a search stream rule.

    ```bash
    go run . delete rule-id
    ```