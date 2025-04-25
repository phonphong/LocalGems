CREATE TABLE cafes (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  name TEXT NOT NULL,
  rating REAL,
  reviews INTEGER,
  price_range TEXT,
  type TEXT,
  address TEXT,
  review_text TEXT
);

INSERT INTO cafes (name, rating, reviews, price_range, type, address, review_text) VALUES ('L''Usine', 4.5, 2400, '₫₫₫', 'Nhà hàng', '19 Lê Thánh Tôn', 'Menu nhiều món đa dạng, vừa vị, decor đẹp mắt.');
INSERT INTO cafes (name, rating, reviews, price_range, type, address, review_text) VALUES ('Cà phê Linh', 4.4, 1200, '1-100.000 ₫', 'Quán cà phê', '1 Trương Định', '');
INSERT INTO cafes (name, rating, reviews, price_range, type, address, review_text) VALUES ('Oromia Coffee & Lounge', 3.9, 556, '1-100.000 ₫', 'Quán cà phê', '85 Phan Kế Bính', '');
INSERT INTO cafes (name, rating, reviews, price_range, type, address, review_text) VALUES ('Soo Kafe', 4.5, 389, '1-100.000 ₫', 'Quán cà phê', '10 Phan Kế Bính', '');
INSERT INTO cafes (name, rating, reviews, price_range, type, address, review_text) VALUES ('THE COFFEE LAB', 4.2, 412, '1-100.000 ₫', 'Quán cà phê', '53A Nguyễn Du', '');