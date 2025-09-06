package main;

public class Main {
        public static void main(String[] args) {
            PersonFactory personFactory = new PersonFactory();
            ExternalStateManager stateManager = new ExternalStateManager();

            Person me = personFactory.getPerson("小明", "男", "1990-01-01");
            String personKey = "小明男1990-01-01";

            // 场景1：下雨被淋湿
            stateManager.setState(personKey, new ExternalState("下雨", "新衣服", "地球", "上海市浦东新区"));
            ExternalState state = stateManager.getState(personKey);
            System.out.printf("%s(%s, %s) 今天出门了，处于%s/%s，天气%s，我被淋湿了，我去服装店买了身%s换上。 [me对象地址: %s]\n",
                    me.getName(), me.getGender(), me.getBirth(), state.areaBig, state.areaSmall, state.weather, state.clothes, System.identityHashCode(me));

            // 场景2：打球流汗
            stateManager.setState(personKey, new ExternalState("晴天", "运动服", "地球", "上海市徐汇区"));
            ExternalState state2 = stateManager.getState(personKey);
            System.out.printf("%s(%s, %s) 去打球，处于%s/%s，天气%s，打球流汗了，我去更衣室换了身%s。 [me对象地址: %s]\n",
                    me.getName(), me.getGender(), me.getBirth(), state2.areaBig, state2.areaSmall, state2.weather, state2.clothes, System.identityHashCode(me));
    }
}
