package main;

import java.util.List;
import java.util.ArrayList;
import java.util.Map;
import java.util.HashMap;

// 抽象组件: DOM 节点
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
