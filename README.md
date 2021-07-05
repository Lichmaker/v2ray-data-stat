# v2ray-data-stat
简单易用的v2ray流量统计工具！

## 为什么需要这样一个工具？
作者原本在使用 PHP 进行流量统计相关处理， 但是PHP使用 grpc 实在不太友好，而且想要实现秒级定时统计还需要借助别的工具，十分麻烦。  
加上初学 golang 有一段时间了，所以来练练手。

如果你也需要进行 v2ray 流量统计，但是不想麻烦的再去编写 grpc 代码，那么本工具必定很适合你！  
而且最终容器大小仅 25 MB左右，超级轻便！

## 使用

1. 拉取仓库代码

```
git clone https://github.com/Lichmaker/v2ray-data-stat
```

2. 修改配置文件

```
cd v2ray-data-stat
cp cmd/stat/.env.example  cmd/stat/.env
vim cmd/stat/.env
```

配置文件 `.env` 说明 ：
```
# 使用 MYSQL 进行数据存储
DB_HOST=172.17.0.1
DB_PORT=3306
DB_DATABASE=yourdatabase
DB_USERNAME=yourusername
DB_PASSWORD=yourpassword

# GPRC HOST，需要提前给您的 v2ray server端进行流量统计配置和GPRC配置，具体可以查看 v2ray 官方文档
GRPC_HOST=your-grpc-host.example.test:80

# 是否对流量进行重置。为TRUE后将会在每次查询后对所有流量数据进行清零。
STAT_RESET=FALSE
```

3. docker 启动

```
docker build -t lichmaker/v2ray-data-stat .
docker run --rm -d lichmaker/v2ray-data-stat --name="v2ray-data-stat"
```

## 数据存储
最终数据会存储在 MYSQL 中，建立2个新表。 `data_summary` 和 `data_statistics_xxxxxx` 中。其中 `data_statistics_xxxxxx` 会根据月份进行分表，例如 2021-07 则会生成 `data_statistics_202107`  

**程序会自动进行数据表迁移，不需要手动操作**
```
CREATE TABLE `data_summary` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `uplink_byte` bigint unsigned DEFAULT NULL,
  `downlink_byte` bigint unsigned DEFAULT NULL,
  `date` date DEFAULT NULL,
  `username` varchar(64) NOT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `date_username` (`date`,`username`),
  KEY `idx_data_summary_created_at` (`created_at`),
  KEY `idx_data_summary_updated_at` (`updated_at`),
  KEY `idx_data_summary_username` (`username`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;


CREATE TABLE `data_statistics_xxxxxx` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `uplink_byte` bigint unsigned DEFAULT NULL,
  `downlink_byte` bigint unsigned DEFAULT NULL,
  `username` varchar(64) NOT NULL,
  `created_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_data_statistics_202107_username` (`username`),
  KEY `idx_data_statistics_202107_created_at` (`created_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
```

## 统计频率
因为懒，没有做成配置项。目前是固定每5秒进行一次统计。

