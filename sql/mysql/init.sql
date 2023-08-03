CREATE TABLE customers (
  [Id] varchar(36) NOT NULL,
  [Name] varchar(100) NOT NULL,
  [Email] varchar(100) NOT NULL,
  [BirthDate] datetime(6) NOT NULL,
  [DateCreated] datetime(6) NOT NULL,
  [DateUpdated] datetime(6) DEFAULT NULL,  
  CONSTRAINT PK_Customers PRIMARY KEY ([ID])
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
