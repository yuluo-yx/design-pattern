package main;

class ProductSelection {
    public void selectProduct() {
        System.out.println("1. 选择商品完成");
    }
}

class Payment {
    public void makePayment() {
        System.out.println("2. 支付完成");
    }
}

class Supplier {
    public void processOrder() {
        System.out.println("3. 供货商确认订单");
    }
}

class Warehouse {
    public void checkStock() {
        System.out.println("4. 仓库备货完成");
    }
}

class Packaging {
    public void packageProduct() {
        System.out.println("5. 商品包装完成");
    }
}

class Delivery {
    public void deliverProduct() {
        System.out.println("6. 开始配送商品");
    }
}

// 2. 创建门面类
public class Facade {

    private ProductSelection productSelection;
    
    private Payment payment;
    
    private Supplier supplier;
    
    private Warehouse warehouse;
    
    private Packaging packaging;
    
    private Delivery delivery;

    public Facade() {
        this.productSelection = new ProductSelection();
        this.payment = new Payment();
        this.supplier = new Supplier();
        this.warehouse = new Warehouse();
        this.packaging = new Packaging();
        this.delivery = new Delivery();
    }

    // 提供一个统一的接口
    public void placeOrder() {

        productSelection.selectProduct();
        payment.makePayment();
        supplier.processOrder();
        warehouse.checkStock();
        packaging.packageProduct();
        delivery.deliverProduct();
    }
}
