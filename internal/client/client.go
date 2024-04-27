package client

import (
	"encoding/json"
	"fmt"
	"net/url"
	"os"
	"strconv"

	"github.com/jfgrea27/hnsnap/internal/models"
	"github.com/jfgrea27/hnsnap/internal/request"
	"github.com/jfgrea27/hnsnap/internal/utils"
	"github.com/rs/zerolog/log"
)

type HNClient struct {
}

func (c *HNClient) TopStories() []models.Story {

	// get top stories ids
	tsIds := getTopStoriesIds()

	// setup
	stories := make([]models.Story, len(tsIds))

	n_workers := getNWorkers(tsIds)
	jobs := make(chan int, len(tsIds))
	results := make(chan models.Story, len(tsIds))

	// start workers
	for w := 1; w <= n_workers; w++ {
		go storyWorker(w, jobs, results)
	}

	// add ids to jobs channel
	for tsId := range tsIds {
		jobs <- tsId
	}
	// all ids added - can close jobs channel
	close(jobs)

	// retrieve results from results channel
	for a := 1; a <= len(tsIds); a++ {
		story, ok := <-results
		if !ok {
			log.Warn().Msg("Not able to retrieve top story.")
		}
		stories = append(stories, story)
	}
	return stories
}

func storyWorker(w_id int, jobs <-chan int, results chan<- models.Story) {
	log.Debug().Msg(fmt.Sprintf("Starging Worker %v fetching story.", w_id))
	for id := range jobs {
		log.Debug().Msg(fmt.Sprintf("Worker %v fetching story %v.", w_id, id))
		story := getTopStory(id)
		results <- story
	}
}

func getNWorkers(ids []int) int {
	// default to 5 workers if no env var supplied.
	n_workers := os.Getenv("HN_N_WORKERS")
	if len(n_workers) == 0 {
		return 5
	} else {
		workers, err := strconv.Atoi(n_workers)
		utils.CheckErr(err, "Could not convert HN_N_WORKERS to string")
		return workers
	}
}

func getTopStory(tsId int) models.Story {
	baseUrl := os.Getenv("HN_BASE_ENDPOINT")
	// construct endpoint
	sUrl, err := url.JoinPath(baseUrl, fmt.Sprintf("item/%v.json", tsId))
	utils.CheckErr(err, "Could not construct story url")

	// fetch top story
	jsonString, err := request.Get(sUrl)

	// unmarshall
	utils.CheckErr(err, "Could not get story from api.")
	story := models.Story{}
	json.Unmarshal([]byte(jsonString), &story)
	return story
}

func getTopStoriesIds() []int {
	baseUrl := os.Getenv("HN_BASE_ENDPOINT")

	// construct endpoint
	tsUrl, err := url.JoinPath(baseUrl, "topstories.json")
	utils.CheckErr(err, "Could not construct top stories url")

	// fetch ids of top stories
	jsonString, err := request.Get(tsUrl)
	if err != nil {
		log.Error().Msg("Could not get top stories from api.")
	}

	// unmarshall
	tsIds := make([]int, 0)
	err = json.Unmarshal([]byte(jsonString), &tsIds)

	if err != nil {
		log.Error().Msg("Could not unmarshall top stories from api.")
	}

	log.Debug().Msg(fmt.Sprintf("Number of top stories retrieved %v.", len(tsIds)))

	return tsIds

}
