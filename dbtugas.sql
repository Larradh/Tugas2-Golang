/*
SQLyog Community v13.1.7 (64 bit)
MySQL - 10.4.6-MariaDB : Database - dbtugas
*********************************************************************
*/

/*!40101 SET NAMES utf8 */;

/*!40101 SET SQL_MODE=''*/;

/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;
CREATE DATABASE /*!32312 IF NOT EXISTS*/`dbtugas` /*!40100 DEFAULT CHARACTER SET latin1 */;

USE `dbtugas`;

/*Table structure for table `mahasiswa` */

DROP TABLE IF EXISTS `mahasiswa`;

CREATE TABLE `mahasiswa` (
  `IDmahasiswa` varchar(12) NOT NULL,
  `Nama` varchar(100) NOT NULL,
  `Jalan` varchar(100) NOT NULL,
  `Kelurahan` varchar(100) NOT NULL,
  `Kecamatan` varchar(100) NOT NULL,
  `Kabupaten` varchar(100) NOT NULL,
  `Provinsi` varchar(100) NOT NULL,
  `Fakultas` varchar(100) NOT NULL,
  `Jurusan` varchar(100) NOT NULL,
  PRIMARY KEY (`IDmahasiswa`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

/*Data for the table `mahasiswa` */

insert  into `mahasiswa`(`IDmahasiswa`,`Nama`,`Jalan`,`Kelurahan`,`Kecamatan`,`Kabupaten`,`Provinsi`,`Fakultas`,`Jurusan`) values 
('1811081021','Larra','Jl. Durian','Jati Makmur','Pondok Gede','Bekasi','Jawa Barat','TI','RPL');

/*Table structure for table `matkul` */

DROP TABLE IF EXISTS `matkul`;

CREATE TABLE `matkul` (
  `IDmatkul` varchar(12) NOT NULL,
  `MataKuliah` varchar(100) NOT NULL,
  PRIMARY KEY (`IDmatkul`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

/*Data for the table `matkul` */

insert  into `matkul`(`IDmatkul`,`MataKuliah`) values 
('P001','P. WEB');

/*Table structure for table `nilai` */

DROP TABLE IF EXISTS `nilai`;

CREATE TABLE `nilai` (
  `IDmahasiswa` varchar(12) NOT NULL,
  `IDmatkul` varchar(12) NOT NULL,
  `Nilai` float NOT NULL,
  `Semester` int(11) NOT NULL,
  KEY `IDmahasiswa` (`IDmahasiswa`,`IDmatkul`),
  KEY `IDmatkul` (`IDmatkul`),
  CONSTRAINT `nilai_ibfk_1` FOREIGN KEY (`IDmahasiswa`) REFERENCES `mahasiswa` (`IDmahasiswa`) ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT `nilai_ibfk_2` FOREIGN KEY (`IDmatkul`) REFERENCES `matkul` (`IDmatkul`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

/*Data for the table `nilai` */

insert  into `nilai`(`IDmahasiswa`,`IDmatkul`,`Nilai`,`Semester`) values 
('1811081021','P001',3.75,4);

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;
