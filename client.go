package main

import (
	"encoding/csv"
	"encoding/gob"
	"fmt"
	"log"
	"math/rand"
	"net"
	"os"
	"strconv"
	"sync"
)

type Data struct {
	Part [][]string
}

type Cluster struct {
	Centroid []float64
	Points   [][]float64
}

func main() {
	parts := LoadAndDivideDataset()

	log.Println("Starting goroutines...")

	// Create a channel for each goroutine
	channels := make([]chan []Cluster, len(parts))
	for i := range channels {
		channels[i] = make(chan []Cluster)
	}

	// Start each goroutine
	var wg sync.WaitGroup
	for i, part := range parts {
		wg.Add(1)
		go func(part [][]string, c chan []Cluster) {
			defer wg.Done()

			log.Println("Connecting to server...")

			// Connect to the server
			conn, err := net.Dial("tcp", "localhost:8080")
			if err != nil {
				fmt.Println(err)
				return
			}
			defer conn.Close()

			log.Println("Sending data to server...")

			// Send the data to the server
			data := &Data{Part: part}
			encoder := gob.NewEncoder(conn)
			encoder.Encode(data)

			log.Println("Receiving results from server...")

			decoder := gob.NewDecoder(conn)
			clusters := make([]Cluster, 0)
			decoder.Decode(&clusters)

			log.Println("Sending clusters to channel...")

			c <- clusters
		}(part, channels[i])
	}

	log.Println("Waiting for goroutines to finish...")

	var totalClusters []Cluster
	wg.Add(1)
	go func() {
		defer wg.Done()
		for _, c := range channels {
			clusters := <-c
			totalClusters = append(totalClusters, clusters...)
		}
	}()

	wg.Wait()

	log.Println("Printing clusters...")

	// Print the clusters
	for i, cluster := range totalClusters {
		log.Printf("Cluster %d: %v\n", i, cluster.Centroid)
	}
}

func LoadAndDivideDataset() [][][]string {
	// Load the dataset
	file, err := os.Open("SocialNetworkDataset.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	// Divide the dataset into parts
	numGoroutines := 10
	partSize := len(records) / numGoroutines
	parts := make([][][]string, numGoroutines)
	for i := 0; i < numGoroutines; i++ {
		start := i * partSize
		end := start + partSize
		if i == numGoroutines-1 {
			end = len(records)
		}
		parts[i] = records[start:end]
	}

	return parts
}

func kmeans(part [][]string, k int) []Cluster {
	clusters := make([]Cluster, k)
	for i := range clusters {
		clusters[i].Centroid = make([]float64, len(part[0])-1)
		for j := range clusters[i].Centroid {
			switch j {
			case 1:
				clusters[i].Centroid[j] = rand.Float64() * 100
			case 3, 4, 5, 6, 7, 8:
				clusters[i].Centroid[j] = rand.Float64() * 10000
			default:
				clusters[i].Centroid[j] = rand.Float64()
			}
		}
	}

	for i := 0; i < 1000; i++ {
		for _, row := range part {
			point := make([]float64, len(row)-1)
			for j := range point {
				point[j], _ = strconv.ParseFloat(row[j+1], 64)
			}

			minDistance := distance(point, clusters[0].Centroid)
			minIndex := 0
			for j := 1; j < k; j++ {
				d := distance(point, clusters[j].Centroid)
				if d < minDistance {
					minDistance = d
					minIndex = j
				}
			}

			clusters[minIndex].Points = append(clusters[minIndex].Points, point)
		}

		for i := range clusters {
			if len(clusters[i].Points) > 0 {
				sum := make([]float64, len(clusters[i].Centroid))
				for _, point := range clusters[i].Points {
					for j, x := range point {
						sum[j] += x
					}
				}
				for j := range clusters[i].Centroid {
					clusters[i].Centroid[j] = sum[j] / float64(len(clusters[i].Points))
				}
			} else {
				for j := range clusters[i].Centroid {
					switch j {
					case 1: // Age
						clusters[i].Centroid[j] = rand.Float64() * 100
					case 3, 4, 5, 6, 7, 8:
						clusters[i].Centroid[j] = rand.Float64() * 10000
					default:
						clusters[i].Centroid[j] = rand.Float64()
					}
				}
			}
			clusters[i].Points = nil
		}
	}

	return clusters
}

func distance(a, b []float64) float64 {
	sum := 0.0
	for i := range a {
		d := a[i] - b[i]
		sum += d * d
	}
	return sum
}
