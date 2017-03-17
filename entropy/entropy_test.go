// +build small

/*
http://www.apache.org/licenses/LICENSE-2.0.txt

Copyright 2017 janczer

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.

*/

package entropy

import (
	"os"
	"testing"

	"github.com/intelsdi-x/snap-plugin-lib-go/v1/plugin"
	. "github.com/smartystreets/goconvey/convey"
)

var srcMockFile = "/tmp/mock_entropy_avail"
var srcMockFileValue = []byte("444\n")

func createMockFile() {
	deleteMockFile()
	f, _ := os.Create(srcMockFile)
	f.Write(srcMockFileValue)
	f.Close()
}

func createEmptyMockFile() {
	srcMockFileValue = []byte("")
	createMockFile()
}

func deleteMockFile() {
	os.Remove(srcMockFile)
}

func TestEntropyCollector(t *testing.T) {
	ec := EntropyCollector{}
	createMockFile()

	Convey("Test EntropyCollector", t, func() {
		Convey("Collect Entropy", func() {
			entropyInfo = srcMockFile
			metrics := []plugin.Metric{plugin.Metric{}}
			mts, err := ec.CollectMetrics(metrics)
			So(mts, ShouldNotBeEmpty)
			So(err, ShouldBeNil)
			So(mts[0].Data, ShouldEqual, 444)
		})
	})

	Convey("Test GetMetricTypes", t, func() {
		Convey("Collect String", func() {
			mt, err := ec.GetMetricTypes(nil)
			So(err, ShouldBeNil)
			So(len(mt), ShouldEqual, 1)
		})
	})

	Convey("Test GetConfigPolicy", t, func() {
		Convey("No error returned", func() {
			_, err := ec.GetConfigPolicy()
			So(err, ShouldBeNil)
		})
	})

	Convey("Test getEntropy", t, func() {
		Convey("No error returned", func() {
			e, err := getEntropy()
			So(err, ShouldBeNil)
			So(e, ShouldBeGreaterThan, 0)
			So(e, ShouldEqual, 444)
		})
		Convey("File not found", func() {
			deleteMockFile()
			e, err := getEntropy()
			So(e, ShouldEqual, 0)
			So(err, ShouldNotBeNil)
		})
		Convey("Empty file", func() {
			createEmptyMockFile()
			e, err := getEntropy()
			So(e, ShouldEqual, 0)
			So(err, ShouldNotBeNil)
		})
	})
	deleteMockFile()
}
