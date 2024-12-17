package blog

import (
	"sort"
	"time"
)

var AllPosts = Posts{
	NonTechGuideToTechDebt,
	HelloWorld,
	LearnGoWithTestsReview,
	LetsGoReview,
	LeetCodeAssessmentsDontMakeSense,
	EventStoreIntroduction,
	EventStoreAppendAndQuery,
	EventStoreSnapshotsAndProjections,
}

var HelloWorld = Post{
	Title:       "Hello, World",
	Description: "First Blog Post",
	Date:        time.Date(2024, time.December, 10, 0, 0, 0, 0, time.UTC),
	FileName:    "hello-world",
}
var NonTechGuideToTechDebt = Post{
	Title:       "The Non-Technical Guide to Technical Debt",
	Description: "Guide for non-developers to grok tech debt",
	Date:        time.Date(2021, time.June, 13, 0, 0, 0, 0, time.UTC),
	FileName:    "non-technical-guide-to-technical-debt",
}
var LearnGoWithTestsReview = Post{
	Title:       "Learn Go with tests - Book Review",
	Description: "Review of Chris James' TDD book",
	Date:        time.Date(2020, time.September, 7, 0, 0, 0, 0, time.UTC),
	FileName:    "review-learn-go-with-tests",
}
var LetsGoReview = Post{
	Title:       "Let's Go - Book Review",
	Description: "Review of Alex Edwards' book of Go web development",
	Date:        time.Date(2020, time.September, 21, 0, 0, 0, 0, time.UTC),
	FileName:    "review-lets-go-book",
}
var LeetCodeAssessmentsDontMakeSense = Post{
	Title:       "LeetCode Assessments Don't Make Sense",
	Description: "Why I think LeetCode assessments aren't the best tool for interviews",
	Date:        time.Date(2022, time.October, 21, 0, 0, 0, 0, time.UTC),
	FileName:    "leetcode-assessments-dont-make-sense",
}
var EventStoreIntroduction = Post{
	Title:       "Building an Event Store with Go: Introduction",
	Description: "Introducing how to build an event store with Go and DynamoDB",
	Date:        time.Date(2023, time.January, 14, 0, 0, 0, 0, time.UTC),
	FileName:    "event-store-introduction",
}

var EventStoreAppendAndQuery = Post{
	Title:       "Building an Event Store with Go: Append and Query",
	Description: "Introducing how to build an event store with Go and DynamoDB",
	Date:        time.Date(2023, time.February, 2, 0, 0, 0, 0, time.UTC),
	FileName:    "event-store-append-and-query",
}

var EventStoreSnapshotsAndProjections = Post{
	Title:       "Building an Event Store with Go: Snapshots and Projections",
	Description: "Introducing how to build an event store with Go and DynamoDB",
	Date:        time.Date(2023, time.February, 25, 0, 0, 0, 0, time.UTC),
	FileName:    "event-store-snapshots-and-projections",
}

// Post is a max heap where the highest date is first
type Posts []Post

func (p Posts) Len() int {
	return len(p)
}

func (p Posts) Less(i, j int) bool {
	return p[i].Date.Unix() > p[j].Date.Unix()
}

func (p Posts) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func OrderedPosts(p Posts) Posts {
	sort.Sort(p)
	return p
}
