<div align="center">

<img src="https://beehiiv-images-production.s3.amazonaws.com/uploads/asset/file/2012aa7c-a939-4262-9647-7ab614e02601/extwis-logo-miessler.png?t=1704502975" alt="extwislogo" width="400" height="400"/>

# `/extractwisdom`

<h4><code>extractwisdom</code> is a <a href="https://github.com/danielmiessler/fabric" target="_blank">Fabric</a> pattern that <em>extracts wisdom</em> from any text.</h4>

[Description](#description) •
[Functionality](#functionality) •
[Usage](#usage) •
[Output](#output) •
[Meta](#meta)

</div>

<br />

## Description

**`extractwisdom` addresses the problem of **too much content** and too little time.**

_Not only that, but it's also too easy to forget the stuff we read, watch, or listen to._

This pattern _extracts wisdom_ from any content that can be translated into text, for example:

- Podcast transcripts
- Academic papers
- Essays
- Blog posts
- Really, anything you can get into text!

## Functionality

When you use `extractwisdom`, it pulls the following content from the input.

- `IDEAS`
  - Extracts the best ideas from the content, i.e., what you might have taken notes on if you were doing so manually.
- `QUOTES`
  - Some of the best quotes from the content.
- `REFERENCES`
  - External writing, art, and other content referenced positively during the content that might be worth following up on.
- `HABITS`
  - Habits of the speakers that could be worth replicating.
- `RECOMMENDATIONS`
  - A list of things that the content recommends Habits of the speakers.

### Use cases

`extractwisdom` output can help you in multiple ways, including:

1. `Time Filtering`<br />
   Allows you to quickly see if content is worth an in-depth review or not.
2. `Note Taking`<br />
   Can be used as a substitute for taking time-consuming, manual notes on the content.

## Usage

You can reference the `extractwisdom` **system** and **user** content directly like so.

### Pull the _system_ prompt directly

```sh
curl -sS https://github.com/danielmiessler/fabric/blob/main/extract-wisdom/dmiessler/extract-wisdom-1.0.0/system.md
```

### Pull the _user_ prompt directly

```sh
curl -sS https://github.com/danielmiessler/fabric/blob/main/extract-wisdom/dmiessler/extract-wisdom-1.0.0/user.md
```

## Output

Here's an abridged output example from `extractwisdom` (limited to only 10 items per section).

```markdown
## SUMMARY:

The content features a conversation between two individuals discussing various topics, including the decline of Western culture, the importance of beauty and subtlety in life, the impact of technology and AI, the resonance of Rilke's poetry, the value of deep reading and revisiting texts, the captivating nature of Ayn Rand's writing, the role of philosophy in understanding the world, and the influence of drugs on society. They also touch upon creativity, attention spans, and the importance of introspection.

## IDEAS:

1. Western culture is perceived to be declining due to a loss of values and an embrace of mediocrity.
2. Mass media and technology have contributed to shorter attention spans and a need for constant stimulation.
3. Rilke's poetry resonates due to its focus on beauty and ecstasy in everyday objects.
4. Subtlety is often overlooked in modern society due to sensory overload.
5. The role of technology in shaping music and performance art is significant.
6. Reading habits have shifted from deep, repetitive reading to consuming large quantities of new material.
7. Revisiting influential books as one ages can lead to new insights based on accumulated wisdom and experiences.
8. Fiction can vividly illustrate philosophical concepts through characters and narratives.
9. Many influential thinkers have backgrounds in philosophy, highlighting its importance in shaping reasoning skills.
10. Philosophy is seen as a bridge between theology and science, asking questions that both fields seek to answer.

## QUOTES:

1. "You can't necessarily think yourself into the answers. You have to create space for the answers to come to you."
2. "The West is dying and we are killing her."
3. "The American Dream has been replaced by mass packaged mediocrity porn, encouraging us to revel like happy pigs in our own meekness."
4. "There's just not that many people who have the courage to reach beyond consensus and go explore new ideas."
5. "I'll start watching Netflix when I've read the whole of human history."
6. "Rilke saw beauty in everything... He sees it's in one little thing, a representation of all things that are beautiful."
7. "Vanilla is a very subtle flavor... it speaks to sort of the sensory overload of the modern age."
8. "When you memorize chapters [of the Bible], it takes a few months, but you really understand how things are structured."
9. "As you get older, if there's books that moved you when you were younger, it's worth going back and rereading them."
10. "She [Ayn Rand] took complicated philosophy and embodied it in a way that anybody could resonate with."

## HABITS:

1. Avoiding mainstream media consumption for deeper engagement with historical texts and personal research.
2. Regularly revisiting influential books from youth to gain new insights with age.
3. Engaging in deep reading practices rather than skimming or speed-reading material.
4. Memorizing entire chapters or passages from significant texts for better understanding.
5. Disengaging from social media and fast-paced news cycles for more focused thought processes.
6. Walking long distances as a form of meditation and reflection.
7. Creating space for thoughts to solidify through introspection and stillness.
8. Embracing emotions such as grief or anger fully rather than suppressing them.
9. Seeking out varied experiences across different careers and lifestyles.
10. Prioritizing curiosity-driven research without specific goals or constraints.

## FACTS:

1. The West is perceived as declining due to cultural shifts away from traditional values.
2. Attention spans have shortened due to technological advancements and media consumption habits.
3. Rilke's poetry emphasizes finding beauty in everyday objects through detailed observation.
4. Modern society often overlooks subtlety due to sensory overload from various stimuli.
5. Reading habits have evolved from deep engagement with texts to consuming large quantities quickly.
6. Revisiting influential books can lead to new insights based on accumulated life experiences.
7. Fiction can effectively illustrate philosophical concepts through character development and narrative arcs.
8. Philosophy plays a significant role in shaping reasoning skills and understanding complex ideas.
9. Creativity may be stifled by cultural nihilism and protectionist attitudes within society.
10. Short-term thinking undermines efforts to create lasting works of beauty or significance.

## REFERENCES:

1. Rainer Maria Rilke's poetry
2. Netflix
3. Underworld concert
4. Katy Perry's theatrical performances
5. Taylor Swift's performances
6. Bible study
7. Atlas Shrugged by Ayn Rand
8. Robert Pirsig's writings
9. Bertrand Russell's definition of philosophy
10. Nietzsche's walks
```

This allows you to quickly extract what's valuable and meaningful from the content for the use cases above.

## Meta

- **Author**: Daniel Miessler
- **Version Information**: Daniel's main `extractwisdom` version.
- **Published**: January 5, 2024
