package Builder.JavaImpl;

class builder {
    public static void main(String[] args) {
        PhoneBuilder phoneBuilder = new PhoneBuilder();
        phoneBuilder.setName("Iphone 15").setMemory(4000).setProcessor("QuadCore").setOS("MAC");

        Phone iphone15 = phoneBuilder.getPhone();

        System.out.println("------ Phone Details ------");
        System.out.println("Name " + iphone15.name);
        System.out.println("OS " + iphone15.os);
        System.out.println("Memory " + iphone15.memory);
        System.out.println("Processor " + iphone15.processor);
        System.out.println("---------------------------");

    }
}
