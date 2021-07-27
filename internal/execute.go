package internal

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"sync"

	"github.com/google/go-github/v37/github"
	"golang.org/x/oauth2"
)

// Repository represents a repository to be add and/or remove users from
type Repository struct {
	Owner string
	Name  string
}

// MapStringItoRepository converts a string slice into a Repository slice
func MapStringItoRepository(repos []string) ([]Repository, error) {
	list := make([]Repository, len(repos))

	for i := range repos {
		s := strings.Split(repos[i], "/")
		if len(s) != 2 {
			return list, fmt.Errorf("repository %s should be in the format {owner}/{name}", repos[i])
		}

		list[i] = Repository{Owner: s[0], Name: s[1]}
	}

	return list, nil
}

// Execute will remove and add user to the repositories
func Execute(
	token string,
	repos []Repository,
	toRm, toAdd []string,
) error {
	if len(token) == 0 {
		return errors.New("github token is invalid")
	}

	if len(repos) == 0 {
		return errors.New("should inform at least one repository")
	}

	if (len(toRm) + len(toAdd)) == 0 {
		return errors.New("should inform at least one user to remove or invite")
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	client := github.NewClient(oauth2.NewClient(ctx, oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)))

	var wg sync.WaitGroup
	done := make(chan error)

	syncWrap := func(fn func(r Repository, u string) error) func(r Repository, u string) {
		return func(r Repository, u string) {
			defer wg.Done()
			if err := fn(r, u); err != nil {
				done <- err
			}
		}
	}
	add := syncWrap(func(r Repository, u string) error {
		_, _, err := client.Repositories.AddCollaborator(ctx, r.Owner, r.Name, u, nil)
		return err
	})
	rm := syncWrap(func(r Repository, u string) error {
		_, err := client.Repositories.RemoveCollaborator(ctx, r.Owner, r.Name, u)
		return err
	})

	wg.Add((len(toRm) + len(toAdd)) * len(repos))
	for _, r := range repos {
		for _, u := range toRm {
			go add(r, u)
		}
		for _, u := range toAdd {
			go rm(r, u)
		}
	}

	go func() {
		wg.Wait()
		done <- nil
		close(done)
	}()

	return <-done
}
