package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/jwcen/mars/config"
	"github.com/jwcen/mars/internal/apiserver/repository/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	totalCount  = 1000000 // 总共插入 100 万条
	batchSize   = 10000   // 每批插入 1 万条
	authorCount = 1000    // 假设有 1000 个作者
)

var titles = []string{
	"Go 语言入门指南",
	"深入理解 GORM",
	"微服务架构实践",
	"Redis 缓存优化",
	"MySQL 性能调优",
	"Docker 容器化部署",
	"Kubernetes 实战",
	"分布式系统设计",
	"消息队列应用",
	"API 设计最佳实践",
}

var contents = []string{
	"这是一篇关于 Go 语言的文章内容...",
	"GORM 是 Go 语言中最流行的 ORM 框架之一...",
	"微服务架构是现代应用开发的重要模式...",
	"Redis 作为内存数据库，提供了高性能的缓存方案...",
	"MySQL 性能优化是数据库管理的重要技能...",
	"Docker 容器化技术改变了应用部署方式...",
	"Kubernetes 是容器编排的事实标准...",
	"分布式系统设计需要考虑很多因素...",
	"消息队列在异步处理中发挥重要作用...",
	"良好的 API 设计能够提升开发效率...",
}

func main() {
	// 初始化数据库连接
	db, err := gorm.Open(mysql.Open(config.Config.DB.DSN), &gorm.Config{})
	if err != nil {
		log.Fatalf("连接数据库失败: %v", err)
	}

	// 设置连接池
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("获取数据库连接失败: %v", err)
	}
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetConnMaxLifetime(time.Hour)

	ctx := context.Background()
	rand.Seed(time.Now().UnixNano())

	startTime := time.Now()
	log.Printf("开始插入 %d 条数据，每批 %d 条...", totalCount, batchSize)

	// 分批插入
	for i := 0; i < totalCount; i += batchSize {
		batch := make([]model.ArticleM, 0, batchSize)
		currentBatch := batchSize
		if i+batchSize > totalCount {
			currentBatch = totalCount - i
		}

		// 生成当前批次的数据
		for j := 0; j < currentBatch; j++ {
			article := model.ArticleM{
				Title:     generateTitle(i + j),
				Content:   generateContent(i + j),
				AuthorId:  int64(rand.Intn(authorCount) + 1), // 随机作者 ID，1-1000
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			}
			batch = append(batch, article)
		}

		// 批量插入
		err := db.WithContext(ctx).CreateInBatches(batch, 1000).Error
		if err != nil {
			log.Fatalf("插入第 %d-%d 条数据失败: %v", i+1, i+currentBatch, err)
		}

		// 显示进度
		progress := float64(i+currentBatch) / float64(totalCount) * 100
		elapsed := time.Since(startTime)
		rate := float64(i+currentBatch) / elapsed.Seconds()
		remaining := float64(totalCount-i-currentBatch) / rate
		log.Printf("进度: %d/%d (%.2f%%) | 已用时: %v | 速度: %.0f 条/秒 | 预计剩余: %v",
			i+currentBatch, totalCount, progress, elapsed, rate, time.Duration(remaining)*time.Second)
	}

	totalTime := time.Since(startTime)
	log.Printf("完成！总共插入 %d 条数据，耗时: %v，平均速度: %.0f 条/秒",
		totalCount, totalTime, float64(totalCount)/totalTime.Seconds())
}

func generateTitle(index int) string {
	base := titles[index%len(titles)]
	return fmt.Sprintf("%s - %d", base, index+1)
}

func generateContent(index int) string {
	base := contents[index%len(contents)]
	return fmt.Sprintf("%s 这是第 %d 篇文章的详细内容。", base, index+1)
}
