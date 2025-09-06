package main;

import java.util.HashMap;
import java.util.Map;

// 享元对象：Person（只存内部状态）
public class Person {
    private final String name;
    private final String gender;
    private final String birth;

    public Person(String name, String gender, String birth) {
        this.name = name;
        this.gender = gender;
        this.birth = birth;
    }

    public String getName() { return name; }
    public String getGender() { return gender; }
    public String getBirth() { return birth; }
}

// 外部状态
class ExternalState {
    String weather;
    String clothes;
    String areaBig;
    String areaSmall;

    public ExternalState(String weather, String clothes, String areaBig, String areaSmall) {
        this.weather = weather;
        this.clothes = clothes;
        this.areaBig = areaBig;
        this.areaSmall = areaSmall;
    }
}

// 外部状态容器（全局管理外部状态）
class ExternalStateManager {
    private final Map<String, ExternalState> stateMap = new HashMap<>();

    public void setState(String personKey, ExternalState state) {
        stateMap.put(personKey, state);
    }

    public ExternalState getState(String personKey) {
        return stateMap.get(personKey);
    }
}

// 享元工厂
class PersonFactory {
    private final Map<String, Person> persons = new HashMap<>();

    public Person getPerson(String name, String gender, String birth) {
        String key = name + gender + birth;
        if (!persons.containsKey(key)) {
            persons.put(key, new Person(name, gender, birth));
        }
        return persons.get(key);
    }
}
