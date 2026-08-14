package main

import (
	"fmt"
	"shop/config"
	"strings"
)

// 写入分类
func main() {
	query, err := config.MysqlQuery("SELECT category FROM video")
	if err != nil {
		fmt.Println("查询video表失败:", err)
		return
	}

	categorySet := make(map[string]bool)
	for _, val := range query {
		if categoryStr, ok := val["category"].(string); ok && categoryStr != "" {
			categories := strings.Split(categoryStr, "/")
			for _, cat := range categories {
				cat = strings.TrimSpace(cat)
				if cat != "" {
					categorySet[cat] = true
				}
			}
		}
	}

	if len(categorySet) == 0 {
		fmt.Println("没有找到任何分类")
		return
	}

	fmt.Printf("找到 %d 个唯一分类: %v\n", len(categorySet), categorySet)

	existingCategories := make(map[string]bool)
	existingQuery, err := config.MysqlQuery("SELECT name FROM category")
	if err != nil {
		fmt.Println("查询category表失败:", err)
		return
	}

	for _, val := range existingQuery {
		if name, ok := val["name"].(string); ok {
			existingCategories[name] = true
		}
	}

	stmt, err := config.Mysql.Prepare("INSERT INTO category (name) VALUES (?)")
	if err != nil {
		fmt.Println("准备SQL语句失败:", err)
		return
	}
	defer stmt.Close()

	insertedCount := 0
	skippedCount := 0
	for category := range categorySet {
		if existingCategories[category] {
			fmt.Printf("分类 '%s' 已存在，跳过\n", category)
			skippedCount++
			continue
		}

		result, err := stmt.Exec(category)
		if err != nil {
			fmt.Printf("插入分类 '%s' 失败: %v\n", category, err)
			continue
		}

		lastId, _ := result.LastInsertId()
		fmt.Printf("成功插入分类: %s (ID: %d)\n", category, lastId)
		insertedCount++
	}

	fmt.Printf("\n完成! 共插入 %d 个分类，跳过 %d 个已存在的分类\n", insertedCount, skippedCount)
}
