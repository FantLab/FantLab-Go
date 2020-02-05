ALTER TABLE `b_subscribers` ADD CONSTRAINT `user_blog` UNIQUE (`user_id`, `blog_id`);
ALTER TABLE `b_topics_subscribers` ADD CONSTRAINT `user_topic` UNIQUE (`user_id`, `topic_id`);
ALTER TABLE `b_topic_likes` ADD CONSTRAINT `topic_user` UNIQUE (`topic_id`, `user_id`);
ALTER TABLE `f_topics_subscribers` ADD CONSTRAINT `user_topic` UNIQUE (`user_id`, `topic_id`);
