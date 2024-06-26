//go:build unit || (database && ram)

package ram

import (
	"context"
	"fmt"
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
	"go.uber.org/atomic"

	v1 "github.com/shortlink-org/shortlink/boundaries/link/link/internal/domain/link/v1"
	"github.com/shortlink-org/shortlink/boundaries/link/link/internal/infrastructure/repository/crud/mock"
	"github.com/shortlink-org/shortlink/pkg/db/options"
)

var linkUniqId atomic.Int64

func BenchmarkRAMSerial(b *testing.B) {
	ctx := context.Background()

	b.Run("Create [single]", func(b *testing.B) {
		b.ReportAllocs()

		// create a db
		store := Store{}

		data := mock.AddLink

		for i := 0; i < b.N; i++ {
			data.Url = fmt.Sprintf("%s/%d", data.Url, i)
			_, err := store.Add(ctx, data)
			require.NoError(b, err)
		}
	})

	b.Run("Create [batch]", func(b *testing.B) {
		b.ReportAllocs()

		// create a db
		store := Store{}

		// Set config
		b.Setenv("STORE_MODE_WRITE", strconv.Itoa(options.MODE_BATCH_WRITE))

		data := mock.AddLink

		for i := 0; i < b.N; i++ {
			data.Url = fmt.Sprintf("%s/%d", data.Url, i)
			_, err := store.Add(ctx, data)
			require.NoError(b, err)
		}
	})
}

func BenchmarkRAMParallel(b *testing.B) {
	ctx := context.Background()

	b.Run("Create [single]", func(b *testing.B) {
		b.ReportAllocs()

		// create a db
		store := Store{}

		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				source, err := getLink()
				require.NoError(b, err)

				_, err = store.Add(ctx, source)
				require.NoError(b, err)
			}
		})
	})

	b.Run("Create [batch]", func(b *testing.B) {
		b.ReportAllocs()

		// Set config
		b.Setenv("STORE_MODE_WRITE", strconv.Itoa(options.MODE_BATCH_WRITE))

		// create a db
		store := Store{}

		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				source, err := getLink()
				require.NoError(b, err)

				_, err = store.Add(ctx, source)
				require.NoError(b, err)
			}
		})
	})
}

func getLink() (*v1.Link, error) {
	id := linkUniqId.Add(1)

	data := &v1.Link{
		Url:      fmt.Sprintf("%s/%d", "http://example.com", id),
		Describe: mock.AddLink.Describe,
	}

	if err := v1.NewURL(data); err != nil {
		return nil, err
	}

	return data, nil
}
