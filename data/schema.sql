create database heymate;

CREATE TABLE `user` (
  `Id` int(11) NOT NULL AUTO_INCREMENT,
  `email` varchar(200) NOT NULL,
  `password` varchar(200) NOT NULL,
  `primary_phone` varchar(20) DEFAULT NULL,
  `last_login` datetime DEFAULT NULL,
  `app_downloaded_date` datetime DEFAULT NULL,
  `push_notification_opt_in` tinyint(1) DEFAULT NULL,
  `created_date` datetime DEFAULT NULL,
  `last_modified` datetime DEFAULT NULL,
  `display_name` varchar(64) DEFAULT NULL,
  `image_url` varchar(200) DEFAULT NULL,
  `email_opt_in` tinyint(1) DEFAULT NULL,
  `message_opt_in` tinyint(1) DEFAULT NULL,
  `is_active` tinyint(1) DEFAULT '1',
  PRIMARY KEY (`Id`),
  UNIQUE KEY `email` (`email`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;

CREATE TABLE `apartment` (
  `Id` int(11) NOT NULL AUTO_INCREMENT,
  `apartment_name` varchar(200) DEFAULT NULL,
  `street` varchar(200) DEFAULT NULL,
  `address` varchar(200) DEFAULT NULL,
  `state` varchar(200) DEFAULT NULL,
  `country` varchar(64) DEFAULT NULL,
  `postal_code` varchar(16) DEFAULT NULL,
  `created_date` datetime NOT NULL,
  `last_modified` datetime NOT NULL,
  `image_url` varchar(200) DEFAULT NULL,
    PRIMARY KEY (`Id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;

CREATE TABLE `apartment_user` (
  `Id` int(10) NOT NULL AUTO_INCREMENT,
  `apartment_id` int(10) DEFAULT NULL,
  `user_id` int(10) DEFAULT NULL,
  `flat_number` varchar(64) NOT NULL,
  `members_count` int(3) NOT NULL,
  `created_date` datetime NOT NULL,
  `last_modified` datetime NOT NULL,
  PRIMARY KEY (`Id`),
  KEY `apartment_id` (`apartment_id`),
  KEY `user_id` (`user_id`),
  CONSTRAINT `apartment_user_ibfk_1` FOREIGN KEY (`apartment_id`) REFERENCES `apartment` (`Id`),
  CONSTRAINT `apartment_user_ibfk_2` FOREIGN KEY (`user_id`) REFERENCES `user` (`Id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;


CREATE TABLE `user_sms` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `user_id` int(11) DEFAULT NULL,
  `phone_number` varchar(64) DEFAULT NULL,
  `sms_otp` varchar(200) DEFAULT NULL,
  `sent` tinyint(1) DEFAULT NULL,
  `sms_type` varchar(20) DEFAULT NULL,
  `created_date` datetime DEFAULT NULL,
  `last_modified` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_sms_usr` (`user_id`),
  KEY `ix_cd` (`created_date`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;


CREATE TABLE `apartment_labour` (
  `Id` int(10) NOT NULL AUTO_INCREMENT,
  `apartment_id` int(10) NOT NULL,
   `labour_name` varchar(64) NOT NULL,
   `phone_number` varchar(20),
   `address` varchar(200),
   `labour_type` varchar(64) ,
   `gender` varchar(20) NOT NULL,
   `image_url` varchar(200),
  `created_date` datetime NOT NULL,
  `last_modified` datetime NOT NULL,
  PRIMARY KEY (`Id`),
  KEY `apartment_id` (`apartment_id`),
  CONSTRAINT `apartment_labour_ibfk_1` FOREIGN KEY (`apartment_id`) REFERENCES `apartment` (`Id`) 
) ENGINE=InnoDB DEFAULT CHARSET=utf8;


CREATE TABLE `apartment_labour_reviews` (
  `Id` int(10) NOT NULL AUTO_INCREMENT,
  `apartment_labour_id` int(10) NOT NULL,
   `reviews` varchar(300) ,
  `created_date` datetime NOT NULL,
  `last_modified` datetime NOT NULL,
  PRIMARY KEY (`Id`),
  KEY `apartment_labour_id` (`apartment_labour_id`),
  CONSTRAINT `apartment_labour_reviews_ibfk_1` FOREIGN KEY (`apartment_labour_id`) REFERENCES `apartment_labour` (`Id`) 
) ENGINE=InnoDB DEFAULT CHARSET=utf8;


CREATE TABLE `apartment_details` (
  `Id` int(10) NOT NULL AUTO_INCREMENT,
  `apartment_id` int(10) NOT NULL,
  `block_number` varchar(64) NOT NULL,
  `v_flat_number` varchar(20) DEFAULT NULL,
  `created_date` datetime NOT NULL,
  `last_modified` datetime NOT NULL,
  PRIMARY KEY (`Id`),
  KEY `apartment_id` (`apartment_id`),
  CONSTRAINT `apartment_details_ibfk_1` FOREIGN KEY (`apartment_id`) REFERENCES `apartment` (`Id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;


CREATE TABLE `apartment_user_entries` (
  `Id` int(10) NOT NULL AUTO_INCREMENT,
  `apartment_details_id` int(10) NOT NULL,
  `user_id` int(10) NOT NULL,
  `family_count` int(10) DEFAULT 0,
  `user_type` varchar(20) NOT NULL,
  `request_validity` tinyint(2) ,
  `author_id` int(10) ,
  `created_date` datetime NOT NULL,
  `last_modified` datetime NOT NULL,
  PRIMARY KEY (`Id`),
  KEY `apartment_details_id` (`apartment_details_id`),
  KEY `user_id` (`user_id`),
  CONSTRAINT `apartment_user_details_ibfk_1` FOREIGN KEY (`apartment_details_id`) REFERENCES `apartment` (`Id`),
  CONSTRAINT `apartment_user_details_bfk_2` FOREIGN KEY (`user_id`) REFERENCES `user` (`Id`)
) ENGINE=InnoDB AUTO_INCREMENT=0 DEFAULT CHARSET=utf8;

CREATE TABLE `apartment_meeting` (
  `Id` int(10) NOT NULL AUTO_INCREMENT,
  `apartment_id` int(10) NOT NULL,
  `meeting_name` varchar(200) NOT NULL,
  `meeting_date` datetime NOT NULL,
  `meeting_place` varchar(64) ,
  `instructions` varchar(200) DEFAULT NULL,
  `created_date` datetime NOT NULL,
  `last_modified` datetime NOT NULL,
  `author_id` int(10) NOT NULL,
  `hide_author` bool ,
  PRIMARY KEY (`Id`),
  KEY `apartment_id` (`apartment_id`),
  KEY `author_id` (`author_id`),
  CONSTRAINT `apartment_meeting_ibfk_1` FOREIGN KEY (`apartment_id`) REFERENCES `apartment` (`Id`),
  CONSTRAINT `apartment_meeting_ibfk_2` FOREIGN KEY (`author_id`) REFERENCES `user` (`Id`)
) ENGINE=InnoDB AUTO_INCREMENT=0 DEFAULT CHARSET=utf8;


CREATE TABLE `apartment_bills` (
  `Id` int(10) NOT NULL AUTO_INCREMENT,
  `apartment_id` int(10) NOT NULL,
  `bill_type` varchar(40) NOT NULL,
  `bill_name` varchar(64) NOT NULL,
  `bill_description` varchar(64) DEFAULT NULL,
  `bill_amount` int(10) NOT NULL,
  `bill_date` datetime NOT NULL,
  `created_date` datetime NOT NULL,
  `last_modified` datetime NOT NULL,
  `author_id` int(10) NOT NULL,
  PRIMARY KEY (`Id`),
  KEY `apartment_id` (`apartment_id`),
  KEY `author_id` (`author_id`),
  CONSTRAINT `apartment_bills_ibfk_1` FOREIGN KEY (`apartment_id`) REFERENCES `apartment` (`Id`),
  CONSTRAINT `apartment_bills_ibfk_2` FOREIGN KEY (`author_id`) REFERENCES `user` (`Id`)
) ENGINE=InnoDB AUTO_INCREMENT=0 DEFAULT CHARSET=utf8;