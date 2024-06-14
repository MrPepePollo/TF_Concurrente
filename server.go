// server.go
package main

import (
	"encoding/gob"
	"fmt"
	"math/rand"
	"net"
	"strconv"
)

type Data struct {
	Part [][]string
}

type Cluster struct {
	Centroid []float64
	Points   [][]float64
}

func handleRequest(conn net.Conn) {
	decoder := gob.NewDecoder(conn)
	data := &Data{}
	decoder.Decode(data)

	clusters := kmeans(data.Part, 3)

	encoder := gob.NewEncoder(conn)
	encoder.Encode(clusters)
}

func main() {
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println(err)
		return
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		go handleRequest(conn)
	}
}

func kmeans(part [][]string, k int) []Cluster {
	clusters := make([]Cluster, k)
	for i := range clusters {
		clusters[i].Centroid = make([]float64, len(part[0])-1)
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
