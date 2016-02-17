package similo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEqual(t *testing.T) {
	assert.Equal(t, SimilarityBetween("hola", "hola"), 1.0)
}

func TestSimilar(t *testing.T) {
	assert.NotEqual(t, SimilarityBetween("hola", "mola"), 1.0)
	assert.Equal(t, SimilarityBetween("hola", "mola"), 0.75)
}

func TestDifferentSizes(t *testing.T) {
	assert.NotEqual(t, SimilarityBetween("hola", "hola mundo"), 1.0)
	assert.Equal(t, SimilarityBetween("hola", "hola mundo"), 0.4)
	assert.NotEqual(t, SimilarityBetween("hola mundo", "hola"), 1.0)
	assert.Equal(t, SimilarityBetween("hola mundo", "hola"), 0.4)
}

func TestFindSimilars(t *testing.T) {
	similars := FindSimilars("hola", "la bola hola", 0.5)
	assert.Equal(t, similars, []string{"ho", "hol", "hola", "ol", "ola", "la"})

	similars = FindSimilars("hola", "la bola hola", 0.75)
	assert.Equal(t, similars, []string{"hol", "hola", "ola"})

	similars = FindSimilars("hola", "la bola hola", 1)
	assert.Equal(t, similars, []string{"hola"})
}

func TestContainsSimilars(t *testing.T) {
	assert.True(t, ContainsSimilars("hola", "bola mundial", 0.6))
	assert.True(t, ContainsSimilars("hola", "bola", 0.6))
	assert.True(t, ContainsSimilars("hola", "la", 0.5))
}
