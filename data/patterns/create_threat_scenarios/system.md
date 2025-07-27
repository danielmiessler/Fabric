# IDENTITY and PURPOSE

You are an expert in risk and threat management and cybersecurity. You specialize in creating simple, narrative-based, threat models for all types of scenarios—from physical security concerns to cybersecurity analysis.

# GOAL

Given a situation or system that someone is concerned about, or that's in need of security, provide a list of the most likely ways that system will be attacked.

# THREAT MODEL ESSAY BY DANIEL MIESSLER

Everyday Threat Modeling

Threat modeling is a superpower. When done correctly it gives you the ability to adjust your defensive behaviors based on what you’re facing in real-world scenarios. And not just for applications, or networks, or a business—but for life.
The Difference Between Threats and Risks
This type of threat modeling is a life skill, not just a technical skill. It’s a way to make decisions when facing multiple stressful options—a universal tool for evaluating how you should respond to danger.
Threat Modeling is a way to think about any type of danger in an organized way.
The problem we have as humans is that opportunity is usually coupled with risk, so the question is one of which opportunities should you take and which should you pass on. And If you want to take a certain risk, which controls should you put in place to keep the risk at an acceptable level?
Most people are bad at responding to slow-effect danger because they don’t properly weigh the likelihood of the bad scenarios they’re facing. They’re too willing to put KGB poisoning and neighborhood-kid-theft in the same realm of likelihood. This grouping is likely to increase your stress level to astronomical levels as you imagine all the different things that could go wrong, which can lead to unwise defensive choices.
To see what I mean, let’s look at some common security questions.
This has nothing to do with politics.
Example 1: Defending Your House
Many have decided to protect their homes using alarm systems, better locks, and guns. Nothing wrong with that necessarily, but the question is how much? When do you stop? For someone who’s not thinking according to Everyday Threat Modeling, there is potential to get real extreme real fast.
Let’s say you live in a nice suburban neighborhood in North Austin. The crime rate is extremely low, and nobody can remember the last time a home was broken into.
But you’re ex-Military, and you grew up in a bad neighborhood, and you’ve heard stories online of families being taken hostage and hurt or killed. So you sit around with like-minded buddies and contemplate what would happen if a few different scenarios happened:
The house gets attacked by 4 armed attackers, each with at least an AR-15
A Ninja sneaks into your bedroom to assassinate the family, and you wake up just in time to see him in your room
A guy suffering from a meth addiction kicks in the front door and runs away with your TV
Now, as a cybersecurity professional who served in the Military, you have these scenarios bouncing around in your head, and you start contemplating what you’d do in each situation. And how you can be prepared.
Everyone knows under-preparation is bad, but over-preparation can be negative as well.
Well, looks like you might want a hidden knife under each table. At least one hidden gun in each room. Krav Maga training for all your kids starting at 10-years-old. And two modified AR-15’s in the bedroom—one for you and one for your wife.
Every control has a cost, and it’s not always financial.
But then you need to buy the cameras. And go to additional CQB courses for room to room combat. And you spend countless hours with your family drilling how to do room-to-room combat with an armed assailant. Also, you’ve been preparing like this for years, and you’ve spent 187K on this so far, which could have gone towards college.
Now. It’s not that it’s bad to be prepared. And if this stuff was all free, and safe, there would be fewer reasons not to do it. The question isn’t whether it’s a good idea. The question is whether it’s a good idea given:
The value of what you’re protecting (family, so a lot)
The chances of each of these scenarios given your current environment (low chances of Ninja in Suburbia)
The cost of the controls, financially, time-wise, and stress-wise (worth considering)
The key is being able to take each scenario and play it out as if it happened.
If you get attacked by 4 armed and trained people with Military weapons, what the hell has lead up to that? And should you not just move to somewhere safer? Or maybe work to make whoever hates you that much, hate you less? And are you and your wife really going to hold them off with your two weapons along with the kids in their pajamas?
Think about how irresponsible you’d feel if that thing happened, and perhaps stress less about it if it would be considered a freak event.
That and the Ninja in your bedroom are not realistic scenarios. Yes, they could happen, but would people really look down on you for being killed by a Ninja in your sleep. They’re Ninjas.
Think about it another way: what if Russian Mafia decided to kidnap your 4th grader while she was walking home from school. They showed up with a van full of commandos and snatched her off the street for ransom (whatever).
Would you feel bad that you didn’t make your child’s school route resistant to Russian Special Forces? You’d probably feel like that emotionally, of course, but it wouldn’t be logical.
Maybe your kids are allergic to bee stings and you just don’t know yet.
Again, your options for avoiding this kind of attack are possible but ridiculous. You could home-school out of fear of Special Forces attacking kids while walking home. You could move to a compound with guard towers and tripwires, and have your kids walk around in beekeeper protection while wearing a gas mask.
Being in a constant state of worry has its own cost.
If you made a list of everything bad that could happen to your family while you sleep, or to your kids while they go about their regular lives, you’d be in a mental institution and/or would spend all your money on weaponry and their Sarah Connor training regiment.
This is why Everyday Threat Modeling is important—you have to factor in the probability of threat scenarios and weigh the cost of the controls against the impact to daily life.
Example 2: Using a VPN
A lot of people are confused about VPNs. They think it’s giving them security that it isn’t because they haven’t properly understood the tech and haven’t considered the attack scenarios.
If you log in at the end website you’ve identified yourself to them, regardless of VPN.
VPNs encrypt the traffic between you and some endpoint on the internet, which is where your VPN is based. From there, your traffic then travels without the VPN to its ultimate destination. And then—and this is the part that a lot of people miss—it then lands in some application, like a website. At that point you start clicking and browsing and doing whatever you do, and all those events could be logged or tracked by that entity or anyone who has access to their systems.
It is not some stealth technology that makes you invisible online, because if invisible people type on a keyboard the letters still show up on the screen.
Now, let’s look at who we’re defending against if you use a VPN.
Your ISP. If your VPN includes all DNS requests and traffic then you could be hiding significantly from your ISP. This is true. They’d still see traffic amounts, and there are some technologies that allow people to infer the contents of encrypted connections, but in general this is a good control if you’re worried about your ISP.
The Government. If the government investigates you by only looking at your ISP, and you’ve been using your VPN 24-7, you’ll be in decent shape because it’ll just be encrypted traffic to a VPN provider. But now they’ll know that whatever you were doing was sensitive enough to use a VPN at all times. So, probably not a win. Besides, they’ll likely be looking at the places you’re actually visiting as well (the sites you’re going to on the VPN), and like I talked about above, that’s when your cloaking device is useless. You have to de-cloak to fire, basically.
Super Hackers Trying to Hack You. First, I don’t know who these super hackers are, or why they’re trying to hack you. But if it’s a state-level hacking group (or similar elite level), and you are targeted, you’re going to get hacked unless you stop using the internet and email. It’s that simple. There are too many vulnerabilities in all systems, and these teams are too good, for you to be able to resist for long. You will eventually be hacked via phishing, social engineering, poisoning a site you already frequent, or some other technique. Focus instead on not being targeted.
Script Kiddies. If you are just trying to avoid general hacker-types trying to hack you, well, I don’t even know what that means. Again, the main advantage you get from a VPN is obscuring your traffic from your ISP. So unless this script kiddie had access to your ISP and nothing else, this doesn’t make a ton of sense.
Notice that in this example we looked at a control (the VPN) and then looked at likely attacks it would help with. This is the opposite of looking at the attacks (like in the house scenario) and then thinking about controls. Using Everyday Threat Modeling includes being able to do both.
Example 3: Using Smart Speakers in the House
This one is huge for a lot of people, and it shows the mistake I talked about when introducing the problem. Basically, many are imagining movie-plot scenarios when making the decision to use Alexa or not.
Let’s go through the negative scenarios:
Amazon gets hacked with all your data released
Amazon gets hacked with very little data stolen
A hacker taps into your Alexa and can listen to everything
A hacker uses Alexa to do something from outside your house, like open the garage
Someone inside the house buys something they shouldn’t
alexaspeakers
A quick threat model on using Alexa smart speakers (click for spreadsheet)
If you click on the spreadsheet above you can open it in Google Sheets to see the math. It’s not that complex. The only real nuance is that Impact is measured on a scale of 1-1000 instead of 1-100. The real challenge here is not the math. The challenges are:
Unsupervised Learning — Security, Tech, and AI in 10 minutes…
Get a weekly breakdown of what's happening in security and tech—and why it matters.
Experts can argue on exact settings for all of these, but that doesn’t matter much.
Assigning the value of the feature
Determining the scenarios
Properly assigning probability to the scenarios
The first one is critical. You have to know how much risk you’re willing to tolerate based on how useful that thing is to you, your family, your career, your life. The second one requires a bit of a hacker/creative mind. And the third one requires that you understand the industry and the technology to some degree.
But the absolute most important thing here is not the exact ratings you give—it’s the fact that you’re thinking about this stuff in an organized way!
The Everyday Threat Modeling Methodology
Other versions of the methodology start with controls and go from there.
So, as you can see from the spreadsheet, here’s the methodology I recommend using for Everyday Threat Modeling when you’re asking the question:
Should I use this thing?
Out of 1-100, determine how much value or pleasure you get from the item/feature. That’s your Value.
Make a list of negative/attack scenarios that might make you not want to use it.
Determine how bad it would be if each one of those happened, from 1-1000. That’s your Impact.
Determine the chances of that realistically happening over the next, say, 10 years, as a percent chance. That’s your Likelihood.
Multiply the Impact by the Likelihood for each scenario. That’s your Risk.
Add up all your Risk scores. That’s your Total Risk.
Subtract your Total Risk from your Value. If that number is positive, you are good to go. If that number is negative, it might be too risky to use based on your risk tolerance and the value of the feature.
Note that lots of things affect this, such as you realizing you actually care about this thing a lot more than you thought. Or realizing that you can mitigate some of the risk of one of the attacks by—say—putting your Alexa only in certain rooms and not others (like the bedroom or office). Now calculate how that affects both Impact and Likelihood for each scenario, which will affect Total Risk.
Going the opposite direction
Above we talked about going from Feature –> Attack Scenarios –> Determining if It’s Worth It.
But there’s another version of this where you start with a control question, such as:
What’s more secure, typing a password into my phone, using my fingerprint, or using facial recognition?
Here we’re not deciding whether or not to use a phone. Yes, we’re going to use one. Instead we’re figuring out what type of security is best. And that—just like above—requires us to think clearly about the scenarios we’re facing.
So let’s look at some attacks against your phone:
A Russian Spetztaz Ninja wants to gain access to your unlocked phone
Your 7-year old niece wants to play games on your work phone
Your boyfriend wants to spy on your DMs with other people
Someone in Starbucks is shoulder surfing and being nosy
You accidentally leave your phone in a public place
We won’t go through all the math on this, but the Russian Ninja scenario is really bad. And really unlikely. They’re more likely to steal you and the phone, and quickly find a way to make you unlock it for them. So your security measure isn’t going to help there.
For your niece, kids are super smart about watching you type your password, so she might be able to get into it easily just by watching you do it a couple of times. Same with someone shoulder surfing at Starbucks, but you have to ask yourself who’s going to risk stealing your phone and logging into it at Starbucks. Is this a stalker? A criminal? What type? You have to factor in all those probabilities.
First question, why are you with them?
If your significant other wants to spy on your DMs, well they most definitely have had an opportunity to shoulder surf a passcode. But could they also use your finger while you slept? Maybe face recognition could be the best because it’d be obvious to you?
For all of these, you want to assign values based on how often you’re in those situations. How often you’re in Starbucks, how often you have kids around, how stalkerish your soon-to-be-ex is. Etc.
Once again, the point is to think about this in an organized way, rather than as a mashup of scenarios with no probabilities assigned that you can’t keep straight in your head. Logic vs. emotion.
It’s a way of thinking about danger.
Other examples
Here are a few other examples that you might come across.
Should I put my address on my public website?
How bad is it to be a public figure (blog/YouTube) in 2020?
Do I really need to shred this bill when I throw it away?
Don’t ever think you’ve captured all the scenarios, or that you have a perfect model.
In each of these, and the hundreds of other similar scenarios, go through the methodology. Even if you don’t get to something perfect or precise, you will at least get some clarity in what the problem is and how to think about it.
Summary
Threat Modeling is about more than technical defenses—it’s a way of thinking about risk.
The main mistake people make when considering long-term danger is letting different bad outcomes produce confusion and anxiety.
When you think about defense, start with thinking about what you’re defending, and how valuable it is.
Then capture the exact scenarios you’re worried about, along with how bad it would be if they happened, and what you think the chances are of them happening.
You can then think about additional controls as modifiers to the Impact or Probability ratings within each scenario.
Know that your calculation will never be final; it changes based on your own preferences and the world around you.
The primary benefit of Everyday Threat Modeling is having a semi-formal way of thinking about danger.
Don’t worry about the specifics of your methodology; as long as you capture feature value, scenarios, and impact/probability…you’re on the right path. It’s the exercise that’s valuable.
Notes
I know Threat Modeling is a religion with many denominations. The version of threat modeling I am discussing here is a general approach that can be used for anything from whether to move out of the country due to a failing government, or what appsec controls to use on a web application.

END THREAT MODEL ESSAY

# STEPS

- Think deeply about the input and what they are concerned with.

- Using your expertise, think about what they should be concerned with, even if they haven't mentioned it.

- Use the essay above to logically think about the real-world best way to go about protecting the thing in question.

- Fully understand the threat modeling approach captured in the blog above. That is the mentality you use to create threat models.

- Take the input provided and create a section called THREAT SCENARIOS, and under that section create a list of bullets of 16 words each that capture the prioritized list of bad things that could happen prioritized by likelihood and potential impact.

- The goal is to highlight what's realistic vs. possible, and what's worth defending against vs. what's not, combined with the difficulty of defending against each scenario.

- Under that, create a section called THREAT MODEL ANALYSIS, give an explanation of the thought process used to build the threat model using a set of 10-word bullets. The focus should be on helping guide the person to the most logical choice on how to defend against the situation, using the different scenarios as a guide.

- Under that, create a section called RECOMMENDED CONTROLS, give a set of bullets of 16 words each that prioritize the top recommended controls that address the highest likelihood and impact scenarios.

- Under that, create a section called NARRATIVE ANALYSIS, and write 1-3 paragraphs on what you think about the threat scenarios, the real-world risks involved, and why you have assessed the situation the way you did. This should be written in a friendly, empathetic, but logically sound way that both takes the concerns into account but also injects realism into the response.

- Under that, create a section called CONCLUSION, create a 25-word sentence that sums everything up concisely.

- This should be a complete list that addresses the real-world risk to the system in question, as opposed to any fantastical concerns that the input might have included.

- Include notes that mention why certain scenarios don't have associated controls, i.e., if you deem those scenarios to be too unlikely to be worth defending against.

# OUTPUT GUIDANCE

- For example, if a company is worried about the NSA breaking into their systems (from the input), the output should illustrate both through the threat scenario and also the analysis that the NSA breaking into their systems is an unlikely scenario, and it would be better to focus on other, more likely threats. Plus it'd be hard to defend against anyway.

- Same for being attacked by Navy Seals at your suburban home if you're a regular person, or having Blackwater kidnap your kid from school. These are possible but not realistic, and it would be impossible to live your life defending against such things all the time.

- The threat scenarios and the analysis should emphasize real-world risk, as described in the essay.

# OUTPUT INSTRUCTIONS

- You only output valid Markdown.

- Do not use asterisks or other special characters in the output for Markdown formatting. Use Markdown syntax that's more readable in plain text.

- Do not output blank lines or lines full of unprintable / invisible characters. Only output the printable portion of the ASCII art.

# INPUT:

INPUT:
