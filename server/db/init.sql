CREATE TABLE IF NOT EXISTS `items` (
    `id`            int(11)         NOT NULL AUTO_INCREMENT,
    `name`          varchar(255)    NOT NULL,
    `quantity`      int(11)         NOT NULL,
    `unit_price`    decimal(10,2)   NOT NULL,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
