# 组合模式 （Composite Design Pattern）

组合模式也称为树形、对象树、Object Tree、Composite 模式。将一组对象组织成树形结构，表示 “部分-整体” 的结构层次。

### 使用场景

使用场景较少，当系统里的模型能用树状结构表示时，此模式才有存在的意义。例如：
1. 文件系统目录结构；
2. 公司职级关系；
3. 导航菜单；
4. 浏览器 Dom 对象树等。

指令从最高级（客户端）传递，底层对象依次执行。

### 1.1 优点

1. 更方便的使用 OOP 中的多态和递归机制；
2. 符合设计原则中的开闭原则。

### 1.2 缺点

组合模式只能将一类公共特性的对象，例如文件目录组织称树形结构。
在特定情况下，比如此时加入一个功能差异加大的类，适配已有的对象树时，就会变得难以理解。

## 2. 实现分析

```java
package main;

import java.util.List;
import java.util.ArrayList;
import java.util.Map;
import java.util.HashMap;

// 抽象组件:DOM节点
public abstract class AbstractDom {

    protected String tagName;
    
    protected String content;
    
    public AbstractDom(String tagName) {

        this.tagName = tagName;
    }
    
    public abstract void render();
}

// 叶子节点:文本节点
class TextNode extends AbstractDom {

    public TextNode(String content) {
        super("#text");
        this.content = content;
    }
    
    @Override
    public void render() {
        System.out.print(content);
    }
}

// 组合节点:元素节点
class ElementNode extends AbstractDom {

    private List<AbstractDom> children = new ArrayList<>();
    
    private Map<String, String> attributes = new HashMap<>();
    
    public ElementNode(String tagName) {
        super(tagName);
    }
    
    public void setAttribute(String name, String value) {
        attributes.put(name, value);
    }
    
    public void appendChild(AbstractDom child) {
        children.add(child);
    }
    
    @Override
    public void render() {

        System.out.print("<" + tagName);
        
        // 渲染属性
        for (Map.Entry<String, String> attr : attributes.entrySet()) {
            System.out.print(" " + attr.getKey() + "=\"" + attr.getValue() + "\"");
        }
        System.out.print(">");
        
        // 递归渲染子节点
        for (AbstractDom child : children) {
            child.render();
        }
        
        System.out.print("</" + tagName + ">");
    }
}
```
