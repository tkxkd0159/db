CREATE TABLE sections (
  section_id INT AUTO_INCREMENT PRIMARY KEY,
  section_name CHAR(3) NOT NULL,
  floor TINYINT NOT NULL
);

INSERT INTO sections (section_name, floor)
VALUES  ('한식', 2),
        ('분식', 2),
        ('중식', 3),
        ('일식', 3),
        ('양식', 3),
        ('카페', 1),
        ('디저트', 1);

CREATE TABLE businesses (
  business_id INT AUTO_INCREMENT PRIMARY KEY,
  fk_section_id INT NOT NULL,
  business_name VARCHAR(10) NOT NULL,
  status CHAR(3) DEFAULT 'OPN' NOT NULL,
  can_takeout TINYINT DEFAULT 1 NOT NULL
);

INSERT INTO businesses (fk_section_id, business_name, status, can_takeout)
VALUES  (3, '화룡각', 'OPN', 1),
        (2, '철구분식', 'OPN', 1),
        (5, '얄코렐라', 'RMD', 1),
        (2, '바른떡볶이', 'OPN', 1),
        (1, '북극냉면', 'OPN', 0),
        (1, '보쌈마니아', 'OPN', 1),
        (5, '에그사라다', 'VCT', 1),
        (6, '달다방', 'OPN', 1),
        (7, '마카오마카롱', 'OPN', 1),
        (2, '김밥마라', 'OPN', 1),
        (7, '소소스윗', 'OPN', 1),
        (4, '사사서셔소쇼스시', 'VCT', 1),
        (3, '린민짬뽕', 'CLS', 1),
        (7, '파시조아', 'OPN', 1),
        (1, '할매장국', 'CLS', 0),
        (5, '노선이탈리아', 'OPN', 1),
        (6, '커피앤코드', 'OPN', 1),
        (2, '신림동백순대', 'VCT', 1);

CREATE TABLE menus (
  menu_id INT AUTO_INCREMENT PRIMARY KEY,
  fk_business_id INT NOT NULL,
  menu_name VARCHAR(20) NOT NULL,
  kilocalories DECIMAL(7,2) NOT NULL,
  price INT NOT NULL,
  likes INT DEFAULT 0 NOT NULL
);

INSERT INTO menus (fk_business_id, menu_name, kilocalories, price, likes)
VALUES  (5, '물냉면', 480.23, 8000, 3),
        (8, '아메리카노', 16.44, 4500, 6),
        (17, '고르곤졸라피자', 1046.27, 12000, 12),
        (6, '보쌈', 1288.24, 14000, 2),
        (15, '장국', 387.36, 8500, -1),
        (17, '까르보나라', 619.11, 9000, 10),
        (9, '바닐라마카롱', 160.62, 1500, 4),
        (16, '백순대', 681.95, 11000, 24),
        (6, '마늘보쌈', 1320.49, 16000, 7),
        (16, '양념순대볶음', 729.17, 12000, 0),
        (14, '단팥빵', 225.88, 1500, 13),
        (1, '간짜장', 682.48, 7000, 3),
        (9, '뚱카롱', 247.62, 2000, 8),
        (5, '비빔냉면', 563.45, 8000, 4),
        (10, '참치김밥', 532.39, 3000, 0),
        (2, '치즈떡볶이', 638.42, 5000, 15),
        (11, '플레인와플', 299.31, 6500, 2),
        (2, '찹쌀순대', 312.76, 3000, -4),
        (15, '육개장', 423.18, 8500, 2),
        (4, '국물떡볶이', 483.29, 4500, 1),
        (10, '돈가스김밥', 562.72, 4000, 0),
        (1, '삼선짬뽕', 787.58, 8000, 32),
        (11, '수플레팬케익', 452.37, 9500, 5),
        (4, '라볶이', 423.16, 5500, 0),
        (8, '모카프라푸치노', 216.39, 6000, 8),
        (14, '옛날팥빙수', 382.35, 8000, 2);

CREATE TABLE ratings (
  rating_id INT AUTO_INCREMENT PRIMARY KEY,
  fk_business_id INT NOT NULL,
  stars TINYINT NOT NULL,
  comment VARCHAR(200) null,
  created timestamp DEFAULT CURRENT_TIMESTAMP NOT NULL
);

INSERT INTO ratings (fk_business_id, stars, comment, created)
VALUES  (2, 4, '치떡이 진리. 순대는 별로', '2021-07-01 12:30:04'),
        (16, 3, '그냥저냥 먹을만해요', '2021-07-01 17:16:07'),
        (14, 5, '인생팥빵. 말이 필요없음', '2021-07-03 11:28:12'),
        (5, 3, '육수는 괜찮은데 면은 그냥 시판면 쓴 것 같네요.', '2021-07-04 19:03:50'),
        (11, 4, '나오는데 넘 오래걸림. 맛은 있어요', '2021-07-04 13:37:42'),
        (9, 2, '빵집에서 파는 마카롱이랑 비슷하거나 못합니다.', '2021-07-06 15:19:23'),
        (16, 5, '신림에서 먹던 맛 완벽재현', '2021-07-06 20:01:39');