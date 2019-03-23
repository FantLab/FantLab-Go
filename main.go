package main

import (
	"fmt"
	"log"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	db, err := gorm.Open("mysql", "root:root@/fantlab?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	db.LogMode(true)

	fdb := &FDB{db}
	userId := 541

	showSubscriptions(fdb, userId)
}

func showSubscriptions(fdb *FDB, userId int) {
	fmt.Println("blogs =", fdb.getSubscribedBlogs(userId))
	fmt.Println("blogTopics =", fdb.getSubscribedBlogTopicIds(userId))

	forumTopics := fdb.getSubscribedForumTopics(userId)
	fmt.Println("forumTopics =", forumTopics)

	forumTopicIds := make([]int, len(forumTopics))
	for index, topic := range forumTopics {
		forumTopicIds[index] = topic.TopicId
	}

	topicMessages := fdb.getSubscribedForumMessages(forumTopicIds, time.Unix(1135987200, 0), 100)
	fmt.Println("messages =", topicMessages)
}
