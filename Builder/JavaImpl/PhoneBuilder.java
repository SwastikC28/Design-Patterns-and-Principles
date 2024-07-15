package Builder.JavaImpl;


public class PhoneBuilder {
    String name;
    String os;
    String processor;
    int memory;

    public PhoneBuilder setName(String name) {
        this.name = name;
        return this;
    }

    public PhoneBuilder setOS(String os) {
        this.os = os;
        return this;
    }

    public PhoneBuilder setProcessor(String processor) {
        this.processor = processor;
        return this;
    }

    public PhoneBuilder setMemory(int memory) {
        this.memory = memory;
        return this;
    }

    public Phone getPhone() {
        return new Phone(this.name, this.os, this.processor, this.memory);
    }
};
