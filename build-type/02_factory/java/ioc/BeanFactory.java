package ioc;

import java.io.InputStream;
import java.lang.annotation.Retention;
import java.lang.annotation.RetentionPolicy;
import java.util.HashMap;
import java.util.Map;
import java.util.Properties;

// IOC 容器实现

public class BeanFactory {

    // IOC 容器
    private final Map<String, Object> beans = new HashMap<>();

    private final Map<String, String> beanDefinitions = new HashMap<>();

    public BeanFactory(String configLocation) {

        loadBeanDefinitions(configLocation);
        instantiateBeans();
        injectDependencies();
    }

    private void loadBeanDefinitions(String configLocation) {

        try (InputStream input = getClass().getClassLoader().getResourceAsStream(configLocation)) {
            Properties properties = new Properties();
            properties.load(input);
            for (String key : properties.stringPropertyNames()) {
                beanDefinitions.put(key, properties.getProperty(key));
            }
        } catch (Exception e) {
            e.printStackTrace();
        }
    }

    private void instantiateBeans() {

        for (Map.Entry<String, String> entry : beanDefinitions.entrySet()) {

            String beanName = entry.getKey();
            String className = entry.getValue();

            try {
                Class<?> clazz = Class.forName(className);
                Object bean = clazz.newInstance();
                beans.put(beanName, bean);
            } catch (Exception e) {
                e.printStackTrace();
            }
        }
    }

    // 依赖注入
    // userController 需要 userService 的注入
    // userService 需要 userDAO 的注入
    private void injectDependencies() {

        for (Object bean : beans.values()) {
            for (java.lang.reflect.Field field : bean.getClass().getDeclaredFields()) {

                if (field.isAnnotationPresent(Autowired.class)) {

                    field.setAccessible(true);
                    String dependencyName = field.getName();
                    Object dependency = getBean(dependencyName);

                    try {
                        field.set(bean, dependency);
                    } catch (IllegalAccessException e) {
                        e.printStackTrace();
                    }
                }
            }
        }
    }

    public Object getBean(String beanName) {
        return beans.get(beanName);
    }

    // 模拟 @Autowired 注解
    @Retention(RetentionPolicy.RUNTIME)
    @interface Autowired {}
}
