sealed class A permits B, C {
    public void sayHello() {
        System.out.println("Hello");
    }
}

final class B extends A {
}

final class C extends A {
}

class X {
}


class Hello {
    public static void main(String[] arguments) {

    }
}
