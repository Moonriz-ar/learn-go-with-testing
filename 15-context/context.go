package context

import (
	"context"
	"fmt"
	"net/http"
)

// software often kicks off long running, resource intensive processes. If the action that caused this get cancelled or fails for some reason, you need to stop these processes in a consistent way through your application.

type Store interface {
	Fetch(ctx context.Context) (string, error)
}

func Server(store Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data, err := store.Fetch(r.Context())

		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Fprint(w, data)
	}
}
