DELETE b1 FROM b_subscribers b1, b_subscribers b2 WHERE b1.subscriber_id > b2.subscriber_id AND b1.user_id = b2.user_id AND b1.blog_id = b2.blog_id;
DELETE b1 FROM b_topics_subscribers b1, b_topics_subscribers b2 WHERE b1.topic_subscriber_id > b2.topic_subscriber_id AND b1.user_id = b2.user_id AND b1.topic_id = b2.topic_id;
DELETE f1 FROM f_messages_preview f1, f_messages_preview f2 WHERE f1.preview_id > f2.preview_id AND f1.user_id = f2.user_id AND f1.topic_id = f2.topic_id;
DELETE f1 FROM f_topics_subscribers f1, f_topics_subscribers f2 WHERE f1.topic_subscriber_id > f2.topic_subscriber_id AND f1.user_id = f2.user_id AND f1.topic_id = f2.topic_id;

ALTER TABLE `b_subscribers` ADD CONSTRAINT `user_blog_pair_unique` UNIQUE (`user_id`, `blog_id`);
ALTER TABLE `b_topics_subscribers` ADD CONSTRAINT `user_topic_pair_unique` UNIQUE (`user_id`, `topic_id`);
ALTER TABLE `b_topic_likes` ADD CONSTRAINT `user_topic_pair_unique` UNIQUE (`user_id`, `topic_id`);
ALTER TABLE `f_messages_preview` ADD CONSTRAINT `user_topic_pair_unique` UNIQUE (`user_id`, `topic_id`);
ALTER TABLE `f_topics_subscribers` ADD CONSTRAINT `user_topic_pair_unique` UNIQUE (`user_id`, `topic_id`);
