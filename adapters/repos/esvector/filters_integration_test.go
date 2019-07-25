// +build integrationTest

package esvector

import (
	"context"
	"fmt"
	"testing"

	"github.com/elastic/go-elasticsearch/v5"
	"github.com/go-openapi/strfmt"
	"github.com/semi-technologies/weaviate/entities/filters"
	"github.com/semi-technologies/weaviate/entities/schema"
	"github.com/semi-technologies/weaviate/entities/schema/kind"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_Filters(t *testing.T) {
	client, err := elasticsearch.NewClient(elasticsearch.Config{
		Addresses: []string{"http://localhost:9201"},
	})
	require.Nil(t, err)
	waitForEsToBeReady(t, client)

	// logger, _ := test.NewNullLogger()
	logger := logrus.New()
	repo := NewRepo(client, logger)
	migrator := NewMigrator(repo)

	t.Run("prepare test schema and data ",
		prepareTestSchemaAndData(repo, migrator))

	type test struct {
		name        string
		filter      *filters.LocalFilter
		expectedLen int
		expectedIDs []strfmt.UUID
	}

	// operators
	eq := filters.OperatorEqual

	// datatypes
	dtInt := schema.DataTypeInt
	tests := []test{
		{
			name:        "horsepower == 130",
			filter:      buildFilter("horsepower", 130, eq, dtInt),
			expectedLen: 1,
			expectedIDs: []strfmt.UUID{carSprinterID},
		},
	}

	for i, test := range tests {
		res, err := repo.VectorClassSearch(context.Background(), kind.Thing,
			carClass.Class, []float32{0.1, 0.1, 0.1, 1.1, 0.1}, 100, test.filter)
		require.Nil(t, err)
		require.Len(t, res, test.expectedLen)
		if len(test.expectedIDs) != test.expectedLen {
			t.Fatalf("wrong test setup at pos %d: lens dont match: %d and %d",
				i, test.expectedLen, len(test.expectedIDs))
		}

		ids := make([]strfmt.UUID, test.expectedLen, test.expectedLen)
		for pos, concept := range res {
			ids[pos] = concept.ID
		}
		assert.ElementsMatch(t, ids, test.expectedIDs, "ids dont match")
	}
}

func prepareTestSchemaAndData(repo *Repo,
	migrator *Migrator) func(t *testing.T) {
	return func(t *testing.T) {
		t.Run("creating the class", func(t *testing.T) {
			require.Nil(t,
				migrator.AddClass(context.Background(), kind.Thing, carClass))
		})

		for i, fixture := range cars {
			t.Run(fmt.Sprintf("importing car %d", i), func(t *testing.T) {
				require.Nil(t,
					repo.PutThing(context.Background(), &fixture, carVectors[i]))
			})
		}
	}
}

func buildFilter(propName string, value interface{}, operator filters.Operator, schemaType schema.DataType) *filters.LocalFilter {
	return &filters.LocalFilter{
		Root: &filters.Clause{
			Operator: operator,
			On: &filters.Path{
				Class:    schema.ClassName(carClass.Class),
				Property: schema.PropertyName(propName),
			},
			Value: &filters.Value{
				Value: value,
				Type:  schemaType,
			},
		},
	}
}