# Sauna

Sauna is a constructed language inspired by Finnish and Old Japanese.

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

It has three long vowels /aa ii uu/ and twelve diphthongs /ai au ei eu oi ou ie uo ea oa iu ui/.

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
イオヱ *iuve* "birch tree" + *-t* partitive suffix -> イオヱツ *iuvet*, but
イヨカツ *eyokat* "magpie" + *-t* partitive suffix -> イヨカチツ *eyokatet*.

## Epenthetic glides

If a stem ending in a short vowel receives a suffix beginning with a short vowel, the two coalesce into a long vowel or diphthong, e.g.
イオヱ *iuve* "birch tree" + *-I* plural suffix -> イオヱイ *iuvei* "birch trees".

However, three vowel moras cannot occur in a row, so to prevent this, an epenthetic /ʋ/ (following a back vowel) or /j/ (otherwise) is inserted between
the stem and the suffix, e.g.
イオヱイ *iuvei* "birch trees" + *-A* genitive suffix -> イオヱイエ *iuveiye* "of birch trees"

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

ナツハ *nappa* "stone" -> ナハア *nahaa* "of (the) stone"

センコ *senku* "river" -> センゴエ *senguo* "of (the) river"

イヨカツ *eyokat* "magpie" -> イヨカヅア *eyokara* "of (the) magpie"

# Grammar

## Syntax and typology

Sauna is a strictly head-final language, with SVO clause order and modifiers coming before nouns.

It is also strictly suffixing, with suffixes stacking in an agglutinative fashion.

Sauna is generally pro-drop, with arbitrary arguments being permitted to be left out of the sentence when clear from context:

```
チオロエニ  ヤトス  
tiuruoni   yatos
tiuru-'AnI yato-s  -∅
home -ALL  go  -PST-S2/3S
"(He) went home."
```

Sauna verbs conjugate for subject, which is often enough to disambiguate, although conjugation only makes a three-way distinction
between first person singular, second or third person singular, and plural subject.

### Parts of speech

There are only two content parts of speech in Sauna, nouns and verbs.
Words with adjectival meanings are usually verbs.
However, some nouns have adjectival meaning; I call these "quality nouns".
Quality nouns are often ideophonic, and often involve reduplication: ヒツリヰツリ *hirrivirri* "sneaky"

Other parts of speech include numerals, noun classifiers, and relational nouns, although
noun classifiers and relational nouns are not entirely distinct from normal nouns.

## Nouns

Nouns are inflected for number, posessor, and case, with suffixes coming in that order.

### Number

The plural suffix is *-I*

イオヱ *iuve* "birch tree" -> イオヱイ *iuvei* "birch trees"

イヨカツ *eyokat* "magpie" -> イヨカツイ *eyokate* "magpies"

ナツハ *nappa* "stone" -> ナツハイ *nappai* "stones"

### Possessor

Nouns are marked for their possessor:

- The first person singular possessor suffix is *-n*: チオロン *tiurun* "my home"

- The second/third person singular possessor suffix is *-t*: チオロツ *tiurut* "your home, his home, her home"

- The first person plural possessor suffix is *-mmIt*: チオロンミツ *tiurummit* "our home"

- The second/third person plural possessor suffix is *-vIt*: チオロヰツ *tiuruvit* "y'all's home, their home"

### Cases

Sauna has nine cases.

Note that there is no dedicated case for direct objects,
which may be marked with the nominative, genitive, or partitive according to usage patterns described below.

#### Topical

Many indicative main clauses have a topic - if not, the topic is probably implied.
The topic is marked with the topical case marker *-vA* and generally comes first in the clause.
The topic may have a variety of semantic roles in the sentence, including agent, patient, experiencer, location, or possessor of some other argument.

```
コニヱ  チオロエニ  ヤトス  
kunive    tiuruoni   yatos
kuni -vA  tiuru-'AnI yato-s  -∅
woman-TOP home -ALL  go  -PST-S2/3S
"The woman went home."
```

```
イヨカツイワ  エレコ  ナライロ  
eyokateva     ereku     narairo
eyokat-I -vA  ereku-∅   narai-rU -∅
magpie-PL-TOP tail -NOM long -NPT-S2/3S
"Magpies have long tails."
```

The topic is often implied to be contrastive, especiialy when the its semantic role is something other than agent or patient.

```
ヱイロンヱ  イオヱイ  イ𛄡ロー  
veirumme   iuvei        iyeruu
veirun-vA  iuve -I -∅   iye  -rU -X
south -TOP birch-PL-NOM small-NPT-SP
"In the south, the birch trees are small."
```

For roles other than agent, patient, location, or possessor, the topical case suffix may be stacked on top of another:

```
チオロエニヱ  コニ  ヤトス  
tiuruonive     kuni      yatos
tiuru-'AnI-vA  kuni -∅   yato-s  -∅
home -ALL -TOP woman-NOM go  -PST-S2/3S
"Home is where the woman went."
```

For some constructions, a particular argument must be the topic.

```
イヨカツイワ  エレコ  ナロ  
eyokateva     ereku     naro
eyokat-I -vA  ereku-∅   na-rU -∅
magpie-PL-TOP tail -NOM be-NPT-S2/3S
"Magpies have tails."
```

```
アリンワ  イヨカツイ  タイロー  
aremma        eyokate       tairou
aren     -vA  eyokat-I -∅   tai -rU -X
1S.polite-TOP magpie-PL-NOM good-NPT-SP
"I like magpies."
```

#### Nominative

The nominative is used for agents and inanimate patients. It is unmarked.

```
イヨカツワ  オラ  イタス  
eyokappa   ora       etas
eyokat-vA  ora  -∅   eta-s  -∅
magpie-TOP acorn-NOM eat-PST-S2/3S
"The magpie ate an acorn."
```

```
オラワ  イヨカツ  イタス  
orava     eyokat     etas
ora  -vA  eyokat-∅   eta-s  -∅
acorn-TOP magpie-NOM eat-PST-S2/3S
"An acorn is what the magpie ate."
```

As a general rule, if the agent and patient are both in the nominative case, the agent should come first.

```
イヨカツ  オラ  イタス  
eyokat     ora       etas
eyokat-∅   ora  -∅   eta-s  -∅
magpie-NOM acorn-NOM eat-PST-S2/3S
"The magpie ate an acorn."
```

#### Genitive

The genitive, marked with *-A* and consonant lenition on the stem, has two distinct functions.

The first is to mark a noun as a possessor or modifier of another noun. Note that nouns also have possession suffixes, so having
a genitive modifier as well is generally not needed except to clarify third person possessors.

```
コニエ  チオロツ  
kunie     tiurut
kuni -'A  tiuru-t
woman-GEN home -P2/3S
"the woman's home"
```

A common type of modification other than possession is to use quality nouns attributively:

```
ヒツリヰツリエ  コニ  
hirrivirrie    kuni
hirrivirri-'A  kuni
sneaky    -GEN woman
"sneaky woman"
```

The second function is to mark animate patients.

```
コニエ  オヘシン  
kunie     uhesin
kuni -'A  uhe-s  -n
woman-GEN see-PST-S1S
"I saw the woman."
```

For the purpose of direct object marking, animals may be marked with either the nominative (treating them as animate)
or the genitive (treating them as inanimate), depending on the degree of emotional proximity to the animal the speaker wants
to convey. This is similar to how it's possible to use either "it" or gendered pronouns with animals in English.

```
イヨカヅア  オヘシン  
eyokara    uhesin
eyokat-'A  uhe-s  -n
magpie-GEN see-PST-S1S
"I saw the magpie."
```

```
イヨカツ  オヘシン  
eyokat     uhesin
eyokat-∅   uhe-s  -n
magpie-NOM see-PST-S1S
"I saw the magpie."
```

#### Partitive

The partitive, marked with *-t*, has two distinct functions. Notably, the partitive cannot co-occur with plural marking.

The first function is to mark a direct object where the verb was unsuccessful or did not reach a state of completion.
The plurality of the partitive noun is ambiguous.

```
イヨカツ  オラツ  イタス  
eyokat     orat      etas
eyokat-∅   ora  -t   eta-s  -∅
magpie-NOM acorn-PAR eat-PST-S2/3S
"The magpie ate some acorns."
```

Contrast this with

```
イヨカツ  オラ  イタス  
eyokat     ora       etas
eyokat-∅   ora  -∅   eta-s  -∅
magpie-NOM acorn-NOM eat-PST-S2/3S
"The magpie ate the acorn."
```

```
イヨカツ  オライ  イタス  
eyokat     orai         etas
eyokat-∅   ora  -I -∅   eta-s  -∅
magpie-NOM acorn-PL-NOM eat-PST-S2/3S
"The magpie ate the acorns."
```

The second function is to adjoin the noun to a cardinal number.

```
イヨカチツ  ソニ  チリ  
eyokatet   sone  tiri
eyokat-t   sone  tiri
magpie-PAR three CL.animal
"three magpies"
```

#### Equative

The equative case, marked with *-ssI* is used in two related ways.

The first way is to mark the complement of a copula:

```
コニスシ  ナロン  
kunissi   naron
kuni -ssI na-rU -n
woman-EQU be-NPT-S1S
"I am a woman."
```

The second way is to indicate manner or similarity:

```
ナツハスシ  コタリロ  
nappasse  kotarero
nappa-ssI kotare-rU -∅
stone-EQU hard  -NPT-S2/3S
"It is hard like a stone."
```

This includes using quality nouns adverbially:

```
ヒツリヰツリスシ  チオロエニ  ヤトシン  
hirrivirrissi  tiuruoni   yatosen
hirrivirri-ssI tiuru-'AnI yato-s  -n
sneaky    -EQU home -ALL  go  -PST-S1S
"I went home sneakily."
```

#### Dative

The dative case, marked with *-nyU*, is used to mark a benefactor or motive.

```
コニンヨ  チオロ  タタシン  
kuninyu   tiuru     tatasen
kuni -nyU tiuru-∅   tata -s  -n
woman-DAT home -NOM build-PST-S1S
"I built the woman a home."
```

```
ナツハツ  ヰチアンヨ  センゴエニ  ヤトシン  
nappat    veteanyo         senguoni   yatosen
nappa-t   vet -tI -'A -nyU senku-'AnI yato-s  -n
stone-PAR take-CNJ-GEN-DAT river-ALL  go  -PST-S1S
"I went to the river to get stones."
```

#### Adpositional cases

There are three cases indicating position or direction: the allative, locative, and ablative.
These three are all marked with suffixes based on the genitive suffix, and so all lenite the stem.

The allative is marked with the genitive suffix plus *-nI*, and indicates motion towards.

```
チオロエニ  ヤトス  
tiuruoni   yatos
tiuru-'AnI yato-s  -∅
home -ALL  go  -PST-S2/3S
"He went home."
```

The allative is also used for recipients.

```
イヨカヅアニ  オラ  オコシン  
eyokarane   ora       okosen
eyokat-'AnI ora  -∅   oko -s  -n
magpie-ALL  acorn-NOM give-PST-S1S
"I gave the magpie an acorn."
```

The locative is marked with the genitive suffix plus *-ttI*, and indicates location at.

```
チオロエツチ  ナロ  
tiuruotti   naro
tiuru-'AttI na-rU -∅
home -LOC   be-NPT-S2/3S
"He is at home."
```

The ablative is marked with the genitive suffix plus *-stI*, and indicates motion away.

```
チオロエスチ  ヤトス  
tiuruosti   yatos
tiuru-'AstI yato-s  -∅
home -ABL   go  -PST-S2/3S
"He left home."
```

### Relational nouns

There is a set of nouns which commonly appear as the possessum of other nouns, indicating a physical position or
some other type of relationship.

An example of such a noun is モ *mu* "inside":

```
チオロエ  モヅエツチ  ナロ  
tiuruo    muretti            naro
tiuru-'A  mu    -t    -'AttI na-rU -∅
home -GEN inside-P2/3S-LOC   be-NPT-S2/3S
"He is inside the house."
```

Some such nouns conventionally only appear with a particular case, such as キ *ki*, which
may mean "companion," but which almost always appears in the equative case with a comitative meaning:

```
コニヱ  キニスシ  ヤトス  
kunive    kinissi           yatos
kuni -vA  ki       -n  -ssI yato-s  -∅
woman-TOP companion-P1S-EQU go  -PST-S2/3S
"The woman went with me."
```

Other positional nouns don't require any case marking to be understood with an oblique meaning, such as トツヱ *tuppe* "without":

```
トツヱンミツ  
tuppemmit
tuppe  -mmIt
without-P1P
"without us"
```

#### Adpositional relationals

The adpositional cases, being based on the genitive case, actually originate as relational nouns.
The possession suffix has been dropped from the adpositional case suffixes, but the same roots can also function
similarly to トツヱ *tuppe* "without", as standalone relational with a possession suffix.
This is preferred over putting first person pronouns in the adpositional cases, and permissible for second and third persons rather than
using a pronoun if there is no risk of ambiguity, but may not be used with an actual noun possessor.

```
*ソ  アリンアニ  オコチ  
*so       arenane        okote
*so  -∅   aren     -'AnI oko -tI
that-NOM 1S.polite-ALL  give-CNJ
"Give it to me."
```

```
ソ  ニン  オコチ  
so       nin     okote
so  -∅   ni -n   oko -tI
that-NOM ALL-P1S give-CNJ
"Give it to me."
```

```
ソ  シエニ  オコロン  
so       sieni          okoron
so  -∅   si       -'AnI oko -rU -n
that-NOM 2S.formal-ALL  give-NPT-S1S
"I'll give it to you."
```

```
ソ  ニツ  オコロン  
so       nit       okoron
so  -∅   ni -t     oko -rU -n
that-NOM ALL-P2/3S give-NPT-S1S
"I'll give it to you/him/her."
```

```
ソ  コニエニ  オコロン  
so       kunieni    okoron
so  -∅   kuni -'AnI oko -rU -n
that-NOM woman-ALL  give-NPT-S1S
"I'll give it to the woman"
```

```
*ソ  コニエ  ニツ  オコロン  
*so       kunie     nit       okoron
*so  -∅   kuni -'A  ni -t     oko -rU -n
that-NOM woman-GEN ALL-P2/3S give-NPT-S1S
"I'll give it to the woman"
```

## Verbs

Verbs can take suffixes for, in order: voice, negation, tense, and subject agreement.

### Voice marking

Voice suffixes are optional and usually not present. The two voice suffixes are the passive/potential *-rA* and the causative *-sA*.

#### Passive/potential

The passive *-rA* is used when there is no specific agent mentioned or implied, to promote the patient to subject, i.e.,
it takes the syntactic position and case marking that the agent would normally take, and the verb conjugates for its person and number.
For animate patients, this means that the case can shift from genitive to nominative:

```
イヨカヅア  オヘロン  
eyokara    uherun
eyokat-'A  uhe-rU -n
magpie-GEN see-NPT-S1S
"I see the magpie."
```

```
イヨカツ  オヘレ  
eyokat     uhere
eyokat-∅   uhe-rA -∅
magpie-NOM see-PSV-S2/3S
"The magpie is seen."
```

However, the patient might not change case, e.g. if it is marked as the topic,
or if it is inanimate (inanimate nouns use the nominative case as both subject and object):

```
イヨカツワ  オヘロン  
eyokappa   uherun
eyokat-vA  uhe-rU -n
magpie-TOP see-NPT-S1S
"I see the magpie."
```

```
イヨカツワ  オヘレ  
eyokappa   uhere
eyokat-vA  uhe-rA -∅
magpie-TOP see-PSV-S2/3S
"The magpie is seen."
```

```
オラ  オヘロン  
ora       uherun
ora  -∅   uhe-rU -n
acorn-NOM see-NPT-S1S
"I see the acorn."
```

```
オラ  オヘレ  
ora       uhere
ora  -∅   uhe-rA -∅
acorn-NOM see-PSV-S2/3S
"The acorn is seen."
```

The same affix may be used with a potential sense.
This still entails promoting the object to subject; there may be ambiguity between the potential and plain passive.

```
イヨカツ  オヘレ  
eyokat     uhere
eyokat-∅   uhe-rA -∅
magpie-NOM see-PSV-S2/3S
"The magpie is seen/visible."
```

However, in the potential sense, an agent may be stated; if so, it is the topic, and the verb is still conjugated for the patient.

```
アリンワ  イヨカツ  オヘレ  
aremma        eyokat     uhere
aren     -vA  eyokat-∅   uhe-rA -∅
1S.polite-TOP magpie-NOM see-PSV-S2/3S
"I can see the magpie."
```

#### Causative

The causative, marked with *-sA*, is used to convey forcing or enabling someone or something to do something. The causee is marked with the dative case.

```
コニンヨ  チオロエニ  ヤトサシン  
kuninyu   tiuruoni   yatosasen
kuni -nyU tiuru-'AnI yato-sA -s  -n
woman-DAT home -ALL  go  -CAU-PST-S1S
"I made the woman go home."
```

### Negation

Verbs are negated with the suffix *-nA*.

```
ソ  イタナシン  
so       etanasen
so  -∅   eta-nA -s  -n
that-NOM eat-NEG-PST-S1S
"I didn't eat it."
```

```
イヨカツ  オヘレネ  
eyokat     uherene
eyokat-∅   uhe-rA -nA -∅
magpie-NOM see-PSV-NEG-S2/3S
"The magpie is not visible."
```

### Tense

Verbs can be marked for past or non-past tense.

#### Non-past

The non-past is versatile in use, being usable for simple present, gnomic, and future senses.

```
イヨカツワ  オラツ  イタロ  
eyokappa   orat      etaro
eyokat-vA  ora  -t   eta-rU -∅
magpie-TOP acorn-PAR eat-NPT-S2/3S
"The magpie eats acorns./The magpie is eating acorns./The magpie will eat acorns."
```

The explicit non-past suffix is *-rU*. However, it is only used if there are no preceding suffixes, i.e., no voice or negation marking.
If there are voice or negation suffixes applied, the non-past is zero-marked.

```
イヨカツワ  オラツ  イタナ  
eyokappa   orat      etana
eyokat-vA  ora  -t   eta-nA -∅
magpie-TOP acorn-PAR eat-NEG-S2/3S
"The magpie doesn't eat acorns./The magpie isn't eating acorns./The magpie won't eat acorns."
```

#### Past

The past tense is marked with *-s*.

```
イヨカツワ  オラツ  イタス  
eyokappa   orat      etas
eyokat-vA  ora  -t   eta-s  -∅
magpie-TOP acorn-PAR eat-PST-S2/3S
"The magpie ate acorns."
```

### Subject agreement

Sauna verbs conjugate for the subject, i.e., the agent of transitive verbs, and the one core argument of intransitive verbs,
including passives and verbs with adjectival meanings.

There is only a three-way distinction made in verb conjugation:

- first person singular is marked *-n*: ナロン *naron* "I am", ナシン *nasen* "I was"

- second and third person singular are unmarked: ナロ *naro* "you are/he, she is", ナス *nas* "you were/he, she was"

- plural subjects are all marked alike, by lengthening a final vowel if the stem ends in a short vowel, and zero-marking otherwise: ナロー *narou* "they are", ナス *nas* "they were"

### Conjunctive form

The conjunctive is non-finite form of verbs which replaces tense and person marking with a suffix *-tI*.
Its core function is as a converb, to convey simultaneous or sequential action by the same subject:

```
ヨロチ  チオロエニ  ヤトシン  
yuruti   tiuruoni   yatosen
yuru-tI  tiuru-'AnI yato-s  -n
sing-CNJ home -ALL  go  -PST-S1S
"I sang as I went home."
```

```
ナツハツ  ヰツチ  チオロ  タタシン  
nappat    vette    tiuru     tatasen
nappa-t   vet -tI  tiuru-∅   tata -s  -n
stone-PAR take-CNJ home -NOM build-PST-S1S
"I got some rocks and built a house."
```

As an extension of this, it is used with ナ *na* "be" to convey a progressive aspect:

```
イタチ  ナロン  
etate   naron
eta-tI  na-rU -n
eat-CNJ be-NPT-S1S
"I'm eating."
```

#### Nominalization

Verbs can be turned into nouns, which can either refer to a single instance of performing the verb or the concept of the verb,
by adding the conjunctive suffix and the genitive case suffix. Other case suffixes then stack on top of the genitive suffix.

```
アリンワ  イタヂア  タイロ  
aremma        etarea          tairo
aren     -vA  eta-tI -'A -∅   tai -rU -∅
1S.polite-TOP eat-CNJ-GEN-NOM good-NPT-S2/3S
"I like eating."
```

Arguments of the nominalized verb can branch leftward from the verb:

```
ナツハツ  ヰチアンヨ  センゴエニ  ヤトシン  
nappat    veteanyo         senguoni   yatosen
nappa-t   vet -tI -'A -nyU senku-'AnI yato-s  -n
stone-PAR take-CNJ-GEN-DAT river-ALL  go  -PST-S1S
"I went to the river to get stones."
```

## Numerals

### Counting system

Sauna's numerals are largely base ten, with multiples of ten up to one hundred grouped into twenties.

オ *u* "one"

イリ *iri* "two"

ソニ *sone* "three"

ニ𛄠 *neye* "four"

トイセ *tuise* "five"

チテ *tite* "six"

ヰキ *viki* "seven"

キキ *kiki* "eight"

オネ *une* "nine"

ヒスヨ *hesko* "ten"

ヒスヨオ *hesko u* "eleven"

ヒスヨイリ *hesko iri* "twelve"

タキ *take* "twenty"

タキヒスヨ *take hesko* "thirty"

イリタキ *iri take* "forty"

イリタキヒスヨ *iri take hesko* "fifty"

テイシス *teisis* "hundred"

オツ *ot* "thousand"

カイラオ *kairau* "ten thousand"

ソニカイラオチテオツトイセテイシスニ𛄠タキヒスヨキキ *sone kairau tite ot tuise teisis neye take hesko kiki* "36,598"

### Cardinal numbers and numeral classifiers

Sauna uses numeral classifiers with cardinal numbers:

```
イヨカチツ  ソニ  チリ  
eyokatet   sone  tiri
eyokat-t   sone  tiri
magpie-PAR three CL.animal
"three magpies"
```

The classifiers are effectively nouns. They take the case marking for the phrase:

```
イヨカチツ  ソニ  チリンヨ  
eyokatet   sone  tirinyu
eyokat-t   sone  tiri     -nyU
magpie-PAR three CL.animal-DAT
"for three magpies"
```

Some nouns can even double as numeral classifiers:

```
キーテツ  ソニ  チ  
kiitet    sone  te
kiite-t   sone  te
cup  -PAR three CL.general
"three cups"
```

```
ヰタツ  ソニ  キーテ  
vetat     sone  kiite
veta -t   sone  kiite
water-PAR three cup
"three cups of water"
```

Some common classifiers include:

- チ *te* can be used for any inanimate noun

- ナ *na* people, or anthropomorphized characters or deities

- チリ *tiri* birds and small to medium-sized mammals

- コイ *kui* small hard objects such as acorns, grains, beads, coins

- ヒセ *hise* pieces of a larger whole

### Ordinal numbers

To make numbers ordinal, they should appear as a modifier of the noun in the genitive:

```
ソニア  イヨカツ  
sonea     eyokat
sone -'A  eyokat
three-GEN magpie
"the third magpie"
```

"First" is irregular, being based on the word インソイ *ensoi* "beginning":

```
インソイア  イヨカツ  
ensoiya       eyokat
ensoi    -'A  eyokat
beginning-GEN magpie
"the first magpie"
```

## Pronouns

### Personal pronouns

There is a range of options for singular address, based on politeness and gender. Plural pronouns are more standard.

#### First person singular

First person pronouns are the most likely to be omitted, since verbs unambiguously conjugate for them as subject,
and the first person may be obvious from context in requests etc.

If an explicit first person pronoun is given, there are three main choices:

- アリン *aren* is the standard first person pronoun, used in formal and some casual situations

- オーラ *oura* is used by masculine people in most casual situations

- サ𛄠 *saye* sounds diminutive and may be used by feminine people casually, though it's not unusual for feminine people to use アリン *aren* in some casual situations or even avoid サ𛄠 *saye* consistently

#### Second person singular

It's common for names or nicknames to be used instead of second person pronouns. Second person pronouns are used in the following situations:

- シ *si* is used when the other person's name is not known, such as when addressing a stranger or in writing or media addressing a generic reader/listener

- アカラ *akara* is used to address someone of significantly higher status

- ワラー *varaa* may be used by a masculine speaker to a masculine listener, sounding somewhat rough or uncouth

#### Third person singular

It's almost universal to use names when talking about single other people.
The only exception is when talking about someone whose name is not known, in which case
ソナ *so na* or ノナ *no na* "that person" can be used.

#### Plural pronouns

Plural pronouns are standard and always used rather than names or some other manner of description.

ナイ *nai* is the exclusive first person plural pronoun ("we" excluding the listener).

キー *kii* is the inclusive equivalent ("we" including the listener).

チ *ti* is the second person plural pronoun.

ニ *ni* is the third person plural pronoun.

### Demonstrative pronouns

Sauna has three levels of distance for demonstratives:

- コ *ko* is proximal ("this")

- ソ *so* is medial ("that by you, that nearby")

- ノ *no* is distal ("that over there, yon")

There is also an indefinite pronoun ト *to* which by itself means "some, any" and is also used to form negative existentials and questions.

Demonstrative pronouns have similar syntax to numerals. They can appear by themselves:

```
ソ  ヰチシン  
so       vetesen
so  -∅   vet -s  -n
that-NOM take-PST-S1S
"I took it."
```

or directly with classifiers:

```
ソ  チ  ヰチシン  
so   te             vetesen
so   te        -∅   vet -s  -n
that CL.general-NOM take-PST-S1S
"I took that thing."
```

```
ノ  ナア  オヘシン  
no  naa          uhesin
no  na      -'A  uhe-s  -n
yon CL.human-GEN see-PST-S1S
"I saw that person."
```

To be attributive, demonstratives are marked with the genitive:

```
ノア  イヨカヅア  オヘシン  
noa     eyokara    uhesin
no -'A  eyokat-'A  uhe-s  -n
yon-GEN magpie-GEN see-PST-S1S
"I saw that magpie."
```

Demonstratives by themselves can mean "here" and "there" using the adpositional cases:

```
ノアニ  ヤトロン  
noane    yatoron
no -'AnI yato-rU -n
yon-ALL  go  -NPT-S1S
"I'm going there."
```

```
チオロン  コアツチ  ナロ  
tiurun        koatte     naro
tiuru-n  -∅   ko  -'AttI na-rU -∅
home -P1S-NOM this-LOC   be-NPT-S2/3S
"My house is here."
```

To retain the sense of a location, they must use the locative case as a topic, even though locations otherwise can be topic without overt locative marking:

```
チオロヱ  コニ  ナロ  
tiuruve   kuni      naro
tiuru-vA  kuni -∅   na-rU -∅
home -TOP woman-NOM be-NPT-S2/3S
"There's a woman in the house."
```

```
ノアツチワ  コニ  ナロ  
noatteva      kuni      naro
no -'AttI-vA  kuni -∅   na-rU -∅
yon-LOC  -TOP woman-NOM be-NPT-S2/3S
"There's a woman over there."
```

```
*ノワ  コニ  ナロ  
*nova    kuni      naro
*no -vA  kuni -∅   na-rU -∅
yon-TOP woman-NOM be-NPT-S2/3S
"There's a woman over there."
```

## Syntax

### General word order

The order of a Sauna clause is generally topic, subject (if it is distinct), oblique phrases, direct object, verb.
However, case marking makes this somewhat flexible; the only rigid rule is that the verb must come last.

```
コハワ  コニ  イヨカヅアニ  オラツ  オコス  
kohava    kuni      eyokarane   orat      okos
koha -vA  kuni -∅   eyokat-'AnI ora  -t   oko -s  -∅
today-TOP woman-NOM magpie-ALL  acorn-PAR give-PST-S2/3S
"Today the woman gave the magpie acorns."
```

If there are multiple predicates due to the use of the conjunctive verb form, their obliques and direct object are ordered
the same way before each respective verb:

```
ヰタ  ネテチ  チオロ  タタシン  
veta      neteti    tiuru     tatasen
veta -∅   nete -tI  tiuru-∅   tata -s  -n
water-NOM drink-CNJ home -NOM build-PST-S1S
"I drank water as I built the house."
```

### Existentials, to have, to be

The verb ナ *na* essentially means "to exist".

```
ノアツチワ  イヨカツ  ナロ  
noatteva      eyokat     naro
no -'AttI-vA  eyokat-∅   na-rU -∅
yon-LOC  -TOP magpie-NOM be-NPT-S2/3S
"There's a magpie over there."
```

```
チオロエ  モツヱ  ト  ナ  ナロ  
tiuruo    muppe            to   na           naro
tiuru-'A  mu    -t    -vA  to   na      -∅   na-rU -∅
home -GEN inside-P2/3S-TOP some CL.human-NOM be-NPT-S2/3S
"Someone is in the house."
```

With a topic, it can be used to mean "have":

```
アリンワ  チオロ  ナロ  
aremma        tiuru     naro
aren     -vA  tiuru-∅   na-rU -∅
1S.polite-TOP home -NOM be-NPT-S2/3S
"I have a house."
```

To equate two nouns, the equative case is used:

```
コニスシ  ナロン  
kunissi   naron
kuni -ssI na-rU -n
woman-EQU be-NPT-S1S
"I am a woman."
```

```
ソ  チオロスシ  ナロ  
so       tiurussi  naro
so  -∅   tiuru-ssI na-rU -∅
that-NOM home -EQU be-NPT-S2/3S
"It is a house."
```

#### Negative existentials

To say something doesn't exist, the indefinite pronoun plus a negative form of ナ *na* is used:

```
チオロエ  モツヱ  ト  ナ  ナナ  
tiuruo    muppe            to   na           nana
tiuru-'A  mu    -t    -vA  to   na      -∅   na-nA -∅
home -GEN inside-P2/3S-TOP some CL.human-NOM be-NEG-S2/3S
"There's no one in the house."
```

### Relative clauses

Relative clauses are expressed without a complementizer or relative pronoun;
some verb and its arguments simply branch leftward from the head noun, with the role of the
coreferent in the subclause being understood from gapping and context.

```
オヘシン  ヱレ  
uhesin      vere
uhe-s  -n   vere
see-PST-S1S man
"the man I saw"
```

```
シ  ニン  オコス  キーテ  
si            nin     okos           kiite
si       -∅   ni -n   oko -s  -∅     kiite
2S.formal-NOM ALL-P1S give-PST-S2/3S cup
"the cup you gave me"
```

The non-past suffix is dropped for the main verb of a relative clause:

The non-past suffix is dropped for the main verb of a relative clause:

```
タンガイ  ヱレ  
tangai       vere
tangai-∅     vere
travel-S2/3S man
"traveling man"
```

```
イ𛄡  チオロ  
iye         tiuru
iye  -∅     tiuru
small-S2/3S home
"small house"
```

However, subject person marking is still included:

```
イ𛄡ー  チオロイ  
iyea     tiurui
iye  -X  tiuru-I
small-SP home -PL
"small houses"
```

### Questions

Sauna uses an interrogative suffix *-kU* which attaches to some element of the sentence to form a question, going after all other suffixes.

For polar questions, this is generally the verb:

```
ノアニ  ヤトスコ  
noane    yatosko
no -'AnI yato-s  -∅    -kU
yon-ALL  go  -PST-S2/3S-INT
"Did you go there?"
```

although the interrogative particle can go on another element to specifically make it the contrastive focus of the question:

```
ノアニコ  ヤトス  
noaneko      yatos
no -'AnI-kU  yato-s  -∅
yon-ALL -INT go  -PST-S2/3S
"Is that where you went?"
```

For content questions, the interrogative particle is placed onto a noun phrase with the indefinite demonstrative.

```
トアニコ  ヤトス  
toaneko       yatos
to  -'AnI-kU  yato-s  -∅
some-ALL -INT go  -PST-S2/3S
"Where did you go?"
```

```
トコ  イタス  
toko         etas
to  -∅  -kU  eta-s  -∅
some-NOM-INT eat-PST-S2/3S
"What did you eat?"
```

```
シヱ  トア  キーテコ  タイロ  
sive          toa      kiiteku       tairo
si       -vA  to  -'A  kiite-∅  -kU  tai -rU -∅
2S.formal-TOP some-GEN cup  -NOM-INT good-NPT-S2/3S
"Which cup do you like?"
```

Combining the indefinite with classifiers is a productive way to more specifically ask "what", "which", or "who"?

```
シヱ  ト  ナアコ  タイロ  
sive          to   naako            tairo
si       -vA  to   na      -'A -kU  tai -rU -∅
2S.formal-TOP some CL.human-GEN-INT good-NPT-S2/3S
"Who do you like?"
```

### Imperatives and cohortatives

Sauna imperatives are formed by using the bare verb stem:

```
ニン  オコ  
nin     oko
ni -n   oko
ALL-P1S give
"Give it to me."
```

This can sound rather blunt. There is a verb キタイコア *ketaikoa* which is specifically used act as a softer imperative,
with the actual verb being requested coming before in conjunctive form:

```
ニン  オコチ  キタイコア  
nin     okote    ketaikoa
ni -n   oko -tI  ketaikoa
ALL-P1S give-CNJ do.please
"Please give it to me."
```

This can actually be overly formal and polite. The most neutral way to pose a request, neither blunt and urgent nor overly formal,
is to omit キタイコア *ketaikoa* and leave the bare conjunctive.

```
ニン  オコチ  
nin     okote
ni -n   oko -tI
ALL-P1S give-CNJ
"Give it to me."
```

The normal way to express a cohortative is with the non-past tense, plural conjugation plus interrogative particle.
Note that this literally just asks about a group, but the implication can be cohortative if it makes sense contextually:

```
ヤトローコ  
yatorouko
yato-rU -X -kU
go  -NPT-SP-INT
"Are they going?/Are we going?/Let's go!"
```

