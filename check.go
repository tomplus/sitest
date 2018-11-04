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
	result.StatusCode = resp.StatusCode

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return result, err
	}

	result.Length = len(body)
	result.Hash = hash(body)
	return result, nil
}

func (sitest Sitest) Run(name string) {

	config := sitest.Sites[name]

	// slow start
	slow_start := time.Duration(rand.Float64()*config.Interval.Seconds()) * time.Second
	log.Printf("[%v] slow start, sleep %v", name, slow_start)
	time.Sleep(slow_start)

	for {
		log.Printf("[%v] querying...", name)
		result, err := checkSite(name)
		if err != nil {
			log.Printf("[%v] error %v (%v)", name, err, result)
		} else {
			log.Printf("[%v] success, result: %+v", name, result)
		}
		sitest.Metrics.Update(name, result, err)
		time.Sleep(config.Interval - result.Duration)
	}
}
