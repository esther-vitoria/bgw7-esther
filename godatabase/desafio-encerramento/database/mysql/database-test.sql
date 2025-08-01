-- DDL
DROP DATABASE IF EXISTS `fantasy_products_test`;

CREATE DATABASE `fantasy_products_test`;

USE `fantasy_products_test`;

-- Table structure for table `customers`
CREATE TABLE `customers` (
    `id` int DEFAULT NULL,
    `first_name` varchar(45) DEFAULT NULL,
    `last_name` varchar(45) DEFAULT NULL,
    `condition` tinyint(1) DEFAULT NULL
);

-- Table structure for table `invoices`
CREATE TABLE `invoices` (
    `id` int DEFAULT NULL,
    `datetime` datetime DEFAULT NULL,
    `customer_id` int DEFAULT NULL,
    `total` float DEFAULT NULL
);

-- Table structure for table `products`
CREATE TABLE `products` (
    `id` int DEFAULT NULL,
    `description` varchar(100) DEFAULT NULL,
    `price` float DEFAULT NULL
);

-- Table structure for table `sales`
CREATE TABLE `sales` (
    `id` int DEFAULT NULL,
    `quantity` int DEFAULT NULL,
    `invoice_id` int DEFAULT NULL,
    `product_id` int DEFAULT NULL
);

-- Data for table `customers`
INSERT INTO `customers` VALUES 
(1,'John','Doe',1),
(2,'Jane','Doe',1),
(3,'Alice','Smith',0),
(4,'Bob','Smith',0),
(5,'Charlie','Brown',1),
(6,'David','Jones',0),
(7,'Eve','Johnson',1),
(8,'Frank','Davis',0);

-- Data for table `invoices`
INSERT INTO `invoices` VALUES
(1,'2023-01-01 10:00:00',1,100.00),
(2,'2023-01-02 11:00:00',2,200.00),
(3,'2023-01-03 12:00:00',3,300.00),
(4,'2023-01-04 13:00:00',4,400.00),
(5,'2023-01-05 14:00:00',5,500.00),
(6,'2023-01-06 15:00:00',6,600.00),
(7,'2023-01-07 16:00:00',7,700.00),
(8,'2023-01-08 17:00:00',8,800.00);

-- Data for table `products`
INSERT INTO `products` VALUES
(1,'Product 1',10.00),
(2,'Product 2',20.00),
(3,'Product 3',30.00),
(4,'Product 4',40.00),
(5,'Product 5',50.00),
(6,'Product 6',60.00);

-- Data for table `sales`
INSERT INTO `sales` VALUES
(1,1,1,1),
(2,2,1,2),
(3,3,2,5),
(4,4,2,2),
(5,5,3,1),
(6,6,3,2),
(7,7,4,1),
(8,8,4,5),
(9,9,5,1),
(10,10,5,2),
(11,11,6,5),
(12,12,6,2),
(13,13,7,1),
(14,14,7,6),
(15,15,8,1),
(16,16,8,2),
(17,17,1,3),
(18,18,1,4),
(19,19,2,3),
(20,20,2,4),
(21,21,3,3),
(22,22,3,4),
(23,23,4,5),
(24,24,4,4),
(25,25,5,3),
(26,26,5,4),
(27,27,6,6),
(28,28,6,4),
(29,29,7,6),
(30,30,7,4),
(31,31,8,3),
(32,32,8,4);