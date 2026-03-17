在 Go 中，结构体（`struct`）的方法封装没有“固定模板”，但核心原则是：**围绕结构体的“数据安全”和“业务逻辑”，封装满足“单一职责”的方法**（避免过度封装）。以下是实际开发中最常用、最核心的封装方法类别，结合业务结构体（如 `User`）给出可直接复用的示例，帮你理清“该封装什么、为什么封装”。

### 一、封装方法的核心原则
1. **数据安全**：结构体私有字段（小写开头）只能通过方法操作，避免外部直接修改导致非法值；
2. **逻辑内聚**：和结构体相关的业务逻辑封装为方法，而非散落在外部函数中；
3. **最小必要**：只封装核心操作，不做无意义的“万能方法”；
4. **接口适配**：按需实现标准接口（如 `fmt.Stringer`、`json.Marshaler`），提升兼容性。

### 二、必选封装方法（基础核心）
#### 1. 构造方法（`NewXXX`）—— 最核心的基础方法
Go 没有内置构造函数，必须自定义 `NewXXX()` 函数创建结构体实例，作用是：
- 初始化字段默认值；
- 校验入参合法性；
- 返回合法的结构体实例（避免创建无效对象）。

```go
// 定义 User 结构体（私有字段用小写，保证数据安全）
type User struct {
    id       int64  // 私有：用户ID（不允许外部直接修改）
    username string // 私有：用户名
    age      int    // 私有：年龄
    email    string // 私有：邮箱
    status   int    // 私有：状态（0-禁用，1-正常）
}

// 构造方法：创建并初始化 User 实例（必选）
// 参数：对外暴露必要参数，内部做校验和默认值填充
func NewUser(id int64, username string, age int, email string) (*User, error) {
    // 1. 入参校验（保证数据合法）
    if id <= 0 {
        return nil, fmt.Errorf("用户ID必须大于0")
    }
    if username == "" {
        return nil, fmt.Errorf("用户名不能为空")
    }
    if age < 0 || age > 150 {
        return nil, fmt.Errorf("年龄必须在0-150之间")
    }
    if !strings.Contains(email, "@") {
        return nil, fmt.Errorf("邮箱格式不合法")
    }

    // 2. 初始化默认值（比如状态默认正常）
    return &User{
        id:       id,
        username: username,
        age:      age,
        email:    email,
        status:   1, // 默认正常状态
    }, nil
}
```

#### 2. Getter 方法（`GetXXX`）—— 读取私有字段
如果结构体字段是**小写私有**（推荐做法），必须封装 Getter 方法暴露字段值，保证“只读不写”的控制：
```go
// GetID：获取用户ID（Getter 方法）
func (u *User) GetID() int64 {
    return u.id
}

// GetUsername：获取用户名
func (u *User) GetUsername() string {
    return u.username
}

// GetAge：获取年龄
func (u *User) GetAge() int {
    return u.age
}

// GetStatusText：状态码转文字（增强型 Getter，更贴合业务）
func (u *User) GetStatusText() string {
    switch u.status {
    case 0:
        return "禁用"
    case 1:
        return "正常"
    default:
        return "未知"
    }
}
```

#### 3. Setter 方法（`SetXXX`）—— 修改私有字段
如需允许外部修改字段，必须封装 Setter 方法（而非直接暴露字段），作用是：
- 校验修改后的值是否合法；
- 可添加额外逻辑（如修改日志、状态联动）。

```go
// SetAge：修改年龄（Setter 方法，带校验）
func (u *User) SetAge(newAge int) error {
    if newAge < 0 || newAge > 150 {
        return fmt.Errorf("年龄必须在0-150之间")
    }
    u.age = newAge
    return nil
}

// SetStatus：修改状态（带业务逻辑）
func (u *User) SetStatus(newStatus int) error {
    if newStatus != 0 && newStatus != 1 {
        return fmt.Errorf("状态只能是0（禁用）或1（正常）")
    }
    // 额外逻辑：记录状态修改时间（示例）
    fmt.Printf("用户[%s]状态从%d修改为%d\n", u.username, u.status, newStatus)
    u.status = newStatus
    return nil
}

// UpdateEmail：修改邮箱（复杂 Setter，可包含更多逻辑）
func (u *User) UpdateEmail(newEmail string) error {
    if !strings.Contains(newEmail, "@") {
        return fmt.Errorf("邮箱格式不合法")
    }
    // 额外逻辑：发送邮箱修改通知（示例）
    fmt.Printf("用户[%s]邮箱已修改，新邮箱：%s\n", u.username, newEmail)
    u.email = newEmail
    return nil
}
```

### 三、可选封装方法（业务常用）
#### 1. 状态判断/转换方法
封装和结构体状态相关的判断逻辑，让业务代码更简洁：
```go
// IsAdult：判断是否成年（业务逻辑方法）
func (u *User) IsAdult() bool {
    return u.age >= 18
}

// IsActive：判断是否是活跃用户（状态正常）
func (u *User) IsActive() bool {
    return u.status == 1
}
```

#### 2. 格式化/序列化方法
实现标准接口（如 `fmt.Stringer`）或自定义序列化方法，方便打印、存储：
```go
// String：实现 fmt.Stringer 接口，打印时自动调用
func (u *User) String() string {
    return fmt.Sprintf("User{ID: %d, 用户名: %s, 年龄: %d, 状态: %s}",
        u.id, u.username, u.age, u.GetStatusText())
}

// ToJSON：自定义序列化为JSON字符串（也可直接用encoding/json，此处为示例）
func (u *User) ToJSON() (string, error) {
    data, err := json.Marshal(map[string]interface{}{
        "id":       u.id,
        "username": u.username,
        "age":      u.age,
        "email":    u.email,
        "status":   u.GetStatusText(),
    })
    if err != nil {
        return "", err
    }
    return string(data), nil
}
```

#### 3. 比较/哈希方法
判断两个结构体实例是否相等，或生成唯一标识（方便集合/存储使用）：
```go
// Equal：判断两个 User 实例是否是同一个用户（按ID判断）
func (u *User) Equal(other *User) bool {
    if other == nil {
        return false
    }
    return u.id == other.id
}

// GetUniqueKey：生成唯一标识（用于集合的key）
func (u *User) GetUniqueKey() string {
    return fmt.Sprintf("user_%d", u.id)
}
```

#### 4. 数据校验方法
校验结构体字段是否合法（可在构造/修改后调用）：
```go
// Validate：校验用户数据是否合法
func (u *User) Validate() error {
    if u.id <= 0 {
        return fmt.Errorf("ID不合法：%d", u.id)
    }
    if u.username == "" {
        return fmt.Errorf("用户名不能为空")
    }
    if u.age < 0 || u.age > 150 {
        return fmt.Errorf("年龄不合法：%d", u.age)
    }
    if !strings.Contains(u.email, "@") {
        return fmt.Errorf("邮箱不合法：%s", u.email)
    }
    return nil
}
```

### 四、完整示例：可直接复用的 User 结构体
```go
package main

import (
    "encoding/json"
    "fmt"
    "strings"
)

// User 业务结构体（私有字段保证数据安全）
type User struct {
    id       int64
    username string
    age      int
    email    string
    status   int // 0-禁用，1-正常
}

// ------------------------------
// 1. 构造方法（必选）
// ------------------------------
func NewUser(id int64, username string, age int, email string) (*User, error) {
    user := &User{
        id:       id,
        username: username,
        age:      age,
        email:    email,
        status:   1, // 默认正常
    }
    // 创建时校验数据合法性
    if err := user.Validate(); err != nil {
        return nil, err
    }
    return user, nil
}

// ------------------------------
// 2. Getter 方法（必选）
// ------------------------------
func (u *User) GetID() int64 {
    return u.id
}

func (u *User) GetUsername() string {
    return u.username
}

func (u *User) GetAge() int {
    return u.age
}

func (u *User) GetEmail() string {
    return u.email
}

func (u *User) GetStatusText() string {
    switch u.status {
    case 0:
        return "禁用"
    case 1:
        return "正常"
    default:
        return "未知"
    }
}

// ------------------------------
// 3. Setter 方法（按需封装）
// ------------------------------
func (u *User) SetAge(newAge int) error {
    if newAge < 0 || newAge > 150 {
        return fmt.Errorf("年龄必须在0-150之间")
    }
    u.age = newAge
    return nil
}

func (u *User) SetStatus(newStatus int) error {
    if newStatus != 0 && newStatus != 1 {
        return fmt.Errorf("状态只能是0（禁用）或1（正常）")
    }
    fmt.Printf("用户[%s]状态从%d修改为%d\n", u.username, u.status, newStatus)
    u.status = newStatus
    return nil
}

// ------------------------------
// 4. 业务逻辑方法（按需封装）
// ------------------------------
func (u *User) IsAdult() bool {
    return u.age >= 18
}

func (u *User) IsActive() bool {
    return u.status == 1
}

func (u *User) Equal(other *User) bool {
    if other == nil {
        return false
    }
    return u.id == other.id
}

func (u *User) GetUniqueKey() string {
    return fmt.Sprintf("user_%d", u.id)
}

// ------------------------------
// 5. 工具方法（按需封装）
// ------------------------------
func (u *User) Validate() error {
    if u.id <= 0 {
        return fmt.Errorf("ID不合法：%d", u.id)
    }
    if u.username == "" {
        return fmt.Errorf("用户名不能为空")
    }
    if u.age < 0 || u.age > 150 {
        return fmt.Errorf("年龄不合法：%d", u.age)
    }
    if !strings.Contains(u.email, "@") {
        return fmt.Errorf("邮箱不合法：%s", u.email)
    }
    return nil
}

func (u *User) String() string {
    return fmt.Sprintf("User{ID: %d, 用户名: %s, 年龄: %d, 状态: %s}",
        u.id, u.username, u.age, u.GetStatusText())
}

func (u *User) ToJSON() (string, error) {
    data, err := json.Marshal(map[string]interface{}{
        "id":       u.id,
        "username": u.username,
        "age":      u.age,
        "email":    u.email,
        "status":   u.GetStatusText(),
    })
    if err != nil {
        return "", err
    }
    return string(data), nil
}

// ------------------------------
// 使用示例
// ------------------------------
func main() {
    // 1. 创建用户（构造方法）
    user, err := NewUser(1001, "张三", 20, "zhangsan@example.com")
    if err != nil {
        fmt.Println("创建用户失败：", err)
        return
    }

    // 2. 调用 Getter 方法
    fmt.Println("用户名：", user.GetUsername()) // 输出：张三
    fmt.Println("是否成年：", user.IsAdult())   // 输出：true

    // 3. 调用 Setter 方法
    if err := user.SetAge(21); err != nil {
        fmt.Println("修改年龄失败：", err)
    }
    fmt.Println("修改后年龄：", user.GetAge()) // 输出：21

    // 4. 打印结构体（自动调用 String() 方法）
    fmt.Println(user) // 输出：User{ID: 1001, 用户名: 张三, 年龄: 21, 状态: 正常}

    // 5. 序列化为JSON
    jsonStr, err := user.ToJSON()
    if err != nil {
        fmt.Println("JSON序列化失败：", err)
    } else {
        fmt.Println("JSON：", jsonStr)
    }
}
```

### 总结
1. **必选方法**：
    - 构造方法 `NewXXX()`：初始化+校验，保证实例合法；
    - Getter 方法 `GetXXX()`：暴露私有字段，控制只读权限；
    - Setter 方法 `SetXXX()`：修改私有字段，带合法性校验。
2. **可选方法**：
    - 状态判断（`IsXXX()`）、数据校验（`Validate()`）：提升业务逻辑内聚性；
    - 格式化（`String()`/`ToJSON()`）、比较（`Equal()`）：提升使用便捷性；
    - 唯一标识（`GetUniqueKey()`）：适配集合/存储场景。
3. **封装原则**：
    - 私有字段（小写）+ 公有方法（大写）：保证数据安全；
    - 方法单一职责：一个方法只做一件事（如 `SetAge` 只改年龄，`Validate` 只做校验）；
    - 按需封装：不做无意义的方法（比如简单结构体无需复杂序列化）。

如果有具体的业务结构体（如订单、商品、支付记录），可以告诉我，我会帮你定制化封装适配的方法。