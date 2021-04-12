package main

import (
	"context"
	"encoding/json"
	"flag"
	"log"
	"net/http"
	"sort"
	"strings"
	"time"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

const (
	importersURLPrefix = "https://api.godoc.org/importers/"
)

type Importer struct {
	Path string
}

type Response struct {
	Results []Importer
}

type Pair struct {
	Key   string
	Value int
}

type PairList []Pair

func (p PairList) Len() int           { return len(p) }
func (p PairList) Less(i, j int) bool { return p[i].Value < p[j].Value }
func (p PairList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func rankByStar(repos map[string]int) PairList {
	pl := make(PairList, len(repos))
	i := 0
	for k, v := range repos {
		pl[i] = Pair{k, v}
		i++
	}
	sort.Sort(sort.Reverse(pl))
	return pl
}

func githubRepoSearch(ctx context.Context, client *github.Client, query string, repoStars map[string]int) {
	result, response, err := client.Search.Repositories(ctx, query, &github.SearchOptions{
		ListOptions: github.ListOptions{
			PerPage: 100,
		},
	})
	if response.Rate.Remaining <= 5 {
		diff := response.Rate.Reset.Sub(time.Now())
		log.Println("Rate limit reached. Sleeping.")
		time.Sleep(diff)
	}
	if err != nil {
		log.Println(err)
		return
	}
	for _, githubRepo := range result.Repositories {
		repoStars[*githubRepo.HTMLURL] = *githubRepo.StargazersCount
	}
}

func main() {
	gopkg := flag.String("pkg", "", "Package to discover")
	githubToken := flag.String("token", "", "Github access token")
	flag.Parse()
	if *gopkg == "" {
		log.Fatal("Package name is required.")
	}
	log.Printf("Fetching importer of pkg: %s", *gopkg)

	resp, err := http.Get(importersURLPrefix + *gopkg)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	response := Response{}
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		log.Fatal(err)
	}

	repos := make(map[string]struct{})
	for _, importer := range response.Results {
		if strings.HasPrefix(importer.Path, "github.com/") {
			split := strings.Split(importer.Path, "/")
			if len(split) >= 3 {
				repos[strings.Join(split[1:3], "/")] = struct{}{}
			}
		}
	}

	log.Printf("Fetched %d github repos which import pkg %s.", len(repos), *gopkg)

	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: *githubToken},
	)
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)

	count := 0
	query := ""
	repoStars := make(map[string]int)
	for repo := range repos {
		query = query + "repo:" + repo + " "
		count++
		if count%100 == 0 {
			githubRepoSearch(ctx, client, query, repoStars)
			query = ""
			log.Printf("Fetched stars of %d repos.", count)
		}
	}
	if query != "" {
		githubRepoSearch(ctx, client, query, repoStars)
	}

	maxRepos := 50
	if len(repoStars) < maxRepos {
		maxRepos = len(repoStars)
	}

	replacer := strings.NewReplacer(".", "+", "/", "+")
	for _, repo := range rankByStar(repoStars)[0:maxRepos] {
		repoName := strings.Join(strings.Split(repo.Key, "/")[3:], "/")
		githubSearchQuery := "https://github.com/search?type=code&q=" + replacer.Replace(*gopkg) + "+in:file+repo:" + repoName
		log.Printf("Repo: %s, Stars: %d, Search query: %v.", repo.Key, repo.Value, githubSearchQuery)
	}
}
