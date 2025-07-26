package ioc;

// 用户服务
// 依赖于 userDAO

public interface UserService {

    void login(String username, String password);
}

class UserServiceImpl implements UserService {

    @BeanFactory.Autowired
    private UserDAO userDao;

    public void login(String username, String password) {

        userDao.login(username, password);
    }

}
