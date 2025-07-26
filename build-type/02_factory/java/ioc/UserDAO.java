package ioc;

// 用户 DAO 层服务

public interface UserDAO {

    // 模型查询数据库操作
    void login(String username, String password);
}

class UserDAOImpl implements UserDAO {

    public void login(String username, String password) {

        // 模拟登录逻辑
        System.out.println("query db, User " + username + " logged in with password: " + password);
    }

}
