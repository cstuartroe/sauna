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

` + glossedSentence("magpie-PL-TOP tail-NOM be-NPT-SP", "Magpies have tails.") + `

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

`)

func GenerateReadme() string {
	content := ""

	content += topSection
	content += phonology
	content += grammar

	return content
}

func GenerateDocs() {
	os.WriteFile("README.md", []byte(GenerateReadme()), 0)
}
