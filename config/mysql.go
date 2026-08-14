package config

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-ini/ini"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"strconv"
	"sync"
	"time"
)

var Mysql *sql.DB

func init() {
	InitMysql()
}

// InitMysql 初始化Mysql连接
func InitMysql() {
	mysql := ReadIniFile("./config/mysql.ini")
	dbAddress := mysql.Section("mysql").Key("dbAddress").Value()
	dbName := mysql.Section("mysql").Key("dbName").Value()
	dbUser := mysql.Section("mysql").Key("dbUser").Value()
	dbPasswd := mysql.Section("mysql").Key("dbPasswd").Value()
	sqlConnStr := fmt.Sprintf("%v:%v@tcp(%v)/%v?charset=utf8mb4&timeout=5s&readTimeout=30s&writeTimeout=30s&interpolateParams=true", dbUser, dbPasswd, dbAddress, dbName)
	var err error
	Mysql, err = sql.Open("mysql", sqlConnStr)
	if err != nil {
		log.Fatal("初始化Mysql时出错：", err)
	}
	// 设置连接池参数
	Mysql.SetMaxOpenConns(100)                 // 最大连接数
	Mysql.SetMaxIdleConns(10)                  // 空闲连接数
	Mysql.SetConnMaxLifetime(30 * time.Minute) // 连接最大存活时间
	Mysql.SetConnMaxIdleTime(10 * time.Minute) // 空闲连接最大存活时间
	if err = Mysql.Ping(); err != nil {
		log.Fatal("连接到Mysql时出错：", err)
	}
	fmt.Println("Mysql连接成功！")
	//warmupConnections(Mysql, 10)
	//go keepAlive(Mysql)
}

// MysqlQuery 执行SQL查询并返回 []map[string]interface{}
func MysqlQuery(query string, args ...any) ([]map[string]interface{}, error) {
	// 使用Prepare语句提高重复查询性能
	stmt, err := Mysql.Prepare(query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	rows, err := stmt.Query(args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	columns, err := rows.Columns()
	if err != nil {
		return nil, err
	}
	count := len(columns)
	// 预分配足够容量的切片，减少内存分配和GC压力
	tableData := make([]map[string]interface{}, 0) // 预估初始容量
	// 复用values和valuePtrs数组
	values := make([]interface{}, count)
	valuePtr := make([]interface{}, count)
	for i := range values {
		valuePtr[i] = &values[i]
	}
	// 预定义columnMap避免每次循环都创建
	columnMap := make(map[string]int, count)
	for i, col := range columns {
		columnMap[col] = i
	}
	for rows.Next() {
		if err := rows.Scan(valuePtr...); err != nil {
			return nil, err
		}
		entry := make(map[string]interface{}, count)
		for col, i := range columnMap {
			val := values[i]
			if b, ok := val.([]byte); ok {
				entry[col] = string(b)
			} else {
				entry[col] = val
			}
		}
		tableData = append(tableData, entry)
	}
	return tableData, nil
}

// ReadIniFile 读取ini文件的数据
func ReadIniFile(fileName string) *ini.File {
	iniObj, err := ini.Load(fileName)
	if err != nil {
		log.Fatal("获取INI文件出错：", err)
	}
	return iniObj
}

// PageLimit 处理Mysql数据分页
func PageLimit(c *gin.Context) string {
	limit, err := strconv.Atoi(c.Query("limit"))
	page, err := strconv.Atoi(c.Query("page"))
	if page == 0 || limit == 0 || err != nil {
		return ""
	} else {
		page = (page - 1) * limit
		res := fmt.Sprintf(" limit %v,%v", page, limit)
		return res
	}
}

// PrintMysqlStats 连接池监控打印
func PrintMysqlStats() {
	stats := Mysql.Stats()
	fmt.Printf("连接池状态:\n")
	fmt.Printf("最大打开连接数: %d\n", stats.MaxOpenConnections)
	fmt.Printf("打开连接数: %d\n", stats.OpenConnections)
	fmt.Printf("使用中连接数: %d\n", stats.InUse)
	fmt.Printf("空闲连接数: %d\n", stats.Idle)
	fmt.Printf("等待连接数: %d\n", stats.WaitCount)
	fmt.Printf("等待时间总计: %v\n", stats.WaitDuration)
	fmt.Printf("最大空闲时间关闭数: %d\n", stats.MaxIdleTimeClosed)
	fmt.Printf("最大生命周期关闭数: %d\n", stats.MaxLifetimeClosed)
}

// 预热指定数量的连接
func warmupConnections(db *sql.DB, count int) {
	start := time.Now()
	var wg sync.WaitGroup
	errChan := make(chan error, count)
	for i := 0; i < count; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			// 执行简单查询来建立连接
			_, err := db.Exec("SELECT 1")
			if err != nil {
				errChan <- err
			}
		}()
	}
	wg.Wait()
	close(errChan)
	// 检查是否有错误
	for err := range errChan {
		if err != nil {
			log.Fatalf("连接池预热失败：%v", err)
		}
	}
	fmt.Printf("Mysql已预热%v个连接，耗时：%v\n", count, time.Since(start))
}

// 定时保活连接池
func keepAlive(db *sql.DB) {
	ticker := time.NewTicker(10 * time.Minute)
	defer ticker.Stop()
	for t := range ticker.C {
		_, err := db.Exec("SELECT 1")
		sprintf := fmt.Sprintf("Mysql保活连接池：%s %v", t.Format("2006-01-02 15:04:05"), err)
		fmt.Println(sprintf)
		LogInfo("%s", sprintf)
	}
}
