# 原型模式 （Prototype Design Pattern）

使用原型实例指定待创建对象的种类，并通过拷贝该原型来创建新的对象。Prototype 模式允许一个原型对象克隆（复制）出多个与其相同的对象，而无需知道任何如何创建的细节。

## 1. 应用

- 对象的创建过程较为复杂且需要频繁创建；
- 期望根据现有的实例来生成新的实例，例如：
  - 对象种类繁多而无法整合到一个类时
  - 难以通过指定类名生成实例时
  - 希望解耦框架与生成的实例时

在实际应用中，Prototype 模式很少单独出现。经常与其他模式混用。

> 1. Java Object 类中有一个 Clone() 方法，可以将 Java 对象复制一份。在对实现了 Cloneable 对象的 Java 类，表示此类可以复制且具有复制的能力。
> 2. Spring 中原型 bean 创建时使用：`@Scope("prototype") // 原型模式 每次注入都会创建一个新的对象实例`。

### 1.1 优点

- 根据客户端要求实现运行时动态创建对象，客户端不需要知道对象的创建细节，便于代码的维护和扩展。
- 当对象创建较为复杂或重复创建大量相似对象时，可简化对象创建，且性能比直接 new 对象更好（new 会自动调用构造链中的所有构造方法，但 clone 不会调用任何类构造方法）。
- Prototype 模式类似工厂模式，但没有工厂模式中的抽象工厂和具体工厂的层级关系，代码结构更清晰和简单。

### 1.2 缺点

- 需要为每个类实现 Cloneable 接口并重写 clone() 方法，改造已有类时必须修改其源码，违背“开闭原则”；
- 单例模式与工厂模式、Prototype 模式是冲突的，尽量不要混用；
- 在实现深拷贝（深克隆）时需要编写较为复杂的代码。

## 2. 实现分析

### 2.1 Java 

Java 中 Object 中提供了 Clone 方法的实现，只需要继承 Cloneable 接口，重写 clone 方法即可。
如果直接调用 Object 父类中的 clone 方法，则实现的是对象浅拷贝。不对引用类型做处理，只是做赋值。

深拷贝需要重写 clone 方法，手动处理引用类型或者使用序列化完成。此处为了避免引入 jackson ，使用 IO 序列化完成，实现了 Serializable 接口。

```java
public class ChatModel implements Serializable, Cloneable {

    @Serial
    private static final long serialVersionUID = 1L;

    private String apiKey;

    private String apiSecret;

    private ChatOption chatOption;

    public ChatModel() {
    }

    public ChatModel(String apiKey, String apiSecret, ChatOption chatOption) {
        this.apiKey = apiKey;
        this.apiSecret = apiSecret;
        this.chatOption = chatOption;
    }

    public String getApiKey() {
        return apiKey;
    }

    public void setApiKey(String apiKey) {
        this.apiKey = apiKey;
    }

    public String getApiSecret() {
        return apiSecret;
    }

    public void setApiSecret(String apiSecret) {
        this.apiSecret = apiSecret;
    }

    public ChatOption getChatOption() {
        return chatOption;
    }

    public void setChatOption(ChatOption chatOption) {
        this.chatOption = chatOption;
    }

    // Object 类中的 clone 实现，不做序列化
    @Override
    public ChatModel clone() {
        try {
            return (ChatModel) super.clone();
        } catch (CloneNotSupportedException e) {
            throw new AssertionError();
        }
    }

    // 深拷贝方法
    // 序列化实现
    public ChatModel deepClone1() {

        var chatModel = new ChatModel();

        try {
            // 序列化
            ByteArrayOutputStream bos = new ByteArrayOutputStream();
            ObjectOutputStream oos = new ObjectOutputStream(bos);
            oos.writeObject(this);
            oos.flush();
            oos.close();

            // 反序列化
            ByteArrayInputStream bis = new ByteArrayInputStream(bos.toByteArray());
            ObjectInputStream ois = new ObjectInputStream(bis);
            chatModel = (ChatModel) ois.readObject();
            return chatModel;
        } catch (IOException | ClassNotFoundException e) {
            throw new RuntimeException("Clone failed", e);
        }
    }

    // 手动处理引用
    public ChatModel deepClone2() {

        var chatModel = new ChatModel();

        // 处理基本类型
        chatModel.setApiKey(this.apiKey);
        chatModel.setApiSecret(this.apiSecret);

        // 处理引用类型
        if (this.chatOption != null) {
            chatModel.setChatOption(this.chatOption.clone());
        }
        return chatModel;
    }

}

class ChatOption implements Serializable, Cloneable {

    @Serial
    private static final long serialVersionUID = 1L;

    private String model;

    public ChatOption() {
    }

    public ChatOption(String model) {
        this.model = model;
    }

    @Override
    public ChatOption clone() {
        try {
            return (ChatOption) super.clone();
        } catch (CloneNotSupportedException e) {
            throw new AssertionError();
        }
    }

}
```

### 2.2 Go

因为 Go 原生支持 JSON 序列化，因此这里直接使用 json 序列化实现深拷贝。

```go
import (
	"encoding/json"
	"time"
)

// 使用对象序列化来完成原型模式的深拷贝

type KeyWord struct {
	Word     string
	Visit    int
	UpdateAt *time.Time
}

// keyword 的 Clone 函数实现
func (k *KeyWord) Clone() *KeyWord {

	var newKeyWords KeyWord

	// 序列化
	b, err := json.Marshal(k)
	if err != nil {
		panic(err)
	}

	// 反序列化
	err = json.Unmarshal(b, &newKeyWords)
	if err != nil {
		panic(err)
	}

	// 新的序列化拷贝之后的对象实例
	return &newKeyWords

}

// keywords 关键字 map
type KeyWords map[string]*KeyWord

// 更新某些热点词

// 深拷贝
// 序列化拷贝所有的引用对象，得到一个新对象，空间地址不同
func (words KeyWords) DeepClone(updateWorks []*KeyWord) KeyWords {

	newWord := KeyWords{}

	for k, v := range words {

		newWord[k] = v.Clone()
	}

	return newWord

}

// 浅拷贝
// 只拷贝了地址，所有的引用对象都是同一个空间地址
func (words KeyWords) ShallowCopy(updateWorks []*KeyWord) KeyWords {

	newWord := KeyWords{}

	for k, v := range words {

		// 浅拷贝
		// 这里只拷贝了地址
		newWord[k] = v

	}

	return newWord

}
```

## 3. 面试题

Q：Java Clone 方法会拷贝引用类型吗？
A：不会，只会拷贝基本数据类型，若要实现深拷贝，必须将 Prototype 模式中的数组、容器对象、引用对象等另行拷贝，（不可变类不用深拷贝，例如 String）。
   例如：深拷贝一个对象时，该对象必须实现 Cloneable 接口并重写 clone() 方法，并在 clone() 方法内部将该对象引用的其他对象也克隆一份。同理，被引用的对象也要做同样的处理。
   
Q：深拷贝与浅拷贝区别
A：浅拷贝：
    1. 对于数据类型是**基本数据类型**的成员变量，浅拷贝会直接进行值传递，也就是将该属性值复制一份给新的对象；
    2. 对于数据类型是**引用数据类型**的成员变量，比如说成员变量是某个数组、某个类的对象等，那么浅拷贝会进行引用传递，也就是只是将该成员变量的引用值（内存地址）复制一份给新的对象。因为实际上两个对象的该成员变量都指向同一个实例。在这种情况下，在一个对象中修改该成员变量会影响到另一个对象的该成员变量值
  深拷贝：
    1. 复制对象的所有基本数据类型的成员变量值；
    2. 为所有引用数据类型的成员变量申请内存，并复制每个引用数据类型成员变量所引用的对象，直到该对象可达的所有对象。既对象深拷贝时要对整个对象（包括对象的引用类型）进行拷贝，
    3. 深拷贝的实现方式：
        1. 实现 Cloneable 接口，重写 Clone 方法；
        2. 对象序列化实现（推荐使用）。

## 4. 扩展

原型模式是 JavaScript 的核心，被称为“语言之魂”。在 JS 中，用于创建重复对象，同时保持性能。它通过复制原型对象来创建新的对象，而不是通过实例化类。
在 JavaScript 中，原型模式是基于原型链实现的，每个对象都有一个指向其原型的指针，而原型对象本身又可以有自己的原型，这样就形成了一个原型链。

```javascript
// 定义父类
function Animal(name) {
    this.name = name;
}

Animal.prototype.speak = function() {
    console.log(`${this.name} makes a noise.`);
};

// 定义子类
function Dog(name) {
    Animal.call(this, name); // 调用父类构造函数
}

// 继承父类的原型
Dog.prototype = Object.create(Animal.prototype);
Dog.prototype.constructor = Dog;

// 重写父类方法
Dog.prototype.speak = function() {
    console.log(`${this.name} barks.`);
};

// 使用示例
const dog = new Dog('Rex');
dog.speak(); // 输出: Rex barks.
```
