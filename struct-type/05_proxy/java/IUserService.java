package main;

import java.util.HashMap;
import java.util.Map;

interface IUserService {

    User getUser(String name);
}

// 服务的原始功能实现
class ServiceImpl implements IUserService {

    // 模拟数据库服务
    private static final Map<String, User> userDB = new HashMap<>();

    static {
        userDB.put("JDK", new User("JDK", 18));
        userDB.put("CGLIB", new User("CGLIB", 20));
    }

    public User getUser(String name) {

        // 原始类从数据库直接获取用户信息
        // 通过代理类从缓存中获取，减少数据库调用次数。
        System.out.println("Get user from database. key is: " + name);
        return userDB.get(name);
    }
}

record User(String name, Integer age) {
}
