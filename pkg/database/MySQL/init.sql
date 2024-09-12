CREATE TABLE IF NOT EXISTS users (
    user_id BIGINT AUTO_INCREMENT PRIMARY KEY,
    username VARCHAR(50) UNIQUE,
    email VARCHAR(70) UNIQUE,
    password VARCHAR(255) NOT NULL,
    status ENUM('active', 'inactive') NOT NULL,
    blocked ENUM('blocked','unblocked') DEFAULT 'unblocked'
    registered_at DATETIME NOT NULL,
    last_login DATETIME NOT NULL,
    role ENUM('user', 'admin') DEFAULT 'user'
) ENGINE=InnoDB;

CREATE TABLE IF NOT EXISTS person(
    person_id BIGINT AUTO_INCREMENT PRIMARY KEY,
    user_id BIGINT NOT NULL,
    first_name VARCHAR(50) NOT NULL,
    last_name VARCHAR(50) NOT NULL,
    birth_day DATETIME NOT NULL,
    gender ENUM('male', 'female','other') NOT NULL,
    country VARCHAR(50),
    city VARCHAR(90),
    image LONGTEXT,
    FOREIGN KEY(user_id) REFERENCES users(user_id) ON DELETE CASCADE
) ENGINE=InnoDB;

CREATE TABLE IF NOT EXISTS favorite_articles (
    user_id BIGINT NOT NULL,
    article_id BIGINT NOT NULL,
    PRIMARY KEY(user_id, article_id),
    FOREIGN KEY(user_id) REFERENCES user(user_id) ON DELETE CASCADE,
    FOREIGN KEY(article_id) REFERENCES articles(article_id) ON DELETE CASCADE
) ENGINE=InnoDB;

CREATE TABLE IF NOT EXISTS articles (
    article_id BIGINT AUTO_INCREMENT PRIMARY KEY,
    title VARCHAR(50) UNIQUE,
    lead_section TEXT,
    image LONGTEXT
) ENGINE=InnoDB;

CREATE TABLE IF NOT EXISTS chapters (
    chapter_id BIGINT AUTO_INCREMENT PRIMARY KEY,
    article_id BIGINT NOT NULL,
    parent_id BIGINT,
    name VARCHAR(100) NOT NULL,
    content LONGTEXT NOT NULL,
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


CREATE TABLE IF NOT EXISTS person_info_box (
    person_info_box_id INT AUTO_INCREMENT PRIMARY KEY,
    full_name VARCHAR(255),
    birth_date DATETIME,
    place_of_birth VARCHAR(255),
    nationality VARCHAR(255),
    death_date DATETIME,
    educations TEXT,
    occupations TEXT,
    parents TEXT,
    children TEXT
) ENGINE=InnoDB;

CREATE TABLE IF NOT EXISTS building_info_box (
    building_info_box_id INT AUTO_INCREMENT PRIMARY KEY,
    full_name VARCHAR(255),
    location VARCHAR(255),
    founded DATETIME,
    architects TEXT,
    height INT,
    floors INT,
    `usage` VARCHAR(255)
) ENGINE=InnoDB;

CREATE TABLE IF NOT EXISTS company_info_box (
    company_info_box_id INT AUTO_INCREMENT PRIMARY KEY,
    full_name VARCHAR(255),
    founded DATETIME,
    industry VARCHAR(255),
    headquarters VARCHAR(255),
    ceo VARCHAR(255),
    revenue VARCHAR(255),
    employees INT
) ENGINE=InnoDB;

CREATE TABLE IF NOT EXISTS event_info_box (
    event_info_box_id INT AUTO_INCREMENT PRIMARY KEY,
    title VARCHAR(255),
    date DATETIME,
    locations TEXT,
    organizers TEXT,
    attendance INT
) ENGINE=InnoDB;

CREATE TABLE IF NOT EXISTS film_info_box (
    film_info_box_id INT AUTO_INCREMENT PRIMARY KEY,
    title VARCHAR(255),
    release_date DATETIME,
    director VARCHAR(255),
    genre VARCHAR(255),
    duration TIME,
    languages TEXT,
    rating DECIMAL(3, 1)
) ENGINE=InnoDB;

CREATE TABLE IF NOT EXISTS book_info_box (
    book_info_box_id INT AUTO_INCREMENT PRIMARY KEY,
    title VARCHAR(255),
    authors TEXT,
    publish_date DATETIME,
    genre VARCHAR(255),
    isbn VARCHAR(255),
    pages INT,
    publisher VARCHAR(255)
) ENGINE=InnoDB;

CREATE TABLE IF NOT EXISTS album_info_box (
    album_info_box_id INT AUTO_INCREMENT PRIMARY KEY,
    title VARCHAR(255),
    artists TEXT,
    release_date DATETIME,
    genre VARCHAR(255),
    label VARCHAR(255),
    tracks TEXT
) ENGINE=InnoDB;


CREATE TABLE IF NOT EXISTS animal_info_box (
    animal_info_box_id INT AUTO_INCREMENT PRIMARY KEY,
    full_name VARCHAR(255),
    species VARCHAR(255),
    habitat VARCHAR(255),
    conservation_status BOOLEAN,
    lifespan VARCHAR(255)
) ENGINE=InnoDB;

CREATE TABLE IF NOT EXISTS award_info_box (
    award_info_box_id INT AUTO_INCREMENT PRIMARY KEY,
    full_name VARCHAR(255),
    year DATETIME,
    category VARCHAR(255),
    recipients TEXT,
    organizations TEXT
) ENGINE=InnoDB;

CREATE TABLE IF NOT EXISTS song_info_box (
    song_info_box_id INT AUTO_INCREMENT PRIMARY KEY,
    title VARCHAR(255),
    artists TEXT,
    album VARCHAR(255),
    release_date DATETIME,
    genre VARCHAR(255),
    duration TIME,
    label VARCHAR(255)
) ENGINE=InnoDB;

CREATE TABLE IF NOT EXISTS country_info_box (
    country_info_box_id  INT AUTO_INCREMENT PRIMARY KEY,
    full_name VARCHAR(255),
    capital VARCHAR(255),
    population INT,
    region VARCHAR(255),
    area FLOAT,
    currency VARCHAR(255),
    official_language VARCHAR(255)
)ENGINE=InnoDB;


CREATE TABLE IF NOT EXISTS university_info_box (
    university_info_box_id INT AUTO_INCREMENT PRIMARY KEY,
    full_name VARCHAR(255),
    founded DATETIME,
    location VARCHAR(255),
    programs TEXT,
    enrollment INT,
    affiliation VARCHAR(255),
    accreditation VARCHAR(255)
) ENGINE=InnoDB;

CREATE TABLE IF NOT EXISTS museum_info_box (
    museum_info_box_id INT AUTO_INCREMENT PRIMARY KEY,
    full_name VARCHAR(255),
    location VARCHAR(255),
    established DATETIME,
    collections TEXT,
    director VARCHAR(255),
    exhibitions TEXT
) ENGINE=InnoDB;

CREATE TABLE IF NOT EXISTS political_info_box (
    political_info_box_id INT AUTO_INCREMENT PRIMARY KEY,
    position VARCHAR(255),
    office_holder VARCHAR(255),
    term_start DATETIME,
    term_end DATETIME,
    party VARCHAR(255),
    constituency VARCHAR(255)
) ENGINE=InnoDB;
