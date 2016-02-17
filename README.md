## Similo

This is a very simple library that provides similarity check between strings, there are 3 functions that allow users to check similarity between strings:

```go
SimilarityBetween(first, second string )
ContainsSimilars(base string, lookup string, acceptance float64) bool
FindSimilars(base string, lookup string, acceptance float64) []string{}
  ```

#### SimilarityBetween

Receives two strings and returns a number between 0 and 1 indicating the similarity between the two strings.

#### FindSimilars

Receives a string (base) that is searched similarities inside another one (lookup) with similarity rate > than passed float value (acceptance), it returns a string slice with all matches > than acceptance float.

#### ContainsSimilars

Receives a string (base) that is searched similarities inside another one (lookup) with similarity rate > than passed float value (acceptance), it returns true if it find at least one match > than acceptance.
