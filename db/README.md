
## 料金ファイル
https://pricing.us-east-1.amazonaws.com/offers/v1.0/aws/AmazonEC2/current/index.json
https://pricing.us-east-1.amazonaws.com/offers/v1.0/aws/AmazonRDS/current/index.json

## スキーマ変更

```
psqldef -U postgres -h 127.0.0.1 cloud_cost --password=adminadmin --skip-view  < schema.sql
xo schema -d template/ -o models postgres://postgres:adminadmin@127.0.0.1/cloud_cost?sslmode=disable
```

```golang
	Vcpu    StringInt     `json:"vcpu"     gorm:"column:vcpu"`
    Memory  StringFloat   `json:"memory"   gorm:"column:memory"`
```


```mermaid
erDiagram
    USERS {
        UserID INT PK
        UserName VARCHAR(255)
    }
    SERVICES {
        ServiceID INT PK
        ServiceType VARCHAR(50)
    }
    AWS_SERVICES {
        ServiceID INT PK
        ServiceName VARCHAR(255)
    }
    AZURE_SERVICES {
        ServiceID INT PK
        ServiceName VARCHAR(255)
    }
    USER_SERVICES {
        UserID INT PK
        ServiceID INT
        ServiceType VARCHAR(50)
    }
    USERS ||--|| USER_SERVICES : selects

```

    USER_SERVICES ||--| SERVICES : refers to
    SERVICES ||--o| AWS_SERVICES : has
    SERVICES ||--o| AZURE_SERVICES : has

        PRIMARY KEY (UserID, ServiceID),
        FOREIGN KEY (UserID) REFERENCES USERS(UserID),
        FOREIGN KEY (ServiceID) REFERENCES SERVICES(ServiceID)