CREATE TABLE articles(
    article_id BIGINT AUTO_INCREMENT PRIMARY KEY,
    title VARCHAR(35) UNIQUE,
    lead_section VARCHAR(150) UNIQUE,
    image VARCHAR(255)
)ENGINE=InnoDB;

CREATE TABLE chapters (
    chapter_id BIGINT AUTO_INCREMENT PRIMARY KEY,
    article_id BIGINT NOT NULL,
    name VARCHAR(100) NOT NULL,
    content MEDIUMTEXT NOT NULL,
    FOREIGN KEY (article_id) REFERENCES article(article_id) ON DELETE CASCADE
) ENGINE=InnoDB;

CREATE TABLE info_box(
    info_box_id BIGINT AUTO_INCREMENT PRIMARY KEY,
    article_id BIGINT UNIQUE,
    type ENUM("person", "building", "company", "event", "film", 
    "book", "album", "animal", "award", "song", "country", 
    "university", "museum", "politicalPosition") NOT NULL,
    object_info_box_id BIGINT NOT NULL,
    FOREIGN KEY (article_id) REFERENCES article(article_id) ON DELETE CASCADE
) ENGINE=InnoDB;


// object_info_box person,building....
