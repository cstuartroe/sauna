package gendocs

import (
	"fmt"
	"os"
	"strings"

	"github.com/cstuartroe/sauna/gloss"
)

func section(s string) string {
	return strings.TrimSpace(s) + "\n\n"
}

func gl(s string) string {
	words, err := gloss.ParseGloss(s)
	if err != nil {
		panic(err)
	}
	return gloss.Text(words)
}

func glossedSentence(s string, translation string) string {
	wrong := false
	if s[0] == '*' {
		s = s[1:]
		wrong = true
	}

	glossParts := strings.Split(s, " ")
	words, err := gloss.ParseGloss(s)
	if err != nil {
		panic(err)
	}

	wordsLine := ""
	glossLine := ""

	if wrong {
		wordsLine += "*"
		glossLine += " "
	}

	if len(words) != len(glossParts) {
		panic("Unequal gloss lengths")
	}

	for i := 0; i < len(words); i++ {
		rom := words[i].Romanization()

		maxLen := len(rom)
		if len(glossParts[i]) > maxLen {
			maxLen = len(glossParts[i])
		}

		wordsLine += fmt.Sprintf("%-*s", maxLen+1, rom)
		glossLine += fmt.Sprintf("%-*s", maxLen+1, glossParts[i])
	}

	return fmt.Sprintf("```\n%s\n%s\n%q\n```", wordsLine, glossLine, translation)
}

var topSection = section(`
# Sauna

Sauna is a constructed language inspired by Finnish and Old Japanese.
`)

var phonology = section(`
# Phonology

## Consonants

| | Labial | Alveolar | Dorsal |
|:--|:-:|:-:|:-:|
|**Nasal**      | m  | n | ŋ<sup>1</sup>|
|**Stop**       | p<sup>2</sup>| t | k |
|**Fricative**  |    | s | (h<sup>2</sup>)|
|**Approximant**| ʋ  | r | j |

<sup>1</sup>/ŋ/ is a rare phoneme, only occurring geminate (/ŋŋ/ spelled *ng*) and usually occurring as a result of morphophonemic processes.

<sup>2</sup>[p] and [h] are in complementary distribution, with [p] occurring geminate or after other consonants and [h] elsewhere.

## Vowels

Sauna has the common five-vowel system /a e i o u/.

It has three long vowels /aa ii uu/ and ten diphthongs /ai au ei eu oi ou ie uo ea oa iu ui/.

## Phonotactics

Syllables are of the form (C)V(V)(C). Syllables with both a second vowel mora (i.e., with a long vowel or diphthong) and a coda consonant are rare.

Only word-initial syllables may lack an onset.

The only consonants which may end a word are /n t s/.
The consonant sequences that can occur word-internally (across syllable boundaries) are
geminates /mm nn ŋŋ pp tt kk ss rr/, nasal + obstruent sequences /mp nt ns ŋk/,
/s/ + stop sequences /sp st sk/, and the additional sequences /rm rn nr nj/.

## Prosody

Sauna is mora-timed, with short vowels having one mora, long vowels and diphthongs having two moras, and coda consonants adding an extra mora.

# Morphophology

## Vowel harmony

Sauna shows a sort of height-based vowel harmony.
Each stem is either "high" or "low", with all vowels in the stem falling into the correct height category,
and the vowels of suffixes also varying according to the height category of the stem.

Some vowels (notably, /e/) and diphthongs can appear in both "high" and "low" words, but this is best thought of as two underlying
phonemes with the same surface realization.

When giving suffixes in this document, a capital-letter archiphoneme notation for vowels may be used to avoid confusion, since giving a surface realization for
a suffix like *-e* could be ambiguous.

The full set of correspondences is:

| Archiphoneme | Low | High |
|:-:|:-:|:-:|
|A|a|e|
|I|e|i|
|U|o|u|
|AA|aa|ea|
|II|ei|ii|
|UU|ou|uu|
|AI|ai|ei|
|AU|au|ou|
|IA|ea|ie|
|UA|oa|uo|
|IU|eu|iu|
|UI|oi|ui|

## Epenthetic vowels

The vowel *I* is inserted between a stem and suffix if needed to preserve phonotactics. Compare
` + fmt.Sprintf(`*%s* "birch tree" + *-t* partitive suffix -> *%s*`, gl("birch"), gl("birch-PAR")) + `, but
` + fmt.Sprintf(`*%s* "magpie" + *-t* partitive suffix -> *%s*`, gl("magpie"), gl("magpie-PAR")) + `.


## Epenthetic glides

If a stem ending in a short vowel receives a suffix beginning with a short vowel, the two coalesce into a long vowel or diphthong, e.g.
` + fmt.Sprintf(`*%s* "birch tree" + *-I* plural suffix -> *%s* "birch trees".`, gl("birch"), gl("birch-PL")) + `

However, three vowel moras cannot occur in a row, so to prevent this, an epenthetic /ʋ/ (following a back vowel) or /j/ (otherwise) is inserted between
the stem and the suffix, e.g.
` + fmt.Sprintf(`*%s* "birch trees" + *-A* genitive suffix -> *%s* "of birch trees"`, gl("birch-PL"), gl("birch-PL-GEN")) + `

## Consonant sandhi

When a word ending in one of the permitted word-final consonants /n t s/ receives a suffix beginning with a consonant, some sandhi rules may apply:

- /n/ assimilates in place to the following sound
  - /n/ + /ʋ/ coalesce into /mm/
- /t/ completely assimilates to a following obstruent
- /ʋ j/ fortify to /p k/ after coda /s t/, with coda /t/ assimilating as usual. that is,
  - /s/ + /ʋ/ becomes /sp/, /s/ + /j/ becomes /sk/
  - /t/ + /ʋ/ becomes /pp/, /t/ + /j/ becomes /kk/
- coda /s t/ become /r/ before a following nasal or /r/

## Consonant lenition

Some suffixes may cause lenition to the last consonant or consonant sequence of the stem.
Lenition never affects word-initial consonants (i.e., if the stem is only one syllable) and only some consonants and sequences are affected.

The full set of lone consonants and consonant sequences is

|No lenition|Lenition|
|:-:|:-:|
|h|v|
|t|r|
|k|v/y|
|pp|h|
|tt|t|
|kk|k|
|mp|mm|
|nt|nn|
|nk|ng|

Some examples of consonant lenition with the leniting genitive suffix *-A* include:

` + fmt.Sprintf(`*%s* "stone" -> *%s* "of (the) stone"`, gl("stone"), gl("stone-GEN")) + `

` + fmt.Sprintf(`*%s* "river" -> *%s* "of (the) river"`, gl("river"), gl("river-GEN")) + `

` + fmt.Sprintf(`*%s* "magpie" -> *%s* "of (the) magpie"`, gl("magpie"), gl("magpie-GEN")) + `
`)

var grammar = section(`
# Grammar

## Syntax and typology

Sauna is a strictly head-final language, with SVO clause order and modifiers coming before nouns.

It is also strictly suffixing, with suffixes stacking in an agglutinative fashion.

Sauna is generally pro-drop, with arbitrary arguments being permitted to be left out of the sentence when clear from context:

` + glossedSentence("home-ALL go-PST-S2/3S", "(He) went home.") + `

Sauna verbs conjugate for subject, which is often enough to disambiguate, although conjugation only makes a three-way distinction
between first person singular, second or third person singular, and plural subject.

### Parts of speech

There are only two content parts of speech in Sauna, nouns and verbs.
Words with adjectival meanings are usually verbs.
However, some nouns have adjectival meaning; I call these "quality nouns".
Quality nouns are often ideophonic, and often involve reduplication: *` + gl("sneaky") + `* "sneaky"

Other parts of speech include numerals, noun classifiers, and relational nouns, although
noun classifiers and relational nouns are not entirely distinct from normal nouns.

## Nouns

Nouns are inflected for number, posessor, and case, with suffixes coming in that order.

### Number

The plural suffix is *-I*

` + fmt.Sprintf(`*%s* "birch tree" -> *%s* "birch trees"`, gl("birch"), gl("birch-PL")) + `

` + fmt.Sprintf(`*%s* "magpie" -> *%s* "magpies"`, gl("magpie"), gl("magpie-PL")) + `

` + fmt.Sprintf(`*%s* "stone" -> *%s* "stones"`, gl("stone"), gl("stone-PL")) + `

### Possessor

Nouns are marked for their possessor:

- The first person singular possessor suffix is *-n*: *` + gl("home-P1S") + `* "my home"

- The second/third person singular possessor suffix is *-t*: *` + gl("home-P2/3S") + `* "your home, his home, her home"

- The first person plural possessor suffix is *-mmIt*: *` + gl("home-P1P") + `* "our home"

- The second/third person plural possessor suffix is *-vIt*: *` + gl("home-P2/3P") + `* "y'all's home, their home"

### Cases

Sauna has nine cases.

Note that there is no dedicated case for direct objects,
which may be marked with the nominative, genitive, or partitive according to usage patterns described below.

#### Topical

Many indicative main clauses have a topic - if not, the topic is probably implied.
The topic is marked with the topical case marker *-vA* and generally comes first in the clause.
The topic may have a variety of semantic roles in the sentence, including agent, patient, experiencer, location, or possessor of some other argument.

` + glossedSentence("woman-TOP home-ALL go-PST-S2/3S", "The woman went home.") + `

` + glossedSentence("magpie-PL-TOP tail-NOM long-NPT-S2/3S", "Magpies have long tails.") + `

The topic is often implied to be contrastive, especiialy when the its semantic role is something other than agent or patient.

` + glossedSentence("south-TOP birch-PL-NOM small-NPT-SP", "In the south, the birch trees are small.") + `

For roles other than agent, patient, location, or possessor, the topical case suffix may be stacked on top of another:

` + glossedSentence("home-ALL-TOP woman-NOM go-PST-S2/3S", "Home is where the woman went.") + `

For some constructions, a particular argument must be the topic.

` + glossedSentence("magpie-PL-TOP tail-NOM be-NPT-S2/3S", "Magpies have tails.") + `

` + glossedSentence("1S.polite-TOP magpie-PL-NOM good-NPT-SP", "I like magpies.") + `

#### Nominative

The nominative is used for agents and inanimate patients. It is unmarked.

` + glossedSentence("magpie-TOP acorn-NOM eat-PST-S2/3S", "The magpie ate an acorn.") + `

` + glossedSentence("acorn-TOP magpie-NOM eat-PST-S2/3S", "An acorn is what the magpie ate.") + `

As a general rule, if the agent and patient are both in the nominative case, the agent should come first.

` + glossedSentence("magpie-NOM acorn-NOM eat-PST-S2/3S", "The magpie ate an acorn.") + `

#### Genitive

The genitive, marked with *-A* and consonant lenition on the stem, has two distinct functions.

The first is to mark a noun as a possessor or modifier of another noun. Note that nouns also have possession suffixes, so having
a genitive modifier as well is generally not needed except to clarify third person possessors.

` + glossedSentence("woman-GEN home-P2/3S", "the woman's home") + `

A common type of modification other than possession is to use quality nouns attributively:

` + glossedSentence("sneaky-GEN woman", "sneaky woman") + `

The second function is to mark animate patients.

` + glossedSentence("woman-GEN see-PST-S1S", "I saw the woman.") + `

For the purpose of direct object marking, animals may be marked with either the nominative (treating them as animate)
or the genitive (treating them as inanimate), depending on the degree of emotional proximity to the animal the speaker wants
to convey. This is similar to how it's possible to use either "it" or gendered pronouns with animals in English.

` + glossedSentence("magpie-GEN see-PST-S1S", "I saw the magpie.") + `

` + glossedSentence("magpie-NOM see-PST-S1S", "I saw the magpie.") + `

#### Partitive

The partitive, marked with *-t*, has two distinct functions. Notably, the partitive cannot co-occur with plural marking.

The first function is to mark a direct object where the verb was unsuccessful or did not reach a state of completion.
The plurality of the partitive noun is ambiguous.

` + glossedSentence("magpie-NOM acorn-PAR eat-PST-S2/3S", "The magpie ate some acorns.") + `

Contrast this with

` + glossedSentence("magpie-NOM acorn-NOM eat-PST-S2/3S", "The magpie ate the acorn.") + `

` + glossedSentence("magpie-NOM acorn-PL-NOM eat-PST-S2/3S", "The magpie ate the acorns.") + `

The second function is to adjoin the noun to a cardinal number.

` + glossedSentence("magpie-PAR three CL.animal", "three magpies") + `

#### Equative

The equative case, marked with *-ssI* is used in two related ways.

The first way is to mark the complement of a copula:

` + glossedSentence("woman-EQU be-NPT-S1S", "I am a woman.") + `

The second way is to indicate manner or similarity:

` + glossedSentence("stone-EQU hard-NPT-S2/3S", "It is hard like a stone.") + `

This includes using quality nouns adverbially:

` + glossedSentence("sneaky-EQU home-ALL go-PST-S1S", "I went home sneakily.") + `

#### Dative

The dative case, marked with *-nyU*, is used to mark a benefactor or motive.

` + glossedSentence("woman-DAT home-NOM build-PST-S1S", "I built the woman a home.") + `

` + glossedSentence("stone-PAR take-CNJ-GEN-DAT river-ALL go-PST-S1S", "I went to the river to get stones.") + `

#### Adpositional cases

There are three cases indicating position or direction: the allative, locative, and ablative.
These three are all marked with suffixes based on the genitive suffix, and so all lenite the stem.

The allative is marked with the genitive suffix plus *-nI*, and indicates motion towards.

` + glossedSentence("home-ALL go-PST-S2/3S", "He went home.") + `

The allative is also used for recipients.

` + glossedSentence("magpie-ALL acorn-NOM give-PST-S1S", "I gave the magpie an acorn.") + `

The locative is marked with the genitive suffix plus *-ttI*, and indicates location at.

` + glossedSentence("home-LOC be-NPT-S2/3S", "He is at home.") + `

The ablative is marked with the genitive suffix plus *-stI*, and indicates motion away.

` + glossedSentence("home-ABL go-PST-S2/3S", "He left home.") + `

### Relational nouns

There is a set of nouns which commonly appear as the possessum of other nouns, indicating a physical position or
some other type of relationship.

An example of such a noun is *` + gl("inside") + `* "inside":

` + glossedSentence("home-GEN inside-P2/3S-LOC be-NPT-S2/3S", "He is inside the house.") + `

Some such nouns conventionally only appear with a particular case, such as *` + gl("companion") + `*, which
may mean "companion," but which almost always appears in the equative case with a comitative meaning:

` + glossedSentence("woman-TOP companion-P1S-EQU go-PST-S2/3S", "The woman went with me.") + `

Other positional nouns don't require any case marking to be understood with an oblique meaning, such as *` + gl("without") + `* "without":

` + glossedSentence("without-P1P", "without us") + `

#### Adpositional relationals

The adpositional cases, being based on the genitive case, actually originate as relational nouns.
The possession suffix has been dropped from the adpositional case suffixes, but the same roots can also function
similarly to *` + gl("without") + `* "without", as standalone relational with a possession suffix.
This is preferred over putting first person pronouns in the adpositional cases, and permissible for second and third persons rather than
using a pronoun if there is no risk of ambiguity, but may not be used with an actual noun possessor.

` + glossedSentence("*that-NOM 1S.polite-ALL give-CNJ", "Give it to me.") + `

` + glossedSentence("that-NOM ALL-P1S give-CNJ", "Give it to me.") + `

` + glossedSentence("that-NOM 2S.formal-ALL give-NPT-S1S", "I'll give it to you.") + `

` + glossedSentence("that-NOM ALL-P2/3S give-NPT-S1S", "I'll give it to you/him/her.") + `

` + glossedSentence("that-NOM woman-ALL give-NPT-S1S", "I'll give it to the woman") + `

` + glossedSentence("*that-NOM woman-GEN ALL-P2/3S give-NPT-S1S", "I'll give it to the woman") + `

## Verbs

Verbs can take suffixes for, in order: voice, negation, tense, and subject agreement.

### Voice marking

Voice suffixes are optional and usually not present. The two voice suffixes are the passive/potential *-rA* and the causative *-sA*.

#### Passive/potential

The passive *-rA* is used when there is no specific agent mentioned or implied, to promote the patient to subject, i.e.,
it takes the syntactic position and case marking that the agent would normally take, and the verb conjugates for its person and number.
For animate patients, this means that the case can shift from genitive to nominative:

` + glossedSentence("magpie-GEN see-NPT-S1S", "I see the magpie.") + `

` + glossedSentence("magpie-NOM see-PSV-S2/3S", "The magpie is seen.") + `

However, the patient might not change case, e.g. if it is marked as the topic,
or if it is inanimate (inanimate nouns use the nominative case as both subject and object):

` + glossedSentence("magpie-TOP see-NPT-S1S", "I see the magpie.") + `

` + glossedSentence("magpie-TOP see-PSV-S2/3S", "The magpie is seen.") + `

` + glossedSentence("acorn-NOM see-NPT-S1S", "I see the acorn.") + `

` + glossedSentence("acorn-NOM see-PSV-S2/3S", "The acorn is seen.") + `

The same affix may be used with a potential sense.
This still entails promoting the object to subject; there may be ambiguity between the potential and plain passive.

` + glossedSentence("magpie-NOM see-PSV-S2/3S", "The magpie is seen/visible.") + `

However, in the potential sense, an agent may be stated; if so, it is the topic, and the verb is still conjugated for the patient.

` + glossedSentence("1S.polite-TOP magpie-NOM see-PSV-S2/3S", "I can see the magpie.") + `

#### Causative

The causative, marked with *-sA*, is used to convey forcing or enabling someone or something to do something. The causee is marked with the dative case.

` + glossedSentence("woman-DAT home-ALL go-CAU-PST-S1S", "I made the woman go home.") + `

### Negation

Verbs are negated with the suffix *-nA*.

` + glossedSentence("that-NOM eat-NEG-PST-S1S", "I didn't eat it.") + `

` + glossedSentence("magpie-NOM see-PSV-NEG-S2/3S", "The magpie is not visible.") + `

### Tense

Verbs can be marked for past or non-past tense.

#### Non-past

The non-past is versatile in use, being usable for simple present, gnomic, and future senses.

` + glossedSentence("magpie-TOP acorn-PAR eat-NPT-S2/3S", "The magpie eats acorns./The magpie is eating acorns./The magpie will eat acorns.") + `

The explicit non-past suffix is *-rU*. However, it is only used if there are no preceding suffixes, i.e., no voice or negation marking.
If there are voice or negation suffixes applied, the non-past is zero-marked.

` + glossedSentence("magpie-TOP acorn-PAR eat-NEG-S2/3S", "The magpie doesn't eat acorns./The magpie isn't eating acorns./The magpie won't eat acorns.") + `

#### Past

The past tense is marked with *-s*.

` + glossedSentence("magpie-TOP acorn-PAR eat-PST-S2/3S", "The magpie ate acorns.") + `

### Subject agreement

Sauna verbs conjugate for the subject, i.e., the agent of transitive verbs, and the one core argument of intransitive verbs,
including passives and verbs with adjectival meanings.

There is only a three-way distinction made in verb conjugation:

- first person singular is marked *-n*: *` + gl("be-NPT-S1S") + `* "I am", *` + gl("be-PST-S1S") + `* "I was"

- second and third person singular are unmarked: *` + gl("be-NPT-S2/3S") + `* "you are/he, she is", *` + gl("be-PST-S2/3S") + `* "you were/he, she was"

- plural subjects are all marked alike, by lengthening a final vowel if the stem ends in a short vowel, and zero-marking otherwise: *` + gl("be-NPT-SP") + `* "they are", *` + gl("be-PST-SP") + `* "they were"

### Conjunctive form

The conjunctive is non-finite form of verbs which replaces tense and person marking with a suffix *-tI*.
Its core function is as a converb, to convey simultaneous or sequential action by the same subject:

` + glossedSentence("sing-CNJ home-ALL go-PST-S1S", "I sang as I went home.") + `

` + glossedSentence("stone-PAR take-CNJ home-NOM build-PST-S1S", "I got some rocks and built a house.") + `

As an extension of this, it is used with *` + gl("be") + `* "be" to convey a progressive aspect:

` + glossedSentence("eat-CNJ be-NPT-S1S", "I'm eating.") + `

#### Nominalization

Verbs can be turned into nouns, which can either refer to a single instance of performing the verb or the concept of the verb,
by adding the conjunctive suffix and the genitive case suffix. Other case suffixes then stack on top of the genitive suffix.

` + glossedSentence("1S.polite-TOP eat-CNJ-GEN-NOM good-NPT-S2/3S", "I like eating.") + `

Arguments of the nominalized verb can branch leftward from the verb:

` + glossedSentence("stone-PAR take-CNJ-GEN-DAT river-ALL go-PST-S1S", "I went to the river to get stones.") + `

## Numerals

### Counting system

Sauna's numerals are largely base ten, with multiples of ten up to one hundred grouped into twenties.

*` + gl("one") + `* "one"

*` + gl("two") + `* "two"

*` + gl("three") + `* "three"

*` + gl("four") + `* "four"

*` + gl("five") + `* "five"

*` + gl("six") + `* "six"

*` + gl("seven") + `* "seven"

*` + gl("eight") + `* "eight"

*` + gl("nine") + `* "nine"

*` + gl("ten") + `* "ten"

*` + gl("ten one") + `* "eleven"

*` + gl("ten two") + `* "twelve"

*` + gl("twenty") + `* "twenty"

*` + gl("twenty ten") + `* "thirty"

*` + gl("two twenty") + `* "forty"

*` + gl("two twenty ten") + `* "fifty"

*` + gl("hundred") + `* "hundred"

*` + gl("thousand") + `* "thousand"

*` + gl("myriad") + `* "ten thousand"

*` + gl("three myriad six thousand five hundred four twenty ten eight") + `* "36,598"

### Cardinal numbers and numeral classifiers

Sauna uses numeral classifiers with cardinal numbers:

` + glossedSentence("magpie-PAR three CL.animal", "three magpies") + `

The classifiers are effectively nouns. They take the case marking for the phrase:

` + glossedSentence("magpie-PAR three CL.animal-DAT", "for three magpies") + `

Some nouns can even double as numeral classifiers:

` + glossedSentence("cup-PAR three CL.general", "three cups") + `

` + glossedSentence("water-PAR three cup", "three cups of water") + `

Some common classifiers include:

- *` + gl("CL.general") + `* can be used for any inanimate noun

- *` + gl("CL.human") + `* people, or anthropomorphized characters or deities

- *` + gl("CL.animal") + `* birds and small to medium-sized mammals

- *` + gl("CL.small.object") + `* small hard objects such as acorns, grains, beads, coins

- *` + gl("CL.piece") + `* pieces of a larger whole

### Ordinal numbers

To make numbers ordinal, they should appear as a modifier of the noun in the genitive:

` + glossedSentence("three-GEN magpie", "the third magpie") + `

"First" is irregular, being based on the word *` + gl("beginning") + `* "beginning":

` + glossedSentence("beginning-GEN magpie", "the first magpie") + `

## Pronouns

### Personal pronouns

There is a range of options for singular address, based on politeness and gender. Plural pronouns are more standard.

#### First person singular

First person pronouns are the most likely to be omitted, since verbs unambiguously conjugate for them as subject,
and the first person may be obvious from context in requests etc.

If an explicit first person pronoun is given, there are three main choices:

- *` + gl("1S.polite") + `* is the standard first person pronoun, used in formal and some casual situations

- *` + gl("1S.masc") + `* is used by masculine people in most casual situations

- *` + gl("1S.fem") + `* sounds diminutive and may be used by feminine people casually, though it's not unusual for feminine people to use *` + gl("1S.polite") + `* in some casual situations or even avoid *` + gl("1S.fem") + `* consistently

#### Second person singular

It's common for names or nicknames to be used instead of second person pronouns. Second person pronouns are used in the following situations:

- *` + gl("2S.formal") + `* is used when the other person's name is not known, such as when addressing a stranger or in writing or media addressing a generic reader/listener

- *` + gl("2S.hon") + `* is used to address someone of significantly higher status

- *` + gl("2S.masc") + `* may be used by a masculine speaker to a masculine listener, sounding somewhat rough or uncouth

#### Third person singular

It's almost universal to use names when talking about single other people.
The only exception is when talking about someone whose name is not known, in which case
*` + gl("that CL.human") + `* or *` + gl("yon CL.human") + `* "that person" can be used.

#### Plural pronouns

Plural pronouns are standard and always used rather than names or some other manner of description.

*` + gl("1P.EXCL") + `* is the exclusive first person plural pronoun ("we" excluding the listener).

*` + gl("1P.INCL") + `* is the inclusive equivalent ("we" including the listener).

*` + gl("2P") + `* is the second person plural pronoun.

*` + gl("3P") + `* is the third person plural pronoun.

### Demonstrative pronouns

Sauna has three levels of distance for demonstratives:

- *` + gl("this") + `* is proximal ("this")

- *` + gl("that") + `* is medial ("that by you, that nearby")

- *` + gl("yon") + `* is distal ("that over there, yon")

There is also an indefinite pronoun *` + gl("some") + `* which by itself means "some, any" and is also used to form negative existentials and questions.

Demonstrative pronouns have similar syntax to numerals. They can appear by themselves:

` + glossedSentence("that-NOM take-PST-S1S", "I took it.") + `

or directly with classifiers:

` + glossedSentence("that CL.general-NOM take-PST-S1S", "I took that thing.") + `

` + glossedSentence("yon CL.human-GEN see-PST-S1S", "I saw that person.") + `

To be attributive, demonstratives are marked with the genitive:

` + glossedSentence("yon-GEN magpie-GEN see-PST-S1S", "I saw that magpie.") + `

Demonstratives by themselves can mean "here" and "there" using the adpositional cases:

` + glossedSentence("yon-ALL go-NPT-S1S", "I'm going there.") + `

` + glossedSentence("home-P1S-NOM this-LOC be-NPT-S2/3S", "My house is here.") + `

To retain the sense of a location, they must use the locative case as a topic, even though locations otherwise can be topic without overt locative marking:

` + glossedSentence("home-TOP woman-NOM be-NPT-S2/3S", "There's a woman in the house.") + `

` + glossedSentence("yon-LOC-TOP woman-NOM be-NPT-S2/3S", "There's a woman over there.") + `

` + glossedSentence("*yon-TOP woman-NOM be-NPT-S2/3S", "There's a woman over there.") + `

## Syntax

### General word order

The order of a Sauna clause is generally topic, subject (if it is distinct), oblique phrases, direct object, verb.
However, case marking makes this somewhat flexible; the only rigid rule is that the verb must come last.

` + glossedSentence("today-TOP woman-NOM magpie-ALL acorn-PAR give-PST-S2/3S", "Today the woman gave the magpie acorns.") + `

If there are multiple predicates due to the use of the conjunctive verb form, their obliques and direct object are ordered
the same way before each respective verb:

` + glossedSentence("water-NOM drink-CNJ home-NOM build-PST-S1S", "I drank water as I built the house.") + `

### Existentials, to have, to be

The verb *` + gl("be") + `* essentially means "to exist".

` + glossedSentence("yon-LOC-TOP magpie-NOM be-NPT-S2/3S", "There's a magpie over there.") + `

With a topic, it can be used to mean "have":

` + glossedSentence("1S.polite-TOP home-NOM be-NPT-S2/3S", "I have a house.") + `

To equate two nouns, the equative case is used:

` + glossedSentence("woman-EQU be-NPT-S1S", "I am a woman.") + `

` + glossedSentence("that-NOM home-EQU be-NPT-S2/3S", "It is a house.") + `

### Questions

Sauna uses an interrogative suffix *-kU* which attaches to some element of the sentence to form a question, going after all other suffixes.

For polar questions, this is generally the verb:

` + glossedSentence("yon-ALL go-PST-S2/3S-INT", "Did you go there?") + `

although the interrogative particle can go on another element to specifically make it the contrastive focus of the question:

` + glossedSentence("yon-ALL-INT go-PST-S2/3S", "Is that where you went?") + `

For content questions, the interrogative particle is placed onto a noun phrase with the indefinite demonstrative.

` + glossedSentence("some-ALL-INT go-PST-S2/3S", "Where did you go?") + `

` + glossedSentence("some-NOM-INT eat-PST-S2/3S", "What did you eat?") + `

` + glossedSentence("2S.formal-TOP some-GEN cup-NOM-INT good-NPT-S2/3S", "Which cup do you like?") + `

Combining the indefinite with classifiers is a productive way to more specifically ask "what", "which", or "who"?

` + glossedSentence("2S.formal-TOP some CL.human-GEN-INT good-NPT-S2/3S", "Who do you like?") + `

### Imperatives and cohortatives

Sauna imperatives are formed by using the bare verb stem:

` + glossedSentence("ALL-P1S give", "Give it to me.") + `

This can sound rather blunt. There is a verb *` + gl("do.please") + `* which is specifically used act as a softer imperative,
with the actual verb being requested coming before in conjunctive form:

` + glossedSentence("ALL-P1S give-CNJ do.please", "Please give it to me.") + `

This can actually be overly formal and polite. The most neutral way to pose a request, neither blunt and urgent nor overly formal,
is to omit *` + gl("do.please") + `* and leave the bare conjunctive.

` + glossedSentence("ALL-P1S give-CNJ", "Give it to me.") + `

The normal way to express a cohortative is with the non-past tense, plural conjugation plus interrogative particle.
Note that this literally just asks about a group, but the implication can be cohortative if it makes sense contextually:

` + glossedSentence("go-NPT-SP-INT", "Are they going?/Are we going?/Let's go!") + `

`)

func GenerateReadme() string {
	content := ""

	content += topSection
	content += phonology
	content += grammar

	return content
}

func GenerateNorthWind() string {
	content := ""

	content += section(glossedSentence(
		"north-GEN wind-GEN companion-P2/3S-EQU sun-TOP some CL.human-NOM-INT ABL strong-CNJ-GEN-DAT fight-PST-SP",
		"The North Wind and the Sun were disputing about who was stronger.",
	))

	content += section(glossedSentence(
		"that-ABL-TOP cozy cloth-NOM carry travel man-NOM near-ALL go-PST-S2/3S",
		"Then a traveling man wearing a warm cloth came by.",
	))

	content += section(glossedSentence(
		"3P-TOP cloth-NOM remove-CAU-CNJ succeed CL.human-NOM ABL strong CL.human-EQU call-PSV-CNJ-LOC one way-NOM be-PST-S2/3S",
		"They agreed that the one who succeeded in making the traveler remove the cloak should be called stronger.",
	))

	content += section(glossedSentence(
		"that-ABL-TOP north-GEN wind-NOM strenuous-EQU blow-PST-S2/3S",
		"Then the North Wind blew as hard as possible.",
	))

	content += section(glossedSentence(
		"but north-GEN wind-TOP some-EQU strenuous-EQU blow time-LOC travel man-NOM that-EQU cuddle-EQU cloth-NOM wrap-PST-S2/3S",
		"But the harder the North Wind blew, the more tightly the traveling man wrapped his cloth.",
	))

	content += section(glossedSentence(
		"ending-P2/3S-TOP north-GEN wind-NOM abate-PST-S2/3S",
		"In the end, the North Wind gave up.",
	))

	content += section(glossedSentence(
		"that-ABL-TOP sun-NOM hot-CNJ shine-PST-S2/3S",
		"Then the Sun shone warmly",
	))

	content += section(glossedSentence(
		"travel man-NOM immediately cloth-NOM remove-PST-S2/3S",
		"The traveling man immediately removed his cloth.",
	))

	content += section(glossedSentence(
		"this way-TOP north-GEN wind-NOM sun-GEN ABL strong CL.human-EQU call-CAU-PSV-PST-S2/3S",
		"In this way, the North Wind had to call the Sun the stronger one.",
	))

	return content
}

func GenerateDocs() {
	os.WriteFile("README.md", []byte(GenerateReadme()), 0)
	os.WriteFile("North Wind.md", []byte(GenerateNorthWind()), 0)
}
