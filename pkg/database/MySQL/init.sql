CREATE TABLE IF NOT EXISTS articles (
    article_id BIGINT AUTO_INCREMENT PRIMARY KEY,
    title VARCHAR(50) UNIQUE,
    lead_section VARCHAR(150),
    image VARCHAR(255)
) ENGINE=InnoDB;

CREATE TABLE IF NOT EXISTS chapters (
    chapter_id BIGINT AUTO_INCREMENT PRIMARY KEY,
    article_id BIGINT NOT NULL,
    name VARCHAR(100) NOT NULL,
    content MEDIUMTEXT NOT NULL,
    FOREIGN KEY (article_id) REFERENCES articles(article_id) ON DELETE CASCADE
) ENGINE=InnoDB;

CREATE TABLE IF NOT EXISTS info_box (
    info_box_id BIGINT AUTO_INCREMENT PRIMARY KEY,
    article_id BIGINT NOT NULL,
    type ENUM('person', 'building', 'company', 'event', 'film', 
              'book', 'album', 'animal', 'award', 'song', 'country', 
              'university', 'museum', 'politicalPosition') NOT NULL,
    object_info_box_id BIGINT NOT NULL,
    FOREIGN KEY (article_id) REFERENCES articles(article_id) ON DELETE CASCADE
) ENGINE=InnoDB;
