package main;

public class Main {
    
    public static void main(String[] args) {
       
        // 创建一个简单的 DOM 树
        ElementNode html = new ElementNode("html");
        
        ElementNode body = new ElementNode("body");
        
        ElementNode div = new ElementNode("div");
        div.setAttribute("class", "container");
        
        ElementNode p = new ElementNode("p");
        p.appendChild(new TextNode("Hello, World!"));
        
        div.appendChild(p);
        body.appendChild(div);
        html.appendChild(body);
        
        // 渲染 DOM 树
        html.render();
    }
}
