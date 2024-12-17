**Originally Published 2023 February 2 on [Dev.To](https://dev.to/cpustejovsky/building-a-dynamodb-event-store-in-go-part-2-append-and-query-3a9c)** 

* [Setup](#Setup)
* [Append](#Append)
* [Query](#Query)

In [Part One](/blog/event-store-introduction), we went over the basics of event stores, DynamoDB, and Test-Driven Development. Now, we'll build our event store and its most basic functionality.

# Setup<a name="Setup"></a>

First, we'll need to [create an AWS Account](https://portal.aws.amazon.com/billing/signup) and [get your access key](https://docs.aws.amazon.com/IAM/latest/UserGuide/id_credentials_access-keys.html) (For better security, follow AWS's advice and [set up AWS credentials in your local environment](https://docs.aws.amazon.com/cli/latest/userguide/cli-configure-files.html)).

Store your access keys in a `.env` file that is included in your `.gitignore`.

Next we'll set up our AWS configuration and create a test for our Event Store.

_(**NOTE**: I am using the `require` and `assert` packages from the [testify](https://github.com/stretchr/testify) repo.)_

```go
var EventStoreTable = "event-store"

func TestEventStore(t *testing.T) {
    //Create Dynamodb Client
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()
    
    cfg, err := config.LoadDefaultConfig(ctx,
        config.WithRegion(os.Getenv("AWS_REGION")),
        config.WithEndpointResolverWithOptions(
            aws.EndpointResolverWithOptionsFunc(
                func(service, region string, options ...interface{}) (aws.Endpoint, error) {
                    return aws.Endpoint{}, &aws.EndpointNotFoundError{}
                },
            ),
        ),
        config.WithCredentialsProvider(
            credentials.StaticCredentialsProvider{
                Value: aws.Credentials{
                    AccessKeyID:     os.Getenv("AWS_ACCESS_ID"),
                    SecretAccessKey: os.Getenv("AWS_SECRET_KEY"),
                },
            },
        ),
    )
    require.Nil(t, err)
    
    //Create Event Store
    client := dynamodb.NewFromConfig(cfg, EventStoreTable)
    es := store.New(client)
    require.NotNil(t, es)
}
```

Once we have a failing test we can write the code needed to make this pass, in this case, the `Event Store` type along with a function constructor:

```go
type EventStore struct {
	DB    *dynamodb.Client
	Table string
}

func New(db *dynamodb.Client, table string) *EventStore {
	return &EventStore{DB: db, Table: table}
}
```

Different environments will likely have different DynamonDB clients and table names. To handle this, we configure the event store to take the table name and DynamoDB client as arguments.

## Event

We described the shape of the event we wanted in Part 1. We'll have that as a struct that holds an id, version, the name of the character, the change to the character's hit points, and a note about what caused this change.

```go
type Event struct {
    Id                      string
    Version                 int
    CharacterName           string
    CharacterHitPoints      int
    Note                    string
}
```

The two fundamental methods of our event store are `Append` and `Query`. We append new events to the event store and query those event from the event store.

# Append<a name="Append"></a>
We first write our test for the event store's `Append` method. Based on DynamoDB's `PutItem` function [definition](https://github.com/aws/aws-sdk-go-v2/blob/0e9721fd7e662a5a0c554249a2e96519258be35b/service/dynamodb/api_op_PutItem.go#L35), our `Append` will need to pass in a `context.Context` along with the event we want to append. This test will also give us three events to query for later on.

```go
//...
id := uuid.NewString()
events := []store.Event{
    {
        Id:                      id,
        Version:                 0,
        CharacterName:           "cpustejovsky",
        CharacterHitPoints: 8,
        Note:                    "Init",
    },
    {
        Id:                      id,
        Version:                 1,
        CharacterName:           "cpustejovsky",
        CharacterHitPoints: -2,
        Note:                    "Slashing damage from goblin",
    },
    {
        Id:                      id,
        Version:                 2,
        CharacterName:           "cpustejovsky",
        CharacterHitPoints: -3,
        Note:                    "bludgeoning damage from bugbear",
    },
}
t.Run("Append Items to Event Store", func(t *testing.T) {
    for _, event := range events {
        err := es.Append(context.Background(), &event)
        if err != nil {
            t.Fatal(err)    
        }
    }
})
```

Now we write an append method to satisfy our test:

```go
func (es *EventStore) Append(ctx context.Context, event Event) error {
	input := &dynamodb.PutItemInput{
		TableName: &es.Table,
		Item: map[string]types.AttributeValue{
            "Id": &types.AttributeValueMemberS{Value: event.Id},
            //AttributeValueMemberN takes a string value, see https://docs.aws.amazon.com/amazondynamodb/latest/APIReference/API_AttributeValue.html
            "Version":                 &types.AttributeValueMemberN{Value: strconv.Itoa(event.Version)},
            "CharacterName":           &types.AttributeValueMemberS{Value: event.CharacterName},
            "CharacterHitPoints": &types.AttributeValueMemberN{Value: strconv.Itoa(event.CharacterHitPoints)},
            "Note":                    &types.AttributeValueMemberS{Value: event.Note},
        },
	}
	_, err := es.DB.PutItem(ctx, input)
	if err != nil {
		return err
	}
	return nil
}
```

Our test is passing but this `Append` method will not work for our event store. We need to have [optimistic concurrency](https://en.wikipedia.org/wiki/Optimistic_concurrency_control) to make sure events can't be overwritten. As it stands now, a new Version 0 event will overwrite the existing event.

Let's first write the test:

```go
t.Run("Attempt to append existing version to event store and fail", func(t *testing.T) {
	e := store.Event{
		Id:      id,
		Version: 0,
	}
	err := es.Append(context.Background(), e)
	assert.NotNil(t, err)
})
```

This test will fail with our current Append method, so we'll need to add a conditional to our DynamoDB input:

```go
cond := "attribute_not_exists(Version)"
input := &dynamodb.PutItemInput{
	TableName: &es.Table,
	Item: map[string]types.AttributeValue{
		"Id": &types.AttributeValueMemberS{Value: event.Id},
		"Version":            &types.AttributeValueMemberN{Value: strconv.Itoa(event.Version)},
		"CharacterName":      &types.AttributeValueMemberS{Value: event.Character.Name},
		"CharacterHitPoints": &types.AttributeValueMemberN{Value: strconv.Itoa(event.Character.HitPoints)},
        "Note":               &types.AttributeValueMemberS{Value: event.Note},
},
	ConditionExpression: &cond,
}
```

Now our test for Append's failure is passing but our Append tests are failing. That's because those three events are already saved to the database. To clean that up, we can write a `t.Cleanup` function at the bottom of our tests to delete those records and reset the state for us between tests.

```go
//...
	t.Cleanup(func() {
		for i := 0; i < len(events); i++ {
			params := dynamodb.DeleteItemInput{
				TableName: &EventStoreTable,
				Key: map[string]types.AttributeValue{
					"Id":      &types.AttributeValueMemberS{Value: id},
					"Version": &types.AttributeValueMemberN{Value: strconv.Itoa(i)},
				},
			}
			_, err := client.DeleteItem(context.Background(), &params)
			if err != nil {
				t.Log("Error deleting items for cleanup:\t", err)
			}
		}
	})
}
```

Now all our tests are passing, but we still have a problem. How will our users distinguish between failures? Our `Append` method could fail because the Event Version already exists. It could also fail if there is an internal problem with DynamoDB.

To provide clarity, we can create a sentinel error if the condition failed and have our test check for that. First we can create the error:

```go
type EventAlreadyExistsError struct {
	ID      string ``
	Version int
}

func (e *EventAlreadyExistsError) Error() string {
	return fmt.Sprintf("event already exists for ID %s and Version %d", e.ID, e.Version)
}
```

Then we add it to our test:

```go
t.Run("Attempt to append existing version to event store and fail", func(t *testing.T) {
	e := store.Event{
		Id:      id,
		Version: 0,
	}
	err := es.Append(context.Background(), e)
	assert.NotNil(t, err)
	checkErr := &store.EventAlreadyExistsError{}
	assert.True(t, errors.As(err, &checkErr))
})
```

And our tests are now failing. To get them passing, we'll need to add the following error handling into our `Append` method:

```go
if err != nil {
	//Using the errors package, the code checks if this is an error specific to the condition being failed and, if so, returns a sentinel error that can be checked
	var errCheck *types.ConditionalCheckFailedException
	if errors.As(err, &errCheck) {
		return &EventAlreadyExistsError{
			ID:      event.Id,
			Version: event.Version,
		}
	}
	return err
}
```

Now our tests are passing again. We have created half of our essential functionality and can move on.

# Query<a name="Query"></a>
Since we set up a state of three events in our event store, we can write a test for our Event Store to query for them. We know it will need to take an Event ID to query and, like `Append`, it will need a `context.Context` to satisfy the DynamoDB API:
```go
t.Run("Query Items from Event Store", func(t *testing.T) {  
   queriedEvents, err := es.Query(ctx, id)  
   assert.Nil(t, err)  
   assert.Equal(t, len(events), len(queriedEvents))  
   for _, event := range events {
       assert.Contains(t, queriedEvents, event)
   }
})
```

Our tests are failing, and we can begin work on making them pass:
```go
// Query takes a context and DynamoDB query parameters and returns a slice of Events and an errorfunc (es *EventStore) Query(ctx context.Context, id string) ([]Event, error) {  
   var events []Event  
   kce := "Id = :uuid"  
   params := &dynamodb.QueryInput{  
      TableName:              &es.Table,  
      KeyConditionExpression: &kce,  
      ExpressionAttributeValues: map[string]types.AttributeValue{  
         ":uuid": &types.AttributeValueMemberS{Value: id},  
      },  
   }  
   // Query paginator provides pagination for queries until there are no more pages for DynamoDB to go through  
   // See: https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/Query.Pagination.htm   p := dynamodb.NewQueryPaginator(es.DB, params)  
   for p.HasMorePages() {  
      out, err := p.NextPage(ctx)  
      if err != nil {  
         return nil, err  
      }  
      // The output is unmarshalled into an Event slice which is appended to the events slice  
      err = attributevalue.UnmarshalListOfMaps(out.Items, &events)  
      if err != nil {  
         return nil, err  
      }  
   }  
   // If the slice is empty, then error is returned  
   if len(events) < 1 {  
      return nil, errors.New("no events found")  
   }  
   return events, nil  
}
```

Most of this code is specific to DynamoDB. In particular, [pagination](https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/Query.Pagination.html) is essential to any queries or scans. They ensure DynamoDB will return more than 1 MB of data back if there is more data than that to return.

At a high level, our Event Store's `Query` method is:
* querying the underlying database
* unmarshalling its items to the `Event` type
* ensuring that at least one `Event` was returned.

Similar to `Append`, we could add a sentinel error here to help a user differentiate between no event being found and an internal DynamoDB error.

First, write the test:

```go
t.Run("Query returns specific error if no Event is found", func(t *testing.T) {
    _, err := es.Query(ctx, uuid.NewString())
    assert.NotNil(t, err)
    checkErr := &store.NoEventFoundError{}
    assert.True(t, errors.As(err, &checkErr))
})
```

To get this test passing, we'll create the error and replace the current `errrors.New()` value with it:

```go
type NoEventFoundError struct{}

func (e *NoEventFoundError) Error() string {
    return "no event found"
}
//...
if len(events) < 1 {
    return nil, &NoEventFoundError{}
}
```

With that we have created the basic functionality of our Event Store. To see what comes next, lets log the output of our query:

```
store.Event{Id:"58f02691-78ed-4ca5-8e59-8f4deb44e063", Version:0, CharacterName:"cpustejovsky", CharacterHitPoints:8, Note:"Init"}
store.Event{Id:"58f02691-78ed-4ca5-8e59-8f4deb44e063", Version:1, CharacterName:"cpustejovsky", CharacterHitPoints:-2, Note:"Slashing damage from goblin"}
store.Event{Id:"58f02691-78ed-4ca5-8e59-8f4deb44e063", Version:2, CharacterName:"cpustejovsky", CharacterHitPoints:-3, Note:"bludgeoning damage from bugbear"}
```

That's not particular helpful and even if it were helpful. We can manually add the 8, -2, and -3 together to conclude that the Character "cpustejovsky" is at 3 hit points. And even if we had the code do that for us, it would take longer and longer to query all the events the more sessions of D&D played.

We will be solving those two issues in the next part when we tackle Projection and Snapshotting.

**Have any questions or comments? Is there anything I missed or got wrong about event stores or DynamoDB? Let me know in the comments or reach out to me on [Twitter](https://twitter.com/CCPustejovsky) or the [Gopher's Slack](https://invite.slack.golangbridge.org/).**