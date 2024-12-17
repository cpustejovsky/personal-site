**Originally Published 2020 October 21 on [Dev.To](https://dev.to/cpustejovsky/do-leetcode-assessments-make-sense-1kp6)** 

I've been a professional developer for a few years now. I've worked  on monoliths and micro-services, debugged legacy code and built greenfield projects. For all these different projects, my day-to-day has looked similar:

* I break down the problem with team members and stakeholders
* I work on the problem using my IDE of choice that I've configured to my liking
* When I run into issues, I turn to:
  * My team members
  * Google
  * StackOverlow
  * Reddit
  * [Gophers Slack](https://gophers.slack.com/messages/general/)
* I often use other people's code as a template or a starting point (even if it's just my past code)
* When appropriate, I import 3rd party libraries so I don't re-invent the wheel

As a back-end developer, these problems almost always relate to the following technologies:
* Servers (REST, gRPC)
* Databases (SQL, MongoDB, DynamoDB)
* Data Streams (Kafka, AWS Kinesis, RabbitMQ)

So when I interview for a back-end developer position, why are technical assessments often:
* closed book tests (no Google, etc.)
* based on data structures and algorithms
* using contrived problems
* using something other than my own text editor or IDE

Why would companies assess my technical skills with data structure and algorithim questions when the job is about databases and severs? It seems like misplaced priorities. To see what I mean, here's a chart showing how long various events would take if we said a CPU cycle takes 1 second and scaled from there.
![Chart of scaled latency](https://dev-to-uploads.s3.amazonaws.com/uploads/articles/rm9z83d196ojzgktbtnd.png)
Internet calls take much, much longer than operations using the CPU and main memory. Given that, why would companies focus on how well I optimize for "second" and "minute" operations and not for the operations that take "years"?

The answers to these questions vary. Some companies want to follow FAANG's example. Others aren't giving themselves the time to create and maintain a [quality technical interview](https://laurieontech.com/posts/interviews/). 

I've decided that I won't entertain interviews that use these kinds of assessments.

The questions in this post show a mismatch between the skills LeetCode assesses me for and the skills I'd be doing at the job.

I prefer doing deeper dives into Golang, SQL, DynamoDB, Kafka, and design patterns for distributed systems. I hone those skills because those will provide the most value for companies I work with.

This is risky. I reject a lot interviews as a result of this decision. I decrease my chances of getting a new job. But I've rarely met developers who like these kinds of assessments or think studying for them is a good use of their time. Most people study LeetCode because they have to, because need a job and most jobs have this as an obstacle. 

It's past time that we as an industry ask ourselves why we are using these kinds of technical assessments.

If you disagree with me, why do you think LeetCode style assessments are valuable, useful, etc.?

If you agree, what would be a better alternative? My suggestion would be to follow [Laurie Barth](https://twitter.com/laurieontech)'s guide for [Designing a Technical Interview](https://laurieontech.com/posts/interviews/)