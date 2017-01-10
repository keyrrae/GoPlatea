package main

import (
	"bufio"
	"crypto/md5"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"runtime"
	"strings"
	"sync"
	"sync/atomic"
	"time"
)

// A struct for configuration file
type Config struct {
	EntryPath         string
	UrlQueuePath      string
	FilterExtPath     string
	LimitedToThisPath string
	DownWebPagePath   string
	DownUrlDataPath   string
	MaxNum            int32
	IntervalTime      string
}

// A struct representing a web page
type UrlInfo struct {
	md5     string // md5 of trimmed url
	url     string // complete url
	path    string // path to stored data as a file
	content string // data of a web page
}

// Spider
type Spider struct {
	conf          *Config         // Reference to configuration
	filterExts    map[string]bool // filter file name extensions
	limitedToThis []string        // crawl urls contain at least on string in this slice(array)
	doneUrls      map[string]bool // md5 of crawled urls
	exceptionUrls map[string]bool // md5 of exceptional urls
	chUrlsInfo    chan *UrlInfo   // Information of urls
	chUrl         chan string     // url
	chHttp        chan bool
	chStopIO      chan bool
	chExit        chan bool
	wg            sync.WaitGroup
	pageNum       int32          // Current crawled web pages
	intervalTime  time.Duration  // interval time
}

// Initialize a spider and return its reference
func NewSpider() *Spider {
	// Only one thread(goroutine) writes to disk
	// Multiple threads to fetch contents from the internet
	// Make full use of network bandwidth
	// Even if disk is blocking, download thread can still download web pages and save to memory
	runtime.GOMAXPROCS(runtime.NumCPU())

	var sp Spider
	sp.conf = NewConfig()
	sp.filterExts = make(map[string]bool)
	sp.limitedToThis = make([]string, 0)
	sp.doneUrls = make(map[string]bool)
	sp.exceptionUrls = make(map[string]bool)
	sp.chUrlsInfo = make(chan *UrlInfo, 100)
	sp.chUrl = make(chan string, 1000000)
	sp.chHttp = make(chan bool, 5)
	sp.chStopIO = make(chan bool)
	sp.chExit = make(chan bool)
	sp.pageNum = 0
	intervalTime, err := time.ParseDuration(sp.conf.IntervalTime)
	if err != nil {
		sp.intervalTime = 500 * time.Millisecond
	} else {
		sp.intervalTime = intervalTime
	}

	sp.readDownUrlData()
	sp.readEntry()
	sp.readUrlQueue()
	sp.readFilterExt()
	sp.readLimitedToThis()

	return &sp
}

// Reading previous crawled web page data
func (this *Spider) readDownUrlData() {
	// make parent directory
	err := os.MkdirAll(filepath.Dir(this.conf.DownUrlDataPath), os.ModePerm)
	if err != nil {
		log.Fatal(err, "\r\n")
	}

	// Open database
	db, err := sql.Open("sqlite3", this.conf.DownUrlDataPath)
	if err != nil {
		log.Fatal(err, "\r\n")
	}
	defer db.Close()

	// Create a table
	// if error, the table already exists
	db.Exec("create table data(md5 varchar(32), url varchar(256), path varchar(256))")

	// read data
	rows, err := db.Query("select * from data")
	if err != nil {
		log.Fatal(err, "\r\n")
	}
	defer rows.Close()

	var md5, url, path string
	for rows.Next() {
		rows.Scan(&md5, &url, &path)
		this.doneUrls[md5] = true
	}
}

// Read endpoint
func (this *Spider) readEntry() {
	file, err := os.Open(this.conf.EntryPath)
	if err != nil {
		log.Fatal(err, "\r\n")
	}
	defer file.Close()

	re := bufio.NewReader(file)
	for {
		urlByte, _, err := re.ReadLine()
		if err != nil {
			break
		}
		if string(urlByte) != "" && !this.doneUrls[getMd5FromUrl(string(urlByte))] {
			this.chUrl <- string(urlByte)
		}
	}
}

// read url queue
func (this *Spider) readUrlQueue() {
	// open database
	db, err := sql.Open("sqlite3", this.conf.UrlQueuePath)
	if err != nil {
		log.Fatal(err, "\r\n")
	}
	defer db.Close()

	// create a table, if exits, fail
	db.Exec("create table data(md5 varchar(32), url varchar(256))")

	// read data
	rows, err := db.Query("select * from data")
	if err != nil {
		log.Fatal(err, "\r\n")
	}
	defer rows.Close()

	var md5, url string
	for rows.Next() {
		rows.Scan(&md5, &url)
		if !this.doneUrls[md5] {
			this.chUrl <- url
			this.doneUrls[md5] = true
		}
	}
}

// filter extensions
func (this *Spider) readFilterExt() {
	file, err := os.OpenFile(this.conf.FilterExtPath, os.O_RDWR|os.O_CREATE|os.O_APPEND, os.ModePerm)
	if err != nil {
		log.Printf("%v\r\n", err)
		return
	}
	defer file.Close()

	re := bufio.NewReader(file)
	for {
		extbyte, _, err := re.ReadLine()
		if err != nil {
			break
		}
		if string(extbyte) != "" {
			this.filterExts[string(extbyte)] = true
		}
	}
}

// read information of an url
func (this *Spider) readLimitedToThis() {
	file, err := os.OpenFile(this.conf.LimitedToThisPath, os.O_RDWR|os.O_CREATE|os.O_APPEND, os.ModePerm)
	if err != nil {
		log.Printf("%v\r\n", err)
		return
	}
	defer file.Close()

	re := bufio.NewReader(file)
	for {
		limbyte, _, err := re.ReadLine()
		if err != nil {
			break
		}
		if string(limbyte) != "" {
			this.limitedToThis = append(this.limitedToThis, string(limbyte))
		}
	}
}

// Save UrlInfo into the database
func (this *Spider) writeUrlInfo() {
	// notify main thread to finish after disk IO thread finished
	defer func() {
		this.chExit <- true
	}()

	// write md5, url, path into this.conf.DownUrlDataPath database
	// Open database
	db, err := sql.Open("sqlite3", this.conf.DownUrlDataPath)
	if err != nil {
		log.Fatal(err, "\r\n")
	}
	defer db.Close()
	
	// create table, fails if exists
	db.Exec("create table data(md5 varchar(32), url varchar(256), path varchar(256))")

	// write content into this.conf.DownWebpagePath/xxx.html
	// create parent directory
	err = os.MkdirAll(this.conf.DownWebPagePath, os.ModePerm)
	if err != nil {
		log.Fatal(err, "\r\n")
	}

	// cannot stop when received canStop instruction, because it needs to wait chUrlsInfo to finish
	var urlInfo *UrlInfo
	var canStop bool = false
	for {
		select {
		case <-this.chStopIO: //got stopIO instruction
			canStop = true
			if len(this.chUrlsInfo) == 0 {
				return
			}
		case urlInfo = <-this.chUrlsInfo:
			fmt.Printf("[%s]---Writing into database...\n", urlInfo.url)
			// Save a web page
			ioutil.WriteFile(urlInfo.path, []byte(urlInfo.content), os.ModePerm)
			// put web page info into the database and ignore errors
			db.Exec("insert into data(md5, url, path) values(?, ?, ?)", urlInfo.md5, urlInfo.url, urlInfo.path)
			this.pageNum = atomic.AddInt32(&this.pageNum, 1)
			fmt.Printf("[%s]---Writing completed.\n", urlInfo.url)
			if canStop && len(this.chUrlsInfo) == 0 {
				return
			}
		}
	}
}

// write url queue to database, next time starts from here
func (this *Spider) writeUrlQueue(urls []string) {
	// Open database
	db, err := sql.Open("sqlite3", this.conf.UrlQueuePath)
	if err != nil {
		log.Printf("%v\r\n", err)
		return
	}
	defer db.Close()

	tx, err := db.Begin() // start a transaction
	if err != nil {
		log.Printf("%v\r\n", err)
	} else {
		for _, vv := range urls {
			vv = trimUrl(vv)
			md5 := getMd5FromUrl(vv)
			if !this.doneUrls[md5] && !this.exceptionUrls[md5] && !this.beFiltered(vv) {
				_, err = tx.Exec("insert into data(md5, url) values(?, ?)", md5, vv)
				if err != nil {
					log.Printf("%v\r\n", err)
					break
				}
			}
			if len(urls) == 0 {
				break
			}
		}
	}
	tx.Commit() // commit the transaction
}

// check if the number of crawled web pages has reached the limit
func (this *Spider) isFinished() bool {
	if atomic.LoadInt32(&this.pageNum) >= this.conf.MaxNum {
		log.Printf("%v\r\n", "Number of crawled pages has reached expected value")
		return true
	}
	return false
}

// Main thread, start crawling
func (this *Spider) Fetch() {
	if len(this.chUrl) == 0 {
		log.Fatal("entry url is empty.\r\n")
		return
	}

	go this.writeUrlInfo()
	this.work()

	this.chStopIO <- true // notify finish writeUrlInfo()
	<-this.chExit         // wait for writeUrlInfo to finish
}

// work thread
func (this *Spider) work() {
	for url := range this.chUrl {
		this.chHttp <- true // Controls number of download threads

		go func(url string) {
			this.wg.Add(1)
			log.Printf("%v\r\n", "[Thread] Download started")
			defer func() {
				<-this.chHttp
				this.wg.Done()
				log.Printf("%v\r\n", "[Thread] Download completed")
			}()
			this.do(url)
		}(url)

		log.Printf("len(chUrlsInfo)==%d --- len(chUrl)==%d --- len(chHttp)==%d\r\n", len(this.chUrlsInfo), len(this.chUrl), len(this.chHttp))
		time.Sleep(this.intervalTime) // slows down the crawling

		if this.isFinished() {
			log.Printf("%v\r\n", "Waiting for all the threads to finish......")
			this.wg.Wait()
			log.Printf("%v\r\n", "All threads finished.")
			if len(this.chUrl) == 0 {
				break
			}

			// save the remain urls in this.chUrl
			urls := make([]string, 0)
			for v := range this.chUrl {
				urls = append(urls, v)
				if len(this.chUrl) == 0 {
					break
				}
			}
			this.writeUrlQueue(urls)
			break
		}
	}
}

// Process a url
func (this *Spider) do(url string) {
	client := &http.Client{
		CheckRedirect: doRedirect,
	}
	// if url redirect, handle with doRedirect() in client.Get(url)
	// process error from doRedirect()
	resp, err := client.Get(url)
	if err != nil {
		log.Printf("%s\r\n", err)
		this.exceptionUrls[getMd5FromUrl(url)] = true
		return
	}
	defer resp.Body.Close()

	// return if not OK, 500 etc.
	if resp.StatusCode != http.StatusOK {
		log.Printf("[%s] resp.StatusCode == [%d]\r\n", url, resp.StatusCode)
		this.exceptionUrls[getMd5FromUrl(url)] = true
		return
	}

	fmt.Printf("[%s]---Downloading...\n", url)
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("%s\r\n", err)
		this.exceptionUrls[getMd5FromUrl(url)] = true
		fmt.Printf("[%s]---An exception，download stopped.\n", url)
		return
	}
	fmt.Printf("[%s]---Download completed.\n", url)

	// save UrlInfo
	md5 := getMd5FromUrl(url)
	path := this.conf.DownWebPagePath + md5 + ".html"
	this.chUrlsInfo <- &UrlInfo{md5: md5, url: url, path: path, content: string(content)}

	fmt.Printf("[%s]---Analyzing...\n", url)
	// got new url
	urls := getURLs(content)
	for i, v := range urls {
		// task finished, write url queue into the database, such that next time it starts from here
		if this.isFinished() {
			this.writeUrlQueue(urls[i:])
			break
		}

		// not finished, continue to crawl
		v = trimUrl(v)
		md5 := getMd5FromUrl(v)
		if !this.doneUrls[md5] && !this.exceptionUrls[md5] {
			if this.beFiltered(v) {
				this.exceptionUrls[md5] = true
			} else {
				this.chUrl <- v
				this.doneUrls[md5] = true
			}
		}
	}
	fmt.Printf("[%s]---Analysis finished.\n", url)
}

// filter
func (this *Spider) beFiltered(url string) bool {
	b1 := this.filterExts[filepath.Ext(url)] // suffix filter
	b2 := true
	// if limitedToThis is not empty and limitedToThis strings is a substring in url
	if len(this.limitedToThis) > 0 {
		for _, v := range this.limitedToThis {
			if strings.Contains(url, v) {
				b2 = false
				break
			}
		}
	} else {
		b2 = false
	}

	return b1 || b2
}

// deal with redirect, StatusCode == 302
func doRedirect(req *http.Request, via []*http.Request) error {
	return errors.New(req.URL.String() + " was as an exception url to do.")
}

// Read configuration file
func NewConfig() *Config {
	argsWithoutProg := os.Args[1:]

	var conf Config

	file, err := ioutil.ReadFile("./spider.conf")
	if err != nil {
		log.Fatal(err, "\r\n")
	}

	err = json.Unmarshal(file, &conf)
	if err != nil {
		log.Fatal(err, "\r\n")
	}

	if len(argsWithoutProg) == 1 {
		// TODO: arguments parsing: number of webpages, interval time
		dataPath := argsWithoutProg[0]
		if !strings.HasSuffix(dataPath, "/") {
			dataPath += "/"
		}

		conf.DownWebPagePath = dataPath
		conf.DownUrlDataPath = dataPath + "downurldata.db"

		confJson, err := json.Marshal(conf)
		if err != nil {
			log.Fatal(err, "\r\n")
		}

		err = ioutil.WriteFile("./spider.conf", confJson, 0644)
		if err != nil {
			log.Fatal(err, "\r\n")
		}
	}

	return &conf
}

// Initialization
// Run before main()
func init() {
	setLogOutput()
}

// setup logger output
func setLogOutput() {
	// short filename for log, convenient for line number
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	logfile, err := os.OpenFile("./spider.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, os.ModePerm)
	// cannot close log file here
	if err != nil {
		log.Printf("%v\r\n", err)
	}
	log.SetOutput(logfile)
}

// trim url, remove chars after #
func trimUrl(url string) string {
	p := strings.Index(url, "#")
	if p != -1 {
		url = url[:p]
	}
	return url
}

// Trim url, calculate its md5 and convert md5 value to a string
func getMd5FromUrl(url string) string {
	url = strings.TrimRight(url, "/")
	url = strings.TrimPrefix(url, "http://")
	url = strings.TrimPrefix(url, "https://")
	md5Obj := md5.New()
	io.WriteString(md5Obj, url)
	str := fmt.Sprintf("%x", md5Obj.Sum(nil)) // convert md5 object to a string
	return str
}

// Extract all the URLs from the crawled HTML's.
func getURLs(content []byte) (urls []string) {
	re := regexp.MustCompile("href\\s*=\\s*['\"]?\\s*(https?://[^'\"\\s]+)\\s*['\"]?")
	allSubmatches := re.FindAllSubmatch([]byte(content), -1)
	for _, v2 := range allSubmatches {
		for k, v := range v2 {
			// ignore k == 0, because k==0 means all the matched strings
			if k > 0 {
				urls = append(urls, string(v))
			}
		}
	}
	return urls
}

func main() {
	log.Printf("%v\r\n", "start......")
	start := time.Now()
	sp := NewSpider()
	sp.Fetch()
	log.Printf("used time: %v", time.Since(start))
}
