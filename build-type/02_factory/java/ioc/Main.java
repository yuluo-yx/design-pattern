package ioc;

// 模型 spring 控制器调用 service，进而调用依赖注入的 DAO
// 将 controller 和 service 和 dao 都交由 IOC 容器管理
class UserController {

    @BeanFactory.Autowired
    private UserService userService;

    public void login(String username, String passwd) {

        userService.login(username, passwd);
    }

}

public class Main {

    public static void main(String[] args) {

        // bean 初始化
        BeanFactory beanFactory = new BeanFactory("beans.properties");

        // bean 注入 在 IOC 容器中完成
        // injectDependencies();

        // 尝试获取 bean
        UserController userController = (UserController) beanFactory.getBean("userController");

        System.out.println("=========================================================================");
        System.out.println("UserController initialized: " + userController);
        UserService userService = (UserService) beanFactory.getBean("userService");
        System.out.println("UserService initialized: " + userService);
        UserDAO userDAO = (UserDAO) beanFactory.getBean("userDao");
        System.out.println("UserDAO initialized: " + userDAO);
        System.out.println("=========================================================================");

        userController.login("admin", "123456");
    }

}
