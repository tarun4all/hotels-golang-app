USE test; 

CREATE TABLE `geolocation` (
  `ipAddress` varchar(20) NOT NULL PRIMARY KEY,
  `countryCode` varchar(2) NOT NULL,
  `country` varchar(60) NOT NULL,
  `city` varchar(60) NOT NULL,
  `latitude` varchar(20) NOT NULL,
  `longitude` varchar(20) NOT NULL,
  `createdAt` varchar(60) NOT NULL
);