-- phpMyAdmin SQL Dump
-- version 5.2.1
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Generation Time: Aug 22, 2024 at 07:07 PM
-- Server version: 10.4.32-MariaDB
-- PHP Version: 8.0.30

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `bookmanagerdb`
--

-- --------------------------------------------------------

--
-- Table structure for table `book`
--

CREATE TABLE `book` (
  `id_book` int(11) NOT NULL,
  `title` varchar(50) NOT NULL,
  `author` varchar(50) NOT NULL,
  `publishedyear` year(4) NOT NULL,
  `genre` varchar(25) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `book`
--

INSERT INTO `book` (`id_book`, `title`, `author`, `publishedyear`, `genre`) VALUES
(2, 'Bumi Manusia', 'Pramoedya Ananta Toer', '1980', 'Historical'),
(3, 'Ayat-Ayat Cinta', 'Habiburrahman El Shirazy', '2004', 'Romance'),
(4, 'Perahu Kertas', 'Dee Lestari', '2009', 'Romance'),
(5, 'Cantik Itu Luka', 'Eka Kurniawan', '2002', 'Fiction'),
(6, 'Supernova: Ksatria, Puteri, dan Bintang Jatuh', 'Dee Lestari', '2001', 'Science Fiction'),
(7, 'Negeri 5 Menara', 'Ahmad Fuadi', '2009', 'Fiction'),
(8, 'Tetralogi Buru: Anak Semua Bangsa', 'Pramoedya Ananta Toer', '1981', 'Historical'),
(9, 'Orang-Orang Biasa', 'Andrea Hirata', '2019', 'Fiction'),
(10, 'Ronggeng Dukuh Paruk', 'Ahmad Tohari', '1982', 'Historical'),
(11, 'Laskar Pelangi', 'Andrea Hirata', '2004', 'Fiction'),
(12, 'tes', 'tes', '2020', 'Fiction'),
(13, 'coba', 'coba', '2029', 'Romance');

--
-- Indexes for dumped tables
--

--
-- Indexes for table `book`
--
ALTER TABLE `book`
  ADD PRIMARY KEY (`id_book`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `book`
--
ALTER TABLE `book`
  MODIFY `id_book` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=14;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
