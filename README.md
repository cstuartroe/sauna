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
*iuve* "birch tree" + *-t* partitive suffix -> *iuvet*, but
*eyokat* "magpie" + *-t* partitive suffix -> *eyokatet*.

## Epenthetic glides

If a stem ending in a short vowel receives a suffix beginning with a short vowel, the two coalesce into a long vowel or diphthong, e.g.
*iuve* "birch tree" + *-I* plural suffix -> *iuvei* "birch trees".

However, three vowel moras cannot occur in a row, so to prevent this, an epenthetic /ʋ/ (following a back vowel) or /j/ (otherwise) is inserted between
the stem and the suffix, e.g.
*iuvei* "birch trees" + *-A* genitive suffix -> *iuveiye* "of birch trees"

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

*nappa* "stone" -> *nahaa* "of (the) stone"

*senku* "river" -> *senguo* "of (the) river"

*eyokat* "magpie" -> *eyokara* "of (the) magpie"

# Grammar

## Syntax and typology

Sauna is a strictly head-final language, with SVO clause order and modifiers coming before nouns.

It is also strictly suffixing, with suffixes stacking in an agglutinative fashion.

Sauna is generally pro-drop, with arbitrary arguments being permitted to be left out of the sentence when clear from context:

```
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
Quality nouns are often ideophonic, and often involve reduplication: *hirrivirri* "sneaky"

Other parts of speech include numerals, noun classifiers, and relational nouns, although
noun classifiers and relational nouns are not entirely distinct from normal nouns.

## Nouns

Nouns are inflected for number, posessor, and case, with suffixes coming in that order.

### Number

The plural suffix is *-I*

*iuve* "birch tree" -> *iuvei* "birch trees"

*eyokat* "magpie" -> *eyokate* "magpies"

*nappa* "stone" -> *nappai* "stones"

### Possessor

Nouns are marked for their possessor:

- The first person singular possessor suffix is *-n*: *tiurun* "my home"

- The second/third person singular possessor suffix is *-t*: *tiurut* "your home, his home, her home"

- The first person plural possessor suffix is *-mmIt*: *tiurummit* "our home"

- The second/third person plural possessor suffix is *-vIt*: *tiuruvit* "y'all's home, their home"

### Cases

Sauna has nine cases.

Note that there is no dedicated case for direct objects,
which may be marked with the nominative, genitive, or partitive according to usage patterns described below.

#### Topical

Many indicative main clauses have a topic - if not, the topic is probably implied.
The topic is marked with the topical case marker *-vA* and generally comes first in the clause.
The topic may have a variety of semantic roles in the sentence, including agent, patient, experiencer, location, or possessor of some other argument.

```
kunive    tiuruoni   yatos            
kuni -vA  tiuru-'AnI yato-s  -∅       
woman-TOP home -ALL  go  -PST-S2/3S   
"The woman went home."
```

```
eyokateva     ereku       narairo           
eyokat-i -vA  ereku-∅     narai-rU -∅       
magpie-PL-TOP tail -NOM   long -NPT-S2/3S   
"Magpies have long tails."
```

The topic is often implied to be contrastive, especiialy when the its semantic role is something other than agent or patient.

```
veirumme   iuvei          iyeruu       
veirun-vA  iuve -i -∅     iye  -rU -X  
south -TOP birch-PL-NOM   small-NPT-SP 
"In the south, the birch trees are small."
```

For roles other than agent, patient, location, or possessor, the topical case suffix may be stacked on top of another:

```
tiuruonive     kuni        yatos            
tiuru-'AnI-vA  kuni -∅     yato-s  -∅       
home -ALL -TOP woman-NOM   go  -PST-S2/3S   
"Home is where the woman went."
```

For some constructions, a particular argument must be the topic.

```
eyokateva     ereku       naro           
eyokat-i -vA  ereku-∅     na-rU -∅       
magpie-PL-TOP tail -NOM   be-NPT-S2/3S   
"Magpies have tails."
```

```
aremma        eyokate         tairou      
aren     -vA  eyokat-i -∅     tai -rU -X  
1S.polite-TOP magpie-PL-NOM   good-NPT-SP 
"I like magpies."
```

#### Nominative

The nominative is used for agents and inanimate patients. It is unmarked.

```
eyokappa   ora         etas            
eyokat-vA  ora  -∅     eta-s  -∅       
magpie-TOP acorn-NOM   eat-PST-S2/3S   
"The magpie ate an acorn."
```

```
orava     eyokat       etas            
ora  -vA  eyokat-∅     eta-s  -∅       
acorn-TOP magpie-NOM   eat-PST-S2/3S   
"An acorn is what the magpie ate."
```

As a general rule, if the agent and patient are both in the nominative case, the agent should come first.

```
eyokat       ora         etas            
eyokat-∅     ora  -∅     eta-s  -∅       
magpie-NOM   acorn-NOM   eat-PST-S2/3S   
"The magpie ate an acorn."
```

#### Genitive

The genitive, marked with *-A* and consonant lenition on the stem, has two distinct functions.

The first is to mark a noun as a possessor or modifier of another noun. Note that nouns also have possession suffixes, so having
a genitive modifier as well is generally not needed except to clarify third person possessors.

```
kunie     tiurut      
kuni -'A  tiuru-t     
woman-GEN home -P2/3S 
"the woman's home"
```

A common type of modification other than possession is to use quality nouns attributively:

```
hirrivirrie    kuni  
hirrivirri-'A  kuni  
sneaky    -GEN woman 
"sneaky woman"
```

The second function is to mark animate patients.

```
kunie     uhesin      
kuni -'A  uhe-s  -n   
woman-GEN see-PST-S1S 
"I saw the woman."
```

For the purpose of direct object marking, animals may be marked with either the nominative (treating them as animate)
or the genitive (treating them as inanimate), depending on the degree of emotional proximity to the animal the speaker wants
to convey. This is similar to how it's possible to use either "it" or gendered pronouns with animals in English.

```
eyokara    uhesin      
eyokat-'A  uhe-s  -n   
magpie-GEN see-PST-S1S 
"I saw the magpie."
```

```
eyokat       uhesin      
eyokat-∅     uhe-s  -n   
magpie-NOM   see-PST-S1S 
"I saw the magpie."
```

#### Partitive

The partitive, marked with *-t*, has two distinct functions. Notably, the partitive cannot co-occur with plural marking.

The first function is to mark a direct object where the verb was unsuccessful or did not reach a state of completion.
The plurality of the partitive noun is ambiguous.

```
eyokat       orat      etas            
eyokat-∅     ora  -t   eta-s  -∅       
magpie-NOM   acorn-PAR eat-PST-S2/3S   
"The magpie ate some acorns."
```

Contrast this with

```
eyokat       ora         etas            
eyokat-∅     ora  -∅     eta-s  -∅       
magpie-NOM   acorn-NOM   eat-PST-S2/3S   
"The magpie ate the acorn."
```

```
eyokat       orai           etas            
eyokat-∅     ora  -i -∅     eta-s  -∅       
magpie-NOM   acorn-PL-NOM   eat-PST-S2/3S   
"The magpie ate the acorns."
```

The second function is to adjoin the noun to a cardinal number.

```
eyokatet   sone  tiri      
eyokat-t   sone  tiri      
magpie-PAR three CL.animal 
"three magpies"
```

#### Equative

The equative case, marked with *-ssI* is used in two related ways.

The first way is to mark the complement of a copula:

```
kunissi   naron      
kuni -ssI na-rU -n   
woman-EQU be-NPT-S1S 
"I am a woman."
```

The second way is to indicate manner or similarity:

```
nappasse  kotarero           
nappa-ssI kotare-rU -∅       
stone-EQU hard  -NPT-S2/3S   
"It is hard like a stone."
```

This includes using quality nouns adverbially:

```
hirrivirrissi  tiuruoni   yatosen      
hirrivirri-ssI tiuru-'AnI yato-s  -n   
sneaky    -EQU home -ALL  go  -PST-S1S 
"I went home sneakily."
```

#### Dative

The dative case, marked with *-nyU*, is used to mark a benefactor or motive.

```
kuninyu   tiuru       tatasen       
kuni -nyU tiuru-∅     tata -s  -n   
woman-DAT home -NOM   build-PST-S1S 
"I built the woman a home."
```

```
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
tiuruoni   yatos            
tiuru-'AnI yato-s  -∅       
home -ALL  go  -PST-S2/3S   
"He went home."
```

The allative is also used for recipients.

```
eyokarane   ora         okosen       
eyokat-'AnI ora  -∅     oko -s  -n   
magpie-ALL  acorn-NOM   give-PST-S1S 
"I gave the magpie an acorn."
```

The locative is marked with the genitive suffix plus *-ttI*, and indicates location at.

```
tiuruotti   naro           
tiuru-'AttI na-rU -∅       
home -LOC   be-NPT-S2/3S   
"He is at home."
```

The ablative is marked with the genitive suffix plus *-stI*, and indicates motion away.

```
tiuruosti   yatos            
tiuru-'AstI yato-s  -∅       
home -ABL   go  -PST-S2/3S   
"He left home."
```

### Relational nouns

There is a set of nouns which commonly appear as the possessum of other nouns, indicating a physical position or
some other type of relationship.

An example of such a noun is *mu* "inside":

```
tiuruo    muretti            naro           
tiuru-'A  mu    -t    -'AttI na-rU -∅       
home -GEN inside-P2/3S-LOC   be-NPT-S2/3S   
"He is inside the house."
```

Some such nouns conventionally only appear with a particular case, such as *ki*, which
may mean "companion," but which almost always appears in the equative case with a comitative meaning:

```
kunive    kinissi           yatos            
kuni -vA  ki       -n  -ssI yato-s  -∅       
woman-TOP companion-P1S-EQU go  -PST-S2/3S   
"The woman went with me."
```

Other positional nouns don't require any case marking to be understood with an oblique meaning, such as *tuppe* "without":

```
tuppemmit    
tuppe  -mmIt 
without-P1P  
"without us"
```

#### Adpositional relationals

The adpositional cases, being based on the genitive case, actually originate as relational nouns.
The possession suffix has been dropped from the adpositional case suffixes, but the same roots can also function
similarly to *tuppe* "without", as standalone relational with a possession suffix.
This is preferred over putting first person pronouns in the adpositional cases, and permissible for second and third persons rather than
using a pronoun if there is no risk of ambiguity, but may not be used with an actual noun possessor.

```
*so         arenane        okote    
*so  -∅     aren     -'AnI oko -tI  
 that-NOM   1S.polite-ALL  give-CNJ 
"Give it to me."
```

```
so         nin     okote    
so  -∅     ni -n   oko -tI  
that-NOM   ALL-P1S give-CNJ 
"Give it to me."
```

```
so         sieni          okoron       
so  -∅     si       -'AnI oko -rU -n   
that-NOM   2S.formal-ALL  give-NPT-S1S 
"I'll give it to you."
```

```
so         nit       okoron       
so  -∅     ni -t     oko -rU -n   
that-NOM   ALL-P2/3S give-NPT-S1S 
"I'll give it to you/him/her."
```

```
so         kunieni    okoron       
so  -∅     kuni -'AnI oko -rU -n   
that-NOM   woman-ALL  give-NPT-S1S 
"I'll give it to the woman"
```

```
*so         kunie     nit       okoron       
*so  -∅     kuni -'A  ni -t     oko -rU -n   
 that-NOM   woman-GEN ALL-P2/3S give-NPT-S1S 
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
eyokara    uherun      
eyokat-'A  uhe-rU -n   
magpie-GEN see-NPT-S1S 
"I see the magpie."
```

```
eyokat       uhere           
eyokat-∅     uhe-rA -∅       
magpie-NOM   see-PSV-S2/3S   
"The magpie is seen."
```

However, the patient might not change case, e.g. if it is marked as the topic,
or if it is inanimate (inanimate nouns use the nominative case as both subject and object):

```
eyokappa   uherun      
eyokat-vA  uhe-rU -n   
magpie-TOP see-NPT-S1S 
"I see the magpie."
```

```
eyokappa   uhere           
eyokat-vA  uhe-rA -∅       
magpie-TOP see-PSV-S2/3S   
"The magpie is seen."
```

```
ora         uherun      
ora  -∅     uhe-rU -n   
acorn-NOM   see-NPT-S1S 
"I see the acorn."
```

```
ora         uhere           
ora  -∅     uhe-rA -∅       
acorn-NOM   see-PSV-S2/3S   
"The acorn is seen."
```

The same affix may be used with a potential sense.
This still entails promoting the object to subject; there may be ambiguity between the potential and plain passive.

```
eyokat       uhere           
eyokat-∅     uhe-rA -∅       
magpie-NOM   see-PSV-S2/3S   
"The magpie is seen/visible."
```

However, in the potential sense, an agent may be stated; if so, it is the topic, and the verb is still conjugated for the patient.

```
aremma        eyokat       uhere           
aren     -vA  eyokat-∅     uhe-rA -∅       
1S.polite-TOP magpie-NOM   see-PSV-S2/3S   
"I can see the magpie."
```

#### Causative

The causative, marked with *-sA*, is used to convey forcing or enabling someone or something to do something. The causee is marked with the dative case.

```
kuninyu   tiuruoni   yatosasen        
kuni -nyU tiuru-'AnI yato-sA -s  -n   
woman-DAT home -ALL  go  -CAU-PST-S1S 
"I made the woman go home."
```

### Negation

Verbs are negated with the suffix *-nA*.

```
so         etanasen        
so  -∅     eta-nA -s  -n   
that-NOM   eat-NEG-PST-S1S 
"I didn't eat it."
```

```
eyokat       uherene             
eyokat-∅     uhe-rA -nA -∅       
magpie-NOM   see-PSV-NEG-S2/3S   
"The magpie is not visible."
```

### Tense

Verbs can be marked for past or non-past tense.

#### Non-past

The non-past is versatile in use, being usable for simple present, gnomic, and future senses.

```
eyokappa   orat      etaro           
eyokat-vA  ora  -t   eta-rU -∅       
magpie-TOP acorn-PAR eat-NPT-S2/3S   
"The magpie eats acorns./The magpie is eating acorns./The magpie will eat acorns."
```

The explicit non-past suffix is *-rU*. However, it is only used if there are no preceding suffixes, i.e., no voice or negation marking.
If there are voice or negation suffixes applied, the non-past is zero-marked.

```
eyokappa   orat      etana           
eyokat-vA  ora  -t   eta-nA -∅       
magpie-TOP acorn-PAR eat-NEG-S2/3S   
"The magpie doesn't eat acorns./The magpie isn't eating acorns./The magpie won't eat acorns."
```

#### Past

The past tense is marked with *-s*.

```
eyokappa   orat      etas            
eyokat-vA  ora  -t   eta-s  -∅       
magpie-TOP acorn-PAR eat-PST-S2/3S   
"The magpie ate acorns."
```

### Subject agreement

Sauna verbs conjugate for the subject, i.e., the agent of transitive verbs, and the one core argument of intransitive verbs,
including passives and verbs with adjectival meanings.

There is only a three-way distinction made in verb conjugation:

- first person singular is marked *-n*: *naron* "I am", *nasen* "I was"

- second and third person singular are unmarked: *naro* "you are/he, she is", *nas* "you were/he, she was"

- plural subjects are all marked alike, by lengthening a final vowel if the stem ends in a short vowel, and zero-marking otherwise: *narou* "they are", *nas* "they were"

### Conjunctive form

The conjunctive is non-finite form of verbs which replaces tense and person marking with a suffix *-tI*.
Its core function is as a converb, to convey simultaneous or sequential action by the same subject:

```
yuruti   tiuruoni   yatosen      
yuru-tI  tiuru-'AnI yato-s  -n   
sing-CNJ home -ALL  go  -PST-S1S 
"I sang as I went home."
```

```
nappat    vette    tiuru       tatasen       
nappa-t   vet -tI  tiuru-∅     tata -s  -n   
stone-PAR take-CNJ home -NOM   build-PST-S1S 
"I got some rocks and built a house."
```

As an extension of this, it is used with *na* "be" to convey a progressive aspect:

```
etate   naron      
eta-tI  na-rU -n   
eat-CNJ be-NPT-S1S 
"I'm eating."
```

#### Nominalization

Verbs can be turned into nouns, which can either refer to a single instance of performing the verb or the concept of the verb,
by adding the conjunctive suffix and the genitive case suffix. Other case suffixes then stack on top of the genitive suffix.

```
aremma        etarea            tairo            
aren     -vA  eta-tI -'A -∅     tai -rU -∅       
1S.polite-TOP eat-CNJ-GEN-NOM   good-NPT-S2/3S   
"I like eating."
```

Arguments of the nominalized verb can branch leftward from the verb:

```
nappat    veteanyo         senguoni   yatosen      
nappa-t   vet -tI -'A -nyU senku-'AnI yato-s  -n   
stone-PAR take-CNJ-GEN-DAT river-ALL  go  -PST-S1S 
"I went to the river to get stones."
```

## Numerals

### Counting system

Sauna's numerals are largely base ten, with multiples of ten up to one hundred grouped into twenties.

*u* "one"

*iri* "two"

*sone* "three"

*neye* "four"

*tuise* "five"

*tite* "six"

*viki* "seven"

*kiki* "eight"

*une* "nine"

*hesko* "ten"

*hesko u* "eleven"

*hesko iri* "twelve"

*take* "twenty"

*take hesko* "thirty"

*iri take* "forty"

*iri take hesko* "fifty"

*teisis* "hundred"

*ot* "thousand"

*kairau* "ten thousand"

*sone kairau tite ot tuise teisis neye take hesko kiki* "36,598"

### Cardinal numbers and numeral classifiers

Sauna uses numeral classifiers with cardinal numbers:

```
eyokatet   sone  tiri      
eyokat-t   sone  tiri      
magpie-PAR three CL.animal 
"three magpies"
```

The classifiers are effectively nouns. They take the case marking for the phrase:

```
eyokatet   sone  tirinyu       
eyokat-t   sone  tiri     -nyU 
magpie-PAR three CL.animal-DAT 
"for three magpies"
```

Some nouns can even double as numeral classifiers:

```
kiitet    sone  te         
kiite-t   sone  te         
cup  -PAR three CL.general 
"three cups"
```

```
vetat     sone  kiite 
veta -t   sone  kiite 
water-PAR three cup   
"three cups of water"
```

Some common classifiers include:

- *te* can be used for any inanimate noun

- *na* people, or anthropomorphized characters or deities

- *tiri* birds and small to medium-sized mammals

- *kui* small hard objects such as acorns, grains, beads, coins

- *hise* pieces of a larger whole

### Ordinal numbers

To make numbers ordinal, they should appear as a modifier of the noun in the genitive:

```
sonea     eyokat 
sone -'A  eyokat 
three-GEN magpie 
"the third magpie"
```

"First" is irregular, being based on the word *ensoi* "beginning":

```
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

- *aren* is the standard first person pronoun, used in formal and some casual situations

- *oura* is used by masculine people in most casual situations

- *saye* sounds diminutive and may be used by feminine people casually, though it's not unusual for feminine people to use *aren* in some casual situations or even avoid *saye* consistently

#### Second person singular

It's common for names or nicknames to be used instead of second person pronouns. Second person pronouns are used in the following situations:

- *si* is used when the other person's name is not known, such as when addressing a stranger or in writing or media addressing a generic reader/listener

- *akara* is used to address someone of significantly higher status

- *varaa* may be used by a masculine speaker to a masculine listener, sounding somewhat rough or uncouth

#### Third person singular

It's almost universal to use names when talking about single other people.
The only exception is when talking about someone whose name is not known, in which case
*so na* or *no na* "that person" can be used.

#### Plural pronouns

Plural pronouns are standard and always used rather than names or some other manner of description.

*nai* is the exclusive first person plural pronoun ("we" excluding the listener).

*kii* is the inclusive equivalent ("we" including the listener).

*ti* is the second person plural pronoun.

*ni* is the third person plural pronoun.

### Demonstrative pronouns

Sauna has three levels of distance for demonstratives:

- *ko* is proximal ("this")

- *so* is medial ("that by you, that nearby")

- *no* is distal ("that over there, yon")

There is also an indefinite pronoun *to* which by itself means "some, any" and is also used to form negative existentials and questions.

Demonstrative pronouns have similar syntax to numerals. They can appear by themselves:

```
so         vetesen      
so  -∅     vet -s  -n   
that-NOM   take-PST-S1S 
"I took it."
```

or directly with classifiers:

```
so   te               vetesen      
so   te        -∅     vet -s  -n   
that CL.general-NOM   take-PST-S1S 
"I took that thing."
```

```
no  naa          uhesin      
no  na      -'A  uhe-s  -n   
yon CL.human-GEN see-PST-S1S 
"I saw that person."
```

To be attributive, demonstratives are marked with the genitive:

```
noa     eyokara    uhesin      
no -'A  eyokat-'A  uhe-s  -n   
yon-GEN magpie-GEN see-PST-S1S 
"I saw that magpie."
```

Demonstratives by themselves can mean "here" and "there" using the adpositional cases:

```
noane    yatoron      
no -'AnI yato-rU -n   
yon-ALL  go  -NPT-S1S 
"I'm going there."
```

```
tiurun          koatte     naro           
tiuru-n  -∅     ko  -'AttI na-rU -∅       
home -P1S-NOM   this-LOC   be-NPT-S2/3S   
"My house is here."
```

To retain the sense of a location, they must use the locative case as a topic, even though locations otherwise can be topic without overt locative marking:

```
tiuruve   kuni        naro           
tiuru-vA  kuni -∅     na-rU -∅       
home -TOP woman-NOM   be-NPT-S2/3S   
"There's a woman in the house."
```

```
noatteva      kuni        naro           
no -'AttI-vA  kuni -∅     na-rU -∅       
yon-LOC  -TOP woman-NOM   be-NPT-S2/3S   
"There's a woman over there."
```

```
*nova    kuni        naro           
*no -vA  kuni -∅     na-rU -∅       
 yon-TOP woman-NOM   be-NPT-S2/3S   
"There's a woman over there."
```

## Syntax

### General word order

The order of a Sauna clause is generally topic, subject (if it is distinct), oblique phrases, direct object, verb.
However, case marking makes this somewhat flexible; the only rigid rule is that the verb must come last.

```
kohava    kuni        eyokarane   orat      okos             
koha -vA  kuni -∅     eyokat-'AnI ora  -t   oko -s  -∅       
today-TOP woman-NOM   magpie-ALL  acorn-PAR give-PST-S2/3S   
"Today the woman gave the magpie acorns."
```

If there are multiple predicates due to the use of the conjunctive verb form, their obliques and direct object are ordered
the same way before each respective verb:

```
veta        neteti    tiuru       tatasen       
veta -∅     nete -tI  tiuru-∅     tata -s  -n   
water-NOM   drink-CNJ home -NOM   build-PST-S1S 
"I drank water as I built the house."
```

### Existentials, to have, to be

The verb *na* essentially means "to exist".

```
noatteva      eyokat       naro           
no -'AttI-vA  eyokat-∅     na-rU -∅       
yon-LOC  -TOP magpie-NOM   be-NPT-S2/3S   
"There's a magpie over there."
```

```
tiuruo    muppe            to   na             naro           
tiuru-'A  mu    -t    -vA  to   na      -∅     na-rU -∅       
home -GEN inside-P2/3S-TOP some CL.human-NOM   be-NPT-S2/3S   
"Someone is in the house."
```

With a topic, it can be used to mean "have":

```
aremma        tiuru       naro           
aren     -vA  tiuru-∅     na-rU -∅       
1S.polite-TOP home -NOM   be-NPT-S2/3S   
"I have a house."
```

To equate two nouns, the equative case is used:

```
kunissi   naron      
kuni -ssI na-rU -n   
woman-EQU be-NPT-S1S 
"I am a woman."
```

```
so         tiurussi  naro           
so  -∅     tiuru-ssI na-rU -∅       
that-NOM   home -EQU be-NPT-S2/3S   
"It is a house."
```

#### Negative existentials

To say something doesn't exist, the indefinite pronoun plus a negative form of *na* is used:

```
tiuruo    muppe            to   na             nana           
tiuru-'A  mu    -t    -vA  to   na      -∅     na-nA -∅       
home -GEN inside-P2/3S-TOP some CL.human-NOM   be-NEG-S2/3S   
"There's no one in the house."
```

### Relative clauses

Relative clauses are expressed without a complementizer or relative pronoun;
some verb and its arguments simply branch leftward from the head noun, with the role of the
coreferent in the subclause being understood from gapping and context.

```
uhesin      vere 
uhe-s  -n   vere 
see-PST-S1S man  
"the man I saw"
```

```
si              nin     okos             kiite 
si       -∅     ni -n   oko -s  -∅       kiite 
2S.formal-NOM   ALL-P1S give-PST-S2/3S   cup   
"the cup you gave me"
```

The non-past suffix is dropped for the main verb of a relative clause:

The non-past suffix is dropped for the main verb of a relative clause:

```
tangai         vere 
tangai-∅       vere 
travel-S2/3S   man  
"traveling man"
```

```
iye           tiuru 
iye  -∅       tiuru 
small-S2/3S   home  
"small house"
```

However, subject person marking is still included:

```
iyea     tiurui   
iye  -X  tiuru-i  
small-SP home -PL 
"small houses"
```

### Questions

Sauna uses an interrogative suffix *-kU* which attaches to some element of the sentence to form a question, going after all other suffixes.

For polar questions, this is generally the verb:

```
noane    yatosko              
no -'AnI yato-s  -∅    -kU    
yon-ALL  go  -PST-S2/3S-INT   
"Did you go there?"
```

although the interrogative particle can go on another element to specifically make it the contrastive focus of the question:

```
noaneko      yatos            
no -'AnI-kU  yato-s  -∅       
yon-ALL -INT go  -PST-S2/3S   
"Is that where you went?"
```

For content questions, the interrogative particle is placed onto a noun phrase with the indefinite demonstrative.

```
toaneko       yatos            
to  -'AnI-kU  yato-s  -∅       
some-ALL -INT go  -PST-S2/3S   
"Where did you go?"
```

```
toko           etas            
to  -∅  -kU    eta-s  -∅       
some-NOM-INT   eat-PST-S2/3S   
"What did you eat?"
```

```
sive          toa      kiiteku         tairo            
si       -vA  to  -'A  kiite-∅  -kU    tai -rU -∅       
2S.formal-TOP some-GEN cup  -NOM-INT   good-NPT-S2/3S   
"Which cup do you like?"
```

Combining the indefinite with classifiers is a productive way to more specifically ask "what", "which", or "who"?

```
sive          to   naako            tairo            
si       -vA  to   na      -'A -kU  tai -rU -∅       
2S.formal-TOP some CL.human-GEN-INT good-NPT-S2/3S   
"Who do you like?"
```

### Imperatives and cohortatives

Sauna imperatives are formed by using the bare verb stem:

```
nin     oko  
ni -n   oko  
ALL-P1S give 
"Give it to me."
```

This can sound rather blunt. There is a verb *ketaikoa* which is specifically used act as a softer imperative,
with the actual verb being requested coming before in conjunctive form:

```
nin     okote    ketaikoa  
ni -n   oko -tI  ketaikoa  
ALL-P1S give-CNJ do.please 
"Please give it to me."
```

This can actually be overly formal and polite. The most neutral way to pose a request, neither blunt and urgent nor overly formal,
is to omit *ketaikoa* and leave the bare conjunctive.

```
nin     okote    
ni -n   oko -tI  
ALL-P1S give-CNJ 
"Give it to me."
```

The normal way to express a cohortative is with the non-past tense, plural conjugation plus interrogative particle.
Note that this literally just asks about a group, but the implication can be cohortative if it makes sense contextually:

```
yatorouko       
yato-rU -X -kU  
go  -NPT-SP-INT 
"Are they going?/Are we going?/Let's go!"
```

