package education

import (
	"html/template"
)

type CTA struct {
	Text string
	URL  template.HTML
}

type ContinuingEducation struct {
	Finished   bool
	Recurring  bool
	Title      string
	Link       string
	Paragraphs []template.HTML
	CTAs       []CTA
}

type ContinuingEducationCategories struct {
	Recurring []ContinuingEducation
	Current   []ContinuingEducation
	Completed []ContinuingEducation
}

func GetContinuingEducationCategories(list []ContinuingEducation) ContinuingEducationCategories {
	var c ContinuingEducationCategories
	for _, item := range list {
		if item.Recurring {
			c.Recurring = append(c.Recurring, item)
			continue
		}
		if item.Finished {
			c.Completed = append(c.Completed, item)
		} else {
			c.Current = append(c.Current, item)
		}
	}
	return c
}

var ContinuingEducationList = []ContinuingEducation{
	//goBook,
	Proofs,
	MITOCW,
	//LittleSchemer,
	dddGo,
	//insideTheMachine,
	kafka,
	dds,
	tddGo,
	concurrencyGo,
	pragProg,
	DesertFathers,
	letsGo,
}

var Proofs = ContinuingEducation{
	Finished:  false,
	Recurring: false,
	Title:     "Proofs: A Long-Form Mathematics Textbook",
	Link:      "https://longformmath.com/proofs-home",
	Paragraphs: []template.HTML{
		`Jay Cummings wrote Proofs with the goal of making mathematical education more accessible. His step-by-step examples help readers grok set theory, induction, proofs, relations, etc.`,
		`Proofs provides an excellent foundation for a computer science education.`,
	},
	CTAs: []CTA{
		{
			Text: "Buy the Book",
			URL:  "https://www.amazon.com/Proofs-Long-Form-Mathematics-Textbook-Math/dp/B08T8JCVF1/ref=sr_1_1?dchild=1&keywords=proofs+a+long-form+mathematics+textbook&sr=8-1",
		},
	},
}

var MITOCW = ContinuingEducation{
	Finished:  false,
	Recurring: false,
	Title:     "MIT Opencourseware",
	Link:      "https://ocw.mit.edu/",
	Paragraphs: []template.HTML{
		`I'm auditing MIT 6.006 (Introduction to Algorithms) to fill in my computer science gaps. The projects that excite me need developers who grok data structures and algorithms.  That knowledge helps us build complex, efficient applications.`,
	},
	CTAs: []CTA{
		{
			Text: "MIT 6.OO6",
			URL:  "https://ocw.mit.edu/courses/6-006-introduction-to-algorithms-spring-2020/",
		},
	},
}

var LittleSchemer = ContinuingEducation{
	Finished:  false,
	Recurring: false,
	Title:     "The Little Schemer",
	Link:      "https://mitpress.mit.edu/9780262560993/the-little-schemer/",
	Paragraphs: []template.HTML{
		`I've always loved Functional Programming (FP). Since <a href="https://www.youtube.com/watch?v=5buaPyJ0XeQ">Go has first class functions</a>,
		grokking FP by learning Scheme is helpful for me as a Go developer.`,
	},
	CTAs: []CTA{},
}

var goBook = ContinuingEducation{
	Finished:  false,
	Recurring: false,
	Title:     "The Go Programming Language",
	Link:      "https://www.gopl.io/",
	Paragraphs: []template.HTML{
		`Brian Kernighan has already written [one classic programming
book](https: //www.amazon.com/Programming-Language-2nd-Brian-Kernighan/dp/0131103628), so I believe I'm in
good hands.`,
	},
}

var letsGo = ContinuingEducation{
	Finished:  true,
	Recurring: false,
	Title:     "Let's Go",
	Link:      "https://lets-go.alexedwards.net/",
	Paragraphs: []template.HTML{
		`This book is absolutely wonderful for any newcomer to Go wanting to dive into web development.`,
		`Alex Edwards shows you how to build scalable, composable, maintainable backends with Go.`,
	},
	CTAs: []CTA{
		{
			Text: "Review", URL: "https://dev.to/cpustejovsky/let-s-go-book-review-1909",
		},
	},
}
var pragProg = ContinuingEducation{
	Finished:  true,
	Recurring: false,
	Title:     "The Pragmatic Programmer",
	Link:      "https://pragprog.com/book/tpp20/the-pragmatic-programmer-20th-anniversary-edition",
	Paragraphs: []template.HTML{
		`This book is filled with wisdom and best practices that any programmer can use to improve their craft and
		better provide value with the software they help create.Now that I've finished reading it, I'm slowly going
		back through it to really grok its lessons.`,
	},
	//CTAs: []CTA{
	//	{
	//		Text: "Review TODO",
	//		URL:  "foobar",
	//	},
	//},
}
var DesertFathers = ContinuingEducation{
	Finished:  true,
	Recurring: true,
	Title:     "The Alphabetical Sayings of the Desert Fathers",
	Link:      "https://svspress.com/give-me-a-word-the-alphabetical-sayings-of-the-desert-fathers/",
	Paragraphs: []template.HTML{
		`
		The Desert Fathers of Christianity hold a special place in my heart.Their lessons of self-discipline,
		humility, perseverance, and not judging others help me not only as a developer, but as a human being.It is
		why I read some of their sayings at the beginning of each day, to center and focus me for the day ahead,
		with all the challenges it may bring.
		`,
	},
}

var tddGo = ContinuingEducation{
	Finished:  true,
	Recurring: false,
	Title:     "Learn Go with tests",
	Link:      "https://quii.gitbook.io/learn-go-with-tests/",
	Paragraphs: []template.HTML{
		`This is the book that taught me Test-Driven Development (TDD). Chris James begins with the basics of Go. 
He then goes step-by-step through the process of building a fully-functioning application. 
All throughout he teaches test-driven development. He points to best practices and anti-patterns to avoid. 
It remains my go-to recommendation for anyone interested in Go or TDD.`,
	},
	CTAs: []CTA{
		{
			Text: "Review", URL: "https://dev.to/cpustejovsky/learn-go-with-tests-book-review-na4",
		},
	},
}

var dds = ContinuingEducation{
	Finished:  true,
	Recurring: false,
	Title:     "Designing Distributed Systems",
	Link:      "https://www.oreilly.com/library/view/designing-distributed-systems/9781491983638",
	Paragraphs: []template.HTML{
		`
	Helping build an event driven distributed systems showed me many of the gaps I have and infrastructure
	conText I lacked. I decided to go through this book to remedy that.
	`,
		`
	So far, it has been a wonderful and practical survey of various ways to use containers and orchestration to
	build a variety of systems.
	`,
	},
	//CTAs: []CTA{
	//	{
	//		Text: "Review TODO",
	//		URL:  "foobar",
	//	},
	//},
}

var kafka = ContinuingEducation{
	Finished:  true,
	Recurring: false,
	Title:     "Kafka: The Definitive Guide",
	Link:      "https://www.oreilly.com/library/view/kafka-the-definitive/9781492043072/",
	Paragraphs: []template.HTML{
		`
	Kafka is a powerful tool with a host of challenges in store for the team that decides to use it.
	`,
		`
	Having now worked with Kafka for a little less than a year, I realize the need to have a deeper
	understanding of the technology.
	`,
		`
	Even if you're using something like Confluent to manage Kafka for you, this book provides a great foundation
	for how to build and maintain high performance and reliable Kafka producers and consumers.
	`,
	},
	//CTAs: []CTA{
	//	{
	//		Text: "Review TODO",
	//		URL:  "foobar",
	//	},
	//},
}

var concurrencyGo = ContinuingEducation{
	Finished:  true,
	Recurring: false,
	Title:     "Concurrency in Go",
	Link:      "https://www.oreilly.com/library/view/concurrency-in-go/9781491941294/",
	Paragraphs: []template.HTML{
		`
	Go is my favorite language for many reasons, but chief is how it handles concurrency. Given that, I thought
	it wise to dive deeper into that.
	`,
		`
	Katherine Cox-Buday has written an amazing introduction into Go's concurrency primitives (goroutines and
	channels), libraries (<code>sync</code> and <code>conText</code>), and best practices.
	`,
		`This is definitely a book I'll be rereading in the future.`,
	},
	//CTAs: []CTA{
	//	{
	//		Text: "Review TODO",
	//		URL:  "foobar",
	//	},
	//},
}

var dddGo = ContinuingEducation{
	Finished:  true,
	Recurring: false,
	Title:     "Domain Driven Design with Golang",
	Link:      "https://www.packtpub.com/product/domain-driven-design-with-golang/9781804613450",
	Paragraphs: []template.HTML{
		`
	Working at Groundfloor introduced me to Domain Driven Design (DDD) in 2022. 
	I wanted to dig deeper. Since Eric Evans wrote the Big Blue Book with classes in mind, a DDD book written specifically for Go developers interested me. 
	Matt Boyle's book does just that, translating DDD concepts into idiomatic Go. `,
	},
	CTAs: []CTA{
		{
			Text: "Review",
			URL:  "https://www.amazon.com/review/R1XEYMP7U6S1AA/ref=cm_cr_srp_d_rdp_perm?ie=UTF8",
		},
	},
}

var insideTheMachine = ContinuingEducation{
	Finished:  false,
	Recurring: false,
	Title:     "Inside the Machine",
	Link:      "https://nostarch.com/insidemachine.htm",
	Paragraphs: []template.HTML{
		`
TEXT GOES HERE
`,
	},
}
