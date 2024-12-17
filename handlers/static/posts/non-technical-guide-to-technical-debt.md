**Originally Published 2021 June 13 on [Dev.To](https://dev.to/cpustejovsky/the-non-technical-guide-to-technical-debt-271h)** 

Technical debt captured my interest the moment I learned about it. I was learning about programming while working in marketing at a startup with technical debt. I saw the cost of technical debt not only for developers but also for marketing, sales, product, and customer success teams.

As a result, I developed a passion for writing clean code and fixing problems where I saw them. I also started thinking of a way to explain this issue to non-developers. Developers feel the pain of technical debt up close and personal. We understand it like the burned hand understands the hot stove. Since technical debt affects the whole software company, everyone should better understand it.

## What is Technical Debt?

![Duck Tape Fix for Car Side Mirror](https://dev-to-uploads.s3.amazonaws.com/uploads/articles/nuanpl9je23aqez415p3.JPG)

Gene Kim's [The Unicorn Project](https://itrevolution.com/the-unicorn-project/) defined explains technical debt like this:
> Every half-measure and cut corner during the life of a piece of software adds to technical debt. When you solve the immediate technical problem without worrying about the future, you add technical debt. 

It's important to realize the "debt" is a debt of time and experience which is very hard to pay back. After talking with enough non-developers, I've come to agree with David Thomas and Andrew Hunt's term for it: "software rot".

Like a neglected tooth can decay or a shoddy house can fall apart, software with too much technical debt can be like a rotten piece of wood.

## Reasons Companies Go Into Technical Debt

If software rot is bad, why would any company go into it? This is why the term "technical debt" is useful. A company goes into technical debt for the same reasons they may take on financial debt.

A software company wants to get a feature or product out before the competition. They want to stay on track for a go-to market plan. A small team can often only make this deadline by cutting some corners and duct taping solutions.

This leads to software rot building up. Because the team is still small it's hard to fix the software rot **and** add new features. Since the code is currently good enough, there's hope that the code will remain that way.

![Homer Simpson Fingers Crossed](https://dev-to-uploads.s3.amazonaws.com/uploads/articles/yz2dm9nt0265rrxluqa4.jpg)

## The Costs of Technical Debt

Intense tooth pain can mean a rotten tooth and a root canal. In the same way, the pain developers feel around adding features or updating software points to a real problem with real costs for a company.

I'll illustrate that problem with some Lego blocks and a knitted hat.

### Knitted Hat

Here is a hat my wife knitted a while ago:

![Knitted hat made by my wife](https://dev-to-uploads.s3.amazonaws.com/uploads/articles/6jcc5grp8exw79qkievm.jpg)
When I asked her how long it would take to replace the purple with royal blue, she said, "I would need... well... I'd need to start over."

For her, it would be easier to use the current hat as a reference and knit a new one. To untangle the knitted hat and replace the thread color would take even longer.

This is why developers often call bad code "spaghetti code". It's code so tangled that it's easier to look at what it's trying to do and rewrite it from scratch.

### Lego Blocks

Here is a structure I made from Lego blocks yesterday:

![Structure Made from Legos; Has multiple colors and a small figurine of Gandalf](https://dev-to-uploads.s3.amazonaws.com/uploads/articles/r0dxpo5ghe5e0wmx1atw.jpg)

It would be easy to change out a color, make it taller, make it wider, or even replace all the foundational blocks.

That is because each Lego brick is it's own thing or **component**. Talk of components, modularity, orthogonality, etc. are all about how to make code like Lego bricks.

### Back to Code

This metaphor breaks down because  I can build something with Lego blocks faster than my wife can knit a hat. Lego block code, however, takes far more time and effort to build than spaghetti code.

When you need a new feature or a change in functionality, you want Lego blocks. You want things you can plug in, remove, combine, and change without hassle. 

But spaghetti code can cause the following symptoms:
* Features taking longer than expected
* Problems no one can diagnose
* Changes breaking unrelated parts of the software.

![Technical Debt Comic by @vincentdnl](https://dev-to-uploads.s3.amazonaws.com/uploads/articles/5b191on4vbio2ok9ot4b.png)

## How to Get Out of Technical Debt

Since technical debt is a debt of time, it takes intentional plans to pay it down. 

Technical debt will never feel as urgent as adding a requested feature or fixing a severe bug. I recommend software companies dedicate X% of a sprint to cleaning up software rot. Don't let that X% decrease unless the whole system is on fire. 

But feature requests don't come from the ether, they come from clients. What are sales and product teams supposed to tell customers who are impatient for feature X? The keywords I would use are **speed** and **reliability**. Untangling spaghetti code and cleaning software rot will almost always make code faster. It will always make bugs easier to spot. In my experience, unexpected errors and lag grind the gears of clients and their users more than a missing feature.

But ultimately, the best way to get out of technical debt is to work together. Together, technology and product teams can balance present needs with long-term stability.

Developers, do you have any analogies or metaphors to explain technical debt? 

Product/project managers, what would help you understand developer's concerns about technical debt?

Leave a comment!