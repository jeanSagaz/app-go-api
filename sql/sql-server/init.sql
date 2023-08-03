IF NOT EXISTS (SELECT * FROM sys.databases WHERE name = 'poc')
BEGIN
  CREATE DATABASE poc;
END;
GO

USE poc
GO

IF OBJECT_ID('Customers', 'U') IS NULL
BEGIN

  CREATE TABLE [dbo].[Customers](
	[Id] [varchar(36)] NOT NULL,
	-- [Id] [uniqueidentifier] NOT NULL,
	[Name] [varchar](100) NOT NULL,
	[Email] [varchar](100) NOT NULL,
	[Created_At] [datetime2](7) NOT NULL,
	[Updated_At] [datetime2](7) NULL,
 CONSTRAINT [PK_Customers] PRIMARY KEY CLUSTERED 
(
	[Id] ASC
)WITH (PAD_INDEX = OFF, STATISTICS_NORECOMPUTE = OFF, IGNORE_DUP_KEY = OFF, ALLOW_ROW_LOCKS = ON, ALLOW_PAGE_LOCKS = ON) ON [PRIMARY]
) ON [PRIMARY]

END;
GO
