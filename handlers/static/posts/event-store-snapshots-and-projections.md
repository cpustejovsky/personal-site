**Originally Published 2023 February 25 on [Dev.To](https://dev.to/cpustejovsky/snapshots-and-projections-df)** 

In [Part 2](/blog/event-store-append-and-query), we built `Append` and `Query` functionality for our event store which left us with this list of events:
```
store.Event{Id:"58f02691-78ed-4ca5-8e59-8f4deb44e063", Version:0, CharacterName:"cpustejovsky", CharacterHitPoints:8, Note:"Init"}
store.Event{Id:"58f02691-78ed-4ca5-8e59-8f4deb44e063", Version:1, CharacterName:"cpustejovsky", CharacterHitPoints:-2, Note:"Slashing damage from goblin"}
store.Event{Id:"58f02691-78ed-4ca5-8e59-8f4deb44e063", Version:2, CharacterName:"cpustejovsky", CharacterHitPoints:-3, Note:"bludgeoning damage from bugbear"}
```
To make use of the long list of events returned from a Query and to manage the increasingly long list of events, we'll be building functionality to `Project` and `Snapshot` our events. 

## Projection

We could manually take the `CharacterHitPoints` for each event in the list and add them together to get the current hit points for our character. (`8 + -2 + -3 = 3`).

To make this easier to use, we can create a projection. A projection takes the stream of events we have and projects it to be read. In this case, you can imagine the value being projected to a UI to show the current hit points of the player character. `Project` will take the `Event` id and return an `Event` with a reconstituted state. Let's first write our test:

```go
t.Run("Project Events from Event Store", func(t *testing.T) {
    aggEvent, err := es.Project(ctx, id)
    require.Nil(t, err)
    assert.Equal(t, 8, aggEvent.CharacterHitPoints)
    assert.Equal(t, "cpustejovsky", aggEvent.CharacterName)
})
```

With a failing test, we can start on making it pass. We'll need to query our event store and range through the events, reconstituting them into a single state.

```go
// Project takes an id, queries events since the last snapshot, and returns a reconstituted Event
func (es *EventStore) Project(ctx context.Context, id string) (*Event, error) {
    var agg Event
	//Query events
	events, err := es.QueryAll(ctx, id)
	if err != nil {
		return nil, err
	}
	//Reconstitute the event
	for i, event := range events {
		if i == len(events)-1 {
			agg.Id = event.Id
			agg.CharacterName = event.CharacterName
			agg.Version = event.Version + 1
		}
		agg.CharacterHitPoints += event.CharacterHitPoints
	}
	return &agg, nil
}
```

Now our tests should be passing. As a refactor, we can split up the reconstitution logic from the `Project` method:

```go
// Project takes an id, queries events since the last snapshot, and returns an aggregated Event
func (es *EventStore) Project(ctx context.Context, id string) (*Event, error) {
	events, err := es.QueryAll(ctx, id)
	if err != nil {
		return nil, err
	}
	agg := reconstituteEvent(events)
	return &agg, nil
}

func reconstituteEvent(events []Event) Event {
	var agg Event
	for i, event := range events {
		if i == len(events)-1 {
			agg.Id = event.Id
			agg.CharacterName = event.CharacterName
			agg.Version = event.Version + 1
		}
		agg.CharacterHitPoints += event.CharacterHitPoints
	}
	return agg
}
```

## Snapshot

As we have more DnD sessions with more and more changes to the hit points of our player character, we will run into trouble handling the growing volume of events. To help with this, we can take snapshots of the state at points in our event log. This will mean the Event Store won't need to query every single event whenever we project it.

To begin with, let's write our test. Our `Snapshot` method will take a context and a reconstituted event `e` and return an error:

```go
t.Run("Snapshot should return no error", func(t *testing.T) {  
   e, err := es.Project(ctx, id)  
   assert.Nil(t, err)  
   err = es.Snapshot(ctx, e)  
   assert.Nil(t, err)  
})
```

There are a variety of ways we could implement this functionality, but for now I'm sticking with DynamoDB's single table design and using the same table for the snapshots as for the standard events. As a result, our `Snapshot` method is using the `Append` method to add the snapshot event to:

```go
// using a constant instead of a hardcoded value for it to be easily used in multiple places 
const SnapshotValue string = "SNAPSHOT"
//...
func (es *EventStore) Snapshot(ctx context.Context, agg *Event) error {  
   e := &Event{  
      Id:                 agg.Id,  
      Version:            agg.Version,  
      CharacterName:      agg.CharacterName,  
      CharacterHitPoints: agg.CharacterHitPoints,  
      Note:               SnapshotValue,  
   }  
   return es.Append(ctx, e)  
}
```

With the `Snapshot` test passing, our next step should be to see if our `QueryAll` still works. Let's write the test:

```go
t.Run("QueryAll should not return the snapshot", func(t *testing.T) {  
   queriedEvents, err := es.QueryAll(ctx, id)  
   assert.Nil(t, err)  
   assert.Equal(t, len(events), len(queriedEvents))  
   for _, event := range events {  
      assert.NotEqual(t, store.SnapshotValue, event.Note)  
   }  
})
```

This test is failing because we haven't modified our `QueryAll` to skip over snapshots. We can make that modification to the DynamoDB parameters to get the test working:

```go
func (es *EventStore) QueryAll(ctx context.Context, id string) ([]Event, error) {
params := dynamodb.QueryInput{  
   TableName:              aws.String(es.Table),  
   KeyConditionExpression: aws.String("Id = :uuid"),  
   FilterExpression:       aws.String("Note <> :note"),  
   ExpressionAttributeValues: map[string]types.AttributeValue{  
      ":uuid": &types.AttributeValueMemberS{Value: id},  
	  ":note": &types.AttributeValueMemberS{Value: SnapshotValue},
   },  
}
//...
```

## Assessing

You can take a look at our current code as of Part 3 [here](https://github.com/cpustejovsky/event-store-template/tree/part-3). There is plenty of functionality to add, but I want to address two key problems over the next two parts.

First, the event store can only keep track of a Player Character's hit points. We need to abstract the `Event` to be able to represent a myriad of different events.

Second, we've tied our Event Store to a DynamoDB client. This is fine for our  integration tests but will burden any unit tests for code relying on our Event Store but not focusing on the Event  Store's behavior. To fix that, we'll be using interfaces and mocks.

**Have any questions or comments? Is there anything I missed or got wrong about event projections or snapshots? Let me know in the comments or reach out to me on [Twitter](https://twitter.com/CCPustejovsky) or the [Gopher's Slack](https://invite.slack.golangbridge.org/).**