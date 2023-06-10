package education

import (
	"html/template"
)

type Link struct {
	Text string
	URL  template.HTML
}

type ContinuingEducation struct {
	Finished   bool
	Recurring  bool
	Title      string
	Paragraphs []template.HTML
	Links      []Link
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
	goBook,
	MITOCW,
	LittleSchemer,
	dddGo,
	insideTheMachine,
	kafka,
	dds,
	ultimateServiceGo,
	tddGo,
	concurrencyGo,
	pragProg,
	scrum,
	DesertFathers,
	letsGo,
}

var MITOCW = ContinuingEducation{
	Finished:  false,
	Recurring: false,
	Title:     "MIT Opencourseware",
	Paragraphs: []template.HTML{
		`I'm filling in the gaps of my Liberal Arts education by auditing MIT Computer Science
						courses.`,
		`I <a target="_blank" rel="noreferrer noopener"  href="https://dev.
to/cpustejovsky/do-leetcode-assessments-make-sense-1kp6">am skeptical of LeetCode assessments </a> for developers. However, 
I appreciate the need for senior developers to understand Data Structures and Algorithms. 
That knowledge helps us build complex applications than run efficiently.`,
	},
	Links: []Link{
		{
			Text: "Mathematics For Computer Science (MIT 6.042J)",
			URL:  "https://ocw.mit.edu/courses/6-042j-mathematics-for-computer-science-fall-2010/",
		},
		{
			Text: "Introduction to Algorithms (MIT 6.006)",
			URL:  "https://ocw.mit.edu/courses/6-006-introduction-to-algorithms-spring-2020/",
		},
	},
}

var LittleSchemer = ContinuingEducation{
	Finished:  false,
	Recurring: false,
	Title:     "The Little Schemer",
	Paragraphs: []template.HTML{
		`I've always loved Functional Programming (FP). Since <a href="https://www.youtube.com/watch?v=5buaPyJ0XeQ">Go has first class functions</a>,
		grokking FP by learning Scheme is helpful for me as a Go developer.`,
	},
	Links: []Link{
		{
			Text: "Visit the Book's Site",
			URL:  "https://mitpress.mit.edu/9780262560993/the-little-schemer/",
		},
	},
}

var goBook = ContinuingEducation{
	Finished:  false,
	Recurring: false,
	Title:     "The Go Programming Language",
	Paragraphs: []template.HTML{
		`Brian Kernighan has already written [one classic programming
book](https: //www.amazon.com/Programming-Language-2nd-Brian-Kernighan/dp/0131103628), so I believe I'm in
good hands.`,
	},
	Links: []Link{
		{
			Text: "Visit the Book's Site", URL: "https://www.gopl.io/",
		},
	},
}

var letsGo = ContinuingEducation{
	Finished:  true,
	Recurring: false,
	Title:     "Let's Go",
	Paragraphs: []template.HTML{
		`
		This book is absolutely wonderful for any newcomer to Go wanting to dive into web development.
		`,
		`
		Alex Edwards shows you how to build scalable, composable, maintainable backends with Go.
		`,
	},
	Links: []Link{
		{
			Text: "Full Review on DEV.to", URL: "https://dev.to/cpustejovsky/let-s-go-book-review-1909",
		},
		{
			Text: "Buy on Alex Edward's Website", URL: "https://lets-go.alexedwards.net/",
		},
	},
}
var pragProg = ContinuingEducation{
	Finished:  true,
	Recurring: false,
	Title:     "The Pragmatic Programmer",
	Paragraphs: []template.HTML{
		`This book is filled with wisdom and best practices that any programmer can use to improve their craft and
		better provide value with the software they help create.Now that I've finished reading it, I'm slowly going
		back through it to really grok its lessons.`,
	},
	Links: []Link{
		{
			Text: "Buy on the Pragmatic Bookshelf", URL: "https://pragprog.com/book/tpp20/the-pragmatic-programmer-20th-anniversary-edition",
		},
	},
}
var DesertFathers = ContinuingEducation{
	Finished:  true,
	Recurring: true,
	Title:     "The Alphabetical Sayings of the Desert Fathers",
	Paragraphs: []template.HTML{
		`
		The Desert Fathers of Christianity hold a special place in my heart.Their lessons of self-discipline,
		humility, perseverance, and not judging others help me not only as a developer, but as a human being.It is
		why I read some of their sayings at the beginning of each day, to center and focus me for the day ahead,
		with all the challenges it may bring.
		`,
	},
	Links: []Link{
		{
			Text: "Buy on Amazon", URL: "https://amazon.com/Sayings-Desert-Fathers-Alphabetical-Collection/dp/0879079592",
		},
	},
}

var scrum = ContinuingEducation{
	Finished:  true,
	Recurring: false,
	Title:     "Learn Go with tests",
	Paragraphs: []template.HTML{
		"I've never been at company that used Agile/Scrum and, as a result, have been able to see first-hand the issues that can arise from not following a system like this.",
		`Jeff Sutherland does not only an amazing job of explaining the "what" and "how" of Scrum, 
but also the "why". Through anecdotes and philosophical asides, 
he lays a foundation for why Scrum can and will help any team be more effective.`,
		"It's very easy to read and is almost certainly worth reading multiple times. There are short summaries at the end of each chapter and an appendix for someone looking to implement Scrum for their team. It's an excellent book that anyone, but especially those working as developers, product managers, and project managers, should read.",
	},
	Links: []Link{
		{
			Text: "Buy on Amazon", URL: "https://www.amazon.com/gp/product/B00JI54HCU/ref=ppx_yo_dt_b_search_asin_title",
		},
	},
}
var tddGo = ContinuingEducation{
	Finished:  true,
	Recurring: false,
	Title:     "Learn Go with tests",
	Paragraphs: []template.HTML{
		"I believe both Golang and TDD are excellent tools for writing scalable, maintainable code so it made sense to improve my Golang skills while also getting into the habit of doing test-driven development.",
	},
	Links: []Link{
		{
			Text: "Full Review on DEV.to", URL: "https://dev.to/cpustejovsky/learn-go-with-tests-book-review-na4",
		},
		{
			Text: "Read on GitBook", URL: "https://quii.gitbook.io/learn-go-with-tests/",
		},
	},
}

var ultimateServiceGo = ContinuingEducation{
	Finished:  true,
	Recurring: false,
	Title:     "Ardan Labs Ultimate Service 3.0",
	Paragraphs: []template.HTML{
		"A friend recommended this course to me and I enjoyed learn package driven design and idiomatic Go patterns for microservice architecture",
	},
	Links: []Link{
		{
			Text: "Check out Ardan Labs", URL: "https://www.ardanlabs.com/",
		},
	},
}

var dds = ContinuingEducation{
	Finished:  true,
	Recurring: false,
	Title:     "Designing Distributed Systems",
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
	Links: []Link{
		{
			Text: "Check it out on O'Reilly Media",
			URL:  "https://www.oreilly.com/library/view/designing-distributed-systems/9781491983638",
		},
	},
}

var kafka = ContinuingEducation{
	Finished:  true,
	Recurring: false,
	Title:     "Kafka: The Definitive Guide",
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
	Links: []Link{
		{
			Text: "Check it out on O'Reilly Media",
			URL:  "https://www.oreilly.com/library/view/kafka-the-definitive/9781492043072/",
		},
	},
}

var concurrencyGo = ContinuingEducation{
	Finished:  true,
	Recurring: false,
	Title:     "Concurrency in Go",
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
	Links: []Link{
		{
			Text: "Check it out on O'Reilly Media",
			URL:  "https://www.oreilly.com/library/view/concurrency-in-go/9781491941294/",
		},
	},
}

var dddGo = ContinuingEducation{
	Finished:  false,
	Recurring: false,
	Title:     "Domain Driven Design with Golang",
	Paragraphs: []template.HTML{
		`
My work with Groundfloor in 2022 introduced me to Domain Driven Design (DDD).
I appreciate DDD's focus on collaboration between the domain experts and the developers.
`,
		`
After I saw <a target="_blank" rel="noreferrer noopener" href="https://quii.dev/">Chris James</a> recommend
this book on Twitter, I decided to give it a read. Eric Evan's original
<a target="_blank" rel="noreferrer noopener" href="https://www.domainlanguage.com/ddd/blue-book/">blue book</a>
on DDD assumes a class-based paradigm. Given that, I appreciate Matt Boyle's translation of those patterns for
Go's paradigms.
`,
	},
	Links: []Link{
		{
			Text: "Visit the Book's Site",
			URL:  "https://www.packtpub.com/product/domain-driven-design-with-golang/9781804613450",
		},
	},
}

var insideTheMachine = ContinuingEducation{
	Finished:  false,
	Recurring: false,
	Title:     "Inside the Machine",
	Paragraphs: []template.HTML{
		`
TEXT GOES HERE
`,
	},
	Links: []Link{
		{
			Text: "Visit the Book's Site",
			URL:  "https://nostarch.com/insidemachine.htm",
		},
	},
}
