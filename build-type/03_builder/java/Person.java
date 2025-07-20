package main;

public class Person {

    private String name;

    private Integer age;

    private String address;

    public String getName() {
        return name;
    }

    public Integer getAge() {
        return age;
    }

    public String getAddress() {
        return address;
    }

    public void setName(String name) {
        this.name = name;
    }

    public void setAge(Integer age) {
        this.age = age;
    }

    public void setAddress(String address) {
        this.address = address;
    }

    public Person(String name, Integer age, String address) {
        this.name = name;
        this.age = age;
        this.address = address;
    }

    // 建造者模式的常用方式是将实例类和建造者类定义在同一个类文件中，
    // 并提供对外可以获取的 builder 方法用于构建对象。
    public static Builder builder() {

        return new Builder();
    }

    public static class Builder {

        private String name;
        private Integer age;
        private String address;

        // 无参构造私有化
        // 只允许使用 builder 方法来创建 Builder 对象
  			private Builder() {
        }

        public Builder name(String name) {
            this.name = name;
            return this;
        }

        public Builder age(int age) {
            this.age = age;
            return this;
        }

        public Builder address(String address) {
            this.address = address;
            return this;
        }

        // 构建方法，返回 Person 对象
        public Person build() {

            // 验证必填字段
            if (this.name == null || this.name.isEmpty()) {
                throw new IllegalArgumentException("Name cannot be null or empty");
            }

            return new Person(name, age, address);
        }

    }

}
