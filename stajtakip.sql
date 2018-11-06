-- --------------------------------------------------------
-- Host:                         localhost
-- Server version:               5.7.18 - MySQL Community Server (GPL)
-- Server OS:                    Win64
-- HeidiSQL Version:             9.4.0.5125
-- --------------------------------------------------------

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET NAMES utf8 */;
/*!50503 SET NAMES utf8mb4 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;


-- Dumping database structure for stajtest
CREATE DATABASE IF NOT EXISTS `stajtest` /*!40100 DEFAULT CHARACTER SET latin1 */;
USE `stajtest`;

-- Dumping structure for table stajtest.gorevliler
CREATE TABLE IF NOT EXISTS `gorevliler` (
  `AdSoyad` varchar(50) NOT NULL,
  `Komisyon` tinyint(1) unsigned NOT NULL DEFAULT '0',
  PRIMARY KEY (`AdSoyad`)
) ENGINE=InnoDB DEFAULT CHARSET=latin5;

-- Data exporting was unselected.
-- Dumping structure for table stajtest.mulakatlar
CREATE TABLE IF NOT EXISTS `mulakatlar` (
  `OgrenciNo` int(11) NOT NULL,
  `MulatakTarihi` datetime NOT NULL,
  `GörevliUye1` varchar(50) NOT NULL,
  `GorevliUye2` varchar(50) NOT NULL,
  PRIMARY KEY (`MulatakTarihi`),
  KEY `OgrNo` (`OgrenciNo`),
  KEY `Görevli` (`GörevliUye1`),
  KEY `Gorevli2` (`GorevliUye2`),
  CONSTRAINT `Gorevli2` FOREIGN KEY (`GorevliUye2`) REFERENCES `gorevliler` (`AdSoyad`),
  CONSTRAINT `Görevli` FOREIGN KEY (`GörevliUye1`) REFERENCES `gorevliler` (`AdSoyad`),
  CONSTRAINT `OgrNo` FOREIGN KEY (`OgrenciNo`) REFERENCES `ogrenciler` (`No`)
) ENGINE=InnoDB DEFAULT CHARSET=latin5;

-- Data exporting was unselected.
-- Dumping structure for table stajtest.ogrenciler
CREATE TABLE IF NOT EXISTS `ogrenciler` (
  `No` int(11) NOT NULL,
  `Ad` varchar(50) NOT NULL,
  `Soyad` varchar(50) NOT NULL,
  `ToplamStaj` tinyint(2) unsigned NOT NULL DEFAULT '0',
  `KabulEdilen` tinyint(2) unsigned NOT NULL DEFAULT '0',
  `Ogretim` tinyint(1) unsigned NOT NULL,
  `SonStajTarihi` date DEFAULT NULL,
  PRIMARY KEY (`No`)
) ENGINE=InnoDB DEFAULT CHARSET=latin5;

-- Data exporting was unselected.
-- Dumping structure for table stajtest.stajlar
CREATE TABLE IF NOT EXISTS `stajlar` (
  `OgrenciNo` int(11) NOT NULL,
  `Kurum` varchar(50) NOT NULL,
  `Sehir` varchar(50) NOT NULL,
  `Baslangic` date NOT NULL,
  `Bitis` date NOT NULL,
  `Konu` varchar(50) NOT NULL,
  `Sinif` tinyint(1) unsigned NOT NULL,
  `KabulEdilen` tinyint(2) unsigned NOT NULL DEFAULT '0',
  `Degerlendirildi` tinyint(1) unsigned NOT NULL DEFAULT '0',
  PRIMARY KEY (`OgrenciNo`,`Baslangic`),
  CONSTRAINT `OgrenciNo` FOREIGN KEY (`OgrenciNo`) REFERENCES `ogrenciler` (`No`)
) ENGINE=InnoDB DEFAULT CHARSET=latin5;

-- Data exporting was unselected.
/*!40101 SET SQL_MODE=IFNULL(@OLD_SQL_MODE, '') */;
/*!40014 SET FOREIGN_KEY_CHECKS=IF(@OLD_FOREIGN_KEY_CHECKS IS NULL, 1, @OLD_FOREIGN_KEY_CHECKS) */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
