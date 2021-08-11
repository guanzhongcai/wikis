package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

// 从文件夹读取日志，分析后，存入influxDB

type Reader interface {
	Read(rc chan []byte)
}

type Writer interface {
	Write(wc chan string)
}

type LogProcess struct {
	rc    chan []byte // 读取chan
	wc    chan string // 写入chan
	read  Reader
	write Writer
}

type SystemInfo struct {
	HandleRequest int    `json:"handleRequest"` // 已处理请求数
	Tps           int    `json:"tps"`           // 系统总吞吐量
	RunTime       string `json:"runTime"`       // 运行时间
	ErrNum        int    `json:"errNum"`        // 错误数
}

type Monitor struct {
	startTime time.Time // 启动时间
	data      SystemInfo
}

func (m *Monitor) start(lp *LogProcess) {
	http.HandleFunc("/monitor", func(writer http.ResponseWriter, request *http.Request) {
		m.data.RunTime = time.Now().Sub(m.startTime).String()
		ret, _ := json.MarshalIndent(m.data, "", "\t")
		io.WriteString(writer, string(ret))
	})

	http.ListenAndServe(":9193", nil)
}

type ReadFromFile struct {
	path string // 读取文件的路径
}

func (r *ReadFromFile) Read(rc chan []byte) {
	// 读取模块
	f, err := os.Open(r.path)
	if err != nil {
		panic(fmt.Sprintf("open file error: %s", err.Error()))
	}

	// 从文件末尾开始逐行读取文件内容
	f.Seek(0, 2)
	rd := bufio.NewReader(f)

	for {
		line, err := rd.ReadBytes('\n')
		if err == io.EOF {
			// TODO：日志轮转时，文件名不会变，此处要重新open这个文件
			time.Sleep(500 * time.Millisecond)
			continue
		} else if err != nil {
			panic(fmt.Sprintf("ReadBytes error: %s", err.Error()))
		}
		rc <- line[:len(line)-1]
	}
}

// Process 解析模块
func (l *LogProcess) Process() {
	// todo 解析为message结构体，以直接传入写入模块
	for v := range l.rc {
		l.wc <- strings.ToUpper(string(v))
	}
}

// WriteToInfluxDB 写入模块
type WriteToInfluxDB struct {
	influxDBDsn string // info data source
}

func (w *WriteToInfluxDB) Write(wc chan string) {
	for v := range wc {
		fmt.Println(v)
	}
}

func main() {

	r := &ReadFromFile{
		path: "access.log",
	}
	w := &WriteToInfluxDB{
		influxDBDsn: "username&password&host:port",
	}
	lp := &LogProcess{
		rc:    make(chan []byte, 200),
		wc:    make(chan string, 200),
		read:  r,
		write: w,
	}

	// 此写法也可以！但golang考虑到可读性，用下面的写法也OK
	// go (*lp).ReadFromFile()
	go lp.read.Read(lp.rc)
	for i := 0; i < 5; i++ {
		go lp.Process()
	}
	for i := 0; i < 20; i++ {
		go lp.write.Write(lp.wc)
	}

	m := &Monitor{
		startTime: time.Now(),
		data:      SystemInfo{},
	}
	m.start(lp)
}
