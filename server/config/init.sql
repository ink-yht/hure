-- 用户表
CREATE TABLE User (
                      ID INT AUTO_INCREMENT PRIMARY KEY,
                      OpenID VARCHAR(50) NOT NULL UNIQUE,
                      Nickname VARCHAR(100),
                      Avatar VARCHAR(255),
                      Phone VARCHAR(11),
                      Gender TINYINT,
                      Role TINYINT DEFAULT 1,
                      Status TINYINT DEFAULT 1,
                      CreatedAt BIGINT,
                      UpdatedAt BIGINT
);

-- 简历表
CREATE TABLE Resume (
                        ID INT AUTO_INCREMENT PRIMARY KEY,
                        UserID INT NOT NULL UNIQUE,  -- 这里有逗号隔开
                        Content TEXT,  -- 这里有逗号隔开
                        FilePath VARCHAR(255),
                        CreatedAt BIGINT,
                        UpdatedAt BIGINT,
                        FOREIGN KEY (UserID) REFERENCES User(ID)
);

-- 招聘者营业执照表
CREATE TABLE BusinessLicense (
                                 ID INT AUTO_INCREMENT PRIMARY KEY,
                                 UserID INT NOT NULL UNIQUE,
                                 Content TEXT,
                                 FilePath VARCHAR(255),
                                 CreatedAt BIGINT,
                                 UpdatedAt BIGINT,
                                 FOREIGN KEY (UserID) REFERENCES User(ID)
);

-- 实名审核表
CREATE TABLE RealNameVerification (
                                      ID INT AUTO_INCREMENT PRIMARY KEY,
                                      UserID INT NOT NULL UNIQUE,
                                      RealName VARCHAR(50),
                                      IDCardNumber VARCHAR(18),
                                      IDCardFrontImg VARCHAR(255),
                                      IDCardBackImg VARCHAR(255),
                                      VerifiedStatus TINYINT DEFAULT 0,
                                      VerifiedTime BIGINT,
                                      CreatedAt BIGINT,
                                      UpdatedAt BIGINT,
                                      FOREIGN KEY (UserID) REFERENCES User(ID)
);

-- 管理员表
CREATE TABLE Admin (
                       ID INT AUTO_INCREMENT PRIMARY KEY,
                       Email VARCHAR(100) NOT NULL UNIQUE,
                       Password VARCHAR(255) NOT NULL,
                       Phone VARCHAR(20) UNIQUE,
                       Avatar VARCHAR(255),
                       Nickname VARCHAR(32),
                       Signature VARCHAR(128),
                       CreatedAt BIGINT,
                       UpdatedAt BIGINT
);

-- 招聘订单表
CREATE TABLE RecruitOrder (
                              ID INT AUTO_INCREMENT PRIMARY KEY,
                              RecruiterID INT NOT NULL,
                              Title VARCHAR(100) NOT NULL,
                              Description TEXT,
                              Status TINYINT(20) DEFAULT 1,
                              ServiceFee BIGINT NOT NULL,
                              CreatedAt BIGINT,
                              UpdatedAt BIGINT,
                              FOREIGN KEY (RecruiterID) REFERENCES User(ID)
);

-- 求职申请表
CREATE TABLE JobApplication (
                                ID INT AUTO_INCREMENT PRIMARY KEY,
                                OrderID INT NOT NULL,
                                JobSeekerID INT NOT NULL,
                                Status TINYINT(20) DEFAULT 1,
                                CreatedAt BIGINT,
                                UpdatedAt BIGINT,
                                FOREIGN KEY (OrderID) REFERENCES RecruitOrder(ID),
                                FOREIGN KEY (JobSeekerID) REFERENCES User(ID)
);

-- 评价表
CREATE TABLE Review (
                        ID INT AUTO_INCREMENT PRIMARY KEY,
                        ReviewerID INT NOT NULL,
                        ReviewedUserID INT NOT NULL,
                        OrderID INT NOT NULL,
                        Content TEXT,
                        Rating TINYINT NOT NULL,
                        CreatedAt BIGINT,
                        UpdatedAt BIGINT,
                        FOREIGN KEY (ReviewerID) REFERENCES User(ID),
                        FOREIGN KEY (ReviewedUserID) REFERENCES User(ID),
                        FOREIGN KEY (OrderID) REFERENCES RecruitOrder(ID)
);

-- 金额操作日志表
CREATE TABLE TransactionLog (
                                ID INT AUTO_INCREMENT PRIMARY KEY,
                                UserID INT NOT NULL,
                                OrderID INT NOT NULL,
                                Amount DOUBLE NOT NULL,
                                Type TINYINT(20),
                                Description VARCHAR(255),
                                CreatedAt BIGINT,
                                UpdatedAt BIGINT,
                                FOREIGN KEY (UserID) REFERENCES User(ID),
                                FOREIGN KEY (OrderID) REFERENCES RecruitOrder(ID)
);