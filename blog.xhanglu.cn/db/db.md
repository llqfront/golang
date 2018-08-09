CREATE TABLE `userinfo` (
  `uid` int(10) NOT NULL AUTO_INCREMENT,
  `username` varchar(64) DEFAULT NULL,
  `departname` varchar(64) DEFAULT NULL,
  `created` date DEFAULT NULL,
  PRIMARY KEY (`uid`)
) ENGINE=InnoDB AUTO_INCREMENT=17 DEFAULT CHARSET=latin1


CREATE TABLE `book` (
  `id` int(10) NOT NULL AUTO_INCREMENT,
  `uname` varchar(64) DEFAULT NULL,
  `email` varchar(32) NOT NULL,
  `content` text,
  `insert_time` int(10) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=77 DEFAULT CHARSET=utf8

CREATE TABLE `userdetail` (
  `uid` int(10) NOT NULL DEFAULT '0',
  `intro` text,
  `profile` text,
  PRIMARY KEY (`uid`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1