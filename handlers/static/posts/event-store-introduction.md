**Originally Published 2023 January 14 on [Dev.To](https://dev.to/cpustejovsky/build-an-event-store-with-dynamodb-and-go-2j7a)** 

* [What is an Event Store for?](#EventStore)
* [DynamoDB](#DynamoDB)
  * [Partition Key](#Partition)
  * [Sort Key](#Sort)
* [Test-Driven Development](#TDD)

I first encountered DynamoDB at a company that used it like SQL. As a result, it ran SO SLOWLY. I didn't understand why anyone would use a non-relational database like this if it could be so easily misused and create such costly reads. Thankfully, I later had the opportunity to use DynamoDB to build an event store and see its power on display.

# What is an Event Store for?<a name="EventStore"></a>

I recently turned thirty. I could represent that in a database by updating an age  value from 29 to 30. But I could also represent it with an initialization event (when I was born) and thirty birthday events each year from 1993 to 2022.

That's a very simple case and wouldn't warrant an event store. However, you may want to consider an event store if you:

* need to manage state asynchronously
* need to be able to audit and observe changes in states
* want to use event-driven architecture

For a simple example, we'll say the Dungeon Master (DM) is a big stickler for the rules of Dungeons and Dragons and wants to make sure all changes to my player character's (PC) hit points (HP) are calculated correctly.

So we'll create an event that has my PC's name, changes to their HP, and a note about what caused the change. But we'll also need:
* a unique ID for these events to distinguish these events from any other events related to other PC's HP.
* a version to keep track of when each event happened.

Now let's move on to DynamoDB.

# DynamoDB <a name="DynamoDB"></a>
[AWS](https://aws.amazon.com/dynamodb/) describes DynamoDB as:
> a fully managed, serverless, key-value NoSQL database designed to run high-performance applications at any scale.

DynamoDB will preform well for you if you design your schema upfront following a [single table design](https://www.alexdebrie.com/posts/dynamodb-single-table/#what-is-single-table-design). But that's its own topic which [Alex Debrie](https://www.alexdebrie.com/) covers much better than I ever could.

Instead we'll focus on the fundamentals of DynamoDB by building a table that stores a specific kind of event.

The key to understanding DynamoDB is understanding its key. The primary key is the essential, required part of each item in the table. it contains a partition key and an optional sort key. Selecting good partition and sort keys is the most decision you'll make when setting up a DynamoDB table.

## Partition Key<a name="Partition"></a>

The partition key is a string, numeric, or binary value. If you don't take advantage of a sort key, you'll need this to be unique. Regardless, you'll want a partition key with as high a cardinality as possible.

An example of cardinatlity are dice in Dungeons and Dragons. 
![Golang gopher and TTRPG dice](https://dev-to-uploads.s3.amazonaws.com/uploads/articles/ubb0qexizxoh3mlis3g5.png)

A D6 die only has six sides and six possible outcomes while a D20 has 20 sides and 20 possible outcomes so a D20 would have a higher cardinality. 

With roughly 5.3×10<sup>36</sup> possibilities, A [uuid](https://en.wikipedia.org/wiki/Universally_unique_identifier)  is a great candidate both for a DynamoDB primary key and for our event store's id.

### Sort Key<a name="Sort"></a>
The optional sort key allows you to search and sort within items that match a given primary key. 

![TTRPG Dice in Order](https://dev-to-uploads.s3.amazonaws.com/uploads/articles/fp3caiuh5qlrf8pv0ogq.jpg)

Anything that can satisfy the `Interface` of Go's [sort package](https://pkg.go.dev/sort#Interface) can be a sort key. This includes strings, floats, and integers. 

For our event store, we'll use an incrementing integer to stand as the event's version.

# Test-Driven Development<a name="TDD"></a>

In the following parts of this series, we'll be building our DynamoDB event store in Go using test-driven development or TDD. The idea behind TDD is to use tests to guide our development. We use them as a specification for the behavior we want for a unit of code or the interaction between different pieces of code. We use the difficulty of tests as a helpful indicator of how rigid our tests may be. If something is hard to test, it's often also not a configurable or modular piece of code to begin with.

For more information on TDD in Go, I highly recommend [Learn Go with Tests](https://quii.gitbook.io/learn-go-with-tests/) by Chris James.

[![Golang Gophers Red Blue Green Refactor](https://dev-to-uploads.s3.amazonaws.com/uploads/articles/anl7r8xftww5tl294grk.png)](https://quii.gitbook.io/learn-go-with-tests/)

----

Ready to start building this event store? Check out [Part 2: Append and Query](https://dev.to/cpustejovsky/building-a-dynamodb-event-store-in-go-part-2-append-and-query-3a9c).

**Have any questions or comments? Is there anything I missed or got wrong about event store or DynamoDB? let me know in the comments.**

