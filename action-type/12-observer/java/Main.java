package main;

import java.util.ArrayList;
import java.util.List;

// 订阅者接口
interface Subscriber {
	void receive(String news);
}

// 报社接口
interface NewspaperOffice {
	void subscribe(Subscriber sub);
	void unsubscribe(Subscriber sub);
	void publish(String news);
}

// 具体报社
class NewsAgency implements NewspaperOffice {
	private List<Subscriber> subscribers = new ArrayList<>();

	@Override
	public void subscribe(Subscriber sub) {
		subscribers.add(sub);
	}

	@Override
	public void unsubscribe(Subscriber sub) {
		subscribers.remove(sub);
	}

	@Override
	public void publish(String news) {
		System.out.println("报社发布新闻: " + news);
		for (Subscriber sub : subscribers) {
			sub.receive(news);
		}
	}
}

// 具体订阅者
class Reader implements Subscriber {
	private String name;

	public Reader(String name) {
		this.name = name;
	}

	@Override
	public void receive(String news) {
		System.out.println(name + " 收到报纸: " + news);
	}
}

public class Main {
	public static void main(String[] args) {
		NewsAgency agency = new NewsAgency();

		Reader readerA = new Reader("yuluo");
		Reader readerB = new Reader("musheng");

		agency.subscribe(readerA);
		agency.subscribe(readerB);

		agency.publish("2025年9月6日头条新闻：Java 观察者模式示例！");

		agency.unsubscribe(readerA);
		agency.publish("2025年9月7日头条新闻：只剩下 musheng 在订阅！");
	}
}
