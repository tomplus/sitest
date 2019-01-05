package main

import (
	"hash/fnv"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"time"
)

func hash(b []byte) uint64 {
	h := fnv.New64a()
	h.Write(b)
	return h.Sum64()
}

func checkSite(name string) (result Result, err error) {

	start := time.Now()

	resp, err := http.Get(name)
	if err != nil {
		return result, err
	}

	defer resp.Body.Close()

	result.Duration = time.Since(start)
	result.Time = time.Now()
	result.StatusCode = resp.StatusCode

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return result, err
	}

	result.Length = len(body)
	result.Hash = hash(body)
	return result, nil
}

// Get last result for site
func (site *Site) GetLastResult() Result {
	site.Mutex.Lock()
	defer site.Mutex.Unlock()
	return site.LastResult
}

func (site *Site) setLastResult(result Result) {
	site.Mutex.Lock()
	defer site.Mutex.Unlock()
	site.LastResult = result
}

// Run tests site forever
func (sitest Sitest) Run(name string) {

	site := sitest.Sites[name]

	// slow start
	slowStart := time.Duration(rand.Float64()*site.Config.Interval.Seconds()) * time.Second
	log.Printf("[%v] slow start, sleep %v", name, slowStart)
	time.Sleep(slowStart)

	for {
		log.Printf("[%v] querying...", name)
		result, err := checkSite(name)
		if err != nil {
			log.Printf("[%v] error %v (%v)", name, err, result)
		} else {
			log.Printf("[%v] success, result: %+v", name, result)
		}

		site.setLastResult(result)
		sitest.Metrics.Update(name, result, err)
		time.Sleep(site.Config.Interval - result.Duration)
	}
}
